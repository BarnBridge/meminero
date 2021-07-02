package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableVotersView, downCreateTableVotersView)
}

func upCreateTableVotersView(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace view voters as
		select user_address,balance_of(user_address) as bond_staked,
		   coalesce(( select locked_until
			 from barn_locks
			 where user_address = barn_users.user_address
			 order by included_in_block desc, log_index desc limit 1 ) ,0)                                          as locked_until,
		   delegated_power(user_address),
		   ( select count(*) from governance_votes where lower(user_id) = lower(barn_users.user_address) ) +
		   ( select count(*) from governance_abrogation_votes where lower(user_id) = lower(barn_users.user_address) ) as votes,
		   ( select count(*) from governance_proposals where lower(proposer) = lower(barn_users.user_address) )         as proposals,
		   voting_power(user_address)                                                                     as voting_power ,
		   has_active_delegation(user_address)
	from barn_users;
	    `)
	return err
}

func downCreateTableVotersView(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop view voters;
	`)
	return err
}
