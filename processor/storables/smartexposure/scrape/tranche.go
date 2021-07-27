package scrape

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) getTranchesDetailsFromChain(ctx context.Context, tranches []ethtypes.EtokenfactoryCreatedETokenEvent) error {
	for _, t := range tranches {
		var newTranche types.TrancheFromChain
		var symbol string

		wg, _ := errgroup.WithContext(ctx)
		wg.Go(eth.CallContractFunction(*ethtypes.Epool.ABI, t.EPool.String(), "getTranche", []interface{}{t.EToken}, &newTranche))
		wg.Go(eth.CallContractFunction(*ethtypes.ERC20.ABI, t.EToken.String(), "symbol", []interface{}{}, &symbol))

		err := wg.Wait()
		if err != nil {
			return errors.Wrap(err, "could not get tranche details from chain")
		}

		factor := decimal.NewFromBigInt(newTranche.SFactorE, 0)
		targetRatio := decimal.NewFromBigInt(newTranche.TargetRatio, 0)
		ratioA, ratioB := s.calculateRatios(factor, targetRatio)
		s.processed.newTranches = append(s.processed.newTranches, types.Tranche{
			EPoolAddress:  utils.NormalizeAddress(t.EPool.String()),
			ETokenAddress: utils.NormalizeAddress(t.EToken.String()),
			ETokenSymbol:  symbol,
			SFactorE:      factor,
			TargetRatio:   targetRatio,
			TokenARatio:   ratioA,
			TokenBRatio:   ratioB,
		})
	}

	return nil

}

func (s *Storable) calculateRatios(factor decimal.Decimal, targetRatio decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	ratioWithDec := targetRatio.Div(factor)
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)
	return tokenARatio, tokenBRatio
}
