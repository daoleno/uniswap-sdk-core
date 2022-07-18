package entities

// Currency is any fungible financial instrument, including Ether, all ERC20 tokens, and other chain-native currencies
type Currency interface {
	IsNative() bool
	IsToken() bool
	ChainId() uint
	Decimals() uint
	Symbol() string
	Name() string
	Equal(other Currency) bool
	Wrapped() *Token
}

// BaseCurrency is an abstract struct, do not use it directly
type BaseCurrency struct {
	currency Currency
	isNative bool   // Returns whether the currency is native to the chain and must be wrapped (e.g. Ether)
	isToken  bool   // Returns whether the currency is a token that is usable in Uniswap without wrapping
	chainId  uint   // The chain ID on which this currency resides
	decimals uint   // The decimals used in representing currency amounts
	symbol   string // The symbol of the currency, i.e. a short textual non-unique identifier
	name     string // The name of the currency, i.e. a descriptive textual non-unique identifier
}

func (c *BaseCurrency) IsNative() bool {
	return c.isNative
}

func (c *BaseCurrency) IsToken() bool {
	return c.isToken
}

func (c *BaseCurrency) ChainId() uint {
	return c.chainId
}

func (c *BaseCurrency) Decimals() uint {
	return c.decimals
}

func (c *BaseCurrency) Symbol() string {
	return c.symbol
}

func (c *BaseCurrency) Name() string {
	return c.name
}

// Equal returns whether the currency is equal to the other currency
func (c *BaseCurrency) Equal(other Currency) bool {
	panic("Equal method has to be overridden")
}

func (c *BaseCurrency) Wrapped() *Token {
	panic("Wrapped method has to be overridden")
}
