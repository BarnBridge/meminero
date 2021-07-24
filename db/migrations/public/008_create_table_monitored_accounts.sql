create table monitored_accounts
(
    address    text not null,
    created_at timestamp default now()
);

create unique index monitored_accounts_address_idx on monitored_accounts (address);
