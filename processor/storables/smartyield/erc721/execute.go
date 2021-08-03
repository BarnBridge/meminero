package erc721

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartYield.IsERC721OfInterest(log.Address.String()) && ethtypes.ERC721.IsTransferEvent(&log) {
				e, err := ethtypes.ERC721.TransferEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode ERC721 Transfer event")
				}

				s.processed.Transfers = append(s.processed.Transfers, e)
			}
		}
	}

	return nil
}
