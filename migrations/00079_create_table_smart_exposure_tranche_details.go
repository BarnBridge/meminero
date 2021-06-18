package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableSmartExposureTrancheDetails, downCreateTableSmartExposureTrancheDetails)
}

func upCreateTableSmartExposureTrancheDetails(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create or replace function get_tranche_details(_e_token_address text)
    returns table
            (
                s_factor_e                          numeric(78),
                target_ratio                        numeric(78),
                token_a_ratio                       double precision,
                token_a_address                     text,
                token_a_symbol                      text,
                token_a_decimals                    bigint,
                token_a_price_usd                   double precision,
                token_a_included_in_block           bigint,
                token_a_block_timestamp             bigint,
                token_b_address                     text,
                token_b_price_usd                   double precision,
                token_b_included_in_block           bigint,
                token_b_block_timestamp             bigint,
                token_b_ratio                       double precision,
                token_b_symbol                      text,
                token_b_decimals                    bigint,
                pool_state_rebalancing_interval     bigint,
                pool_state_rebalancing_condition    numeric(78),
                pool_state_last_rebalance           bigint,
                tranche_state_token_a_liquidity     double precision,
                tranche_state_token_b_liquidity     double precision,
                tranche_state_etoken_price          double precision,
                tranche_state_current_ratio         double precision,
                tranche_state_token_a_current_ratio double precision,
                tranche_state_token_b_current_ratio double precision,
                tranche_state_included_in_block     bigint,
                tranche_state_block_timestamp       timestamp
            )
    language plpgsql
as
$$
declare
    _pool_address                       text;
    s_factor_e                          numeric(78);
    target_ratio                        numeric(78);
    token_a_ratio                       double precision;
    token_b_ratio                       double precision;
    token_a_address                     text;
    token_a_symbol                      text;
    token_a_decimals                    bigint;
    token_a_price_usd                   double precision;
    token_a_included_in_block           bigint;
    token_a_block_timestamp             bigint;
    token_b_address                     text;
    token_b_symbol                      text;
    token_b_decimals                    bigint;
    token_b_price_usd                   double precision;
    token_b_included_in_block           bigint;
    token_b_block_timestamp             bigint;
    pool_state_rebalancing_interval     bigint;
    pool_state_rebalancing_condition    numeric(78);
    pool_state_last_rebalance           bigint;
    tranche_state_token_a_liquidity     double precision;
    tranche_state_token_b_liquidity     double precision;
    tranche_state_etoken_price          double precision;
    tranche_state_current_ratio         double precision;
    tranche_state_token_a_current_ratio double precision;
    tranche_state_token_b_current_ratio double precision;
    tranche_state_included_in_block     bigint;
    tranche_state_block_timestamp       timestamp;
begin
    select into _pool_address,s_factor_e,target_ratio,token_a_ratio,token_b_ratio t.pool_address,
                                                                                  t.s_factor_e,
                                                                                  t.target_ratio,
                                                                                  t.token_a_ratio,
                                                                                  t.token_b_ratio
    from smart_exposure_tranches t
    where t.etoken_address = _e_token_address;

    select into token_a_address,token_a_symbol,token_a_decimals,token_b_address,token_b_symbol,token_b_decimals p.token_a_address,
                                                                                                                p.token_a_symbol,
                                                                                                                p.token_a_decimals,
                                                                                                                p.token_b_address,
                                                                                                                p.token_b_symbol,
                                                                                                                p.token_b_decimals
    from smart_exposure_pools p
    where p.pool_address = _pool_address;

    select into token_a_price_usd,token_a_included_in_block,token_a_block_timestamp price.price_usd,
                                                                                    price.included_in_block,
                                                                                    price.block_timestamp
    from tokens_prices price
    where price.token_address = token_a_address
    order by block_timestamp desc
    limit 1;

    select into token_b_price_usd,token_b_included_in_block,token_b_block_timestamp price.price_usd,
                                                                                    price.included_in_block,
                                                                                    price.block_timestamp
    from tokens_prices price
    where price.token_address = token_b_address
    order by block_timestamp desc
    limit 1;


    select into pool_state_rebalancing_interval,pool_state_rebalancing_condition,pool_state_last_rebalance ps.rebalancing_interval,
                                                                                                           ps.rebalancing_condition,
                                                                                                           ps.last_rebalance
    from smart_exposure_pool_state ps
    where ps.pool_address = _pool_address
    order by block_timestamp desc
    limit 1;

    select into tranche_state_token_a_liquidity, tranche_state_token_b_liquidity,tranche_state_etoken_price,tranche_state_current_ratio,
        tranche_state_token_a_current_ratio,tranche_state_token_b_current_ratio,tranche_state_included_in_block,tranche_state_block_timestamp ts.token_a_liquidity,
                                                                                                                                              ts.token_b_liquidity,
                                                                                                                                              ts.etoken_price,
                                                                                                                                              ts.current_ratio,
                                                                                                                                              ts.token_a_current_ratio,
                                                                                                                                              ts.token_b_current_ratio,
                                                                                                                                              ts.included_in_block,
                                                                                                                                              ts.block_timestamp
    from smart_exposure_tranche_state ts
    where ts.e_token_address = _e_token_address
    order by block_timestamp desc
    limit 1;
    return query
        select s_factor_e,
               target_ratio,
               token_a_ratio,
               token_a_address,
               token_a_symbol,
               token_a_decimals,
               token_a_price_usd,
               token_a_included_in_block,
               token_a_block_timestamp,
               token_b_address,
               token_b_price_usd,
               token_b_included_in_block,
               token_b_block_timestamp,
               token_b_ratio,
               token_b_symbol,
               token_b_decimals,
               pool_state_rebalancing_interval,
               pool_state_rebalancing_condition,
               pool_state_last_rebalance,
               tranche_state_token_a_liquidity,
               tranche_state_token_b_liquidity,
               tranche_state_etoken_price,
               tranche_state_current_ratio,
               tranche_state_token_a_current_ratio,
               tranche_state_token_b_current_ratio,
               tranche_state_included_in_block,
               tranche_state_block_timestamp;
end
$$;
`)
	return err
}

func downCreateTableSmartExposureTrancheDetails(tx *sql.Tx) error {
	_, err := tx.Exec(`drop function if exists  get_tranche_details;`)
	return err
}
