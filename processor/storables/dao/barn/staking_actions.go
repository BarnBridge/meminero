package barn

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) handleStakingActions(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Barn.IsDepositEvent(&log) {
			deposit, err := ethtypes.Barn.DepositEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode deposit event")
			}
			s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
				UserAddress:      utils.NormalizeAddress(deposit.User.String()),
				Amount:           deposit.Amount,
				BalanceAfter:     deposit.NewBalance,
				ActionType:       DEPOSIT,
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
				Amount:           withdraw.AmountWithdrew,
				BalanceAfter:     withdraw.AmountLeft,
				ActionType:       WITHDRAW,
				TransactionHash:  utils.NormalizeAddress(withdraw.Raw.TxHash.String()),
				TransactionIndex: int64(withdraw.Raw.TxIndex),
				LogIndex:         int64(withdraw.Raw.Index),
			})
		}
	}
	return nil
}
