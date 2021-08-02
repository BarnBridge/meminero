create table if not exists smart_alpha.epoch_end_events
(
    pool_address      text        not null,

    epoch_id          bigint      not null,
    junior_profits    numeric(78) not null,
    senior_profits    numeric(78) not null,

    block_timestamp   bigint      not null,
    included_in_block bigint      not null,
    tx_hash           text        not null,
    tx_index          integer     not null,
    log_index         integer     not null,
    created_at        timestamp default now()
)
