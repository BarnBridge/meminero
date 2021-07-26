package scrape

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.Amount,
				AmountA:         t.AmountA,
				AmountB:         t.AmountB,
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

			s.processed.seTransactions = append(s.processed.seTransactions, SETransaction{
				ETokenAddress:   utils.NormalizeAddress(t.EToken.String()),
				UserAddress:     utils.NormalizeAddress(t.User.String()),
				Amount:          t.Amount,
				AmountA:         t.AmountA,
				AmountB:         t.AmountB,
				TransactionType: "WITHDRAW",
				TxHash:          utils.NormalizeAddress(t.Raw.TxHash.String()),
				TxIndex:         int64(t.Raw.TxIndex),
				LogIndex:        int64(t.Raw.Index),
			})
		}
	}
	return nil
}
