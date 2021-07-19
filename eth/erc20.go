package eth

import (
	"context"

	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
)

func GetERC20TokenFromChain(addr string) (*types.Token, error) {
	var symbol string
	var decimals uint8

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(CallContractFunction(*ethtypes.ERC20.ABI, addr, "symbol", []interface{}{}, &symbol))
	wg.Go(CallContractFunction(*ethtypes.ERC20.ABI, addr, "decimals", []interface{}{}, &decimals))

	err := wg.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "could not get token info")
	}

	return &types.Token{
		Address:  utils.NormalizeAddress(addr),
		Symbol:   symbol,
		Decimals: int64(decimals),
	}, nil
}
