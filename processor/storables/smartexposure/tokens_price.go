package smartexposure

import (
	"context"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/types"
)

func GetTokensPrice(ctx context.Context, state *state.Manager, blockNumber int64) (map[string]decimal.Decimal, error) {
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}

	tokensPrices := make(map[string]decimal.Decimal)
	tokens := make(map[string]types.Token)

	for _, pool := range state.SmartExposure.Pools {
		tokens[pool.ATokenAddress] = state.Tokens[pool.ATokenAddress]
		tokens[pool.BTokenAddress] = state.Tokens[pool.BTokenAddress]
	}

	for _, t := range tokens {
		t := t
		wg.Go(func() error {
			var tokenPrice *big.Int
			err := eth.CallContractFunction(*ethtypes.Ethaggregator.ABI, t.AggregatorAddress, "latestAnswer", []interface{}{}, &tokenPrice, blockNumber)()
			if err != nil {
				return err
			}

			mu.Lock()
			// NOTE there might be times when decimals might be different than 8
			tokensPrices[t.Address] = decimal.NewFromBigInt(tokenPrice, -int32(8))
			mu.Unlock()

			return nil
		})
	}

	err := wg.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "failed to call latestAnswer")
	}

	return tokensPrices, nil
}
