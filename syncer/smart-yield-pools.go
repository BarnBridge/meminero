package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/utils"
)

type SmartYieldPool struct {
	ProtocolId          string `json:"protocolId"`
	ControllerAddress   string `json:"controllerAddress"`
	ModelAddress        string `json:"modelAddress"`
	ProviderAddress     string `json:"providerAddress"`
	SmartYieldAddress   string `json:"smartYieldAddress"`
	OracleAddress       string `json:"oracleAddress"`
	JuniorBondAddress   string `json:"juniorBondAddress"`
	SeniorBondAddress   string `json:"seniorBondAddress"`
	ReceiptTokenAddress string `json:"receiptTokenAddress"`
	UnderlyingAddress   string `json:"underlyingAddress"`
	UnderlyingSymbol    string `json:"underlyingSymbol"`
	UnderlyingDecimals  int64  `json:"underlyingDecimals"`
	StartAtBlock        int64  `json:"startAtBlock"`
}

type SmartYieldPools []SmartYieldPool

func (p SmartYieldPools) Sync(tx pgx.Tx) error {
	if len(p) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(p)).Info("syncing smart yield pools")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing smart yield pools")
	}()

	for _, pool := range p {
		_, err := tx.Exec(context.Background(), `
			insert into smart_yield.pools
			(protocol_id, pool_address, controller_address, model_address, provider_address, oracle_address, junior_bond_address, senior_bond_address, receipt_token_address, underlying_address, underlying_symbol, underlying_decimals, start_at_block)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
			on conflict (pool_address) do
			update set protocol_id = $1, controller_address = $3, model_address = $4, provider_address = $5, oracle_address = $6, junior_bond_address = $7,
		               senior_bond_address = $8, receipt_token_address = $9, underlying_address = $10, underlying_symbol = $11, underlying_decimals = $12, start_at_block = $13
		`,
			pool.ProtocolId,
			utils.NormalizeAddress(pool.SmartYieldAddress),
			utils.NormalizeAddress(pool.ControllerAddress),
			utils.NormalizeAddress(pool.ModelAddress),
			utils.NormalizeAddress(pool.ProviderAddress),
			utils.NormalizeAddress(pool.OracleAddress),
			utils.NormalizeAddress(pool.JuniorBondAddress),
			utils.NormalizeAddress(pool.SeniorBondAddress),
			utils.NormalizeAddress(pool.ReceiptTokenAddress),
			utils.NormalizeAddress(pool.UnderlyingAddress),
			pool.UnderlyingSymbol,
			pool.UnderlyingDecimals,
			pool.StartAtBlock,
		)
		if err != nil {
			return errors.Wrap(err, "could not insert smart yield pool")
		}
	}

	return nil
}
