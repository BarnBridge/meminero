package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableGovernanceVotes, downCreateTableGovernanceVotes)
}

func upCreateTableGovernanceVotes(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create table governance_votes
	(
		proposal_id				   bigint not null ,
		user_id					   text not null ,
		support 				   bool not null,
		power 					   bigint not null,
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

func downCreateTableGovernanceVotes(tx *sql.Tx) error {
	_, err := tx.Exec("drop table governance_votes")
	return err
}
