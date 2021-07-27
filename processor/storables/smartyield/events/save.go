package events

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/notifications"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Debug("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done storing")
	}()

	err := s.saveJuniorEntryEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save junior entry events")
	}

	err = s.saveJuniorInstantWithdrawEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save junior instant withdraw events")
	}

	err = s.saveJunior2StepWithdrawEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save junior 2 step withdraw events")
	}

	err = s.saveJunior2StepRedeemEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save junior 2 step redeem events")
	}

	err = s.saveSeniorEntryEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save senior entry events")
	}

	err = s.saveSeniorRedeemEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save senior redeem events")
	}

	err = s.saveTransactionHistory(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save transaction history")
	}

	err = s.saveControllerHarvestEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save controller harvest events")
	}

	err = s.saveProviderTransferFeesEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save provider transfer fees events")
	}

	return nil
}

func (s *Storable) saveJuniorEntryEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorEntryEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job

	for _, e := range s.processed.JuniorEntryEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Buyer.String()),
			e.UnderlyingInDecimal(0),
			e.TokensOutDecimal(0),
			e.FeeDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})

		jd := notifications.SmartYieldJobData{
			StartTime:             s.block.BlockCreationTime,
			Pool:                  *p,
			Buyer:                 utils.NormalizeAddress(e.Buyer.String()),
			Amount:                e.TokensOutDecimal(0),
			IncludedInBlockNumber: s.block.Number,
		}
		j, err := notifications.NewSmartYieldTokenBoughtJob(&jd)
		if err != nil {
			return errors.Wrap(err, "could not create notification job")
		}

		jobs = append(jobs, j)
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "junior_entry_events"},
		[]string{
			"pool_address",
			"buyer_address",
			"underlying_in",
			"tokens_out",
			"fee",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into junior_entry_events")
	}

	if config.Store.Storable.SmartYield.Notifications && len(jobs) > 0 {
		ctx, cancel := context.WithTimeout(ctx, time.Second*2)
		defer cancel()

		err = notifications.ExecuteJobsWithTx(ctx, tx, jobs...)
		if err != nil && err != context.DeadlineExceeded {
			return errors.Wrap(err, "could not execute notification jobs")
		}
	}

	return nil
}

func (s *Storable) saveJuniorInstantWithdrawEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.JuniorInstantWithdrawEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.JuniorInstantWithdrawEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Seller.String()),
			e.TokensInDecimal(0),
			e.UnderlyingOutDecimal(0),
			e.ForfeitsDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "junior_instant_withdraw_events"},
		[]string{
			"pool_address",
			"seller_address",
			"tokens_in",
			"underlying_out",
			"forfeits",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into junior_instant_withdraw_events")
	}

	return nil
}

func (s *Storable) saveJunior2StepWithdrawEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Junior2StepWithdrawEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.Junior2StepWithdrawEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Buyer.String()),
			p.JuniorBondAddress,
			e.JuniorBondId.Int64(),
			e.TokensInDecimal(0),
			e.MaturesAt.Int64(),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "junior_2step_withdraw_events"},
		[]string{
			"pool_address",
			"buyer_address",
			"junior_bond_address",
			"junior_bond_id",
			"tokens_in",
			"matures_at",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into junior_2step_withdraw_events")
	}

	return nil
}

func (s *Storable) saveJunior2StepRedeemEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Junior2StepRedeemEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.Junior2StepRedeemEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Owner.String()),
			p.JuniorBondAddress,
			e.JuniorBondId.Int64(),
			e.UnderlyingOutDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "junior_2step_redeem_events"},
		[]string{
			"pool_address",
			"owner_address",
			"junior_bond_address",
			"junior_bond_id",
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
		return errors.Wrap(err, "could not copy into junior_2step_redeem_events")
	}

	return nil
}

func (s *Storable) saveSeniorEntryEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.SeniorEntryEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.SeniorEntryEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Buyer.String()),
			p.SeniorBondAddress,
			e.SeniorBondId.Int64(),
			e.UnderlyingInDecimal(0),
			e.GainDecimal(0),
			e.ForDays.Int64(),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "senior_entry_events"},
		[]string{
			"pool_address",
			"buyer_address",
			"senior_bond_address",
			"senior_bond_id",
			"underlying_in",
			"gain",
			"for_days",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into senior_entry_events")
	}

	return nil
}

func (s *Storable) saveSeniorRedeemEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.SeniorRedeemEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.SeniorRedeemEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		rows = append(rows, []interface{}{
			p.PoolAddress,
			utils.NormalizeAddress(e.Owner.String()),
			p.SeniorBondAddress,
			e.SeniorBondId.Int64(),
			e.FeeDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "senior_redeem_events"},
		[]string{
			"pool_address",
			"owner_address",
			"senior_bond_address",
			"senior_bond_id",
			"fee",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into senior_redeem_events")
	}

	return nil
}

