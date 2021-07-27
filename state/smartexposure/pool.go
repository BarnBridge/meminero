package smartexposure

import (
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/utils"
)

func (se *SmartExposure) SEPoolByAddress(address string) *types.Pool {
	pool := se.Pools[utils.NormalizeAddress(address)]
	return &pool
}