package state

import (
	"context"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/pkg/errors"
)

func (m *Manager) loadAllAccounts() error {
	rows, err := m.db.Query(context.Background(), `select address from monitored_accounts`)
	if err != nil {
		return errors.Wrap(err, "could not query database for monitored accounts")
	}

	accounts := make(map[string]bool)
	for rows.Next() {
		var a string
		err := rows.Scan(&a)
		if err != nil {
			return errors.Wrap(err, "could no scan monitored accounts from database")
		}
		a = utils.NormalizeAddress(a)
		accounts[a] = true
	}

	m.monitoredAccounts = accounts

	return nil
}

func (m *Manager) IsMonitoredAccount(addr string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.monitoredAccounts[utils.NormalizeAddress(addr)]{
		return true
	}

	return false
}
