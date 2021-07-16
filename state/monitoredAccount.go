package state

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/utils"
)

func (m *Manager) loadAllAccounts() error {
	rows, err := m.db.Query(context.Background(), `select address from monitored_accounts`)
	if err != nil {
		return errors.Wrap(err, "could not query database for monitored accounts")
	}

	var accounts []string
	for rows.Next() {
		var a string
		err := rows.Scan(&a)
		if err != nil {
			return errors.Wrap(err, "could no scan monitored accounts from database")
		}

		accounts = append(accounts, utils.NormalizeAddress(a))
	}

	m.monitoredAccounts = accounts

	return nil
}

func (m *Manager) IsMonitoredAccount(addr string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, a := range m.monitoredAccounts {
		if utils.NormalizeAddress(a) == utils.NormalizeAddress(addr) {
			return true
		}
	}

	return false
}
