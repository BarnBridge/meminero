package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpCreateFunctions, DownCreateFunctions)
}

func UpCreateFunctions(tx *sql.Tx) error {
	_, err := tx.Exec(`
	create or replace function balance_of(addr text) returns numeric(78)
		language plpgsql as
	$$
	declare
		result numeric(78);
	begin
		select balance_after 
		into result
		from barn_staking_actions
		where user_address = addr
		order by included_in_block desc, log_index desc
		limit 1;
	
		return coalesce(result, 0);
	end;
	$$;
	
	create or replace function has_active_delegation(addr text) returns bool
		language plpgsql as
	$$
	declare
		action text;
	begin
		select action_type
		into action
		from barn_delegate_actions
		where sender = addr
		order by included_in_block desc, log_index desc
		limit 1;
	
		if action = 'START' then return true; end if;
	
		return false;
	end;
	$$;
	
	create or replace function user_multiplier(addr text) returns numeric(78)
		language plpgsql as
	$$
	declare
		locked_until_ts bigint;
		time_now        bigint;
		time_left       bigint;
		multiplier      numeric(78);
	begin
		multiplier = 1 * 10 ^ 18;
	
		select locked_until into locked_until_ts from barn_locks where user_address = addr order by included_in_block desc, log_index desc limit 1;
	
		if not found then return multiplier; end if;
	
		select floor(extract(epoch from now())) into time_now;
	
		if locked_until_ts <= time_now then return multiplier; end if;
	
		time_left = locked_until_ts - time_now;
	
		multiplier = multiplier + (time_left::numeric * 10 ^ 18 / 31536000::numeric);
	
		return multiplier;
	end;
	$$;
	
	create or replace function delegated_power(addr text) returns numeric(78)
		language plpgsql as
	$$
	declare
		result numeric(78);
	begin
		select receiver_new_delegated_power
		into result
		from barn_delegate_changes
		where receiver = addr
		order by included_in_block desc, log_index desc;
	
		return coalesce(result, 0);
	end;
	$$;
	
	create or replace function voting_power(addr text) returns numeric(78)
		language plpgsql as
	$$
	declare
		is_delegating   bool;
		delegated_power numeric(78);
		self_power      numeric(78);
	begin
		select has_active_delegation(addr) into is_delegating;
	
		if is_delegating then
			self_power = 0;
		else
			select balance_of(addr) * user_multiplier(addr) / 10 ^ 18 into self_power;
		end if;
	
		select delegated_power(addr) into delegated_power;
	
		return self_power + delegated_power;
	end;
	$$;
	`)
	return err
}

func DownCreateFunctions(tx *sql.Tx) error {
	_, err := tx.Exec(`
	drop function delegated_power(addr text);
	drop function voting_power(addr text);
	drop function balance_of(addr text);
	drop function user_multiplier(addr text);
	drop function has_active_delegation(addr text);
	`)
	return err
}
