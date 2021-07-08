create table smart_yield.rewards_claims
(
    user_address         text    not null,
    amount               numeric(78),
    pool_address         text    not null,
    tx_hash              text    not null,
    tx_index             integer not null,
    log_index            integer not null,
    block_timestamp      bigint  not null,
    included_in_block    bigint  not null,
    reward_token_address text
);

---- create above / drop below ----

drop table if exists smart_yield.rewards_claims;