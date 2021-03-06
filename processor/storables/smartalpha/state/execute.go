package events

import (
	"context"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartalpha"
)

func (s *Storable) Execute(ctx context.Context) error {
	abi := *ethtypes.SmartAlpha.ABI

	wg, ctx1 := errgroup.WithContext(ctx)
	mu := new(sync.Mutex)

	for _, p := range s.state.SmartAlpha.Pools {
		p := p

		if s.block.Number < p.StartAtBlock {
			continue
		}

		wg.Go(func() error {
			var epoch *big.Int

			var poolState = &smartalpha.State{
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

			mu.Lock()
			s.processed.States = append(s.processed.States, *poolState)
			mu.Unlock()

			return nil
		})
	}

	return wg.Wait()
}
