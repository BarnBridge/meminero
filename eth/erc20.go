package eth

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	abiStore "github.com/barnbridge/smartbackend/abi"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
)

func GetERC20TokenFromChain(addr string) (*types.Token, error) {
	a, err := abiStore.Get("erc20")
	if err != nil {
		return nil, errors.Wrap(err, "could not find abi")
	}

	var symbol string
	var decimals uint8

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(CallContractFunction(*a, addr, "symbol", []interface{}{}, &symbol))
	wg.Go(CallContractFunction(*a, addr, "decimals", []interface{}{}, &decimals))

	err = wg.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "could not get token info")
	}

	return &types.Token{
		Address:  utils.NormalizeAddress(addr),
		Symbol:   symbol,
		Decimals: int64(decimals),
	}, nil
}
