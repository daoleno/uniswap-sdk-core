package entities

type Native struct {
	*BaseCurrency
	wrapped *Token
}

func NewNative(wrapped *Token, symbol, name string) Currency {
	native := &Native{
		BaseCurrency: &BaseCurrency{
			isNative: true,
			isToken:  false,
			chainId:  wrapped.ChainId(),
			decimals: wrapped.Decimals(),
			symbol:   symbol,
			name:     name,
		},
		wrapped: wrapped,
	}
	native.BaseCurrency.currency = native
	return native
}

func (n *Native) Equal(other Currency) bool {
	v, isNative := other.(*Native)
	if isNative {
		return v.isNative && v.chainId == n.chainId

	}
	return false
}

func (n *Native) Wrapped() *Token {
	return n.wrapped
}
