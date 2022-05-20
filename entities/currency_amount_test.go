package entities

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromRawAmount(t *testing.T) {
	token := NewToken(1, AddressOne, 18, "", "")
	amount := FromRawAmount(token, big.NewInt(100))
	assert.Equal(t, amount.Quotient(), big.NewInt(100), "works")
}

func TestMultiply(t *testing.T) {
	token := NewToken(1, AddressOne, 18, "", "")
	amount := FromRawAmount(token, big.NewInt(100)).Multiply(NewPercent(big.NewInt(15), big.NewInt(100)).Fraction)
	assert.Equal(t, amount.Quotient(), big.NewInt(15), "returns the amount after multiplication")
}

func TestEtherAmount(t *testing.T) {
	amount := FromRawAmount(EtherOnChain(1), big.NewInt(100))
	assert.Equal(t, amount.Quotient(), big.NewInt(100), "produces ether amount.quotient")
	assert.Equal(t, amount.Currency, EtherOnChain(1), "produces ether amount.currency")
}

func TestMaxTokenAmount(t *testing.T) {
	amount := FromRawAmount(NewToken(1, AddressOne, 18, "", ""), MaxUint256)
	assert.Equal(t, amount.Quotient(), MaxUint256, "token amount can be max uint256")

	assert.Panics(t, func() {
		FromRawAmount(NewToken(1, AddressOne, 18, "", ""), new(big.Int).Add(MaxUint256, big.NewInt(1)))
	}, "token amount cannot exceed max uint256'")

	assert.Panics(t, func() {
		FromFractionalAmount(NewToken(1, AddressOne, 18, "", ""), new(big.Int).Add(new(big.Int).Mul(MaxUint256, big.NewInt(2)), big.NewInt(2)), big.NewInt(2))
	}, "token amount quotient cannot exceed max uint256")

	amount1 := FromFractionalAmount(NewToken(1, AddressOne, 18, "", ""), new(big.Int).Add(MaxUint256, big.NewInt(2)), big.NewInt(2))
	assert.Equal(t, amount1.Fraction.Numerator, new(big.Int).Add(MaxUint256, big.NewInt(2)), "token amount numerator can be gt. uint256 if denominator is gt. 1")
}

func TestToFixed(t *testing.T) {
	token0 := NewToken(1, AddressOne, 0, "", "")
	amount0 := FromRawAmount(token0, big.NewInt(1000))
	assert.Panics(t, func() { amount0.ToFixed(3) }, "panics for decimals > currency.decimals")

	token1 := NewToken(1, AddressOne, 0, "", "")
	amount1 := FromRawAmount(token1, big.NewInt(123456))
	assert.Equal(t, amount1.ToFixed(0), "123456", "is correct for 0 decimals'")

	token2 := NewToken(1, AddressOne, 18, "", "")
	amount2 := FromRawAmount(token2, big.NewInt(1e15))
	assert.Equal(t, amount2.ToFixed(9), "0.001000000", "is correct for 18 decimals")
}

func TestToSignificant(t *testing.T) {
	token0 := NewToken(1, AddressOne, 0, "", "")
	amount0 := FromRawAmount(token0, big.NewInt(1000))
	assert.Equal(t, amount0.ToSignificant(3), "1000", "does not panic for sig figs > currency.decimals'")

	token1 := NewToken(1, AddressOne, 0, "", "")
	amount1 := FromRawAmount(token1, big.NewInt(123456))
	// TODO: support rounding, here in v3-sdk is 123400
	assert.Equal(t, amount1.ToSignificant(4), "123500", "is correct for 0 decimals")

	token2 := NewToken(1, AddressOne, 18, "", "")
	amount2 := FromRawAmount(token2, big.NewInt(1e15))
	assert.Equal(t, amount2.ToSignificant(9), "0.001", "is correct for 18 decimals")
}

func TestToExact(t *testing.T) {
	token0 := NewToken(1, AddressOne, 0, "", "")
	amount0 := FromRawAmount(token0, big.NewInt(1000))
	assert.Equal(t, amount0.ToExact(), "1000", "does not panic for sig figs > currency.decimals'")

	token1 := NewToken(1, AddressOne, 0, "", "")
	amount1 := FromRawAmount(token1, big.NewInt(123456))
	assert.Equal(t, amount1.ToExact(), "123456", "is correct for 0 decimals")

	token2 := NewToken(1, AddressOne, 18, "", "")
	amount2 := FromRawAmount(token2, big.NewInt(123e13))
	assert.Equal(t, amount2.ToExact(), "0.00123", "is correct for 18 decimals")
}
