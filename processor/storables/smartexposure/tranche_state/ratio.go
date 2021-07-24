package tranche_state

import (
	"github.com/barnbridge/meminero/types"
	"github.com/shopspring/decimal"
)

func (s *Storable) getETokenPrice(pool types.SEPool, state TrancheState, tranche types.SETranche) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	ratioWithDec := decimal.NewFromBigInt(state.CurrentRatio, 0).Div(tranche.SFactorE)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	tokenAConvRate := decimal.NewFromBigInt(state.ConversionRate.AmountAConversion, int32(-(pool.ATokenDecimals))).Mul(s.processed.tokenPrices[pool.ATokenAddress])
	tokenBConvRate := decimal.NewFromBigInt(state.ConversionRate.AmountBConversion, int32(-(pool.BTokenDecimals))).Mul(s.processed.tokenPrices[pool.BTokenAddress])
	eTokenPrice := tokenAConvRate.Add(tokenBConvRate)

	return eTokenPrice, tokenARatio, tokenBRatio
}
