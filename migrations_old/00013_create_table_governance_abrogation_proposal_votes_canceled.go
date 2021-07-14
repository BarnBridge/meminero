package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableGovernanceAbrogationVotesCanceled, downCreateTableGovernanceAbrogationVotesCanceled)
}

func upCreateTableGovernanceAbrogationVotesCanceled(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table governance_abrogation_votes_canceled
	(
		proposal_id				   bigint not null ,
		user_id					   text not null ,
		block_timestamp				   bigint,
		
		tx_hash                    text    not null,
		tx_index                   integer not null,
		log_index                  integer not null,
		logged_by                  text    not null,
		
		included_in_block          bigint  not null,
		created_at                 timestamp default now()
	);
	`)
	return err
}

func downCreateTableGovernanceAbrogationVotesCanceled(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists governance_abrogation_votes_canceled")
	return err
}
