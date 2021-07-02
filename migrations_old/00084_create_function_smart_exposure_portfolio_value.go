package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateFunctionSmartExposurePortfolioValue, downCreateFunctionSmartExposurePortfolioValue)
}

func upCreateFunctionSmartExposurePortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
create or replace function etoken_active_position_at_ts(user_address text, ts bigint)
    returns TABLE
            (
                etoken_address text,
                balance        numeric(78,18)
            )
    language plpgsql
as
$$
begin
    return query with transfers as (select e_token_address as address, -value as amount
                                    from etoken_transfers
                                    where sender = user_address
                                      and block_timestamp <= ts
                                    union all
                                    select e_token_address as address, value as amount
                                    from etoken_transfers
                                    where receiver = user_address
                                      and block_timestamp <= ts)
                 select address, sum(amount) as balance
                 from transfers
                 group by address;
end;
$$;

create or replace function se_user_portfolio_value(addr text, ts bigint) returns double precision
    language plpgsql
as
$$
declare
    value double precision;
begin
    select into value sum(coalesce(a.balance / pow(10, 18) * (select etoken_price
                                                              from smart_exposure_tranche_state s
                                                              where s.e_token_address = a.etoken_address
                                                                and s.block_timestamp <= to_timestamp(ts)
                                                              limit 1),0))
    from etoken_active_position_at_ts(addr, ts) a;
    
    return value;
end;
$$;

create or replace function se_user_portfolio_value_by_pool(addr text, ts bigint,_pool_address text) returns double precision
    language plpgsql
as
$$
declare
    value double precision;
begin
    select into value sum(coalesce(a.balance / pow(10, 18) * (select etoken_price
                                                              from smart_exposure_tranche_state s
                                                              where s.e_token_address = a.etoken_address
                                                                and s.block_timestamp <= to_timestamp(ts)
                                                              limit 1),0))
    from etoken_active_position_at_ts(addr, ts) a
    where a.etoken_address in (select etoken_address from smart_exposure_tranches where pool_address = _pool_address);
    return value;
end;
$$;

	`)
	return err
}

func downCreateFunctionSmartExposurePortfolioValue(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop function if exists etoken_active_position_at_ts;
		drop function if exists se_user_portfolio_value;
		drop function if exists se_user_portfolio_value_by_pool;
	`)
	return err
}
