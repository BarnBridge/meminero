create type staking_action_type as enum('DEPOSIT','WITHDRAW');
create type transfer_type as enum('IN','OUT');
---- create above / drop below ----

drop type if exists staking_actions_type;
drop type if exists transfer_type;
