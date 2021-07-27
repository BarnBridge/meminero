package smartexposure

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/utils"
)

func (se *SmartExposure) LoadPools(ctx context.Context, db *pgxpool.Pool) error {
	rows, err := db.Query(ctx, `select pool_address, pool_name, token_a_address, token_a_symbol, token_a_decimals, token_b_address, token_b_symbol, token_b_decimals,start_at_block from smart_exposure.pools;`)
	if err != nil {
		return errors.Wrap(err, "could not query database for SmartExposure pools")
	}

	for rows.Next() {
		var p types.Pool
		err := rows.Scan(&p.EPoolAddress, &p.ProtocolId, &p.ATokenAddress, &p.ATokenSymbol, &p.ATokenDecimals, &p.BTokenAddress, &p.BTokenSymbol, &p.BTokenDecimals, &p.StartAtBlock)
		if err != nil {
			return errors.Wrap(err, "could not scan pools from database")
		}

		p.EPoolAddress = utils.NormalizeAddress(p.EPoolAddress)
		p.ATokenAddress = utils.NormalizeAddress(p.ATokenAddress)
		p.BTokenAddress = utils.NormalizeAddress(p.BTokenAddress)

		se.Pools[p.EPoolAddress] = p
	}

	return nil
}

func (se *SmartExposure) LoadTranches(ctx context.Context, db *pgxpool.Pool) error {
	rows, err := db.Query(ctx, `select pool_address,etoken_address,etoken_symbol,s_factor_e,target_ratio,token_a_ratio,token_b_ratio from smart_exposure.tranches;`)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "could not query database for SmartExposure tranches")
	}

	for rows.Next() {
		var p types.Tranche
		err := rows.Scan(&p.EPoolAddress, &p.ETokenAddress, &p.ETokenSymbol, &p.SFactorE, &p.TargetRatio, &p.TokenARatio, &p.TokenBRatio)
		if err != nil {
			return errors.Wrap(err, "could not scan tranches from database")
		}

		p.EPoolAddress = utils.NormalizeAddress(p.EPoolAddress)
		p.ETokenAddress = utils.NormalizeAddress(p.ETokenAddress)
		se.Tranches[p.ETokenAddress] = p
	}

	return nil
}
