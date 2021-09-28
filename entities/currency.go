package entities

// A currency is any fungible financial instrument, including Ether, all ERC20 tokens, and other chain-native currencies
type Currency struct {
	IsNative bool   // Returns whether the currency is native to the chain and must be wrapped (e.g. Ether)
	IsToken  bool   // Returns whether the currency is a token that is usable in Uniswap without wrapping
	ChainID  uint   // The chain ID on which this currency resides
	Decimals uint   // The decimals used in representing currency amounts
	Symbol   string // The symbol of the currency, i.e. a short textual non-unique identifier
	Name     string // The name of the currency, i.e. a descriptive textual non-unique identifier
}

// NewBaseCurrency constructs an instance of the `BaseCurrency`.
func NewBaseCurrency(chainID uint, decimals uint, symbol string, name string) *Currency {
	return &Currency{
		ChainID:  chainID,
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}

// NewNativeCurrency constructs an instrance of the `NativeCurrency`
func NewNativeCurrency(chainID uint, decimals uint, symbol string, name string) *Currency {
	return &Currency{
		IsNative: true,
		ChainID:  chainID,
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}

// NewTokenCurrency constructs an instance of the `TokenCurrency`
func NewTokenCurrency(chainID uint, decimals uint, symbol string, name string) *Currency {
	return &Currency{
		IsToken:  true,
		ChainID:  chainID,
		Decimals: decimals,
		Symbol:   symbol,
		Name:     name,
	}
}

// Equal returns whether the currency is equal to the other currency
func (c *Currency) Equal(other *Currency) bool {
	return c.ChainID == other.ChainID && c.Decimals == other.Decimals && c.Symbol == other.Symbol && c.Name == other.Name
}
