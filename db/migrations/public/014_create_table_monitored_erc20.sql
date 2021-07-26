create table public.monitored_erc20
(
    address    text not null,
    created_at timestamp default now()
);

create unique index monitored_erc20_address_idx on monitored_erc20 (address);
