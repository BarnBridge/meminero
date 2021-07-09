package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUpdateSyTxType, downUpdateSyTxType)
}

func upUpdateSyTxType(tx *sql.Tx) error {
	_, err := tx.Exec(`
			alter type sy_tx_history_tx_type add value if not exists 'JUNIOR_UNSTAKE';
			alter type sy_tx_history_tx_type add value if not exists 'JUNIOR_STAKE';
			`)
	return err
}

func downUpdateSyTxType(tx *sql.Tx) error {
	return nil
}
