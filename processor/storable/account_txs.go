package storable

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/types"
)

type AccountTxsStorable struct {
	Block *types.Block

	logger *logrus.Entry

	accountTxs []*AccountTx
}

// AccountTx is a type of entity that represents a transaction between two accounts
// For each transaction, 2 AccountTx entities will be created, one for each direction of the transaction
// This helps with querying an account's transactions history in a specific order (e.g. chronological) and paginated
type AccountTx struct {
	Address         string
	Counterparty    string
	TxHash          string
	Out             bool
	IncludedInBlock int64
	TxIndex         int64
}

func NewStorableAccountTxs(block *types.Block) *AccountTxsStorable {
	return &AccountTxsStorable{
		Block:  block,
		logger: logrus.WithField("module", "storable(account-txs)"),
	}
}

func (s *AccountTxsStorable) Result() interface{} {
	return s.accountTxs
}

func (s *AccountTxsStorable) Rollback(tx *sql.Tx) error {
	_, err := tx.Exec(`delete from account_txs where included_in_block = $1`, s.Block.Number)
	return err
}

func (s *AccountTxsStorable) Execute() error {
	if len(s.Block.Txs) == 0 {
		return nil
	}

	s.logger.Trace("storing account transactions")
	start := time.Now()
	defer func() {
		s.logger.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"count":    len(s.accountTxs),
		}).Debug("done storing account transactions")
	}()

	for _, tx := range s.Block.Txs {
		storableAccountTxOut, err := s.buildStorableAccountTx(tx, true)
		if err != nil {
			return err
		}
		s.accountTxs = append(s.accountTxs, storableAccountTxOut)

		storableAccountTxIn, err := s.buildStorableAccountTx(tx, false)
		if err != nil {
			return err
		}
		s.accountTxs = append(s.accountTxs, storableAccountTxIn)
	}

	return nil
}

func (s *AccountTxsStorable) SaveToDatabase(tx *sql.Tx) error {
	stmt, err := tx.Prepare(pq.CopyIn("account_txs", "address", "counterparty", "tx_hash", "out", "included_in_block", "tx_index"))
	if err != nil {
		return err
	}

	for _, at := range s.accountTxs {
		_, err = stmt.Exec(at.Address, at.Counterparty, at.TxHash, at.Out, at.IncludedInBlock, at.TxIndex)
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

func (s *AccountTxsStorable) buildStorableAccountTx(tx types.Tx, out bool) (*AccountTx, error) {
	at := &AccountTx{}
	at.IncludedInBlock = tx.IncludedInBlock
	at.TxIndex = tx.TxIndex
	at.TxHash = tx.TxHash
	at.Out = out

	if out {
		at.Address = string(tx.From)
		at.Counterparty = string(tx.To)
	} else {
		at.Address = string(tx.To)
		at.Counterparty = string(tx.From)
	}

	return at, nil
}
