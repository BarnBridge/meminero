create type sy_tx_history_tx_type as enum ( 'JUNIOR_DEPOSIT', 'JUNIOR_INSTANT_WITHDRAW', 'JUNIOR_REGULAR_WITHDRAW','JUNIOR_REDEEM','SENIOR_DEPOSIT','SENIOR_REDEEM','JTOKEN_SEND', 'JTOKEN_RECEIVE','JBOND_SEND', 'JBOND_RECEIVE', 'SBOND_SEND', 'SBOND_RECEIVE','JUNIOR_UNSTAKE','JUNIOR_STAKE');

create table smart_yield.transaction_history
(
    protocol_id              text                  not null,
    pool_address             text                  not null,

    underlying_token_address text                  not null,
    user_address             text                  not null,
    amount                   numeric(78),
    tranche                  text                  not null,
    transaction_type         sy_tx_history_tx_type not null,

    block_timestamp          bigint                not null,
    included_in_block        bigint                not null,
    tx_hash                  text                  not null,
    tx_index                 integer               not null,
    log_index                integer               not null,
    created_at               timestamp default now()
);

create index tx_history_user_address_idx on smart_yield.transaction_history (user_address asc, block_timestamp desc, tx_index desc, log_index desc);
