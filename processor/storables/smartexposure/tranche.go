package smartexposure

import (
	"context"
	"math/big"
	"sync"

	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
)

func (s *Storable) getTranchesDetailsFromChain(ctx context.Context, tranches []ethtypes.EtokenfactoryCreatedETokenEvent) error {
	wg, _ := errgroup.WithContext(ctx)
	var mu sync.Mutex

	for _, tranche := range tranches {
		t := tranche

		a := ethtypes.Epool.ABI
		wg.Go(func() error {
			subwg, _ := errgroup.WithContext(ctx)
			var newTranche types.SETranche

			subwg.Go(eth.CallContractFunction(*a, utils.NormalizeAddress(t.EPool.String()), "getTranche", []interface{}{utils.NormalizeAddress(t.EToken.String())}, &newTranche))
			subwg.Go(eth.CallContractFunction(*ethtypes.ERC20.ABI, utils.NormalizeAddress(t.EToken.String()), "symbol", []interface{}{}, &newTranche.ETokenSymbol))
			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			newTranche.EPoolAddress = utils.NormalizeAddress(t.EPool.String())
			newTranche.ETokenAddress = utils.NormalizeAddress(t.EToken.String())
			newTranche.TokenARatio, newTranche.TokenBRatio = s.calculateRatios(newTranche.SFactorE, newTranche.TargetRatio)
			s.processed.newTranches = append(s.processed.newTranches, newTranche)
			mu.Unlock()

			return nil
		})
	}

	err := wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not get new tranche info")
	}

	return nil
}

func (s *Storable) calculateRatios(factor *big.Int, targetRatio *big.Int) (decimal.Decimal, decimal.Decimal) {
	ratioWithDec := decimal.NewFromBigInt(targetRatio, 0).Div(decimal.NewFromBigInt(factor, 0))
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	return tokenARatio, tokenBRatio
}
