package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEther(t *testing.T) {
	assert.Equal(t, EtherOnChain(1), EtherOnChain(1))
	assert.NotEqual(t, EtherOnChain(1), EtherOnChain(2))
	assert.False(t, EtherOnChain(1).Equal(EtherOnChain(2)))
	assert.True(t, EtherOnChain(1).Equal(EtherOnChain(1)))
}