func (s *Storable) saveTransactionHistory(ctx context.Context, tx pgx.Tx) error {
	var rows [][]interface{}

	for _, e := range s.processed.JuniorEntryEvents {
		rows = append(rows, s.txHistory(e.Buyer.String(), e.UnderlyingIn, smartyield.JuniorTranche, smartyield.JuniorDeposit, e.Raw))
	}

	for _, e := range s.processed.JuniorInstantWithdrawEvents {
		rows = append(rows, s.txHistory(e.Seller.String(), e.TokensIn, smartyield.JuniorTranche, smartyield.JuniorInstantWithdraw, e.Raw))
	}

	for _, e := range s.processed.Junior2StepWithdrawEvents {
		rows = append(rows, s.txHistory(e.Buyer.String(), e.TokensIn, smartyield.JuniorTranche, smartyield.JuniorRegularWithdraw, e.Raw))
	}

	for _, e := range s.processed.Junior2StepRedeemEvents {
		rows = append(rows, s.txHistory(e.Owner.String(), e.UnderlyingOut, smartyield.JuniorTranche, smartyield.JuniorRedeem, e.Raw))
	}

	for _, e := range s.processed.SeniorEntryEvents {
		rows = append(rows, s.txHistory(e.Buyer.String(), e.UnderlyingIn, smartyield.SeniorTranche, smartyield.SeniorDeposit, e.Raw))
	}

	for _, e := range s.processed.SeniorRedeemEvents {
		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		var underlyingIn, gain decimal.Decimal
		err := tx.QueryRow(ctx,
			`select underlying_in, gain from smart_yield.senior_entry_events where senior_bond_address = $1 and senior_bond_id = $2`,
			p.SeniorBondAddress, e.SeniorBondId.Int64(),
		).Scan(&underlyingIn, &gain)
		if err != nil {
			return errors.Wrap(err, "could not find senior bond details")
		}

		amount := underlyingIn.Add(gain).Sub(e.FeeDecimal(0))

		rows = append(rows, s.txHistory(e.Owner.String(), amount.Coefficient(), smartyield.SeniorTranche, smartyield.SeniorRedeem, e.Raw))
	}

	for _, e := range s.processed.Transfers {
		// we don't want to store Mint & Burn events in the transaction history table because there will already be another
		// corresponding action (eg: SeniorDeposit)
		if e.From.String() == utils.ZeroAddress || e.To.String() == utils.ZeroAddress {
			continue
		}

		p := s.state.SmartYield.PoolByAddress(e.Raw.Address.String())

		// ignore any transfers related to a 'BuyJuniorBond' event
		if utils.NormalizeAddress(e.To.String()) == utils.NormalizeAddress(p.PoolAddress) {
			continue
		}

		// Ignore any transfer to and from a reward pool
		if s.state.SmartYield.RewardPoolByAddress(e.To.String()) != nil || s.state.SmartYield.RewardPoolByAddress(e.From.String()) != nil {
			continue
		}

		rows = append(rows, s.txHistory(e.From.String(), e.Value, smartyield.JuniorTranche, smartyield.JtokenSend, e.Raw))
		rows = append(rows, s.txHistory(e.To.String(), e.Value, smartyield.JuniorTranche, smartyield.JtokenReceive, e.Raw))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "transaction_history"},
		[]string{
			"protocol_id",
			"pool_address",
			"underlying_token_address",
			"user_address",
			"amount",
			"tranche",
			"transaction_type",
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

func (s *Storable) saveControllerHarvestEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.ControllerHarvests) == 0 {
		return nil
	}

	for _, e := range s.processed.ControllerHarvests {
		p := s.state.SmartYield.PoolByControllerAddress(e.Raw.Address.String())

		_, err := tx.Exec(ctx, `
			insert into smart_yield.controller_harvests (protocol_id,
														 controller_address,
														 caller_address,
														 comp_reward_total,
														 comp_reward_sold,
														 underlying_pool_share,
														 underlying_reward,
														 harvest_cost,
														 block_timestamp,
														 included_in_block,
														 tx_hash,
														 tx_index,
														 log_index)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		`,
			p.ProtocolId,
			p.ControllerAddress,
			utils.NormalizeAddress(e.Caller.String()),
			e.CompRewardTotalDecimal(0),
			e.CompRewardSoldDecimal(0),
			e.UnderlyingPoolShareDecimal(0),
			e.UnderlyingRewardDecimal(0),
			e.HarvestCostDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		)
		if err != nil {
			return errors.Wrap(err, "could not execute insert")
		}
	}

	return nil
}

func (s *Storable) saveProviderTransferFeesEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.ProviderTransferFees) == 0 {
		return nil
	}

	for _, e := range s.processed.ProviderTransferFees {
		p := s.state.SmartYield.PoolByProviderAddress(e.Raw.Address.String())

		_, err := tx.Exec(ctx, `
			insert into smart_yield.provider_transfer_fees (protocol_id,
															provider_address,
															caller_address,
															fees_owner,
															fees,
															block_timestamp,
															included_in_block,
															tx_hash,
															tx_index,
															log_index)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`,
			p.ProtocolId,
			p.ProviderAddress,
			utils.NormalizeAddress(e.Caller.String()),
			utils.NormalizeAddress(e.FeesOwner.String()),
			e.FeesDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		)
		if err != nil {
			return errors.Wrap(err, "could not execute insert")
		}
	}

	return nil
}
