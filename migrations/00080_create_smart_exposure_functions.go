package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateSmartExposureFunctions, downCreateSmartExposureFunctions)
}

func upCreateSmartExposureFunctions(tx *sql.Tx) error {
	_, err := tx.Exec(`
create or replace function get_token_price_chart(_token_address text, start bigint, _date_trunc text)
    returns table
            (
                point       timestamp,
                token_price double precision
            )
    language plpgsql
as
$$
begin
    return query
        select date_trunc(_date_trunc, to_timestamp(block_timestamp)::date)::timestamp as point,
               avg(price_usd)                          as token_price
        from tokens_prices
        where token_address = _token_address
          and block_timestamp > start
        group by point
        order by point;
end
$$;

create or replace function get_ratio_deviation(_etoken_address text, start bigint, _date_trunc text)
    returns table
            (
                point       timestamp,
                deviation numeric(78,18)
            )
    language plpgsql
as
$$
    declare 
        target_ratio numeric(78,18);
begin
        select into target_ratio t.target_ratio::numeric(78,18) from smart_exposure_tranches t where t.etoken_address = _etoken_address ;
    return query
        select date_trunc(_date_trunc, block_timestamp)::timestamp as point,
                    avg(1 - ((ts.current_ratio::numeric(78,18))/target_ratio)) as deviation
        from smart_exposure_tranche_state ts
			 where ts.e_token_address = _etoken_address
          and extract(epoch from block_timestamp)::bigint > start
        group by point
        order by point;
end
$$;
`)
	return err
}

func downCreateSmartExposureFunctions(tx *sql.Tx) error {
	_, err := tx.Exec(`
			drop function if exists get_token_price_chart;
			drop function if exists get_ratio_deviation
`)
	return err
}
