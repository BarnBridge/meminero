package accountERC20Transfers

import (
	"database/sql"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/lib/pq"
)
const (
	AmountIn  = "IN"
	AmountOut = "OUT"
)

func (s *Storable) storeTransfers(tx *sql.Tx) error {
	if len(s.processed.transfers) == 0 {
		return nil
	}

	stmt, err := tx.Prepare(pq.CopyIn("account_erc20_transfers", "token_address", "account", "counterparty", "amount", "tx_direction", "tx_hash", "tx_index", "log_index", "included_in_block", "block_timestamp"))
	if err != nil {
		return err
	}

	for _, t := range s.processed.transfers {
		_, err = stmt.Exec(utils.NormalizeAddress(t.Raw.Address.String()), utils.NormalizeAddress(t.From.String()), utils.NormalizeAddress(t.To.String()), t.Value.String(), AmountOut, utils.NormalizeAddress(t.Raw.TxHash.String()), t.Raw.TxIndex, t.Raw.Index, s.processed.blockNumber, s.processed.blockTimestamp)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(utils.NormalizeAddress(t.Raw.Address.String()), utils.NormalizeAddress(t.To.String()),utils.NormalizeAddress(t.From.String()), t.Value.String(), AmountIn, utils.NormalizeAddress(t.Raw.TxHash.String()), t.Raw.TxIndex, t.Raw.Index, s.processed.blockNumber, s.processed.blockTimestamp)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

