create table smart_yield.smart_yield_junior_buy
(
    sy_address          text    not null,
    buyer_address       text    not null,
    junior_bond_address text    not null,
    junior_bond_id      bigint  not null,
    tokens_in           numeric(78),
    matures_at          bigint,
    tx_hash             text    not null,
    tx_index            integer not null,
    log_index           integer not null,
    block_timestamp     bigint  not null,
    included_in_block   bigint  not null,
    created_at          timestamp default now()
);

create index smart_yield_junior_buy_junior_bond_address_id_idx
    on smart_yield.smart_yield_junior_buy (junior_bond_address, junior_bond_id);

---- create above / drop below ----

drop table if exists smart_yield.smart_yield_junior_buy;
drop index if exists smart_yield.smart_yield_junior_buy_junior_bond_address_id_idx;