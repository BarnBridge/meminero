create table public.erc20_transfers
(
    token_address     text    not null,
    sender            text    not null,
    receiver          text    not null,
    value             numeric(78),
    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

create index erc20_transfers_sender_idx
    on public.erc20_transfers (sender asc, block_timestamp desc);

create index erc20_transfers_receiver_idx
    on public.erc20_transfers (receiver asc, block_timestamp desc);


---- create above / drop below ----

drop table if exists public.erc20_transfers;
drop index if exists public.erc20_transfers_sender_idx;
drop index if exists public.erc20_transfers_receiver_idx;
