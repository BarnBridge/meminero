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
    return query WITH transfers AS (
        SELECT public.erc20_transfers.sender  AS address,
               - public.erc20_transfers.value AS amount
        FROM public.erc20_transfers
        where token_address = _address
        UNION ALL
        SELECT public.erc20_transfers.receiver AS address,
               public.erc20_transfers.value    AS amount
        FROM public.erc20_transfers
        where token_address = _address
    )
                 SELECT transfers.address,
                        sum(transfers.amount) AS balance
                 FROM transfers
                 GROUP BY transfers.address;

    return query
        select address, balance;
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
    return query WITH transfers AS (
        SELECT public.erc20_transfers.sender  AS address,
               - public.erc20_transfers.value AS amount
        FROM public.erc20_transfers
        WHERE (public.erc20_transfers.sender not in (excluded_addresses))
          AND (public.erc20_transfers.receiver not in (excluded_addresses))
          AND token_address = _address
        UNION ALL
        SELECT public.erc20_transfers.receiver AS addrress,
               public.erc20_transfers.value    AS amount
        FROM public.erc20_transfers
        WHERE (public.erc20_transfers.sender not in (excluded_addresses))
          AND (public.erc20_transfers.receiver not in (excluded_addresses))
          AND token_address = _address
    )
                 SELECT transfers.address,
                        sum(transfers.amount) AS balance
                 FROM transfers
                 GROUP BY transfers.address;
end
$$;

---- create above / drop below ----

drop function if exists public.erc20_user_with_balance(_address text)
drop function if exists public.erc20_users_with_balance_excluded_transfers(_address text, excluded_addresses text[])