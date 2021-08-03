package yieldfarming

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.YieldFarming.Address) {
				if ethtypes.YieldFarming.IsDepositEvent(&log) {
					d, err := ethtypes.YieldFarming.DepositEvent(log)
					if err != nil {
						return errors.Wrap(err, "could nod decode deposit event")
					}

					s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
						UserAddress:      utils.NormalizeAddress(d.User.String()),
						TokenAddress:     utils.NormalizeAddress(d.TokenAddress.String()),
						Amount:           d.AmountDecimal(0),
						ActionType:       DEPOSIT,
						TransactionHash:  utils.NormalizeAddress(d.Raw.TxHash.String()),
						TransactionIndex: int64(d.Raw.TxIndex),
						LogIndex:         int64(d.Raw.Index),
					})
				}

				if ethtypes.YieldFarming.IsWithdrawEvent(&log) {
					w, err := ethtypes.YieldFarming.WithdrawEvent(log)
					if err != nil {
						return errors.Wrap(err, "could nod decode withdraw event")
					}

					s.processed.stakingActions = append(s.processed.stakingActions, StakingAction{
						UserAddress:      utils.NormalizeAddress(w.User.String()),
						TokenAddress:     utils.NormalizeAddress(w.TokenAddress.String()),
						Amount:           w.AmountDecimal(0),
						ActionType:       WITHDRAW,
						TransactionHash:  utils.NormalizeAddress(w.Raw.TxHash.String()),
						TransactionIndex: int64(w.Raw.TxIndex),
						LogIndex:         int64(w.Raw.Index),
					})
				}
			}
		}
	}

	return nil
}
