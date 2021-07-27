package tranche_state

import (
	"github.com/barnbridge/meminero/state/smartexposure"
	"github.com/shopspring/decimal"
)

func (s *Storable) getETokenPrice(pool smartexposure.SEPool, state TrancheState, tranche smartexposure.SETranche) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	ratioWithDec := state.CurrentRatio.Div(tranche.SFactorE)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	tokenAConvRate := state.ConversionRate.AmountAConversion.Shift(int32(-(pool.ATokenDecimals))).Mul(s.processed.tokenPrices[pool.ATokenAddress])
	tokenBConvRate := state.ConversionRate.AmountBConversion.Shift(int32(-(pool.BTokenDecimals))).Mul(s.processed.tokenPrices[pool.BTokenAddress])
	eTokenPrice := tokenAConvRate.Add(tokenBConvRate)

	return eTokenPrice, tokenARatio, tokenBRatio
}
