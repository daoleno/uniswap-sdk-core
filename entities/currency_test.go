package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	assert.True(t, EtherOnChain(1).Equals(EtherOnChain(1).Currency), "ether on same chains is ether")
	assert.False(t, EtherOnChain(1).Equals(t0.Currency), "ether is not token0")
	assert.False(t, t1.Equals(t0), "token1 is not token0")
	assert.True(t, t0.Equals(t0), "token0 is token0")
	assert.True(t, t0.Equals(NewToken(1, AddressZero, 18, "symbol", "name")), "token0 is equal to another token0")
}
