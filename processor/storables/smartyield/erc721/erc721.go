package erc721

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		Transfers []ethtypes.ERC721TransferEvent
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartYield.erc721)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartYield.IsERC721OfInterest(log.Address.String()) && ethtypes.ERC721.IsERC721TransferEvent(&log) {
				e, err := ethtypes.ERC721.ERC721TransferEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode ERC721 Transfer event")
				}

				s.processed.Transfers = append(s.processed.Transfers, e)
			}
		}
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
