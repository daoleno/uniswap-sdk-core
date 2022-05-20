package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	assert.Panics(t, func() { NewToken(3, AddressOne, 256, "", "") }, "fails with 256 decimals")
}

func TestTokenEquals(t *testing.T) {
	assert.False(t, NewToken(1, AddressOne, 18, "", "").Equal(NewToken(1, AddressTwo, 18, "", "")), "fails if address differs")
	assert.False(t, NewToken(3, AddressOne, 18, "", "").Equal(NewToken(1, AddressOne, 18, "", "")), "fails if chain id differs")
	assert.True(t, NewToken(1, AddressOne, 9, "", "").Equal(NewToken(1, AddressOne, 18, "", "")), "true if only decimals differs")
	assert.True(t, NewToken(1, AddressOne, 18, "", "").Equal(NewToken(1, AddressOne, 18, "", "")), "true if address is the same")

	token := NewToken(1, AddressOne, 18, "", "")
	assert.True(t, token.Equal(token), "true on reference equality")

	tokenA := NewToken(1, AddressOne, 9, "abc", "def")
	tokenB := NewToken(1, AddressOne, 18, "ghi", "jkl")
	assert.True(t, tokenA.Equal(tokenB), "true even if name/symbol/decimals differ")
}
