package entities

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var AddressOne = common.HexToAddress("0x0000000000000000000000000000000000000001")

func TestFromRawAmount(t *testing.T) {
	token := NewToken(1, AddressOne, 18, "", "")
	amount := FromRawAmount(token.Currency, big.NewInt(100))
	assert.Equal(t, amount.Quotient(), big.NewInt(100), "works")
}

func TestMultiply(t *testing.T) {
	token := NewToken(1, AddressOne, 18, "", "")
	amount := FromRawAmount(token.Currency, big.NewInt(100)).Multiply(NewPercent(big.NewInt(15), big.NewInt(100)).Fraction)
	assert.Equal(t, amount.Quotient(), big.NewInt(15), "returns the amount after multiplication")
}

func TestEther(t *testing.T) {
	amount := FromRawAmount(EtherOnChain(1).Currency, big.NewInt(100))
	assert.Equal(t, amount.Quotient(), big.NewInt(100), "produces ether amount.quotient")
	assert.Equal(t, amount.Currency, EtherOnChain(1).Currency, "produces ether amount.currency")
}

func TestMaxTokenAmount(t *testing.T) {
	amount := FromRawAmount(NewToken(1, AddressOne, 18, "", "").Currency, MaxUint256)
	assert.Equal(t, amount.Quotient(), MaxUint256, "token amount can be max uint256")

	assert.Panics(t, func() {
		FromRawAmount(NewToken(1, AddressOne, 18, "", "").Currency, new(big.Int).Add(MaxUint256, big.NewInt(1)))
	}, "token amount cannot exceed max uint256'")

	assert.Panics(t, func() {
		FromFractionalAmount(NewToken(1, AddressOne, 18, "", "").Currency, new(big.Int).Add(new(big.Int).Mul(MaxUint256, big.NewInt(2)), big.NewInt(2)), big.NewInt(2))
	}, "token amount quotient cannot exceed max uint256")

	amount1 := FromFractionalAmount(NewToken(1, AddressOne, 18, "", "").Currency, new(big.Int).Add(MaxUint256, big.NewInt(2)), big.NewInt(2))
	assert.Equal(t, amount1.Fraction.Numerator, new(big.Int).Add(MaxUint256, big.NewInt(2)), "token amount numerator can be gt. uint256 if denominator is gt. 1")
}
