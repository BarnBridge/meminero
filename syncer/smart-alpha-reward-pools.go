package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

type SmartAlphaRewardPool struct {
	PoolType             string   `json:"poolType"`
	PoolAddress          string   `json:"poolAddress"`
	PoolTokenAddress     string   `json:"poolTokenAddress"`
	RewardTokenAddresses []string `json:"rewardTokenAddresses"`
	StartAtBlock         int64    `json:"startAtBlock"`
}

type SmartAlphaRewardPools []SmartAlphaRewardPool

func (p SmartAlphaRewardPools) Sync(tx pgx.Tx) error {
	if len(p) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(p)).Info("syncing smart alpha reward pools")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing smart alpha reward pools")
	}()

	for _, pool := range p {
		_, err := tx.Exec(context.Background(), `
			insert into smart_alpha.reward_pools
			(pool_type, pool_address, pool_token_address, reward_token_addresses, start_at_block)
			values ($1, $2, $3, $4, $5) 
			on conflict (pool_address) do
			update set pool_type = $1, pool_token_address = $3, reward_token_addresses = $4, start_at_block = $5
		`,
			pool.PoolType,
			utils.NormalizeAddress(pool.PoolAddress),
			utils.NormalizeAddress(pool.PoolTokenAddress),
			utils.NormalizeAddresses(pool.RewardTokenAddresses),
			pool.StartAtBlock,
		)
		if err != nil {
			return errors.Wrap(err, "could not insert smart alpha reward pool")
		}
	}

	return nil
}
