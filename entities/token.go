package entities

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrDifferentChain = errors.New("different chain")
	ErrSameAddress    = errors.New("same address")
)

// Represents an ERC20 token with a unique address and some metadata.
type Token struct {
	*Currency
	Address common.Address // The contract address on the chain on which this token lives
}

// NewToken creates a new token with the given currency and address.
func NewToken(chainID uint, address common.Address, decimals uint, symbol string, name string) *Token {
	return &Token{
		Currency: NewTokenCurrency(chainID, decimals, symbol, name),
		Address:  address,
	}
}

// NewNativeToken creates a new native token with the given currency and address.
func NewNativeToken(chainID uint, address common.Address, decimals uint, symbol string, name string) *Token {
	return &Token{
		Currency: NewBaseCurrency(chainID, decimals, symbol, name),
		Address:  address,
	}
}

/**
 * Returns true if the two tokens are equivalent, i.e. have the same chainId and address.
 * @param other other token to compare
 */

func (t *Token) Equals(other *Token) bool {
	return t.ChainID == other.ChainID && t.Address == other.Address
}

/**
 * Returns true if the address of this token sorts before the address of the other token
 * @param other other token to compare
 * @throws if the tokens have the same address
 * @throws if the tokens are on different chains
 */

func (t *Token) SortsBefore(other *Token) (bool, error) {
	if t.ChainID != other.ChainID {
		return false, ErrDifferentChain
	}
	if t.Address == other.Address {
		return false, ErrSameAddress
	}
	return strings.ToLower(t.Address.Hex()) < strings.ToLower(other.Address.Hex()), nil
}
