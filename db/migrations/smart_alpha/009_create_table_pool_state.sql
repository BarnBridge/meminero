create table if not exists smart_alpha.pool_state
(
    pool_address                  text        not null,
    queued_juniors_underlying_in  numeric(78) not null,
    queued_juniors_underlying_out numeric(78) not null,
    queued_junior_tokens_burn     numeric(78) not null,
    queued_seniors_underlying_in  numeric(78) not null,
    queued_seniors_underlying_out numeric(78) not null,
    queued_senior_tokens_burn     numeric(78) not null,
    estimated_senior_liquidity    numeric(78) not null,
    estimated_junior_liquidity    numeric(78) not null,
    estimated_senior_token_price  numeric(78) not null,
    estimated_junior_token_price  numeric(78) not null,

    block_timestamp               bigint      not null,
    included_in_block             bigint      not null,
    created_at                    timestamp default now()
)
