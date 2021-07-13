create table smart_yield.senior_entry_events
(
    sy_address          text        not null,
    buyer_address       text        not null,
    senior_bond_address text        not null,
    senior_bond_id      numeric(78) not null,
    underlying_in       numeric(78),
    gain                numeric(78),
    for_days            bigint,
    block_timestamp     bigint      not null,
    included_in_block   bigint      not null,
    tx_hash             text        not null,
    tx_index            integer     not null,
    log_index           integer     not null,
    created_at          timestamp default now()
);

create index senior_entry_events_senior_bond_address_id_idx
    on smart_yield.senior_entry_events (senior_bond_address, senior_bond_id);


---- create above / drop below ----

drop table if exists smart_yield.senior_entry_events;
drop index if exists smart_yield.senior_entry_events_senior_bond_address_id_idx;
