package yieldfarming

import (
	"context"

	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *Storable) decodeStakingActions(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.YieldFarming.IsYieldFarmingDepositEvent(&log) {
			d, err := ethtypes.YieldFarming.YieldFarmingDepositEvent(log)
			if err != nil {
				return errors.Wrap(err, "could nod decode deposit event")
			}
			s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
				UserAddress:      utils.NormalizeAddress(d.User.String()),
				TokenAddress:     utils.NormalizeAddress(d.TokenAddress.String()),
				Amount:           d.Amount,
				ActionType:       DEPOSIT,
				TransactionHash:  utils.NormalizeAddress(d.Raw.TxHash.String()),
				TransactionIndex: int64(d.Raw.TxIndex),
				LogIndex:         int64(d.Raw.Index),
			})
		}

		if ethtypes.YieldFarming.IsYieldFarmingWithdrawEvent(&log) {
			w, err := ethtypes.YieldFarming.YieldFarmingWithdrawEvent(log)
			if err != nil {
				return errors.Wrap(err, "could nod decode withdraw event")
			}
			s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
				UserAddress:      utils.NormalizeAddress(w.User.String()),
				TokenAddress:     utils.NormalizeAddress(w.TokenAddress.String()),
				Amount:           w.Amount,
				ActionType:       WITHDRAW,
				TransactionHash:  utils.NormalizeAddress(w.Raw.TxHash.String()),
				TransactionIndex: int64(w.Raw.TxIndex),
				LogIndex:         int64(w.Raw.Index),
			})
		}
	}
	return nil
}

func (s *Storable) storeStakingActions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.stakingActions) == 0 {
		s.logger.WithField("module", "staking_actions").Debug("no events found")
		return nil
	}
	var rows [][]interface{}

	for _, t := range s.processed.stakingActions {
		value := decimal.NewFromBigInt(t.Amount, 0)
		rows = append(rows, []interface{}{
			t.UserAddress,
			t.TokenAddress,
			value,
			t.ActionType,
			s.block.BlockCreationTime,
			s.block.Number,
			t.TransactionHash,
			t.TransactionIndex,
			t.LogIndex,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"yield_farming", "transactions"},
		[]string{"user_address", "token_address", "amount", "action_type", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	return err
}