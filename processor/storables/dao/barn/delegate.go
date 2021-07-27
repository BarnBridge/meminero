package barn

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) handleDelegateEvents(logs []gethtypes.Log, ctx context.Context) error {

	for _, log := range logs {
		if ethtypes.Barn.IsDelegateEvent(&log) {
			var action DelegateAction
			a, err := ethtypes.Barn.DelegateEvent(log)
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
				ActionType:          DELEGATE_INCREASE,
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
				ActionType:          DELEGATE_DECREASE,
				TransactionHash:     utils.NormalizeAddress(decrease.Raw.TxHash.String()),
				TransactionIndex:    int64(decrease.Raw.TxIndex),
				LogIndex:            int64(decrease.Raw.Index),
			})
		}
	}
	return nil
}
