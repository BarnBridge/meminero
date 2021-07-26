package scrape

import (
	"context"
	"math/big"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *Storable) getTranchesDetailsFromChain(ctx context.Context, tranches []ethtypes.EtokenfactoryCreatedETokenEvent) error {
	for _, t := range tranches {

		var newTranche smartexposure.TrancheFromChain
		a := ethtypes.Epool.ABI
		err := eth.CallContractFunction(*a, t.EPool.String(), "getTranche", []interface{}{t.EToken}, &newTranche)()
		if err != nil {
			return errors.Wrap(err, "could not get tranche details from chain")
		}

		var symbol string
		err = eth.CallContractFunction(*ethtypes.ERC20.ABI, t.EToken.String(), "symbol", []interface{}{}, &symbol)()
		if err != nil {
			return errors.Wrap(err, "could not get tranche symbol from chain")
		}

		ratioA, ratioB := s.calculateRatios(newTranche.SFactorE, newTranche.TargetRatio)
		s.processed.newTranches = append(s.processed.newTranches, types.SETranche{
			EPoolAddress:  utils.NormalizeAddress(t.EPool.String()),
			ETokenAddress: utils.NormalizeAddress(t.EToken.String()),
			ETokenSymbol:  symbol,
			SFactorE:      newTranche.SFactorE,
			TargetRatio:   newTranche.TargetRatio,
			TokenARatio:   ratioA,
			TokenBRatio:   ratioB,
		})
	}
	return nil
}

func (s *Storable) calculateRatios(factor *big.Int, targetRatio *big.Int) (decimal.Decimal, decimal.Decimal) {
	ratioWithDec := decimal.NewFromBigInt(targetRatio, 0).Div(decimal.NewFromBigInt(factor, 0))
	tokenBRatio := decimal.NewFromInt(1).Div(ratioWithDec.Add(decimal.NewFromInt(1)))
	tokenARatio := decimal.NewFromInt(1).Sub(tokenBRatio)

	return tokenARatio, tokenBRatio
}
