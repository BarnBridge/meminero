package events

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.saveEpochInfo(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save epoch infos")
	}

	err = s.savePoolState(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save pool state")
	}

	return nil
}

func (s *Storable) saveEpochInfo(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.EpochInfos) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, ei := range s.processed.EpochInfos {
		upR, _ := decimal.NewFromBigInt(ei.UpsideExposureRate, -18).Float64()
		downR, _ := decimal.NewFromBigInt(ei.DownsideProtectionRate, -18).Float64()
		rows = append(rows, []interface{}{
			ei.PoolAddress,
			ei.Epoch.Int64(),
			decimal.NewFromBigInt(ei.SeniorLiquidity, 0),
			decimal.NewFromBigInt(ei.JuniorLiquidity, 0),
			upR,
			downR,
			decimal.NewFromBigInt(ei.JuniorTokenPrice, 0),
			decimal.NewFromBigInt(ei.SeniorTokenPrice, 0),
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "pool_epoch_info"},
		[]string{
			"pool_address",
			"epoch_id",
			"senior_liquidity",
			"junior_liquidity",
			"upside_exposure_rate",
			"downside_protection_rate",
			"junior_token_price_start",
			"senior_token_price_start",
			"block_timestamp",
			"included_in_block",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into pool_epoch_info")
	}

	return nil
}

func (s *Storable) savePoolState(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.States) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, ps := range s.processed.States {
		rows = append(rows, []interface{}{
			ps.PoolAddress,
			decimal.NewFromBigInt(ps.QueuedJuniorsUnderlyingIn, 0),
			decimal.NewFromBigInt(ps.QueuedJuniorsUnderlyingOut, 0),
			decimal.NewFromBigInt(ps.QueuedJuniorTokensBurn, 0),
			decimal.NewFromBigInt(ps.QueuedSeniorsUnderlyingIn, 0),
			decimal.NewFromBigInt(ps.QueuedSeniorsUnderlyingOut, 0),
			decimal.NewFromBigInt(ps.QueuedSeniorTokensBurn, 0),
			decimal.NewFromBigInt(ps.EstimatedSeniorLiquidity, 0),
			decimal.NewFromBigInt(ps.EstimatedJuniorLiquidity, 0),
			decimal.NewFromBigInt(ps.EstimatedSeniorTokenPrice, 0),
			decimal.NewFromBigInt(ps.EstimatedJuniorTokenPrice, 0),
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "pool_state"},
		[]string{
			"pool_address",
			"queued_juniors_underlying_in",
			"queued_juniors_underlying_out",
			"queued_junior_tokens_burn",
			"queued_seniors_underlying_in",
			"queued_seniors_underlying_out",
			"queued_senior_tokens_burn",
			"estimated_senior_liquidity",
			"estimated_junior_liquidity",
			"estimated_senior_token_price",
			"estimated_junior_token_price",
			"block_timestamp",
			"included_in_block",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into pool_epoch_info")
	}

	return nil
}
