package events

import (
	"github.com/shopspring/decimal"
)

func (s *Storable) calculateRatios(factor decimal.Decimal, targetRatio decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	ratioWithDec := targetRatio.Div(factor)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)
	return tokenARatio, tokenBRatio
}
