package entities

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	AddressZero = common.HexToAddress("0x0000000000000000000000000000000000000000")
	AddressOne  = common.HexToAddress("0x0000000000000000000000000000000000000001")
	AddressTwo  = common.HexToAddress("0x0000000000000000000000000000000000000002")

	t0   = NewToken(1, AddressZero, 18, "", "")
	t0_6 = NewToken(1, AddressZero, 6, "", "")
	t1   = NewToken(1, AddressOne, 18, "", "")
)

func TestNewPrice(t *testing.T) {
	price0 := NewPrice(t0, t1, big.NewInt(1), big.NewInt(54321))
	assert.Equal(t, price0.ToSignificant(5), "54321")
	assert.True(t, price0.BaseCurrency.Equal(t0))
	assert.True(t, price0.QuoteCurrency.Equal(t1))
}

func TestQuote(t *testing.T) {
	price := NewPrice(t0, t1, big.NewInt(1), big.NewInt(5))
	q, err := price.Quote(FromRawAmount(t0, big.NewInt(10)))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, q, FromRawAmount(t1, big.NewInt(50)))
}

func TestPriceToSignificant(t *testing.T) {
	price0 := NewPrice(t0, t1, big.NewInt(123), big.NewInt(456))
	assert.Equal(t, price0.ToSignificant(4), "3.707", "no decimals")

	price1 := NewPrice(t0, t1, big.NewInt(456), big.NewInt(123))
	assert.Equal(t, price1.ToSignificant(4), "0.2697", "no decimals flip ratio")

	price2 := NewPrice(t0_6, t1, big.NewInt(123), big.NewInt(456))
	assert.Equal(t, price2.ToSignificant(4), "0.000000000003707", "with decimal difference")

	price3 := NewPrice(t0_6, t1, big.NewInt(456), big.NewInt(123))
	assert.Equal(t, price3.ToSignificant(4), "0.0000000000002697", "with decimal difference flipped")

	price4 := NewPrice(t1, t0_6, big.NewInt(456), big.NewInt(123))
	assert.Equal(t, price4.ToSignificant(4), "269700000000", "with decimal difference flipped base quote flipped")
}
