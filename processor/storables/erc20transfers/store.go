package erc20transfers

import (
	"context"

	"github.com/barnbridge/meminero/utils"
	"github.com/jackc/pgx/v4"
)

func (s *Storable) storeERC20Transfers(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.transfers) == 0 {
		s.logger.WithField("handler", "erc20Transfers").Debug("no events found")
		return nil
	}
	var rows [][]interface{}

	for _, t := range s.processed.transfers {
		rows = append(rows, []interface{}{
			utils.NormalizeAddress(t.Raw.Address.String()),
			utils.NormalizeAddress(t.From.String()),
			utils.NormalizeAddress(t.To.String()),
			t.ValueDecimal(0),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(t.Raw.TxHash.String()),
			t.Raw.TxIndex,
			t.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"erc20_transfers"},
		[]string{"token_address", "sender", "receiver", "value", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

	return nil
}
