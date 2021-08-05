create or replace function public.erc20_balances_at_ts(user_address text, tokens text[], ts bigint)
    returns TABLE
            (
                token_address text,
                balance       numeric
            )
    language plpgsql
as
$$
begin
    return query with transfers as (select t.token_address as address, -value as amount
                                    from public.erc20_transfers t
                                    where sender = user_address
                                      and t.token_address = any (tokens)
                                      and block_timestamp <= ts
                                    union all
                                    select t.token_address as address, value as amount
                                    from public.erc20_transfers t
                                    where receiver = user_address
                                      and t.token_address = any (tokens)
                                      and block_timestamp <= ts)
                 select address, sum(amount) as balance
                 from transfers
                 group by address;
end;
$$;