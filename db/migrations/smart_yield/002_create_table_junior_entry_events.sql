create table smart_yield.junior_entry_events
(
    pool_address      text    not null,

    buyer_address     text    not null,
    underlying_in     numeric(78),
    tokens_out        numeric(78),
    fee               numeric(78),

    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    created_at        timestamp default now()
);

---- create above / drop below ----

drop table if exists smart_yield.junior_entry_events;
