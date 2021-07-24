package tranche_state

import (
	"context"

	"github.com/barnbridge/meminero/utils"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func (s *Storable) storeTranchesState(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.trancheState) == 0 {
		s.logger.WithField("handler", "tranche state").Debug("no events found")
		return nil
	}
	var rows [][]interface{}

	for trancheAddress, t := range s.processed.trancheState {
		pool := s.state.SEPoolByAddress(t.EPoolAddress)
		tranche := s.state.SETrancheByETokenAddress(trancheAddress)

		tokenAPrice, err := utils.GetTokenPrice(ctx, tx, pool.ATokenAddress, s.block.Number)
		if err != nil {
			return err
		}

		tokenBPrice, err := utils.GetTokenPrice(ctx, tx, pool.BTokenAddress, s.block.Number)
		if err != nil {
			return err
		}

		tokenALiquidity := decimal.NewFromBigInt(t.TokenALiquidity, -int32(pool.ATokenDecimals)).Mul(tokenAPrice)
		tokenBLiquidity := decimal.NewFromBigInt(t.TokenBLiquidity, -int32(pool.BTokenDecimals)).Mul(tokenBPrice)

		eTokenPrice, tokenARatio, tokenBRatio := s.getETokenPrice(*pool, *t, *tranche, tokenAPrice, tokenBPrice)
		rows = append(rows, []interface{}{
			t.EPoolAddress,
			trancheAddress,
			tokenALiquidity,
			tokenBLiquidity,
			t.ConversionRate.AmountAConversion,
			t.ConversionRate.AmountBConversion,
			eTokenPrice,
			tokenARatio,
			tokenBRatio,
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "tranche_state"},
		[]string{"pool_address", "etoken_address", "token_a_liquidity", "token_b_liquidity", "current_ratio", "amount_a_conversion", "amount_b_conversion", "etoken_price", "token_a_current_ratio", "token_b_current_ratio", "block_timestamp" , "included_in_block"},
		pgx.CopyFromRows(rows),
	)
	return err
}
