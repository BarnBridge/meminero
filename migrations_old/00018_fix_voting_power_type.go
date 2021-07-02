package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upFixVotingPowerType, downFixVotingPowerType)
}

func upFixVotingPowerType(tx *sql.Tx) error {
	_, err := tx.Exec(`
	alter table governance_votes alter column power type numeric(78);
	alter table governance_abrogation_votes alter column power type numeric(78);
	`)

	return err
}

func downFixVotingPowerType(tx *sql.Tx) error {
	_, err := tx.Exec(`
	alter table governance_votes alter column power type bigint;
	alter table governance_abrogation_votes alter column power type bigint;
	`)

	return err
}
