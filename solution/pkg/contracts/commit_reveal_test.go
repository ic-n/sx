package contracts

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ic-n/sx/pkg/tools"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var _ = Describe("Contract test", Label("sol"), func() {
	var (
		auth *bind.TransactOpts
		cli  *backends.SimulatedBackend
		c    *CommitReveal
	)

	BeforeEach(func() {
		var err error
		auth, cli, c, err = testContract()
		Expect(err)
	})

	It("should record a proper commit", func(ctx SpecContext) {
		vote := tools.B32("1~mybigsecret")
		_, err := c.CommitVote(auth, vote)
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		data, err := c.GetVoteCommitsArray(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(tools.UnB32(data[0])).Should(Equal("1~mybigsecret"))

		voteStatus, err := c.VoteStatuses(nil, vote)
		Expect(err).NotTo(HaveOccurred())
		Expect(voteStatus).Should(Equal("Committed"))

		numberOfVotes, err := c.NumberOfVotesCast(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(numberOfVotes.Int64()).Should(BeEquivalentTo(1))
	})

	It("should record proper reveals", func(ctx SpecContext) {
		const vote1 = "1~mybigsecret"
		const vote2 = "2~mybigsecret"
		commit1 := crypto.Keccak256([]byte(vote1))
		commit2 := crypto.Keccak256([]byte(vote2))

		_, err := c.CommitVote(auth, [32]byte(commit1))
		Expect(err).NotTo(HaveOccurred())
		_, err = c.CommitVote(auth, [32]byte(commit2))
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		err = cli.AdjustTime(time.Duration(phaseSeconds) * time.Second)
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		_, err = c.RevealVote(auth, vote1, [32]byte(commit1))
		Expect(err).NotTo(HaveOccurred())
		_, err = c.RevealVote(auth, vote2, [32]byte(commit2))
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		votesFor1, err := c.VotesForChoice1(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(votesFor1.Int64()).Should(BeEquivalentTo(1))

		votesFor2, err := c.VotesForChoice2(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(votesFor2.Int64()).Should(BeEquivalentTo(1))

		voteStatusCommit1, err := c.VoteStatuses(nil, [32]byte(commit1))
		Expect(err).NotTo(HaveOccurred())
		Expect(voteStatusCommit1).Should(Equal("Revealed"))

		voteStatusCommit2, err := c.VoteStatuses(nil, [32]byte(commit1))
		Expect(err).NotTo(HaveOccurred())
		Expect(voteStatusCommit2).Should(Equal("Revealed"))
	})

	It("should not record an incorrect reveal", func(ctx SpecContext) {
		const badVote = "3~mybigsecret"
		_, err := c.CommitVote(auth, tools.B32(badVote))
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		err = cli.AdjustTime(time.Duration(phaseSeconds) * time.Second)
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		tx, err := c.RevealVote(auth, badVote, tools.B32(badVote))
		Expect(err).To(MatchError("execution reverted: Vote hash does not match vote commit."))
		Expect(tx).To(BeNil())
	})

	It("should correctly get winner", func(ctx SpecContext) {
		const vote = "1~mybigsecret"
		commit := crypto.Keccak256([]byte(vote))
		c.CommitVote(auth, [32]byte(commit))
		cli.Commit()

		err := cli.AdjustTime(time.Duration(phaseSeconds) * time.Second)
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		_, err = c.RevealVote(auth, vote, [32]byte(commit))
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		winner, err := c.GetWinner(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(winner).Should(Equal("YES"))
	})
})

const (
	simChainID = 1337
	gasLim     = 4712388

	phaseSeconds = 120
)

func testContract() (*bind.TransactOpts, *backends.SimulatedBackend, *CommitReveal, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to generate key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(simChainID))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create transactor: %w", err)
	}
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: tools.Wei(1000)}
	client := backends.NewSimulatedBackend(alloc, gasLim)

	_, _, contract, err := DeployCommitReveal(auth, client, big.NewInt(phaseSeconds), "YES", "NO")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to deploy contract: %w", err)
	}

	client.Commit()

	return auth, client, contract, nil
}
