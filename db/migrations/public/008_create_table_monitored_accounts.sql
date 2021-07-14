create table monitored_accounts
(
    address text not null,
    created_at timestamp default now()
);

---- create above / drop below ----

drop table if exists monitored_accounts;
