create table tokens
(
    address text not null,
    symbol text not null,
    decimals bigint not null,
    aggregator_address text not null,
    price_provider_type text not null,
    created_at timestamp default now()
);

---- create above / drop below ----

drop table if exists tokens;
