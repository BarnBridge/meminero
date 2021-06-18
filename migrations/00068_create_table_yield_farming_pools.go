package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableYieldFarmingPools, downCreateTableYieldFarmingPools)
}

func upCreateTableYieldFarmingPools(tx *sql.Tx) error {
	_, err := tx.Exec(`
create or replace function yf_sum_actions_by_token(a_type staking_action_type, _token_address text,
                                                   _start_time bigint,_end_time bigint,scale text)
    returns table
            (
                amount numeric(78),
                point   timestamp with time zone
            )
    language plpgsql
as
$$
begin
    return query select sum(yf.amount),
                        date_trunc(scale, to_timestamp(block_timestamp)) as wk
                 from yield_farming_actions yf
                 where action_type = a_type
                   and token_address = _token_address
    				and block_timestamp between _start_time and _end_time
                 group by wk
                 order by wk;
end;
$$;

create or replace function yf_stats_by_token(_token_address text,
                                               _start_time bigint,_end_time bigint,scale text)
    returns table
            (
                point            timestamp with time zone,
                sum_deposits   numeric(78),
                sum_withdrawals numeric(78)
            )
    language plpgsql
as
$$
begin
    return query select coalesce(d.point, w.point), coalesce(d.amount, 0) as sum_deposits, coalesce(w.amount, 0) as sum_withdrawals
                 from yf_sum_actions_by_token('DEPOSIT', _token_address, _start_time,_end_time,scale ) d
                          full outer join yf_sum_actions_by_token('WITHDRAW', _token_address,  _start_time,_end_time,scale ) w
                                          on d.point = w.point;
end
$$;

	`)

	return err
}

func downCreateTableYieldFarmingPools(tx *sql.Tx) error {
	_, err := tx.Exec(`
drop function if exists yf_sum_actions_by_token;
drop function if exists yf_stats_by_token;
	`)

	return err
}
