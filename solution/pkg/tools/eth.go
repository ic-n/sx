package tools

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func B32(s string) [32]byte {
	v := [32]byte{}
	copy(v[:], []byte(s))
	return v
}

func UnB32(v [32]byte) string {
	return strings.TrimRight(string(v[:]), "\x00")
}

func Keccak256(s string) [32]byte {
	return [32]byte(crypto.Keccak256([]byte(s)))
}

func Wei(eth float64) *big.Int {
	weiMultiplier := new(big.Int)
	weiMultiplier.Exp(big.NewInt(10), big.NewInt(18), nil)

	ethInWei := new(big.Float).Mul(big.NewFloat(eth), new(big.Float).SetInt(weiMultiplier))
	wei := new(big.Int)
	ethInWei.Int(wei)

	return wei
}
