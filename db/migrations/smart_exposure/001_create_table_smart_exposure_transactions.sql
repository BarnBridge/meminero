create table smart_exposure.transaction_history
(
    user_address      text,
    etoken_address    text                       not null,
    amount            numeric(78),
    amount_a          numeric(78),
    amount_b          numeric(78),
    transaction_type  public.staking_action_type not null,

    block_timestamp   bigint                     not null,
    included_in_block bigint                     not null,
    tx_hash           text                       not null,
    tx_index          integer                    not null,
    log_index         integer                    not null,
    created_at        timestamp default now()
);

create index transaction_history_user_address_chronological_idx
    on smart_exposure.transaction_history (user_address asc, included_in_block desc, tx_index desc, log_index desc);

create index transaction_history_etoken_address_chronological_idx
    on smart_exposure.transaction_history (etoken_address asc, included_in_block desc, tx_index desc, log_index desc);


---- create above / drop below ----

drop table if exists smart_exposure.transaction_history;
