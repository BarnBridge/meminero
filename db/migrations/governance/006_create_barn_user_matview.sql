create materialized view governance.barn_users as
select distinct user_address
from governance.barn_staking_actions
union
select distinct receiver
from governance.barn_delegate_changes;

create unique index on governance.barn_users (user_address);

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on governance.barn_staking_actions
    for each statement
execute procedure governance.refresh_barn_users();

create trigger refresh_barn_users
    after insert or update or delete or truncate
    on governance.barn_delegate_changes
    for each statement
execute procedure governance.refresh_barn_users();

---- create above / drop below ----

drop function governance.refresh_barn_users;
drop materialized view governance.barn_users;
