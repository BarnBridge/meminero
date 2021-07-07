create table smart_yield.smart_yield_senior_redeem
(
    sy_address          text        not null,
    owner_address       text        not null,
    senior_bond_address text        not null,
    senior_bond_id      bigint      not null,
    fee                 numeric(78) not null,
    tx_hash             text        not null,
    tx_index            integer     not null,
    log_index           integer     not null,
    block_timestamp     bigint      not null,
    included_in_block   bigint      not null,
    created_at          timestamp default now()
);

create index smart_yield_senior_redeem_junior_bond_address_id_idx
    on smart_yield.smart_yield_senior_redeem (senior_bond_address asc, senior_bond_id asc, block_timestamp desc);

create index sy_senior_redeem_user_address_idx
    on smart_yield.smart_yield_senior_redeem (owner_address);

---- create above / drop below ----

drop table if exists smart_yield.smart_yield_senior_redeem;
drop index if exists smart_yield.smart_yield_senior_redeem_junior_bond_address_id_idx;
drop index if exists smart_yield.sy_senior_redeem_user_address_idx;
