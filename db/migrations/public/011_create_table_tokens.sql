create table tokens
(
    address    text   not null,
    symbol     text   not null,
    decimals   bigint not null,
    prices     jsonb  not null,
    created_at timestamp default now()
);

create unique index tokens_address_idx on tokens (address);
