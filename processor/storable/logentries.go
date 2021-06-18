package storable

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/smartbackend/types"
)

type LogEntriesStorable struct {
	Block *types.Block

	logger *logrus.Entry

	logs []types.LogEntry
}

func NewStorableLogEntries(block *types.Block) *LogEntriesStorable {
	return &LogEntriesStorable{
		Block:  block,
		logger: logrus.WithField("module", "storable(logentries)"),
	}
}

func (s *LogEntriesStorable) Result() interface{} {
	return s.logs
}

func (s *LogEntriesStorable) Rollback(tx *sql.Tx) error {
	_, err := tx.Exec(`delete from log_entries where included_in_block = $1`, s.Block.Number)
	return err
}

func (s *LogEntriesStorable) Execute() error {
	for _, tx := range s.Block.Txs {
		for _, log := range tx.LogEntries {
			s.logs = append(s.logs, log)
		}
	}

	return nil
}

func (s *LogEntriesStorable) SaveToDatabase(tx *sql.Tx) error {
	if len(s.Block.Txs) == 0 {
		return nil
	}

	s.logger.Trace("storing log entries")
	start := time.Now()
	count := 0
	defer func() {
		s.logger.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"count":    count,
		}).Debug("done storing log entries")
	}()

	stmt, err := tx.Prepare(pq.CopyIn("log_entries", "tx_hash", "log_index", "log_data", "logged_by", "topic_0", "topic_1", "topic_2", "topic_3", "included_in_block"))
	if err != nil {
		return err
	}

	for _, log := range s.logs {
		count++
		_, err = stmt.Exec(log.TxHash, log.LogIndex, log.LogData, log.LoggedBy, log.Topic0, log.Topic1, log.Topic2, log.Topic3, log.IncludedInBlock)
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
