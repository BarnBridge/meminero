create or replace view governance.voters as
select user_address,
       barn.balance_of(user_address)                                                                                 as bond_staked,
       coalesce((select locked_until
                 from barn.barn_locks
                 where user_address = barn_users.user_address
                 order by included_in_block desc, log_index desc
                 limit 1),
                0)                                                                                                   as locked_until,
       barn.delegated_power(user_address),
       (select count(*) from governance.governance_votes where lower(user_id) = lower(barn_users.user_address)) +
       (select count(*)
        from governance.governance_abrogation_votes
        where lower(user_id) = lower(barn_users.user_address))                                                       as votes,
       (select count(*)
        from governance.governance_proposals
        where lower(proposer) = lower(barn_users.user_address))                                                      as proposals,
       barn.voting_power(user_address)                                                                               as voting_power,
       barn.has_active_delegation(user_address)
from barn.barn_users;

---- create above / drop below ----

drop view governance.voters