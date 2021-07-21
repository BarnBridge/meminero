package barn

import (
	"context"

	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (s *Storable) handleDelegateEvents(logs []gethtypes.Log, ctx context.Context) error {

	for _, log := range logs {
		if ethtypes.Barn.IsBarnDelegateEvent(&log) {
			var action DelegateAction
			a, err := ethtypes.Barn.BarnDelegateEvent(log)
			if err != nil {
				return err
			}
			if utils.NormalizeAddress(a.To.String()) == ZeroAddress {
				action = DelegateAction{
					BarnDelegateEvent: a,
					ActionType:        DELEGATE_STOP,
				}
			} else {
				action = DelegateAction{
					BarnDelegateEvent: a,
					ActionType:        DELEGATE_START,
				}
			}
			s.processed.delegateActions = append(s.processed.delegateActions, action)
		}

		if ethtypes.Barn.IsBarnDelegatedPowerIncreasedEvent(&log) {
			increase, err := ethtypes.Barn.BarnDelegatedPowerIncreasedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode delegate power increased event")
			}

			s.processed.delegateChanges = append(s.processed.delegateChanges, DelegateChange{
				Sender:              increase.From.String(),
				Receiver:            increase.To.String(),
				Amount:              increase.Amount,
				ToNewDelegatedPower: increase.ToNewDelegatedPower,
				ActionType:          DELEGATE_INCREASE,
				TransactionHash:     increase.Raw.TxHash.String(),
				TransactionIndex:    int64(increase.Raw.TxIndex),
				LogIndex:            int64(increase.Raw.Index),
			})
		}

		if ethtypes.Barn.IsBarnDelegatedPowerDecreasedEvent(&log) {
			decrease, err := ethtypes.Barn.BarnDelegatedPowerDecreasedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode delegate power increased event")
			}
			s.processed.delegateChanges = append(s.processed.delegateChanges, DelegateChange{
				Sender:              decrease.From.String(),
				Receiver:            decrease.To.String(),
				Amount:              decrease.Amount,
				ToNewDelegatedPower: decrease.ToNewDelegatedPower,
				ActionType:          DELEGATE_DECREASE,
				TransactionHash:     decrease.Raw.TxHash.String(),
				TransactionIndex:    int64(decrease.Raw.TxIndex),
				LogIndex:            int64(decrease.Raw.Index),
			})
		}
	}
	return nil
}
