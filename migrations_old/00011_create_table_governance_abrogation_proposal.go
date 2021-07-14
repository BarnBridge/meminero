package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableGovernanceAbrogationProposals, downCreateTableGovernanceAbrogationProposals)
}

func upCreateTableGovernanceAbrogationProposals(tx *sql.Tx) error {
	_, err := tx.Exec(`

	create table governance_abrogation_proposals
	(
		proposal_id					bigint not null ,
		creator						text not null,
		create_time					bigint not null,
		description					text not null,
		
		tx_hash                    text    not null,
		tx_index                   integer not null,
		log_index                  integer not null,
		logged_by                  text    not null,
		
		included_in_block           bigint  not null,
		created_at                  timestamp default now()
	);
	
	`)
	return err
}

func downCreateTableGovernanceAbrogationProposals(tx *sql.Tx) error {
	_, err := tx.Exec("drop table if exists governance_abrogation_proposals")
	return err
}
