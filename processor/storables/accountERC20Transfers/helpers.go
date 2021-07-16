package accountERC20Transfers

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/smartbackend/utils"
)

const (
	AmountIn  = "IN"
	AmountOut = "OUT"
)

func (s *Storable) storeTransfers(tx pgx.Tx) error {
	if len(s.processed.transfers) == 0 {
		return nil
	}
	var rows [][]interface{}

	for _, t := range s.processed.transfers{
		value :=decimal.NewFromBigInt(t.Value,0)

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(t.Raw.Address.String()),
			utils.NormalizeAddress(t.From.String()),
			utils.NormalizeAddress(t.To.String()),
			value,
			AmountOut,
			utils.NormalizeAddress(t.Raw.TxHash.String()),
			t.Raw.TxIndex,
			t.Raw.Index,
			s.block.Number,
			s.block.BlockCreationTime,
		},[]interface{}{
			utils.NormalizeAddress(t.Raw.Address.String()),
			utils.NormalizeAddress(t.To.String()),
			utils.NormalizeAddress(t.From.String()),
			value,
			AmountIn,
			utils.NormalizeAddress(t.Raw.TxHash.String()),
			t.Raw.TxIndex,
			t.Raw.Index,
			s.block.Number,
			s.block.BlockCreationTime,
		})
	}

	_, err := tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"account_erc20_transfers"},
		[]string{"token_address", "account", "counterparty", "amount", "tx_direction", "tx_hash", "tx_index", "log_index", "included_in_block", "block_timestamp"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]interface{}, error) {
			return rows[i], nil
		}),
	)
	if err != nil {
		return err
	}

	return nil
}
