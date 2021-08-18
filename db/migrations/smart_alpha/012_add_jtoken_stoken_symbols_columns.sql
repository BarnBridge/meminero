alter table smart_alpha.pools
    add column if not exists junior_token_symbol text;
alter table smart_alpha.pools
    add column if not exists senior_token_symbol text;
alter table smart_alpha.pools
    add column if not exists senior_rate_model_address text;
alter table smart_alpha.pools
    add column if not exists accounting_model_address text;
