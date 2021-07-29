create table if not exists smart_alpha.user_redeem_tokens_events
(
    pool_address      text                not null,
    tranche           smart_alpha.tranche not null,

    user_address      text                not null,
    epoch_id          bigint              not null,
    tokens_out        numeric(78)         not null,

    block_timestamp   bigint              not null,
    included_in_block bigint              not null,
    tx_hash           text                not null,
    tx_index          integer             not null,
    log_index         integer             not null,
    created_at        timestamp default now()
)
