package events

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	globalTypes "github.com/barnbridge/meminero/types"
)

type Storable struct {
	block *globalTypes.Block

	state  *state.Manager
	logger *logrus.Entry

	processed struct {
		States     []State
		EpochInfos []EpochInfo
	}
}

type State struct {
	PoolAddress                string
	QueuedJuniorsUnderlyingIn  *big.Int
	QueuedJuniorsUnderlyingOut *big.Int
	QueuedJuniorTokensBurn     *big.Int
	QueuedSeniorsUnderlyingIn  *big.Int
	QueuedSeniorsUnderlyingOut *big.Int
	QueuedSeniorTokensBurn     *big.Int
	EstimatedJuniorLiquidity   *big.Int
	EstimatedSeniorLiquidity   *big.Int
	EstimatedJuniorTokenPrice  *big.Int
	EstimatedSeniorTokenPrice  *big.Int
}

type EpochInfo struct {
	PoolAddress            string
	Epoch                  *big.Int
	SeniorLiquidity        *big.Int
	JuniorLiquidity        *big.Int
	UpsideExposureRate     *big.Int
	DownsideProtectionRate *big.Int
	JuniorTokenPrice       *big.Int
	SeniorTokenPrice       *big.Int
}

func New(block *globalTypes.Block, state *state.Manager) *Storable {
	return &Storable{
		block:  block,
		state:  state,
		logger: logrus.WithField("module", "storable(smartAlpha.state)"),
	}
}

func (s *Storable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	abi := *ethtypes.SmartAlpha.ABI

	wg, ctx1 := errgroup.WithContext(ctx)
	mu := new(sync.Mutex)

	for _, p := range s.state.SmartAlpha.Pools {
		p := p

		wg.Go(func() error {
			var epoch *big.Int

			var poolState = &State{
				PoolAddress: p.PoolAddress,
			}

			subwg, _ := errgroup.WithContext(ctx1)
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "epoch", []interface{}{}, &epoch, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedJuniorsUnderlyingIn", []interface{}{}, &poolState.QueuedJuniorsUnderlyingIn, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedJuniorsUnderlyingOut", []interface{}{}, &poolState.QueuedJuniorsUnderlyingOut, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedJuniorTokensBurn", []interface{}{}, &poolState.QueuedJuniorTokensBurn, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedSeniorsUnderlyingIn", []interface{}{}, &poolState.QueuedSeniorsUnderlyingIn, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedSeniorsUnderlyingOut", []interface{}{}, &poolState.QueuedSeniorsUnderlyingOut, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "queuedSeniorTokensBurn", []interface{}{}, &poolState.QueuedSeniorTokensBurn, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "estimateCurrentJuniorLiquidity", []interface{}{}, &poolState.EstimatedJuniorLiquidity, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "estimateCurrentSeniorLiquidity", []interface{}{}, &poolState.EstimatedSeniorLiquidity, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "estimateCurrentJuniorTokenPrice", []interface{}{}, &poolState.EstimatedJuniorTokenPrice, s.block.Number))
			subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "estimateCurrentSeniorTokenPrice", []interface{}{}, &poolState.EstimatedSeniorTokenPrice, s.block.Number))

			err := subwg.Wait()
			if err != nil {
				return errors.Wrap(err, "could not get pool state")
			}

			has, err := s.state.SmartAlpha.HasEpochInfo(ctx, p.PoolAddress, epoch.Int64())
			if err != nil {
				return errors.Wrap(err, "could not check epoch info")
			}
			if !has {
				var epochInfo = &EpochInfo{
					PoolAddress: p.PoolAddress,
					Epoch:       epoch,
				}

				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "epochSeniorLiquidity", []interface{}{}, &epochInfo.SeniorLiquidity, s.block.Number))
				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "epochJuniorLiquidity", []interface{}{}, &epochInfo.JuniorLiquidity, s.block.Number))
				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "epochUpsideExposureRate", []interface{}{}, &epochInfo.UpsideExposureRate, s.block.Number))
				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "epochDownsideProtectionRate", []interface{}{}, &epochInfo.DownsideProtectionRate, s.block.Number))
				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "getEpochJuniorTokenPrice", []interface{}{}, &epochInfo.JuniorTokenPrice, s.block.Number))
				subwg.Go(eth.CallContractFunction(abi, p.PoolAddress, "getEpochSeniorTokenPrice", []interface{}{}, &epochInfo.SeniorTokenPrice, s.block.Number))

				err := subwg.Wait()
				if err != nil {
					return errors.Wrap(err, "could not epoch info from state")
				}

				mu.Lock()
				s.processed.EpochInfos = append(s.processed.EpochInfos, *epochInfo)
				mu.Unlock()
			}

			mu.Lock()
			s.processed.States = append(s.processed.States, *poolState)
			mu.Unlock()

			return nil
		})
	}

	return wg.Wait()
}

func (s *Storable) Result() interface{} {
	return s.processed
}
