create table smart_exposure.tranche_state
(
    included_in_block     bigint not null,
    block_timestamp       bigint not null,
    pool_address          text   not null,
    etoken_address        text   not null,
    token_a_liquidity     double precision,
    token_b_liquidity     double precision,
    current_ratio         numeric(78),
    amount_a_conversion   numeric(78),
    amount_b_conversion   numeric(78),
    etoken_price          double precision,
    token_a_current_ratio double precision,
    token_b_current_ratio double precision
);

---- create above / drop below ----

drop table if exists smart_exposure.tranche_state;