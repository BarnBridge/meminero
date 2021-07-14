package accountERC20Transfers

import (
	"context"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

const (
	AmountIn  = "IN"
	AmountOut = "OUT"
)

func (s *Storable) storeTransfers(tx pgx.Tx) error {
	if len(s.processed.transfers) == 0 {
		return nil
	}

	_, err := tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"account_erc20_transfers"},
		[]string{"token_address", "account", "counterparty", "amount", "tx_direction", "tx_hash", "tx_index", "log_index", "included_in_block", "block_timestamp"},
		pgx.CopyFromSlice(len(s.processed.transfers), func(i int) ([]interface{}, error) {
			return []interface{}{utils.NormalizeAddress(s.processed.transfers[i].Raw.Address.String()), utils.NormalizeAddress(s.processed.transfers[i].From.String()), utils.NormalizeAddress(s.processed.transfers[i].To.String()),
				 decimal.NewFromBigInt(s.processed.transfers[i].Value,0), AmountOut, utils.NormalizeAddress(s.processed.transfers[i].Raw.TxHash.String()), int(s.processed.transfers[i].Raw.TxIndex), int(s.processed.transfers[i].Raw.Index),
				s.processed.blockNumber, s.processed.blockTimestamp}, nil
		}),
	)
	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"account_erc20_transfers"},
		[]string{"token_address", "account", "counterparty", "amount", "tx_direction", "tx_hash", "tx_index", "log_index", "included_in_block", "block_timestamp"},
		pgx.CopyFromSlice(len(s.processed.transfers), func(i int) ([]interface{}, error) {
			return []interface{}{utils.NormalizeAddress(s.processed.transfers[i].Raw.Address.String()), utils.NormalizeAddress(s.processed.transfers[i].To.String()), utils.NormalizeAddress(s.processed.transfers[i].From.String()),
				decimal.NewFromBigInt(s.processed.transfers[i].Value,0), AmountIn, utils.NormalizeAddress(s.processed.transfers[i].Raw.TxHash.String()), s.processed.transfers[i].Raw.TxIndex, s.processed.transfers[i].Raw.Index,
				s.processed.blockNumber, s.processed.blockTimestamp}, nil
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
