create table public.monitored_erc20
(
    address text not null,
    created_at timestamp default now()
);

---- create above / drop below ----

drop table if exists public.monitored_erc20