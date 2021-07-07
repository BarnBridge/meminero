create or replace function smart_yield.current_owner_of_token_at_ts(address text, id bigint, ts bigint) returns text
    language plpgsql as
$$
declare
    receiver text;
begin
    select into receiver t.receiver
    from smart_yield.erc721_transfers t
    where token_address = address
      and token_id = id
      and block_timestamp <= ts
    order by included_in_block desc, tx_index desc, log_index desc
    limit 1;

    return receiver;
end;
$$;

create or replace function smart_yield.senior_bond_value_at_ts(token_address text, token_id bigint, ts bigint) returns numeric(78)
    language plpgsql as
$$
declare
    value numeric(78);
begin
    select into value case
                          when block_timestamp + for_days * 24 * 60 * 60 <= ts
                              then underlying_in + gain
                          else underlying_in end
    from smart_yield.smart_yield_senior_buy
    where senior_bond_address = token_address
      and senior_bond_id = token_id;

    return value;
end;
$$;

create or replace function smart_yield.senior_bond_redeemed_at_ts(token_address text, token_id bigint, ts bigint) returns boolean
    language plpgsql as
$$
declare
    redeemed boolean;
begin
    select into redeemed count(*) > 0
    from smart_yield.smart_yield_senior_redeem
    where senior_bond_address = token_address
      and senior_bond_id = token_id
      and block_timestamp <= ts;

    return redeemed;
end;
$$;

create or replace function smart_yield.senior_underlying_price_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    price double precision;
begin
    select into price price_usd
    from public.token_prices p
    where p.token_address =
          (select underlying_address from smart_yield.smart_yield_pools where senior_bond_address = addr)
      and block_timestamp <= ts
    order by block_timestamp desc
    limit 1;

    return price;
end;
$$;

create or replace function smart_yield.senior_portfolio_at_ts(addr text, ts bigint)
    returns table
            (
                token_address text,
                token_id      bigint
            )
    language plpgsql
as
$$
begin
    return query select distinct t.token_address, t.token_id
                 from smart_yield.erc721_transfers t
                 where token_type = 'senior'
                   and receiver = addr
                   and block_timestamp <= ts
                   and smart_yield.current_owner_of_token_at_ts(t.token_address, t.token_id, ts) = receiver
                   and not smart_yield.senior_bond_redeemed_at_ts(t.token_address, t.token_id, ts);
end;
$$;

create or replace function smart_yield.senior_portfolio_value_at_ts(addr text, ts bigint) returns double precision
    language plpgsql as
$$
declare
    value double precision;
begin
    select into value coalesce(
                              sum(
                                              smart_yield.senior_bond_value_at_ts(token_address, token_id, ts)::numeric(78, 18) /
                                              pow(10, (select underlying_decimals
                                                       from smart_yield.smart_yield_pools
                                                       where senior_bond_address = token_address)) *
                                              smart_yield.senior_underlying_price_at_ts(token_address, ts)
                                  ),
                              0
                          ) as senior_portfolio_value
    from smart_yield.senior_portfolio_at_ts(addr, ts);

    return value;
end;
$$;

---- create above / drop below ----

drop function if exists smart_yield.current_owner_of_token_at_ts(address text, id bigint, ts bigint);
drop function if exists smart_yield.senior_bond_value_at_ts(token_address text, token_id bigint, ts bigint);
drop function if exists smart_yield.senior_bond_redeemed_at_ts(token_address text, token_id bigint, ts bigint);
drop function if exists smart_yield.senior_underlying_price_at_ts(addr text, ts bigint);
drop function if exists smart_yield.senior_portfolio_at_ts(addr text, ts bigint);
drop function if exists smart_yield.senior_bond_value_at_ts(token_address text, token_id bigint, ts bigint);