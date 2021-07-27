package tranche_state

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	seTypes "github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/processor/storables/tokenprices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *types.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		trancheState map[string]TrancheState
		tokenPrices  map[string]decimal.Decimal
	}
}

func New(block *types.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartExposure.trancheState)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	s.processed.trancheState = make(map[string]TrancheState)
	tokens, err := smartexposure.BuildTokensSliceForSE(s.state)
	if err != nil {
		return err
	}

	s.processed.tokenPrices, err = tokenprices.GetTokensPrices(ctx, tokens, s.block.Number)
	if err != nil {
		return err
	}

	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}

	for trancheAddress, tranche := range s.state.SmartExposure.Tranches {
		if s.block.Number < tranche.StartAtBlock {
			s.logger.WithField("tranche", trancheAddress).Info("skipping tranche due to StartAtBlock property")
			continue
		}

		trancheAddress := trancheAddress
		tranche := tranche

		wg.Go(func() error {
			var t seTypes.TrancheFromChain
			var currentRatio *big.Int
			var conversionRate struct {
				AmountA *big.Int `json:"amountA"`
				AmountB *big.Int `json:"amountB"`
			}

			subwg, _ := errgroup.WithContext(ctx)
			subwg.Go(eth.CallContractFunction(*ethtypes.Epool.ABI, tranche.EPoolAddress, "getTranche", []interface{}{common.HexToAddress(trancheAddress)}, &t, s.block.Number))
			subwg.Go(eth.CallContractFunction(*ethtypes.Epoolhelper2.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "currentRatio", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress)}, &currentRatio, s.block.Number))
			subwg.Go(eth.CallContractFunction(*ethtypes.Epoolhelper2.ABI, config.Store.Storable.SmartExposure.EPoolHelperAddress, "tokenATokenBForEToken", []interface{}{common.HexToAddress(tranche.EPoolAddress), common.HexToAddress(trancheAddress), tranche.SFactorE.BigInt()}, &conversionRate, s.block.Number))

			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			s.processed.trancheState[trancheAddress] = TrancheState{
				EPoolAddress:    tranche.EPoolAddress,
				CurrentRatio:    decimal.NewFromBigInt(currentRatio, 0),
				TokenALiquidity: decimal.NewFromBigInt(t.ReserveA, 0),
				TokenBLiquidity: decimal.NewFromBigInt(t.ReserveB, 0),
				ConversionRate: ConversionRate{
					AmountAConversion: decimal.NewFromBigInt(conversionRate.AmountA, 0),
					AmountBConversion: decimal.NewFromBigInt(conversionRate.AmountB, 0),
				},
			}
			mu.Unlock()

			return nil
		})
	}

	err = wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not get data from chain")
	}

	return nil
}

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	_, err := tx.Exec(ctx, `delete from smart_exposure.tranche_state where included_in_block = $1`, s.block.Number)

	return err
}

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	err := s.storeTranchesState(ctx, tx)
	return err
}

func (s *Storable) Result() interface{} {
	return s.processed

}
