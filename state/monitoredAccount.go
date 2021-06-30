package state

import (
	web3types "github.com/alethio/web3-go/types"
	"github.com/barnbridge/smartbackend/utils"
)

func  IsMonitoredAccount(log web3types.Log) bool {
	for _, a := range instance.monitoredAccounts{
		if len(log.Topics) >= 3 {
			if utils.NormalizeAddress(a) == utils.Topic2Address(log.Topics[1]) ||
				utils.NormalizeAddress(a) == utils.Topic2Address(log.Topics[2]) {
				return true
			}
		}
	}
	return false
}
