package smartexposure

import (
	"context"
	"math/big"
	"sync"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func GetTokensPrice(ctx context.Context, state *state.Manager, blockNumber int64) (error, map[string]decimal.Decimal) {
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	tokensPrices := make(map[string]decimal.Decimal)
	tokens := make(map[string]types.Token)
	for _, pool := range state.SmartExposure.SEPools() {
		tokens[pool.ATokenAddress] = state.Tokens[pool.ATokenAddress]
		tokens[pool.BTokenAddress] = state.Tokens[pool.BTokenAddress]
	}

	for _, t := range tokens {
		t := t
		var tokenPrice *big.Int
		wg.Go(eth.CallContractFunction(*ethtypes.Ethaggregator.ABI, t.AggregatorAddress, "latestAnswer", []interface{}{}, &tokenPrice))

		err := wg.Wait()
		if err != nil {
			return err, nil
		}

		mu.Lock()
		tokensPrices[t.Address] = decimal.NewFromBigInt(tokenPrice, -int32(8))
		mu.Unlock()

	}

	return nil, tokensPrices
}
