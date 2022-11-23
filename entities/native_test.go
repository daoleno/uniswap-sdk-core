package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNative(t *testing.T) {
	wrapped0 := NewToken(1, common.Address{}, 18, "WCNATIVE", "Wrapped Custom Native Token")
	wrapped1 := NewToken(2, common.Address{}, 18, "WCNATIVE", "Wrapped Custom Native Token")
	n0 := NewNative(wrapped0, "CNATIVE", "Custom Native Token")
	n0Copy := NewNative(wrapped0, "CNATIVE", "Custom Native Token")
	n1 := NewNative(wrapped1, "CNATIVE", "Custom Native Token")
	assert.Equal(t, n0, n0Copy)
	assert.NotEqual(t, n0, n1)
	assert.False(t, n0.Equal(n1))
	assert.True(t, n0.Equal(n0Copy))
}
