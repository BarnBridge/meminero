alter table smart_alpha.pools
    add column if not exists junior_token_symbol text;
alter table smart_alpha.pools
    add column if not exists senior_token_symbol text;
alter table smart_alpha.pools
    add column if not exists senior_rate_model_address text;
alter table smart_alpha.pools
    add column if not exists accounting_model_address text;

---- create above / drop below ----

alter table smart_alpha.pools drop column if exists junior_token_symbol;
alter table smart_alpha.pools drop column if exists senior_token_symbol;
alter table smart_alpha.pools drop column if exists senior_rate_model_address;
alter table smart_alpha.pools drop column if exists accounting_model_address;