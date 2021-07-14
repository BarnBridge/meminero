create table smart_yield.senior_redeem_events
(
    pool_address        text        not null,

    owner_address       text        not null,
    senior_bond_address text        not null,
    senior_bond_id      bigint      not null,
    fee                 numeric(78) not null,

    block_timestamp     bigint      not null,
    included_in_block   bigint      not null,
    tx_hash             text        not null,
    tx_index            integer     not null,
    log_index           integer     not null,
    created_at          timestamp default now()
);

create index senior_redeem_events_senior_bond_address_id_idx on smart_yield.senior_redeem_events (senior_bond_address asc, senior_bond_id asc, block_timestamp desc);

create index senior_redeem_events_user_address_idx on smart_yield.senior_redeem_events (owner_address);

---- create above / drop below ----

drop table if exists smart_yield.senior_redeem_events;
drop index if exists smart_yield.senior_redeem_events_senior_bond_address_id_idx;
drop index if exists smart_yield.senior_redeem_events_user_address_idx;
