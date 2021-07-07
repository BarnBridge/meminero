create materialized view barn.barn_users as
select distinct user_address
from barn.barn_staking_actions
union
select distinct receiver
from barn.barn_delegate_changes;

create unique index on barn.barn_users (user_address);

create or replace function barn.refresh_barn_users() returns TRIGGER
    language plpgsql as
$$
begin
    refresh materialized view concurrently barn.barn_users;
    return null;
end
$$;

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on barn.barn_staking_actions
    for each statement
execute procedure barn.refresh_barn_users();

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on barn.barn_delegate_changes
    for each statement
execute procedure barn.refresh_barn_users();

---- create above / drop below ----

drop function barn.refresh_barn_users;
drop materialized view barn.barn_users;
