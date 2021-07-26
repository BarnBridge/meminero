create table yield_farming.transactions
(
    user_address      text                       not null,
    token_address     text                       not null,
    amount            numeric(78),
    action_type       public.staking_action_type not null,

    block_timestamp   bigint                     not null,
    included_in_block bigint                     not null,
    tx_hash           text                       not null,
    tx_index          integer                    not null,
    log_index         integer                    not null,
    created_at        timestamp default now()
);
