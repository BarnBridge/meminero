package events

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	var epoolTxs []gethtypes.Log
	var newETokens []ethtypes.ETokenFactoryCreatedETokenEvent
	for _, tx := range s.block.Txs {
		for _, log := range tx.LogEntries {
			if s.state.SmartExposure.PoolByAddress(log.Address.String()) != nil ||
				utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.EPoolPeripheryAddress) {
				epoolTxs = append(epoolTxs, log)
			}

			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.ETokenFactoryAddress) &&
				ethtypes.ETokenFactory.IsCreatedETokenEvent(&log) {
				eToken, err := ethtypes.ETokenFactory.CreatedETokenEvent(log)
				if err != nil {
					return errors.Wrap(err, "could not decode Created EToken event")
				}
				newETokens = append(newETokens, eToken)
			}
		}

	}

	err := s.decodePoolTransactions(epoolTxs)
	if err != nil {
		return err
	}

	err = s.getTranchesDetailsFromChain(ctx, newETokens)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storable) decodePoolTransactions(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.EPool.IsIssuedETokenEvent(&log) {
			t, err := ethtypes.EPool.IssuedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode IssuedEtoken event")
			}

			// ignore events that where userAddress = epoolPeripheryAddress because there exists another events that comes from epoolPeriphery
			if utils.NormalizeAddress(t.User.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.EPoolPeripheryAddress) {
				continue
			}

			s.processed.seTransactions = append(s.processed.seTransactions, Transaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.AmountDecimal(0),
				AmountA:         t.AmountADecimal(0),
				AmountB:         t.AmountBDecimal(0),
				TransactionType: "DEPOSIT",
				TxHash:          utils.NormalizeAddress(t.Raw.TxHash.String()),
				TxIndex:         int64(t.Raw.TxIndex),
				LogIndex:        int64(t.Raw.Index),
			})
		}

		if ethtypes.EPool.IsRedeemedETokenEvent(&log) {
			t, err := ethtypes.EPool.RedeemedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode RedeemedEToken event")
			}

			// ignore events that where userAddress = epoolPeripheryAddress because there exists another events that comes from epoolPeriphery
			if utils.NormalizeAddress(t.User.String()) == utils.NormalizeAddress(config.Store.Storable.SmartExposure.EPoolPeripheryAddress) {
				continue
			}

			s.processed.seTransactions = append(s.processed.seTransactions, Transaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.AmountDecimal(0),
				AmountA:         t.AmountADecimal(0),
				AmountB:         t.AmountBDecimal(0),
				TransactionType: "WITHDRAW",
				TxHash:          utils.NormalizeAddress(t.Raw.TxHash.String()),
				TxIndex:         int64(t.Raw.TxIndex),
				LogIndex:        int64(t.Raw.Index),
			})
		}

		if ethtypes.EPoolPeriphery.IsIssuedETokenEvent(&log) {
			t, err := ethtypes.EPoolPeriphery.IssuedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode issuedEToken event from epoolperiphery contract")
			}

			s.processed.seTransactions = append(s.processed.seTransactions, Transaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.AmountDecimal(0),
				AmountA:         t.AmountADecimal(0),
				AmountB:         t.AmountBDecimal(0),
				TransactionType: "DEPOSIT",
				TxHash:          utils.NormalizeAddress(t.Raw.TxHash.String()),
				TxIndex:         int64(t.Raw.TxIndex),
				LogIndex:        int64(t.Raw.Index),
			})
		}

		if ethtypes.EPoolPeriphery.IsRedeemedETokenEvent(&log) {
			t, err := ethtypes.EPoolPeriphery.RedeemedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode RedeemedEToken event from epoolperiphery contract")
			}

			s.processed.seTransactions = append(s.processed.seTransactions, Transaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.AmountDecimal(0),
				AmountA:         t.AmountADecimal(0),
				AmountB:         t.AmountBDecimal(0),
				TransactionType: "WITHDRAW",
				TxHash:          utils.NormalizeAddress(t.Raw.TxHash.String()),
				TxIndex:         int64(t.Raw.TxIndex),
				LogIndex:        int64(t.Raw.Index),
			})
		}
	}
	return nil
}

func (s *Storable) getTranchesDetailsFromChain(ctx context.Context, tranches []ethtypes.ETokenFactoryCreatedETokenEvent) error {
	for _, t := range tranches {
		var newTranche types.TrancheFromChain
		var symbol string

		wg, _ := errgroup.WithContext(ctx)
		wg.Go(eth.CallContractFunction(*ethtypes.EPool.ABI, t.EPool.String(), "getTranche", []interface{}{t.EToken}, &newTranche))
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
