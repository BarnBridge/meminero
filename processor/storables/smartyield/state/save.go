package state

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	var rows [][]interface{}

	for _, ps := range s.processed.PoolStates {
		rows = append(rows, []interface{}{
			ps.PoolAddress,
			ps.TotalLiquidity.Sub(ps.JuniorLiquidity),
			ps.JuniorLiquidity,
			ps.JTokenPrice,
			decimal.NewFromBigInt(ps.Abond.Principal, 0),
			decimal.NewFromBigInt(ps.Abond.Gain, 0),
			decimal.NewFromBigInt(ps.Abond.IssuedAt, -18).IntPart(),
			decimal.NewFromBigInt(ps.Abond.MaturesAt, -18).IntPart(),
			ps.AbondAPY,
			ps.SeniorAPY,
			ps.JuniorAPY,
			ps.OriginatorApy,
			ps.OriginatorNetApy,
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "pool_state"},
		[]string{
			"pool_address",
			"senior_liquidity",
			"junior_liquidity",
			"jtoken_price",
			"abond_principal",
			"abond_gain",
			"abond_issued_at",
			"abond_matures_at",
			"abond_apy",
			"senior_apy",
			"junior_apy",
			"originator_apy",
			"originator_net_apy",
			"block_timestamp",
			"included_in_block",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into smart_yield.pool_state")
	}

	return nil
}
