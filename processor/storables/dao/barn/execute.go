package barn

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	var barnLogs []gethtypes.Log
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Barn.Address) {
				barnLogs = append(barnLogs, log)
			}
		}
	}

	if len(barnLogs) == 0 {
		return nil
	}

	err := s.handleDelegateEvents(barnLogs, ctx)
	if err != nil {
		return err
	}

	err = s.handleLockEvents(barnLogs)
	if err != nil {
		return err
	}

	err = s.handleStakingActions(barnLogs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) handleDelegateEvents(logs []gethtypes.Log, ctx context.Context) error {
	for _, log := range logs {
		if ethtypes.Barn.IsDelegateEvent(&log) {
			var action DelegateAction
			a, err := ethtypes.Barn.DelegateEvent(log)
			if err != nil {
				return err
			}
			if utils.NormalizeAddress(a.To.String()) == utils.ZeroAddress {
				action = DelegateAction{
					BarnDelegateEvent: a,
					ActionType:        DelegateStop,
				}
			} else {
				action = DelegateAction{
					BarnDelegateEvent: a,
					ActionType:        DelegateStart,
				}
			}
			s.processed.delegateActions = append(s.processed.delegateActions, action)
		}

		if ethtypes.Barn.IsDelegatedPowerIncreasedEvent(&log) {
			increase, err := ethtypes.Barn.DelegatedPowerIncreasedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode delegate power increased event")
			}

			s.processed.delegateChanges = append(s.processed.delegateChanges, DelegateChange{
				Sender:              utils.NormalizeAddress(increase.From.String()),
				Receiver:            utils.NormalizeAddress(increase.To.String()),
				Amount:              increase.AmountDecimal(0),
				ToNewDelegatedPower: increase.ToNewDelegatedPowerDecimal(0),
				ActionType:          DelegateIncrease,
				TransactionHash:     utils.NormalizeAddress(increase.Raw.TxHash.String()),
				TransactionIndex:    int64(increase.Raw.TxIndex),
				LogIndex:            int64(increase.Raw.Index),
			})
		}

		if ethtypes.Barn.IsDelegatedPowerDecreasedEvent(&log) {
			decrease, err := ethtypes.Barn.DelegatedPowerDecreasedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode delegate power increased event")
			}
			s.processed.delegateChanges = append(s.processed.delegateChanges, DelegateChange{
				Sender:              utils.NormalizeAddress(decrease.From.String()),
				Receiver:            utils.NormalizeAddress(decrease.To.String()),
				Amount:              decrease.AmountDecimal(0),
				ToNewDelegatedPower: decrease.ToNewDelegatedPowerDecimal(0),
				ActionType:          DelegateDecrease,
				TransactionHash:     utils.NormalizeAddress(decrease.Raw.TxHash.String()),
				TransactionIndex:    int64(decrease.Raw.TxIndex),
				LogIndex:            int64(decrease.Raw.Index),
			})
		}
	}
	return nil
}

func (s *Storable) handleLockEvents(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Barn.IsLockEvent(&log) {
			lock, err := ethtypes.Barn.LockEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode lock event")
			}

			s.processed.locks = append(s.processed.locks, lock)
		}
	}
	return nil
}

func (s *Storable) handleStakingActions(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Barn.IsDepositEvent(&log) {
			deposit, err := ethtypes.Barn.DepositEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode deposit event")
			}
			s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
				UserAddress:      utils.NormalizeAddress(deposit.User.String()),
				Amount:           deposit.AmountDecimal(0),
				BalanceAfter:     deposit.NewBalanceDecimal(0),
				ActionType:       Deposit,
				TransactionHash:  utils.NormalizeAddress(deposit.Raw.TxHash.String()),
				TransactionIndex: int64(deposit.Raw.TxIndex),
				LogIndex:         int64(deposit.Raw.Index),
			})
		}

		if ethtypes.Barn.IsWithdrawEvent(&log) {
			withdraw, err := ethtypes.Barn.WithdrawEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode withdraw event")
			}
			s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
				UserAddress:      utils.NormalizeAddress(withdraw.User.String()),
				Amount:           withdraw.AmountWithdrewDecimal(0),
				BalanceAfter:     withdraw.AmountLeftDecimal(0),
				ActionType:       Withdraw,
				TransactionHash:  utils.NormalizeAddress(withdraw.Raw.TxHash.String()),
				TransactionIndex: int64(withdraw.Raw.TxIndex),
				LogIndex:         int64(withdraw.Raw.Index),
			})
		}
	}
	return nil
}
