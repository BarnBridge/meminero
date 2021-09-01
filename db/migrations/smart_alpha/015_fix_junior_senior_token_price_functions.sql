create or replace function smart_alpha.junior_token_to_usd_at_ts(token_address text, amount numeric(78), ts bigint) returns double precision
    language plpgsql as
$$
declare
    oracle_asset_symbol text;
    pool_token_decimals integer;
    pool_token_address  text;
    _pool_address       text;
begin
    select into oracle_asset_symbol,pool_token_decimals,pool_token_address, _pool_address p.oracle_asset_symbol,
                                                                                          p.pool_token_decimals,
                                                                                          p.pool_token_address,
                                                                                          p.pool_address
    from smart_alpha.pools p
    where p.junior_token_address = token_address;

    return ( select amount::numeric(78, 18) / pow(10, pool_token_decimals) *
                    ( select estimated_junior_token_price::numeric(78, 18) / pow(10, 18)
                      from smart_alpha.pool_state
                      where pool_address = _pool_address
                        and block_timestamp <= ts
                      order by block_timestamp desc
                      limit 1 ) * ( select token_usd_price_at_ts(pool_token_address, ts) ) );
end;
$$;

create or replace function smart_alpha.senior_token_to_usd_at_ts(token_address text, amount numeric(78), ts bigint) returns double precision
    language plpgsql as
$$
declare
    oracle_asset_symbol text;
    pool_token_decimals integer;
    pool_token_address  text;
    _pool_address       text;
begin
    select into oracle_asset_symbol,pool_token_decimals,pool_token_address, _pool_address p.oracle_asset_symbol,
                                                                                          p.pool_token_decimals,
                                                                                          p.pool_token_address,
                                                                                          p.pool_address
    from smart_alpha.pools p
    where p.senior_token_address = token_address;

    return ( select amount::numeric(78, 18) / pow(10, pool_token_decimals) *
                    ( select estimated_senior_token_price::numeric(78, 18) / pow(10, 18)
                      from smart_alpha.pool_state
                      where pool_address = _pool_address and block_timestamp <= ts
                      order by block_timestamp desc
                      limit 1 ) * ( select token_usd_price_at_ts(pool_token_address, ts) ) );
end;
$$;
