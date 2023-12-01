package contracts

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

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

func TestDeployInbox(t *testing.T) {
	t.Run("record a proper commit", func(t *testing.T) {
		auth, cli, c, err := testContract()
		require.NoError(t, err)

		vote := b32("1~mybigsecret")
		tx, err := c.CommitVote(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: auth.GasLimit,
		}, vote)
		require.NoError(t, err, tx)
		cli.Commit()

		data, err := c.GetVoteCommitsArray(nil)
		require.NoError(t, err)
		require.EqualValues(t, "1~mybigsecret", unb32(data[0]))

		voteStatus, err := c.VoteStatuses(nil, vote)
		require.NoError(t, err)
		require.Equal(t, voteStatus, "Committed")

		numberOfVotes, err := c.NumberOfVotesCast(nil)
		require.NoError(t, err)
		require.EqualValues(t, 1, numberOfVotes.Int64())
	})
}

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
