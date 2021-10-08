package smartyield

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (sy *SmartYield) LoadPools(ctx context.Context, db *pgxpool.Pool) error {
	sy.Pools = []smartyield.Pool{}

	rows, err := db.Query(ctx, `
		select protocol_id,
			   pool_address,
			   controller_address,
			   model_address,
			   provider_address,
			   oracle_address,
			   junior_bond_address,
			   senior_bond_address,
			   receipt_token_address,
			   underlying_address,
			   underlying_symbol,
			   underlying_decimals,
			   start_at_block
		from smart_yield.pools
	`)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "could not load smart yield pools")
	}

	for rows.Next() {
		var p smartyield.Pool

		err := rows.Scan(
			&p.ProtocolId,
			&p.PoolAddress,
			&p.ControllerAddress,
			&p.ModelAddress,
			&p.ProviderAddress,
			&p.OracleAddress,
			&p.JuniorBondAddress,
			&p.SeniorBondAddress,
			&p.ReceiptTokenAddress,
			&p.UnderlyingAddress,
			&p.UnderlyingSymbol,
			&p.UnderlyingDecimals,
			&p.StartAtBlock,
		)
		if err != nil {
			return errors.Wrap(err, "could not scan smart yield pools from database")
		}

		p.PoolAddress = utils.NormalizeAddress(p.PoolAddress)
		p.ControllerAddress = utils.NormalizeAddress(p.ControllerAddress)
		p.ModelAddress = utils.NormalizeAddress(p.ModelAddress)
		p.ProviderAddress = utils.NormalizeAddress(p.ProviderAddress)
		p.OracleAddress = utils.NormalizeAddress(p.OracleAddress)
		p.JuniorBondAddress = utils.NormalizeAddress(p.JuniorBondAddress)
		p.SeniorBondAddress = utils.NormalizeAddress(p.SeniorBondAddress)
		p.ReceiptTokenAddress = utils.NormalizeAddress(p.ReceiptTokenAddress)
		p.UnderlyingAddress = utils.NormalizeAddress(p.UnderlyingAddress)

		sy.Pools = append(sy.Pools, p)
	}

	return nil
}

func (sy *SmartYield) LoadRewardPools(ctx context.Context, db *pgxpool.Pool) error {
	sy.RewardPools = []types.RewardPool{}

	rows, err := db.Query(ctx, `
		select pool_type, pool_address, pool_token_address, reward_token_addresses, start_at_block
		from smart_yield.reward_pools
	`)
	if err != nil && err != pgx.ErrNoRows {
		return errors.Wrap(err, "could not execute query")
	}

	for rows.Next() {
		var p types.RewardPool

		err := rows.Scan(&p.PoolType, &p.PoolAddress, &p.PoolTokenAddress, &p.RewardTokenAddresses, &p.StartAtBlock)
		if err != nil {
			return errors.Wrap(err, "could not scan reward pool")
		}

		sy.RewardPools = append(sy.RewardPools, p)
	}

	return nil
}
