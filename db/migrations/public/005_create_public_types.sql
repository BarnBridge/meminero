create type staking_action_type as enum('DEPOSIT','WITHDRAW');
---- create above / drop below ----

drop type if exists staking_actions_type;
