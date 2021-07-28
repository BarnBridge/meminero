package scrape

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
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
