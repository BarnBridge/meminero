package types

import (
	"github.com/shopspring/decimal"
)

type SETranche struct {
	EPoolAddress  string
	ETokenAddress string
	ETokenSymbol  string
	SFactorE      decimal.Decimal
	TargetRatio   decimal.Decimal
	TokenARatio   decimal.Decimal
	TokenBRatio   decimal.Decimal
}
