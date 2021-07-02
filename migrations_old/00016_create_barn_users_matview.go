package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpBarnUsersMatView, DownBarnUsersMatView)
}

func UpBarnUsersMatView(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create materialized view barn_users as
	select distinct user_address
	from barn_staking_actions
	union
	select distinct receiver
	from barn_delegate_changes;

	create unique index on barn_users(user_address);
	
	create or replace function refresh_barn_users() returns TRIGGER
		language plpgsql as
	$$
	begin
		refresh materialized view concurrently barn_users;
		return null;
	end
	$$;
	
	create trigger refresh_barn_users
		after insert or update or delete or truncate
		on barn_staking_actions
		for each statement
	execute procedure refresh_barn_users();
	
	create trigger refresh_barn_users
		after insert or update or delete or truncate
		on barn_delegate_changes
		for each statement
	execute procedure refresh_barn_users();
	`)
	return err
}

func DownBarnUsersMatView(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop trigger refresh_barn_users on barn_staking_actions;
		drop trigger refresh_barn_users on barn_delegate_changes;
		drop function refresh_barn_users;
		drop materialized view barn_users;
	`)
	return err
}
