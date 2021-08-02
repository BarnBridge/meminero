create type smart_alpha_transaction_type as enum ( 'JUNIOR_ENTRY', 'JUNIOR_REDEEM_TOKENS', 'JUNIOR_EXIT', 'JUNIOR_REDEEM_UNDERLYING', 'SENIOR_ENTRY', 'SENIOR_REDEEM_TOKENS', 'SENIOR_EXIT', 'SENIOR_REDEEM_UNDERLYING', 'JTOKEN_SEND', 'JTOKEN_RECEIVE', 'STOKEN_SEND', 'STOKEN_RECEIVE');

create table if not exists smart_alpha.transaction_history
(
    pool_address      text                         not null,
    tranche           text                         not null,
    transaction_type  smart_alpha_transaction_type not null,

    user_address      text                         not null,
    amount            numeric(78)                  not null,

    block_timestamp   bigint                       not null,
    included_in_block bigint                       not null,
    tx_hash           text                         not null,
    tx_index          integer                      not null,
    log_index         integer                      not null,
    created_at        timestamp default now()
);

create index tx_history_user_address_idx on smart_alpha.transaction_history (user_address asc, block_timestamp desc, tx_index desc, log_index desc);
create index tx_history_pool_address_idx on smart_alpha.transaction_history (pool_address asc, block_timestamp desc, tx_index desc, log_index desc);
