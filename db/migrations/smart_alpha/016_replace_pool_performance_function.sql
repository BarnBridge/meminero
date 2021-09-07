create or replace function smart_alpha.performance_at_ts(pool text, ts bigint)
    returns table
            (
                senior_without_sa double precision,
                senior_with_sa    double precision,
                junior_without_sa double precision,
                junior_with_sa    double precision
            )
    language plpgsql
as
$$
declare
    token_price                 double precision;
    token_address               text;
    token_decimals              integer;
    quote_asset_symbol          text;
    jtoken_price_start          numeric(78, 18);
    stoken_price_start          numeric(78, 18);
    jtoken_price_estimate       numeric(78, 18);
    stoken_price_estimate       numeric(78, 18);
begin
    select into token_address, quote_asset_symbol, token_decimals p.pool_token_address,
                                                                  p.oracle_asset_symbol,
                                                                  p.pool_token_decimals
    from smart_alpha.pools p
    where pool_address = pool;

    select into token_price public.token_price_at_ts(token_address, quote_asset_symbol, ts);

    select into jtoken_price_start, stoken_price_start junior_token_price_start, senior_token_price_start
    from smart_alpha.pool_epoch_info
    where pool_address = pool
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    select into jtoken_price_estimate, stoken_price_estimate estimated_junior_token_price, estimated_senior_token_price
    from smart_alpha.pool_state
    where pool_address = pool
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return query select 1::numeric(78, 18) * token_price                                              as senior_without_sa,
                        1::numeric(78, 18) / stoken_price_start * stoken_price_estimate * token_price as senior_with_sa,
                        1::numeric(78, 18) * token_price                                              as junior_without_sa,
                        1::numeric(78, 18) / jtoken_price_start * jtoken_price_estimate * token_price as junior_with_sa;
end
$$;

