create table smart_yield.junior_2step_redeem_events
(
    pool_address        text    not null,

    owner_address       text    not null,
    junior_bond_address text    not null,
    junior_bond_id      bigint  not null,
    underlying_out      numeric(78),

    block_timestamp     bigint  not null,
    included_in_block   bigint  not null,
    tx_hash             text    not null,
    tx_index            integer not null,
    log_index           integer not null,
    created_at          timestamp default now()
);

create index junior_2step_redeem_junior_bond_address_id_idx on smart_yield.junior_2step_redeem_events (junior_bond_address asc, junior_bond_id asc, block_timestamp desc);

create index junior_2step_redeem_user_address_idx on smart_yield.junior_2step_redeem_events (owner_address);
