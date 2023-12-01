package contracts

import (
	"fmt"
	"math/big"
	"strings"
	"time"

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
		vote := b32("1~mybigsecret")
		_, err := c.CommitVote(auth, vote)
		Expect(err).NotTo(HaveOccurred())
		cli.Commit()

		data, err := c.GetVoteCommitsArray(nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(unb32(data[0])).Should(Equal("1~mybigsecret"))

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

		err = cli.AdjustTime(time.Duration(PhaseLengthInSeconds.Int64()) * time.Second)
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

	//   it("should not record an incorrect reveal", async () => {
	It("should not record an incorrect reveal", func(ctx SpecContext) {})
	// 	const badVote = "3~mybigsecret";
	// 	const commit = ethers.encodeBytes32String(badVote);
	// 	await commitReveal.commitVote(commit);
	// 	await ethers.provider.send("evm_increaseTime", [PHASE_LENGTH_IN_SECONDS]);
	// 	const tx = commitReveal.revealVote(badVote, commit);
	// 	await expect(tx).to.be.rejectedWith(Error);
	//   });

	//   it("should correctly get winner", async () => {
	It("should correctly get winner", func(ctx SpecContext) {})
	// 	const vote = "1~mybigsecret";
	// 	const commit = ethers.keccak256(ethers.toUtf8Bytes(vote));
	// 	await commitReveal.commitVote(commit);
	// 	await ethers.provider.send("evm_increaseTime", [PHASE_LENGTH_IN_SECONDS]);
	// 	await commitReveal.revealVote(vote, commit);
	// 	const winner = await commitReveal.getWinner();
	// 	expect(winner).to.equal("YES");
	//   });
})

func b32(s string) [32]byte {
	v := [32]byte{}
	copy(v[:], []byte(s))
	return v
}

func unb32(v [32]byte) string {
	return strings.TrimRight(string(v[:]), "\x00")
}

func ConvertEthToWei(ethAmount float64) *big.Int {
	weiMultiplier := new(big.Int)
	weiMultiplier.Exp(big.NewInt(10), big.NewInt(18), nil)

	ethInWei := new(big.Float).Mul(big.NewFloat(ethAmount), new(big.Float).SetInt(weiMultiplier))
	wei := new(big.Int)
	ethInWei.Int(wei)

	return wei
}

var PhaseLengthInSeconds = big.NewInt(120)

const (
	simChainID = 1337
	gasLim     = 4712388
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
	alloc[auth.From] = core.GenesisAccount{Balance: ConvertEthToWei(1000)}
	client := backends.NewSimulatedBackend(alloc, gasLim)

	_, _, contract, err := DeployCommitReveal(auth, client, PhaseLengthInSeconds, "YES", "NO")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to deploy contract: %w", err)
	}

	client.Commit()

	return auth, client, contract, nil
}
