package entities

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPercent(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(1)), toPercent(NewFraction(big.NewInt(1), big.NewInt(1))))
}

func TestPercentAdd(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(100)).Add(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(3), big.NewInt(100)), "returns a percent")
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(25)).Add(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(150), big.NewInt(2500)), "different denominators")
}

func TestPercentSubtract(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(100)).Subtract(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(-1), big.NewInt(100)), "returns a percent")
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(25)).Subtract(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(50), big.NewInt(2500)), "different denominators")
}

func TestPercentMultiply(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(100)).Multiply(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(2), big.NewInt(10000)), "returns a percent")
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(25)).Multiply(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(2), big.NewInt(2500)), "different denominators")
}

func TestPercentDivide(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(100)).Divide(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(100), big.NewInt(200)), "returns a percent")
	assert.Equal(t, NewPercent(big.NewInt(1), big.NewInt(25)).Divide(NewPercent(big.NewInt(2), big.NewInt(100))), NewPercent(big.NewInt(100), big.NewInt(50)), "different denominators")
}

func TestPercentToSignificant(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(154), big.NewInt(10000)).ToSignificant(3), "1.54", "returns the value scaled by 100")
}

func TestPercentToFixed(t *testing.T) {
	assert.Equal(t, NewPercent(big.NewInt(154), big.NewInt(10000)).ToFixed(2), "1.54", "returns the value scaled by 100")
}
