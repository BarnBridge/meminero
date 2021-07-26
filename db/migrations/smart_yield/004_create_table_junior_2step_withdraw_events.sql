create table smart_yield.junior_2step_withdraw_events
(
    pool_address        text    not null,

    buyer_address       text    not null,
    junior_bond_address text    not null,
    junior_bond_id      bigint  not null,
    tokens_in           numeric(78),
    matures_at          bigint,

    block_timestamp     bigint  not null,
    included_in_block   bigint  not null,
    tx_hash             text    not null,
    tx_index            integer not null,
    log_index           integer not null,
    created_at          timestamp default now()
);

create index junior_2step_withdraw_junior_bond_address_id_idx on smart_yield.junior_2step_withdraw_events (junior_bond_address, junior_bond_id);
