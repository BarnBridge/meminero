package state

import (
	"context"
	"database/sql"

	"github.com/barnbridge/meminero/processor/storables/smartexposure"
	"github.com/barnbridge/meminero/utils"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (m *Manager) SETranches() map[string]*smartexposure.SETranche {
	return m.seTranches
}

func (m *Manager) SETrancheByETokenAddress(address string) *smartexposure.SETranche {
	return m.seTranches[address]
}

func (m *Manager) loadAllSETranches(ctx context.Context) error {
	rows, err := m.db.Query(ctx, `select pool_address,etoken_address,etoken_symbol,s_factor_e,target_ratio,token_a_ratio,token_b_ratio from smart_exposure.tranches;`)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "could not query database for SmartExposure tranches")
	}

	m.seTranches = make(map[string]*smartexposure.SETranche)
	for rows.Next() {
		var p smartexposure.SETranche
		var factor, targetRatio decimal.Decimal
		err := rows.Scan(&p.EPoolAddress, &p.ETokenAddress, &p.ETokenSymbol, &factor, &targetRatio, &p.TokenARatio, &p.TokenBRatio)
		if err != nil {
			return errors.Wrap(err, "could not scan tranches from database")
		}

		p.EPoolAddress = utils.NormalizeAddress(p.EPoolAddress)
		p.ETokenAddress = utils.NormalizeAddress(p.ETokenAddress)
		p.SFactorE = factor.BigInt()
		p.TargetRatio = targetRatio.BigInt()
		m.seTranches[p.ETokenAddress] = &p
	}

	return nil
}

func (m *Manager) AddNewTrancheToState(tranche smartexposure.SETranche) {
	m.seTranches[tranche.ETokenAddress] = &tranche
}
