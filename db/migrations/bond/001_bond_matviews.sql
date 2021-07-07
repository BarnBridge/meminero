create view bond.bond_users_with_balance(address, balance) as
WITH transfers AS (
    SELECT public.erc20_transfers.sender  AS address,
           - public.erc20_transfers.value AS amount
    FROM public.erc20_transfers
    where token_address = '0x0391d2021f89dc339f60fff84546ea23e337750f'
    UNION ALL
    SELECT public.erc20_transfers.receiver AS address,
           public.erc20_transfers.value    AS amount
    FROM public.erc20_transfers
    where token_address = '0x0391d2021f89dc339f60fff84546ea23e337750f'
)
SELECT transfers.address,
       sum(transfers.amount) AS balance
FROM transfers
GROUP BY transfers.address;

create view bond.bond_users_with_balance_no_staking(address, balance) as
WITH transfers AS (
    SELECT public.erc20_transfers.sender  AS address,
           - public.erc20_transfers.value AS amount
    FROM public.erc20_transfers
    WHERE (public.erc20_transfers.sender <> ALL
           (ARRAY ['0xb0fa2beee3cf36a7ac7e99b885b48538ab364853'::text, '0x10e138877df69ca44fdc68655f86c88cde142d7f'::text]))
      AND (public.erc20_transfers.receiver <> ALL
           (ARRAY ['0xb0fa2beee3cf36a7ac7e99b885b48538ab364853'::text, '0x10e138877df69ca44fdc68655f86c88cde142d7f'::text]))
      AND token_address = '0x0391d2021f89dc339f60fff84546ea23e337750f'
    UNION ALL
    SELECT public.erc20_transfers.receiver AS addrress,
           public.erc20_transfers.value    AS amount
    FROM public.erc20_transfers
    WHERE (public.erc20_transfers.sender <> ALL
           (ARRAY ['0xb0fa2beee3cf36a7ac7e99b885b48538ab364853'::text, '0x10e138877df69ca44fdc68655f86c88cde142d7f'::text]))
      AND (public.erc20_transfers.receiver <> ALL
           (ARRAY ['0xb0fa2beee3cf36a7ac7e99b885b48538ab364853'::text, '0x10e138877df69ca44fdc68655f86c88cde142d7f'::text]))
      AND token_address = '0x0391d2021f89dc339f60fff84546ea23e337750f'
)
SELECT transfers.address,
       sum(transfers.amount) AS balance
FROM transfers
GROUP BY transfers.address;

---- create above / drop below ----

drop view if exists bond.bond_users_with_balance;
drop view if exists bond.bond_users_with_balance_no_staking;