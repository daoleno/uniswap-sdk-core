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

// Token represents an ERC20 token with a unique address and some metadata.
type Token struct {
	*BaseCurrency
	Address common.Address // The contract address on the chain on which this token lives
}

// NewToken creates a new token with the given currency and address.
func NewToken(chainID uint, address common.Address, decimals uint, symbol string, name string) *Token {
	if decimals >= 255 {
		panic("Token currency decimals must be less than 255")
	}
	token := &Token{
		BaseCurrency: &BaseCurrency{
			isNative: false,
			isToken:  true,
			chainId:  chainID,
			decimals: decimals,
			symbol:   symbol,
			name:     name,
		},
		Address: address,
	}
	token.BaseCurrency.currency = token
	return token
}

// Equal
/**
 * Returns true if the two tokens are equivalent, i.e. have the same chainId and address.
 * @param other token to compare
 */
func (t *Token) Equal(other Currency) bool {
	if other != nil {
		v, isToken := other.(*Token)
		if isToken {
			return v.isToken && t.chainId == v.chainId && t.Address == v.Address
		}
	}
	return false
}

// SortsBefore
/**
 * Returns true if the address of this token sorts before the address of the other token
 * @param other other token to compare
 * @throws if the tokens have the same address
 * @throws if the tokens are on different chains
 */
func (t *Token) SortsBefore(other *Token) (bool, error) {
	if t.chainId != other.chainId {
		return false, ErrDifferentChain
	}
	if t.Address == other.Address {
		return false, ErrSameAddress
	}
	return strings.ToLower(t.Address.Hex()) < strings.ToLower(other.Address.Hex()), nil
}

func (t *Token) Wrapped() *Token {
	return t
}
