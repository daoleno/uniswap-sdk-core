package entities

type Ether struct {
	*Currency
}

// Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
func newEther(chainID uint) *Ether {
	return &Ether{
		Currency: NewNativeCurrency(chainID, 18, "ETH", "Ether"),
	}
}

func (e *Ether) Wrapped() *Token {
	return WETH9[e.ChainID]
}

func EtherOnChain(chainID uint) *Ether {
	return newEther(chainID)
}

// Equals compares two Ethers for equality
func (e *Ether) Equals(other *Currency) bool {
	return other.IsNative && other.ChainID == e.ChainID
}
