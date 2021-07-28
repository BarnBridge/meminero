package rewards

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	factories []string

	processed struct {
		Pools       []smartyield.RewardPool
		Claims      []ethtypes.SmartYieldPoolSingleClaimEvent
		ClaimsMulti []ethtypes.SmartYieldPoolMultiClaimRewardTokenEvent
		Deposits    []ethtypes.SmartYieldPoolSingleDepositEvent
		Withdrawals []ethtypes.SmartYieldPoolSingleWithdrawEvent
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:     block,
		state:     state,
		logger:    logrus.WithField("module", "storable(smartYield.rewards)"),
		factories: strings.Split(config.Store.Storable.SmartYield.Rewards.Factories, ","),
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
			if s.isFactory(log.Address.String()) {
				if ethtypes.SmartYieldPoolFactorySingle.IsPoolCreatedEvent(&log) {
					p, err := ethtypes.SmartYieldPoolFactorySingle.PoolCreatedEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode PoolSingle.PoolCreated event")
					}

					err = s.processNewPoolSingle(ctx, p)
					if err != nil {
						return errors.Wrap(err, "could not process new PoolSingle")
					}
				}

				if ethtypes.SmartYieldPoolFactoryMulti.IsPoolMultiCreatedEvent(&log) {
					p, err := ethtypes.SmartYieldPoolFactoryMulti.PoolMultiCreatedEvent(log)
					if err != nil {
						return errors.Wrap(err, "could not decode PoolMulti.PoolMultiCreatedEvent")
					}

					err = s.processNewPoolMulti(ctx, p)
					if err != nil {
						return errors.Wrap(err, "could not process new PoolMulti")
					}
				}
			}

			if s.state.SmartYield.RewardPoolByAddress(log.Address.String()) != nil {
				err := s.processRewardPoolEvent(log)
				if err != nil {
					return errors.Wrapf(err, "could not process reward pool event %s")
				}
			}
		}
	}

	return nil
}

func (s Storable) Result() interface{} {
	return s.processed
}
