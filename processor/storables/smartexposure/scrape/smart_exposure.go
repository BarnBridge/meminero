package scrape

import (
	"context"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	types2 "github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		seTransactions []SETransaction
		newTranches    []types2.Tranche
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartExposure.scrape)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var epoolTxs []gethtypes.Log
	var newETokens []ethtypes.EtokenfactoryCreatedETokenEvent
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartExposure.SEPoolByAddress(log.Address.String()) != nil ||
				utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.EPoolPeripheryAddress) {
				epoolTxs = append(epoolTxs, log)
			}

			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.ETokenFactoryAddress) &&
				ethtypes.Etokenfactory.IsCreatedETokenEvent(&log) {
				eToken, err := ethtypes.Etokenfactory.CreatedETokenEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode Created EToken event")
				}
				newETokens = append(newETokens, eToken)
			}
		}

	}

	err := s.decodePoolTransactions(epoolTxs)
	if err != nil {
		return err
	}

	err = s.getTranchesDetailsFromChain(ctx, newETokens)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	_, err := tx.Exec(ctx, `delete from smart_exposure.transaction_history where included_in_block = $1`, s.block.Number)
	if err != nil {
		return errors.Wrap(err, "could not execute rollback on smart_exposure.transaction_history")
	}
	_, err = tx.Exec(ctx, `delete from smart_exposure.tranches where start_at_block = $1`, s.block.Number)
	return errors.Wrap(err, "could not execute rollback on smart_exposure.tranches table")

}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	err := s.storeEPoolTransactions(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store smart exposure transactions")
	}

	err = s.storeNewTranches(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store new smart exposure tranches")
	}

	return nil
}

func (s *Storable) Result() interface{} {
	return s.processed
}
