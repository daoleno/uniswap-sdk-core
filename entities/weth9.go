package entities

import "github.com/ethereum/go-ethereum/common"

// Known WETH9 implementation addresses, used in our implementation of Ether#wrapped
var WETH9 = map[uint]*Token{
	1:  NewNativeToken(1, common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), 18, "WETH", "Wrapped Ether"),
	3:  NewNativeToken(3, common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"), 18, "WETH", "Wrapped Ether"),
	4:  NewNativeToken(4, common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"), 18, "WETH", "Wrapped Ether"),
	5:  NewNativeToken(5, common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"), 18, "WETH", "Wrapped Ether"),
	42: NewNativeToken(42, common.HexToAddress("0xd0A1E359811322d97991E03f863a0C30C2cF029C"), 18, "WETH", "Wrapped Ether"),

	10: NewNativeToken(10, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	69: NewNativeToken(69, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),

	42161:  NewNativeToken(42161, common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"), 18, "WETH", "Wrapped Ether"),
	421611: NewNativeToken(421611, common.HexToAddress("0xB47e6A5f8b33b3F17603C83a0535A9dcD7E32681"), 18, "WETH", "Wrapped Ether"),
}
