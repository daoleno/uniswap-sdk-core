package entities

type Ether struct {
	*Currency
}

// Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
func NewEther(chainID int) *Ether {
	return &Ether{
		Currency: NewNativeCurrency(chainID, 18, "ETH", "Ether"),
	}
}

// Equals compares two Ethers for equality
func (e *Ether) Equals(other *Ether) bool {
	return other.IsNative && other.ChainID == e.ChainID
}
