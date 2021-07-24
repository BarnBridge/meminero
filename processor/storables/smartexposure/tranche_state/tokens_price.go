package tranche_state

import (
	"context"
	"math/big"
	"sync"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func (s *Storable) getTokensPrice(ctx context.Context) error {
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}

	tokens := make(map[string]types.Token)
	for _, pool := range s.state.SEPools() {
		tokens[pool.ATokenAddress] = s.state.Tokens[pool.ATokenAddress]
		tokens[pool.BTokenAddress] = s.state.Tokens[pool.BTokenAddress]
	}

	for _, t := range tokens {
		t := t
		var tokenPrice *big.Int
		wg.Go(eth.CallContractFunction(*ethtypes.Ethaggregator.ABI, t.AggregatorAddress, "latestAnswer", []interface{}{}, &tokenPrice, s.block.Number))

		err := wg.Wait()
		if err != nil {
			return err
		}

		mu.Lock()
		s.processed.tokenPrices[t.Address] = decimal.NewFromBigInt(tokenPrice, -int32(8))
		mu.Unlock()

	}

	return nil
}
