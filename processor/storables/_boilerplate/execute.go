package boilerplate

import (
	"context"

	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			// do something
			if log.Address.String() == utils.ZeroAddress {

			}
		}
	}

	return nil
}
