create or replace function public.erc20_user_with_balance(_address text)
    returns table
            (
                address text,
                balance numeric(78)
            )
    language plpgsql
as
$$
begin
    return query with transfers as ( select public.erc20_transfers.sender  as address,
                                            - public.erc20_transfers.value as amount
                                     from public.erc20_transfers
                                     where token_address = _address
                                     union all
                                     select public.erc20_transfers.receiver as address,
                                            public.erc20_transfers.value    as amount
                                     from public.erc20_transfers
                                     where token_address = _address )
                 select transfers.address, sum(transfers.amount) as balance
                 from transfers
                 group by transfers.address;

    return query select address, balance;
end
$$;

create or replace function public.erc20_users_with_balance_excluded_transfers(_address text, excluded_addresses text[])
    returns table
            (
                address text,
                balance numeric(78)
            )
    language plpgsql
as
$$
begin
    return query with transfers as ( select public.erc20_transfers.sender  as address,
                                            - public.erc20_transfers.value as amount
                                     from public.erc20_transfers
                                     where (public.erc20_transfers.sender not in (excluded_addresses))
                                       and (public.erc20_transfers.receiver not in (excluded_addresses))
                                       and token_address = _address
                                     union all
                                     select public.erc20_transfers.receiver as addrress,
                                            public.erc20_transfers.value    as amount
                                     from public.erc20_transfers
                                     where (public.erc20_transfers.sender not in (excluded_addresses))
                                       and (public.erc20_transfers.receiver not in (excluded_addresses))
                                       and token_address = _address )
                 select transfers.address, sum(transfers.amount) as balance
                 from transfers
                 group by transfers.address;
end
$$;

---- create above / drop below ----

drop function if exists public.erc20_user_with_balance(_address text);
drop function if exists public.erc20_users_with_balance_excluded_transfers(_address text, excluded_addresses text[]);
