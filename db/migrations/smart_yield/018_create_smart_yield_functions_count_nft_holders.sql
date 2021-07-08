create or replace function smart_yield.number_of_nft_holders(addr text) returns bigint
    language plpgsql as
$$
declare
    value bigint;
begin
    select into value count(*)
    from (with transfers as (select sender as address, -1 as amount
                             from smart_yield.erc721_transfers
                             where token_address = addr
                             union all
                             select receiver as address, 1 as amount
                             from smart_yield.erc721_transfers
                             where token_address = addr)
          select address, sum(amount) as balance
          from transfers
          group by address) x
    where balance > 0;

    return value;
end;
$$;

create or replace function smart_yield.number_of_seniors(pool_address text) returns bigint
    language plpgsql as
$$
declare
    value bigint;
begin
    select into value smart_yield.number_of_nft_holders(
                              (select senior_bond_address
                               from smart_yield.pools
                               where sy_address = pool_address));

    return value;
end;
$$;

create or replace function smart_yield.number_of_juniors_locked(pool_address text) returns bigint
    language plpgsql as
$$
declare
    value bigint;
begin
    select into value smart_yield.number_of_nft_holders(
                              (select junior_bond_address
                               from smart_yield.pools
                               where sy_address = pool_address));

    return value;
end;
$$;

create or replace function smart_yield.number_of_jtoken_holders(addr text) returns bigint
    language plpgsql as
$$
declare
    holdersCount bigint;
begin
    select into holdersCount count(*)
    from (with transfers as (select sender as address, -value as amount
                             from public.erc20_transfers
                             where token_address = addr
                             union all
                             select receiver as address, value as amount
                             from public.erc20_transfers
                             where token_address = addr)
          select address, sum(amount) as balance
          from transfers
          group by address) x
    where balance > 0;

    return holdersCount;
end;
$$;

create or replace function smart_yield.junior_liquidity_locked(pool_address text) returns numeric(78)
    language plpgsql as
$$
declare
    value numeric(78);
begin
    select into value sum(case
                              when (select count(*)
                                    from smart_yield.junior_2step_redeem_events as r
                                    where r.junior_bond_address = b.junior_bond_address
                                      and r.junior_bond_id = b.junior_bond_id) = 0 then tokens_in
                              else 0 end)
    from smart_yield.junior_2step_withdraw_events as b
    where b.sy_address = pool_address;

    return value;
end
$$;

---- create above / drop below ----

drop function if exists smart_yield.number_of_nft_holders(addr text);
drop function if exists smart_yield.number_of_seniors(pool_address text);
drop function if exists smart_yield.number_of_juniors_locked(pool_address text);
drop function if exists smart_yield.number_of_jtoken_holders(addr text);
drop function if exists smart_yield.junior_liquidity_locked(pool_address text);
