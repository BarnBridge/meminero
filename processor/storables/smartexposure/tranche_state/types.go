package tranche_state

import (
	"github.com/shopspring/decimal"
)

type TrancheState struct {
	EPoolAddress string
	CurrentRatio decimal.Decimal

	TokenALiquidity decimal.Decimal
	TokenBLiquidity decimal.Decimal

	ConversionRate ConversionRate
}

type ConversionRate struct {
	AmountAConversion decimal.Decimal
	AmountBConversion decimal.Decimal
}
