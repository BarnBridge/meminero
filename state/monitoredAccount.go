package state

import (
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func  IsMonitoredAccount(log gethtypes.Log) bool {
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
