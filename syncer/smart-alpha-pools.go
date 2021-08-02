package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/processor/storables/smartalpha"
	"github.com/barnbridge/meminero/utils"
)

type SmartAlphaPools []smartalpha.Pool

func (p SmartAlphaPools) Sync(tx pgx.Tx) error {
	if len(p) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(p)).Info("syncing smart alpha pools")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing smart alpha pools")
	}()

	for _, pool := range p {
		_, err := tx.Exec(context.Background(), `
			insert into smart_alpha.pools (pool_name, pool_address, pool_token_address, pool_token_symbol, pool_token_decimals,
										   junior_token_address, senior_token_address, oracle_address, oracle_asset_symbol,
										   start_at_block)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			on conflict (pool_address) do update set pool_name            = $1,
													 pool_token_address   = $3,
													 pool_token_symbol    = $4,
													 pool_token_decimals  = $5,
													 junior_token_address = $6,
													 senior_token_address = $7,
													 oracle_address       = $8,
													 oracle_asset_symbol  = $9,
													 start_at_block       = $10;
		`,
			pool.PoolName,
			utils.NormalizeAddress(pool.PoolAddress),
			utils.NormalizeAddress(pool.PoolToken.Address),
			pool.PoolToken.Symbol,
			pool.PoolToken.Decimals,
			utils.NormalizeAddress(pool.JuniorTokenAddress),
			utils.NormalizeAddress(pool.SeniorTokenAddress),
			utils.NormalizeAddress(pool.OracleAddress),
			pool.OracleAssetSymbol,
			pool.StartAtBlock,
		)
		if err != nil {
			return errors.Wrap(err, "could not insert smart alpha pool")
		}
	}

	return nil
}
