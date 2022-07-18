package entities

// Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
type Ether struct {
	*BaseCurrency
}

func EtherOnChain(chainId uint) *Ether {
	ether := &Ether{
		BaseCurrency: &BaseCurrency{
			isNative: true,
			isToken:  false,
			chainId:  chainId,
			decimals: 18,
			symbol:   "ETH",
			name:     "Ether",
		},
	}
	ether.BaseCurrency.currency = ether
	return ether
}

func (e *Ether) Equal(other Currency) bool {
	v, isEther := other.(*Ether)
	if isEther {
		return v.isNative && v.chainId == e.chainId

	}
	return false
}

func (e *Ether) Wrapped() *Token {
	return WETH9[e.ChainId()]
}
