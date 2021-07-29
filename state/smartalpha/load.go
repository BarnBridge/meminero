package smartalpha

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/utils"
)

func (sa *SmartAlpha) Load(ctx context.Context, db *pgxpool.Pool) error {
	rows, err := db.Query(ctx, `
		select pool_name,
			   pool_address,
			   pool_token_address,
			   pool_token_symbol,
			   pool_token_decimals,
			   junior_token_address,
			   senior_token_address,
			   oracle_address,
			   oracle_asset_symbol,
			   start_at_block
		from smart_alpha.pools;
	`)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "could not query database")
	}

	for rows.Next() {
		var p smartalpha.Pool

		err := rows.Scan(
			&p.PoolName,
			&p.PoolAddress,
			&p.PoolToken.Address,
			&p.PoolToken.Symbol,
			&p.PoolToken.Decimals,
			&p.JuniorTokenAddress,
			&p.SeniorTokenAddress,
			&p.OracleAddress,
			&p.OracleAssetSymbol,
			&p.StartAtBlock,
		)
		if err != nil {
			return errors.Wrap(err, "could not scan pool")
		}

		p.PoolAddress = utils.NormalizeAddress(p.PoolAddress)
		p.PoolToken.Address = utils.NormalizeAddress(p.PoolToken.Address)
		p.JuniorTokenAddress = utils.NormalizeAddress(p.JuniorTokenAddress)
		p.SeniorTokenAddress = utils.NormalizeAddress(p.SeniorTokenAddress)
		p.OracleAddress = utils.NormalizeAddress(p.OracleAddress)

		sa.Pools = append(sa.Pools, p)
	}

	return nil
}
