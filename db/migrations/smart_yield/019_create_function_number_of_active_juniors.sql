create or replace function smart_yield.number_of_active_juniors(_pool_address text) returns bigint
    language plpgsql as
$$
declare
    holdersCount bigint;
begin
    select into holdersCount count(*)
    from ( with transfers as ( select sender as address, -value as amount
                               from public.erc20_transfers
                               where token_address = _pool_address
                                 and receiver not in ( select pool_address
                                                       from smart_yield.reward_pools
                                                       where pool_token_address = _pool_address )
                               union all
                               select receiver as address, value as amount
                               from public.erc20_transfers
                               where token_address = _pool_address
                                 and sender not in ( select pool_address
                                                     from smart_yield.reward_pools
                                                     where pool_token_address = _pool_address ) )
           select address, sum(amount) as balance
           from transfers
           group by address ) x
    where balance > 0;

    return holdersCount;
end;
$$;
