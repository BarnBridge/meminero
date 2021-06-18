package storable

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/types"
)

type TxStorable struct {
	Block  *types.Block
	logger *logrus.Entry
}

func NewStorableTxs(block *types.Block) *TxStorable {
	return &TxStorable{
		Block:  block,
		logger: logrus.WithField("module", "storable(txs)"),
	}
}

func (s *TxStorable) Result() interface{} {
	return s.Block.Txs
}

func (s *TxStorable) Rollback(tx *sql.Tx) error {
	_, err := tx.Exec(`delete from txs where included_in_block = $1`, s.Block.Number)
	return err
}

func (s *TxStorable) Execute() error {
	return nil
}

func (s *TxStorable) SaveToDatabase(dbTx *sql.Tx) error {
	s.logger.Trace("storing txs")
	start := time.Now()
	defer func() {
		s.logger.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"count":    len(s.Block.Txs),
		}).Debug("done storing txs")
	}()

	stmt, err := dbTx.Prepare(pq.CopyIn("txs", "tx_hash", "included_in_block", "tx_index", "from", "to", "value", "tx_nonce", "msg_gas_limit", "tx_gas_used", "tx_gas_price", "cumulative_gas_used", "msg_payload", "msg_status", "creates", "tx_logs_bloom", "block_creation_time", "log_entries_triggered"))
	if err != nil {
		return err
	}

	for _, tx := range s.Block.Txs {
		_, err = stmt.Exec(tx.TxHash, tx.IncludedInBlock, tx.TxIndex, tx.From, tx.To, tx.Value, tx.TxNonce, tx.MsgGasLimit, tx.TxGasUsed, tx.TxGasPrice, tx.CumulativeGasUsed, tx.MsgPayload, tx.MsgStatus, tx.Creates, tx.TxLogsBloom, tx.BlockCreationTime, len(tx.LogEntries))
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
