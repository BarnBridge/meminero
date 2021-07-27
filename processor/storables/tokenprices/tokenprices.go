package tokenprices

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Storable struct {
	block *types.Block

	logger *logrus.Entry
	state  *state.Manager

	processed struct {
		prices map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(tokensPrices)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	s.processed.prices = make(map[string]decimal.Decimal)

	for _, t := range s.state.Tokens {
		t := t
		wg.Go(func() error {
			var tokenPrice *big.Int
			err := eth.CallContractFunction(*ethtypes.Ethaggregator.ABI, t.AggregatorAddress, "latestAnswer", []interface{}{}, &tokenPrice, s.block.Number)()
			if err != nil {
				return err
			}

			mu.Lock()
			s.processed.prices[t.Address] = decimal.NewFromBigInt(tokenPrice, -int32(8))
			mu.Unlock()

			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		return errors.Wrap(err, "failed to call latestAnswer")
	}
	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from token_prices where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Debug("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done storing")
	}()

	err := s.storeTokensPrice(ctx, tx)
	return err
}

func (s *Storable) Result() interface{} {
	return s.processed
}
