package smartexposure

import (
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/utils"
)

func (se *SmartExposure) PoolByAddress(address string) *types.Pool {
	pool, exists := se.Pools[utils.NormalizeAddress(address)]
	if !exists {
		return nil
	}

	return &pool
}
