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

		if ethtypes.Epool.IsIssuedETokenEvent(&log) {
			t, err := ethtypes.Epool.IssuedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode IssuedEtoken event")
			}

			if utils.NormalizeAddress(t.User.String()) == config.Store.Storable.SmartExposure.EPoolPeripheryAddress {
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

		if ethtypes.Epool.IsRedeemedETokenEvent(&log) {
			t, err := ethtypes.Epool.RedeemedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode RedeemedEToken event")
			}
			if utils.NormalizeAddress(t.User.String()) == config.Store.Storable.SmartExposure.EPoolPeripheryAddress {
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

		if ethtypes.Epoolperiphery.IsIssuedETokenEvent(&log) {
			t, err := ethtypes.Epoolperiphery.IssuedETokenEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode issuedEToken event from epoolperiphery contract")
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

		if ethtypes.Epoolperiphery.IsRedeemedETokenEvent(&log) {
			t, err := ethtypes.Epoolperiphery.RedeemedETokenEvent(log)
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
