package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionNumberOfJuniors, downCreateFunctionNumberOfJuniors)
}

func upCreateFunctionNumberOfJuniors(tx *sql.Tx) error {
	_, err := tx.Exec(`
create or replace function number_of_active_juniors(addr text) returns bigint
    language plpgsql as
$$
declare
    holdersCount bigint;
begin
    select into holdersCount count(*)
    from ( with transfers as ( select sender as address, -value as amount
                               from jtoken_transfers
                               where sy_address = addr
                                 and receiver not in
                                     ( select pool_address from smart_yield_reward_pools where sy_address = addr )
                               union all
                               select receiver as address, value as amount
                               from jtoken_transfers
                               where sy_address = addr
                                 and sender not in
                                     ( select pool_address from smart_yield_reward_pools where sy_address = addr ) )
           select address, sum(amount) as balance
           from transfers
           group by address ) x
    where balance > 0;

    return holdersCount;
end;
$$;
`)

	return err
}

func downCreateFunctionNumberOfJuniors(tx *sql.Tx) error {
	_, err := tx.Exec(`
drop function if exists number_of_active_juniors;
`)

	return err
}
