package tranche_state

import (
	"math/big"
)

type TrancheState struct {
	EPoolAddress string
	CurrentRatio *big.Int

	TokenALiquidity *big.Int
	TokenBLiquidity *big.Int

	ConversionRate ConversionRate
}

type ConversionRate struct {
	AmountAConversion *big.Int
	AmountBConversion *big.Int
}
