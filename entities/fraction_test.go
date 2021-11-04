package entities

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractionQuotient(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(8), big.NewInt(3)).Quotient(), big.NewInt(2), "floor division - one below")
	assert.Equal(t, NewFraction(big.NewInt(12), big.NewInt(4)).Quotient(), big.NewInt(3), "floor division - exact")
	assert.Equal(t, NewFraction(big.NewInt(16), big.NewInt(5)).Quotient(), big.NewInt(3), "floor division - one above")
}

func TestFractionRemainder(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(8), big.NewInt(3)).Remainder(), NewFraction(big.NewInt(2), big.NewInt(3)), "returns fraction after division - 8/3")
	assert.Equal(t, NewFraction(big.NewInt(12), big.NewInt(4)).Remainder(), NewFraction(big.NewInt(0), big.NewInt(4)), "returns fraction after division - 12/4")
	assert.Equal(t, NewFraction(big.NewInt(16), big.NewInt(5)).Remainder(), NewFraction(big.NewInt(1), big.NewInt(5)), "returns fraction after division - 16/5")
}

func TestFractionInvert(t *testing.T) {
	// flips num and denom
	assert.Equal(t, NewFraction(big.NewInt(5), big.NewInt(10)).Invert().Numerator, big.NewInt(10), "flips num and denom - num")
	assert.Equal(t, NewFraction(big.NewInt(5), big.NewInt(10)).Invert().Denominator, big.NewInt(5), "flips num and denom - denom")
}

func TestFractionAdd(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(10)).Add(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(52), big.NewInt(120)), "multiples denoms and adds nums")
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(5)).Add(NewFraction(big.NewInt(2), big.NewInt(5))), NewFraction(big.NewInt(3), big.NewInt(5)), "same denom")
}

func TestFractionSubtract(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(10)).Subtract(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(-28), big.NewInt(120)), "multiples denoms and subtracts nums")
	assert.Equal(t, NewFraction(big.NewInt(3), big.NewInt(5)).Subtract(NewFraction(big.NewInt(2), big.NewInt(5))), NewFraction(big.NewInt(1), big.NewInt(5)), "same denom")
}

func TestFractionLessThan(t *testing.T) {
	assert.True(t, NewFraction(big.NewInt(1), big.NewInt(10)).LessThan(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.False(t, NewFraction(big.NewInt(1), big.NewInt(3)).LessThan(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.False(t, NewFraction(big.NewInt(5), big.NewInt(12)).LessThan(NewFraction(big.NewInt(4), big.NewInt(12))))
}

func TestFractionEqualTo(t *testing.T) {
	assert.False(t, NewFraction(big.NewInt(1), big.NewInt(10)).EqualTo(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.True(t, NewFraction(big.NewInt(1), big.NewInt(3)).EqualTo(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.False(t, NewFraction(big.NewInt(5), big.NewInt(12)).EqualTo(NewFraction(big.NewInt(4), big.NewInt(12))))
}

func TestFractionGreaterThan(t *testing.T) {
	assert.False(t, NewFraction(big.NewInt(1), big.NewInt(10)).GreaterThan(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.False(t, NewFraction(big.NewInt(1), big.NewInt(3)).GreaterThan(NewFraction(big.NewInt(4), big.NewInt(12))))
	assert.True(t, NewFraction(big.NewInt(5), big.NewInt(12)).GreaterThan(NewFraction(big.NewInt(4), big.NewInt(12))))
}

func TestFractionMultiply(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(10)).Multiply(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(4), big.NewInt(120)))
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(3)).Multiply(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(4), big.NewInt(36)))
	assert.Equal(t, NewFraction(big.NewInt(5), big.NewInt(12)).Multiply(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(20), big.NewInt(144)))
}

func TestFractionDivide(t *testing.T) {
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(10)).Divide(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(12), big.NewInt(40)))
	assert.Equal(t, NewFraction(big.NewInt(1), big.NewInt(3)).Divide(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(12), big.NewInt(12)))
	assert.Equal(t, NewFraction(big.NewInt(5), big.NewInt(12)).Divide(NewFraction(big.NewInt(4), big.NewInt(12))), NewFraction(big.NewInt(60), big.NewInt(48)))
}

func TestFractionToSignificant(t *testing.T) {
	f := NewFraction(big.NewInt(126), big.NewInt(10))
	assert.Equal(t, f.ToSignificant(3), "12.6")
	assert.Equal(t, f.ToSignificant(1), "10")
}

func TestFractionToFixed(t *testing.T) {
	f := NewFraction(big.NewInt(126), big.NewInt(10))
	assert.Equal(t, f.ToFixed(0), "13")
	assert.Equal(t, f.ToFixed(1), "12.6")
}
