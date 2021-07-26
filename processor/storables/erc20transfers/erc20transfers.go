package erc20transfers

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry
	state  *state.Manager

	processed struct {
		transfers []ethtypes.ERC20TransferEvent
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(erc20_transfers)"),
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
			if s.state.IsMonitoredERC20(log.Address.String()) && len(log.Topics) == 3 && ethtypes.ERC20.IsERC20TransferEvent(&log) {
				erc20Transfer, err := ethtypes.ERC20.ERC20TransferEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not decode erc20 transfer in tx %s", log.TxHash.String())
				}

				s.processed.transfers = append(s.processed.transfers, erc20Transfer)
			}
		}
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from public.erc20_transfers where included_in_block = $1`, s.block.Number)
	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Debug("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done storing")
	}()

	err := s.storeERC20Transfers(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}
	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
