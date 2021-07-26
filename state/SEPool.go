package state

import (
	"context"

	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	"github.com/barnbridge/meminero/utils"

	"github.com/pkg/errors"
)

func (m *Manager) SEPools() map[string]*smartexposure.SEPool {
	return m.sePools
}

func (m *Manager) SEPoolByETokenAddress(address string) *smartexposure.SEPool {
	tranche := m.seTranches[utils.NormalizeAddress(address)]
	if tranche.ETokenAddress == "" {
		return nil
	}

	return m.sePools[tranche.EPoolAddress]
}

func (m *Manager) SEPoolByAddress(address string) *smartexposure.SEPool {
	return m.sePools[utils.NormalizeAddress(address)]
}

func (m *Manager) loadAllSEPools(ctx context.Context) error {
	rows, err := m.db.Query(ctx, `select pool_address, pool_name, token_a_address, token_a_symbol, token_a_decimals, token_b_address, token_b_symbol, token_b_decimals,start_at_block from smart_exposure.pools;`)
	if err != nil {
		return errors.Wrap(err, "could not query database for SmartExposure pools")
	}

	m.sePools = make(map[string]*smartexposure.SEPool)
	for rows.Next() {
		var p smartexposure.SEPool
		err := rows.Scan(&p.EPoolAddress, &p.ProtocolId, &p.ATokenAddress, &p.ATokenSymbol, &p.ATokenDecimals, &p.BTokenAddress, &p.BTokenSymbol, &p.BTokenDecimals, &p.StartAtBlock)
		if err != nil {
			return errors.Wrap(err, "could not scan pools from database")
		}

		p.EPoolAddress = utils.NormalizeAddress(p.EPoolAddress)
		p.ATokenAddress = utils.NormalizeAddress(p.ATokenAddress)
		p.BTokenAddress = utils.NormalizeAddress(p.BTokenAddress)

		m.sePools[p.EPoolAddress] = &p
	}

	return nil
}
