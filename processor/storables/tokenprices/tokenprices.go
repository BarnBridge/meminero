package tokenprices

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
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
		logger: logrus.WithField("module", "storable(tokenPrices)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()
	var err error
	s.processed.prices, err = GetTokensPrices(ctx, s.state.Tokens, s.block.Number)
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

	_, err := tx.Exec(ctx, `delete from token_prices where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Trace("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Trace("done storing")
	}()

	err := s.storeTokensPrice(ctx, tx)
	return err
}

func (s *Storable) Result() interface{} {
	return s.processed
}
