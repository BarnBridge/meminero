package tokenprices

import (
	"context"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
)

func GetTokensPrices(ctx context.Context, tokens map[string]types.Token, blockNumber int64) (map[string]decimal.Decimal, error) {
	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	prices := make(map[string]decimal.Decimal)

	for _, t := range tokens {
		if t.AggregatorAddress == "" {
			continue
		}

		t := t
		wg.Go(func() error {
			var tokenPrice *big.Int
			err := eth.CallContractFunction(*ethtypes.ETHAggregator.ABI, t.AggregatorAddress, "latestAnswer", []interface{}{}, &tokenPrice, blockNumber)()
			if err != nil {
				return errors.Wrapf(err, "could not call latestAnswer on contract(%s)", t.AggregatorAddress)
			}

			mu.Lock()
			prices[t.Address] = decimal.NewFromBigInt(tokenPrice, -int32(8))
			mu.Unlock()

			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "failed to call latestAnswer")
	}
	return prices, nil
}
