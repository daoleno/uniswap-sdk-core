package entities

import "math/big"

type TradeType int

const (
	ExactInput TradeType = iota
	ExactOutput
)

type Rounding int

const (
	RoundDown Rounding = iota
	RoundHalfUp
	RoundUp
)

var MaxUint256, _ = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
