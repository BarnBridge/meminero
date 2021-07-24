create table public.account_erc20_transfers
(
    token_address     text    not null,
    account           text    not null,
    counterparty      text    not null,
    amount            numeric(78),
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    tx_direction      transfer_type
);

create index account_erc20_transfers_account_addr_idx on public.account_erc20_transfers (account asc, included_in_block desc, tx_index desc, log_index desc);
