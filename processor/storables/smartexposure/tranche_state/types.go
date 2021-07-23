package tranche_state

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type TrancheState struct {
	EPoolAddress string
	CurrentRatio *big.Int

	TokenALiquidity decimal.Decimal
	TokenBLiquidity decimal.Decimal

	ConversionRate     ConversionRate
	TokenACurrentRatio decimal.Decimal
	TokenBCurrentRatio decimal.Decimal

	ETokenPrice decimal.Decimal
}

type ConversionRate struct {
	AmountAConversion *big.Int
	AmountBConversion *big.Int
}
