package accounterc20transfers

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/barnbridge/meminero/utils"
)

const (
	AmountIn  = "IN"
	AmountOut = "OUT"
)

func (s *Storable) storeTransfers(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.transfers) == 0 {
		return nil
	}
	var rows [][]interface{}

	for _, t := range s.processed.transfers {
		rows = append(rows, []interface{}{
			utils.NormalizeAddress(t.Raw.Address.String()),
			utils.NormalizeAddress(t.From.String()),
			utils.NormalizeAddress(t.To.String()),
			t.ValueDecimal(0),
			AmountOut,
			utils.NormalizeAddress(t.Raw.TxHash.String()),
			t.Raw.TxIndex,
			t.Raw.Index,
			s.block.Number,
			s.block.BlockCreationTime,
		}, []interface{}{
			utils.NormalizeAddress(t.Raw.Address.String()),
			utils.NormalizeAddress(t.To.String()),
			utils.NormalizeAddress(t.From.String()),
			t.ValueDecimal(0),
			AmountIn,
			utils.NormalizeAddress(t.Raw.TxHash.String()),
			t.Raw.TxIndex,
			t.Raw.Index,
			s.block.Number,
			s.block.BlockCreationTime,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"account_erc20_transfers"},
		[]string{"token_address", "account", "counterparty", "amount", "tx_direction", "tx_hash", "tx_index", "log_index", "included_in_block", "block_timestamp"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

	return nil
}
