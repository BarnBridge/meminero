create table smart_yield.smart_yield_token_sell
(
    sy_address        text    not null,
    seller_address    text    not null,
    tokens_in         numeric(78),
    underlying_out    numeric(78),
    forfeits          numeric(78),
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    block_timestamp   bigint  not null,
    included_in_block bigint  not null,
    created_at        timestamp default now()
);

---- create above / drop below ----

drop table if exists smart_yield.smart_yield_token_sell;