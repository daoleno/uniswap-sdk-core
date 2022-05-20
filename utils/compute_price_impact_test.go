package utils

import (
	"math/big"
	"testing"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestComputePriceImpact(t *testing.T) {
	var (
		AddressZero = common.HexToAddress("0x0000000000000000000000000000000000000000")
		AddressOne  = common.HexToAddress("0x0000000000000000000000000000000000000001")

		t0 = entities.NewToken(1, AddressZero, 18, "", "")
		t1 = entities.NewToken(1, AddressOne, 18, "", "")
	)
	impact0, err := ComputePriceImpact(
		entities.NewPrice(entities.EtherOnChain(1), t0, big.NewInt(10), big.NewInt(100)),
		entities.FromRawAmount(entities.EtherOnChain(1), big.NewInt(10)),
		entities.FromRawAmount(t0, big.NewInt(100)))

	if err != nil {
		panic(err)
	}
	assert.Equal(t, impact0, entities.NewPercent(big.NewInt(0), big.NewInt(10000)), "is correct for zero")

	impact1, err := ComputePriceImpact(
		entities.NewPrice(t0, t1, big.NewInt(10), big.NewInt(100)),
		entities.FromRawAmount(t0, big.NewInt(10)),
		entities.FromRawAmount(t1, big.NewInt(50)))

	if err != nil {
		panic(err)
	}
	assert.Equal(t, impact1, entities.NewPercent(big.NewInt(5000), big.NewInt(10000)), "is correct for half output")

	impact2, err := ComputePriceImpact(
		entities.NewPrice(t0, t1, big.NewInt(10), big.NewInt(100)),
		entities.FromRawAmount(t0, big.NewInt(10)),
		entities.FromRawAmount(t1, big.NewInt(200)))

	if err != nil {
		panic(err)
	}
	assert.Equal(t, impact2, entities.NewPercent(big.NewInt(-10000), big.NewInt(10000)), "is negative for more output")

}
