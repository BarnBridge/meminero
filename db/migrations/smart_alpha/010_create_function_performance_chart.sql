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
    epoch_senior_liquidity      numeric(78, 18);
    _estimated_senior_liquidity numeric(78, 18);
    epoch_junior_liquidity      numeric(78, 18);
    _estimated_junior_liquidity numeric(78, 18);
    token_price                 double precision;
    token_address               text;
    token_decimals              integer;
    quote_asset_symbol          text;
begin
    select into token_address, quote_asset_symbol, token_decimals p.pool_token_address,
                                                                  p.oracle_asset_symbol,
                                                                  p.pool_token_decimals
    from smart_alpha.pools p
    where pool_address = pool;

    select into token_price public.token_price_at_ts(token_address, quote_asset_symbol, ts);

    select into epoch_junior_liquidity, epoch_senior_liquidity junior_liquidity, senior_liquidity
    from smart_alpha.pool_epoch_info
    where pool_address = pool
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    select into _estimated_junior_liquidity, _estimated_senior_liquidity estimated_junior_liquidity,
                                                                         estimated_senior_liquidity
    from smart_alpha.pool_state
    where pool_address = pool
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return query select epoch_senior_liquidity / pow(10, token_decimals) * token_price      as senior_without_sa,
                        _estimated_senior_liquidity / pow(10, token_decimals) * token_price as senior_with_sa,
                        epoch_junior_liquidity / pow(10, token_decimals) * token_price      as junior_without_sa,
                        _estimated_junior_liquidity / pow(10, token_decimals) * token_price as junior_with_sa;
end
$$;

