package accounterc20transfers

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if len(log.Topics) == 3 && ethtypes.ERC20.IsTransferEvent(&log) {
				erc20Transfer, err := ethtypes.ERC20.TransferEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not decode erc20 transfer in tx %s", log.TxHash.String())
				}

				if !s.state.IsMonitoredAccount(erc20Transfer.From.String()) &&
					!s.state.IsMonitoredAccount(erc20Transfer.To.String()) {
					continue
				}

				s.processed.transfers = append(s.processed.transfers, erc20Transfer)

				exists := s.state.CheckTokenExists(log.Address.String())
				if !exists {
					token, err := eth.GetERC20TokenFromChain(log.Address.String())
					if err != nil {
						return err
					}

					err = s.state.StoreToken(ctx, *token)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
