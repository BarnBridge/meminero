package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(UpAlterDeleteBlock3Function, DownAlterDeleteBlock3Function)
}

func UpAlterDeleteBlock3Function(tx *sql.Tx) error {
	_, err := tx.Exec(`
	drop function  delete_block(in block_number bigint);
	create or replace function delete_block(in block_number bigint, tables character varying[]) returns void as
	$body$
	declare
		tbl    varchar;
	begin
		foreach tbl in array tables
			loop
				perform __delete_entity(tbl, block_number);
			end loop;

		delete from blocks where number = block_number;
	end;
	$body$ language 'plpgsql';
	`)
	return err
}

func DownAlterDeleteBlock3Function(tx *sql.Tx) error {
	return nil
}
