package smartexposure

import (
	"context"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		seTransactions []SETransaction
		newTranches    []types.SETranche
	}
}

func (s *Storable) New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smart exposure)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	var epoolTxs []gethtypes.Log
	var newETokens []ethtypes.EtokenfactoryCreatedETokenEvent
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SEPoolByAddress(log.Address.String()) != nil ||
				utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.EPoolPeripheryAddress) {
				epoolTxs = append(epoolTxs, log)
			}

			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.ETokenFactoryAddress) {
				if ethtypes.Etokenfactory.IsEtokenfactoryCreatedETokenEvent(&log) {
					eToken, err := ethtypes.Etokenfactory.EtokenfactoryCreatedETokenEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode Created EToken event")
					}
					newETokens = append(newETokens, eToken)
				}
			}

		}
	}
	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	return nil
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
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
