create table smart_exposure.tranches
(
    pool_address   text             not null,
    etoken_address text             not null,
    etoken_symbol  text,
    s_factor_e     numeric(78)      not null,
    target_ratio   numeric(78)      not null,
    token_a_ratio  double precision not null,
    token_b_ratio  double precision not null,
    start_at_block bigint           not null,

    created_at     timestamp default now()
);
