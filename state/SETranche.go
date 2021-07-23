package state

import (
	"context"
	"database/sql"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/pkg/errors"
)

func (m *Manager) SETrancheByETokenAddress(address string) *types.SETranche {
	return m.seTranches[address]
}

func (m *Manager) loadAllSETranches(ctx context.Context) error {
	rows, err := m.db.Query(ctx, `select pool_address,etoken_address,etoken_symbol,s_factor_e,target_ratio,token_a_ratio,token_b_ratio from smart_exposure.tranches;`)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "could not query database for SmartExposure tranches")
	}

	m.seTranches = make(map[string]*types.SETranche)
	for rows.Next() {
		var p types.SETranche
		err := rows.Scan(&p.EPoolAddress, &p.ETokenAddress, &p.ETokenSymbol, &p.SFactorE, &p.TargetRatio, &p.TokenARatio, &p.TokenBRatio)
		if err != nil {
			return errors.Wrap(err, "could not scan tranches from database")
		}

		p.EPoolAddress = utils.NormalizeAddress(p.EPoolAddress)
		p.ETokenAddress = utils.NormalizeAddress(p.ETokenAddress)

		m.seTranches[p.ETokenAddress] = &p
	}

	return nil
}

func (m *Manager) AddNewTrancheToState(tranche types.SETranche) {
	m.seTranches[tranche.ETokenAddress] = &tranche
}
