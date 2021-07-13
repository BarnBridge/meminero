create table smart_yield.provider_transfer_fees
(
    protocol_id       text    not null,
    provider_address  text    not null,
    caller_address    text    not null,
    fees_owner        text    not null,
    fees              numeric(78),
    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

---- create above / drop below ----

drop table if exists smart_yield.provider_transfer_fees;