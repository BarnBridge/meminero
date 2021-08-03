package tokenprices

import (
	"context"
	"math/big"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func GetTokensPrices(ctx context.Context, tokens map[string]types.Token, blockNumber int64) (map[string]map[string]decimal.Decimal, error) {
	var calls = make(map[string]types.PriceProvider)
	for _, t := range tokens {
		for _, v := range t.Prices {
			if v.Provider != "chainlink" {
				return nil, errors.New("invalid provider")
			}

			if blockNumber < v.StartAtBlock {
				continue
			}

			for _, p := range v.Path {
				calls[utils.NormalizeAddress(p.Address)] = p
			}
		}
	}

	wg, _ := errgroup.WithContext(ctx)
	var mu = &sync.Mutex{}
	results := make(map[string]decimal.Decimal)

	for _, c := range calls {
		c := c
		wg.Go(func() error {
			var price *big.Int
			err := eth.CallContractFunction(*ethtypes.ETHAggregator.ABI, c.Address, "latestAnswer", []interface{}{}, &price, blockNumber)()
			if err != nil {
				return errors.Wrapf(err, "could not call latestAnswer on contract(%s)", c.Address)
			}

			mu.Lock()
			priceScaled := decimal.NewFromBigInt(price, -int32(c.Decimals))
			results[utils.NormalizeAddress(c.Address)] = priceScaled
			mu.Unlock()

			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get prices from oracles")
	}

	prices := make(map[string]map[string]decimal.Decimal)
	for _, t := range tokens {
		for _, v := range t.Prices {
			if blockNumber < v.StartAtBlock {
				continue
			}

			if len(v.Path) > 0 {
				price := decimal.NewFromInt(1)

				for _, p := range v.Path {
					res := results[utils.NormalizeAddress(p.Address)]
					if p.Reverse {
						res = decimal.NewFromInt(1).Div(res)
					}

					price = price.Mul(res)
				}

				if prices[t.Address] == nil {
					prices[t.Address] = make(map[string]decimal.Decimal)
				}

				prices[t.Address][strings.ToUpper(v.Quote)] = price
			}
		}
	}

	return prices, nil
}

type Price struct {
	QuoteAsset string
	Value      decimal.Decimal
}
