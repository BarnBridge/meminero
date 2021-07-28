package yieldfarming

import (
	"context"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

type Storable struct {
	block     *types.Block
	logger    *logrus.Entry
	processed struct {
		stakingActions []StakingAction
	}
}

func New(block *types.Block) *Storable {
	return &Storable{
		block:  block,
		logger: logrus.WithField("module", "storable(yieldFarming)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	start := time.Now()
	s.logger.Trace("executing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done")
	}()

	var logs []gethtypes.Log
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.YieldFarming.Address) {
				logs = append(logs, log)
			}
		}
	}
	if len(logs) == 0 {
		return nil
	}

	err := s.decodeStakingActions(logs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Trace("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from yield_farming.transactions where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Trace("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done storing")
	}()

	err := s.storeStakingActions(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
