package state

import (
	"context"

	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (m *Manager) IsMonitoredAccount(log gethtypes.Log) bool {
	if len(instance.monitoredAccounts) == 0 {
		return false
	}

	for _, a := range instance.monitoredAccounts{
		if len(log.Topics) >= 3 {
			if utils.NormalizeAddress(a) == utils.Topic2Address(log.Topics[1].String()) ||
				utils.NormalizeAddress(a) == utils.Topic2Address(log.Topics[2].String()) {
				return true
			}
		}
	}
	return false
}

func (m *Manager) loadAllAccounts() error {
	rows, err := instance.db.Query(context.Background(),`select address from monitored_accounts`)
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

		accounts = append(accounts, a)
	}

	instance.monitoredAccounts = accounts

	return nil
}