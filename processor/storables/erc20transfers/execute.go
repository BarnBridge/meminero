package erc20transfers

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.IsMonitoredERC20(log.Address.String()) && len(log.Topics) == 3 && ethtypes.ERC20.IsTransferEvent(&log) {
				erc20Transfer, err := ethtypes.ERC20.TransferEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not decode erc20 transfer in tx %s", log.TxHash.String())
				}

				s.processed.transfers = append(s.processed.transfers, erc20Transfer)
			}
		}
	}

	return nil
}
