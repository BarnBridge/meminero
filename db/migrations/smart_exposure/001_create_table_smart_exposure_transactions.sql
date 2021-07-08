create table smart_exposure.transactions
(
    user_address      text,
    tx_hash           text                not null,
    tx_index          integer             not null,
    log_index         integer             not null,
    etoken_address   text                not null,
    amount            numeric(78),
    amount_a          numeric(78),
    amount_b          numeric(78),
    transaction_type  public.staking_action_type not null,
    block_timestamp   bigint              not null,
    included_in_block bigint              not null,
    created_at        timestamp default now()
);


---- create above / drop below ----

drop table if exists smart_exposure.transactions;
