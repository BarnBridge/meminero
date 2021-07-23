package yieldfarming

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
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
	var logs []gethtypes.Log
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.YieldFarming.Address) {
				logs = append(logs, log)
			}
		}
	}
	if len(logs) == 0 {
		s.logger.WithField("module", "yield farming").Debug("no events found")
	}

	err := s.decodeStakingActions(logs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from yield_farming.transactions where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeStakingActions(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store erc20transfers")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
