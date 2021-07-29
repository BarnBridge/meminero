package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

type SmartExposurePool struct {
	PoolName     string      `json:"poolName"`
	EPoolAddress string      `json:"ePoolAddress"`
	TokenA       types.Token `json:"tokenA"`
	TokenB       types.Token `json:"tokenB"`
	StartAtBlock int64       `json:"startAtBlock"`
}

type SmartExposurePools []SmartExposurePool

func (p SmartExposurePools) Sync(tx pgx.Tx) error {
	if len(p) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(p)).Info("syncing smart exposure pools")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing smart exposure pools")
	}()

	for _, pool := range p {
		_, err := tx.Exec(context.Background(), `
			insert into smart_exposure.pools 
			(pool_address, pool_name, token_a_address, token_a_symbol, token_a_decimals, token_b_address, token_b_symbol, token_b_decimals, start_at_block)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
			on conflict (pool_address) do
			update set pool_name = $2, token_a_address = $3, token_a_symbol = $4, token_a_decimals = $5, token_b_address = $6, token_b_symbol = $7, token_b_decimals = $8, start_at_block = $9
		`,
			utils.NormalizeAddress(pool.EPoolAddress),
			pool.PoolName,
			utils.NormalizeAddress(pool.TokenA.Address), pool.TokenA.Symbol, pool.TokenA.Decimals,
			utils.NormalizeAddress(pool.TokenB.Address), pool.TokenB.Symbol, pool.TokenB.Decimals,
			pool.StartAtBlock)
		if err != nil {
			return errors.Wrap(err, "could not insert smart exposure pool")
		}
	}

	return nil
}
