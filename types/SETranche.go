package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type SETranche struct {
	EPoolAddress  string
	ETokenAddress string
	ETokenSymbol  string
	SFactorE      *big.Int
	TargetRatio   *big.Int
	TokenARatio   decimal.Decimal
	TokenBRatio   decimal.Decimal
}
