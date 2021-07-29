package tranche_state

import (
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
)

func (s *Storable) getETokenPrice(pool types.Pool, state TrancheState, tranche types.Tranche) (decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	ratioWithDec := state.CurrentRatio.Div(tranche.SFactorE)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	tokenAConvRate := state.ConversionRate.AmountAConversion.Shift(int32(-(pool.TokenA.Decimals))).Mul(s.processed.tokenPrices[pool.TokenA.Address]["USD"])
	tokenBConvRate := state.ConversionRate.AmountBConversion.Shift(int32(-(pool.TokenB.Decimals))).Mul(s.processed.tokenPrices[pool.TokenB.Address]["USD"])
	eTokenPrice := tokenAConvRate.Add(tokenBConvRate)

	return eTokenPrice, tokenARatio, tokenBRatio
}
