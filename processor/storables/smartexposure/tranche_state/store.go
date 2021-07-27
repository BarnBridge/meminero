package tranche_state

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) storeTranchesState(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.trancheState) == 0 {
		s.logger.WithField("handler", "tranche state").Debug("no events found")
		return nil
	}
	var rows [][]interface{}

	for trancheAddress, t := range s.processed.trancheState {
		pool := s.state.SmartExposure.SEPoolByAddress(t.EPoolAddress)
		tranche := s.state.SmartExposure.SETrancheByETokenAddress(trancheAddress)

		tokenAPrice := s.processed.tokenPrices[pool.ATokenAddress]
		tokenBPrice := s.processed.tokenPrices[pool.BTokenAddress]

		tokenALiquidity, _ := (t.TokenALiquidity.Shift(-int32(pool.ATokenDecimals)).Mul(tokenAPrice)).Float64()
		tokenBLiquidity, _ := (t.TokenBLiquidity.Shift(-int32(pool.BTokenDecimals)).Mul(tokenBPrice)).Float64()

		eTokenPrice, tokenARatio, tokenBRatio := s.getETokenPrice(*pool, *t, *tranche)
		price, _ := eTokenPrice.Float64()

		ratioA, _ := tokenARatio.Float64()
		ratioB, _ := tokenBRatio.Float64()

		rows = append(rows, []interface{}{
			t.EPoolAddress,
			trancheAddress,
			tokenALiquidity,
			tokenBLiquidity,
			t.CurrentRatio,
			t.ConversionRate.AmountAConversion,
			t.ConversionRate.AmountBConversion,
			price,
			ratioA,
			ratioB,
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "tranche_state"},
		[]string{"pool_address", "etoken_address", "token_a_liquidity", "token_b_liquidity", "current_ratio", "amount_a_conversion", "amount_b_conversion", "etoken_price", "token_a_current_ratio", "token_b_current_ratio", "block_timestamp", "included_in_block"},
		pgx.CopyFromRows(rows),
	)
	return err
}
