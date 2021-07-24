package tranche_state

import (
	"github.com/barnbridge/meminero/types"
	"github.com/shopspring/decimal"
)

func (s Storable) getETokenPrice(pool types.SEPool, state TrancheState, tranche types.SETranche, tokenAPrice decimal.Decimal, tokenBPrice decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	ratioWithDec := decimal.NewFromBigInt(state.CurrentRatio, 0).Div(tranche.SFactorE)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	tokenAConvRate := decimal.NewFromBigInt(state.ConversionRate.AmountAConversion, int32(-(pool.ATokenDecimals))).Mul(tokenAPrice)
	tokenBConvRate := decimal.NewFromBigInt(state.ConversionRate.AmountBConversion, int32(-(pool.BTokenDecimals))).Mul(tokenBPrice)
	eTokenPrice := tokenAConvRate.Add(tokenBConvRate)

	return eTokenPrice, tokenARatio, tokenBRatio
}
