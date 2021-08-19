package events

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.saveJoinEntryQueueEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store join entry queue events")
	}

	err = s.saveJoinExitQueueEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store join exit queue events")
	}

	err = s.saveRedeemTokensEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store redeem tokens events")
	}

	err = s.saveRedeemUnderlyingEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store redeem underlying events")
	}

	err = s.saveTxHistory(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store tx history")
	}

	err = s.saveEpochEndEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store epoch end events")
	}

	err = s.saveEpochInfo(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save epoch infos")
	}

	return nil
}

func (s *Storable) saveJoinEntryQueueEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorJoinEntryQueueEvents) == 0 && len(s.processed.SeniorJoinEntryQueueEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.JuniorJoinEntryQueueEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.JuniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.UnderlyingInDecimal(0),
			e.CurrentQueueBalanceDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	for _, e := range s.processed.SeniorJoinEntryQueueEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.SeniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.UnderlyingInDecimal(0),
			e.CurrentQueueBalanceDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "user_join_entry_queue_events"},
		[]string{
			"pool_address",
			"tranche",
			"user_address",
			"epoch_id",
			"underlying_in",
			"queue_balance_after",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into user_join_entry_queue_events")
	}

	return nil
}

func (s *Storable) saveJoinExitQueueEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorJoinExitQueueEvents) == 0 && len(s.processed.SeniorJoinExitQueueEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.JuniorJoinExitQueueEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.JuniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.TokensInDecimal(0),
			e.CurrentQueueBalanceDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	for _, e := range s.processed.SeniorJoinExitQueueEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.SeniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.TokensInDecimal(0),
			e.CurrentQueueBalanceDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "user_join_exit_queue_events"},
		[]string{
			"pool_address",
			"tranche",
			"user_address",
			"epoch_id",
			"tokens_in",
			"queue_balance_after",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into user_join_exit_queue_events")
	}

	return nil
}

func (s *Storable) saveRedeemTokensEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorRedeemTokensEvents) == 0 && len(s.processed.SeniorRedeemTokensEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.JuniorRedeemTokensEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.JuniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.TokensOutDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	for _, e := range s.processed.SeniorRedeemTokensEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.SeniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.TokensOutDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "user_redeem_tokens_events"},
		[]string{
			"pool_address",
			"tranche",
			"user_address",
			"epoch_id",
			"tokens_out",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into user_redeem_tokens_events")
	}

	return nil
}

func (s *Storable) saveRedeemUnderlyingEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorRedeemUnderlyingEvents) == 0 && len(s.processed.SeniorRedeemUnderlyingEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.JuniorRedeemUnderlyingEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.JuniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.UnderlyingOutDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	for _, e := range s.processed.SeniorRedeemUnderlyingEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress,
			smartalpha.SeniorTranche,
			utils.NormalizeAddress(e.User.String()),
			e.EpochId.Int64(),
			e.UnderlyingOutDecimal(0),
		}

		rows = append(rows, append(row, s.generateEventData(e.Raw)...))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "user_redeem_underlying_events"},
		[]string{
			"pool_address",
			"tranche",
			"user_address",
			"epoch_id",
			"underlying_out",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into user_redeem_underlying_events")
	}

	return nil
}

func (s *Storable) saveTxHistory(ctx context.Context, tx pgx.Tx) error {
	var rows [][]interface{}

	for _, e := range s.processed.JuniorJoinEntryQueueEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.JuniorTranche,
			smartalpha.JuniorEntry,
			e.User.String(),
			e.UnderlyingInDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.SeniorJoinEntryQueueEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.SeniorTranche,
			smartalpha.SeniorEntry,
			e.User.String(),
			e.UnderlyingInDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.JuniorJoinExitQueueEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.JuniorTranche,
			smartalpha.JuniorExit,
			e.User.String(),
			e.TokensInDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.SeniorJoinExitQueueEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.SeniorTranche,
			smartalpha.SeniorExit,
			e.User.String(),
			e.TokensInDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.JuniorRedeemTokensEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.JuniorTranche,
			smartalpha.JuniorRedeemTokens,
			e.User.String(),
			e.TokensOutDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.SeniorRedeemTokensEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.SeniorTranche,
			smartalpha.SeniorRedeemTokens,
			e.User.String(),
			e.TokensOutDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.JuniorRedeemUnderlyingEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.JuniorTranche,
			smartalpha.JuniorRedeemUnderlying,
			e.User.String(),
			e.UnderlyingOutDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.SeniorRedeemUnderlyingEvents {
		rows = append(rows, s.txHistory(
			s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String()).PoolAddress,
			smartalpha.SeniorTranche,
			smartalpha.SeniorRedeemUnderlying,
			e.User.String(),
			e.UnderlyingOutDecimal(0),
			e.Raw,
		))
	}

	for _, e := range s.processed.TokenTransferEvents {
		// ignore mint and burn events
		if e.From.String() == utils.ZeroAddress || e.To.String() == utils.ZeroAddress {
			continue
		}

		// ignore any transfer to and from the pools because they have another event in the transaction history (e.g. JuniorRedeemTokens)
		if s.state.SmartAlpha.PoolByAddress(e.From.String()) != nil || s.state.SmartAlpha.PoolByAddress(e.To.String()) != nil {
			continue
		}

		p := s.state.SmartAlpha.PoolByJuniorTokenAddress(e.Raw.Address.String())

		if p != nil {
			rows = append(rows, s.txHistory(
				p.PoolAddress,
				smartalpha.JuniorTranche,
				smartalpha.JtokenSend,
				e.From.String(),
				e.ValueDecimal(0),
				e.Raw,
			), s.txHistory(
				p.PoolAddress,
				smartalpha.JuniorTranche,
				smartalpha.JtokenReceive,
				e.To.String(),
				e.ValueDecimal(0),
				e.Raw,
			))
		} else {
			p := s.state.SmartAlpha.PoolBySeniorTokenAddress(e.Raw.Address.String())
			if p == nil {
				continue
			}

			rows = append(rows, s.txHistory(
				p.PoolAddress,
				smartalpha.SeniorTranche,
				smartalpha.StokenSend,
				e.From.String(),
				e.ValueDecimal(0),
				e.Raw,
			), s.txHistory(
				p.PoolAddress,
				smartalpha.SeniorTranche,
				smartalpha.StokenReceive,
				e.To.String(),
				e.ValueDecimal(0),
				e.Raw,
			))
		}
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_alpha", "transaction_history"},
		[]string{
			"pool_address",
			"tranche",
			"transaction_type",
			"user_address",
			"amount",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into transaction_history")
	}

	return nil
}

func (s *Storable) saveEpochEndEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.EpochEndEvents) == 0 {
		return nil
	}

	for _, e := range s.processed.EpochEndEvents {
		p := s.state.SmartAlpha.PoolByAddress(e.Raw.Address.String())

		row := []interface{}{
			p.PoolAddress, e.EpochId.Int64(), e.JuniorProfitsDecimal(0), e.SeniorProfitsDecimal(0),
		}
		row = append(row, s.generateEventData(e.Raw)...)

		_, err := tx.Exec(ctx, `
			insert into smart_alpha.epoch_end_events (pool_address, epoch_id, junior_profits, senior_profits, block_timestamp,
													  included_in_block, tx_hash, tx_index, log_index)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9);
		`, row...)
		if err != nil {
			return errors.Wrap(err, "could not insert epoch end event")
		}
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
			decimal.NewFromBigInt(ei.EpochEntryPrice, 0),
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
			"epoch_entry_price",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into pool_epoch_info")
	}

	return nil
}
