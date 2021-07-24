create table smart_exposure.pools
(
    pool_address     text   not null,
    pool_name        text   not null,
    token_a_address  text   not null,
    token_a_symbol   text   not null,
    token_a_decimals bigint not null,
    token_b_address  text   not null,
    token_b_symbol   text   not null,
    token_b_decimals bigint not null,
    start_at_block   bigint not null,

    created_at       timestamp default now()
);
