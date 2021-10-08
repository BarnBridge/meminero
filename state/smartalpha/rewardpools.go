package smartalpha

import (
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (sa *SmartAlpha) RewardPoolByAddress(addr string) *types.RewardPool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sa.RewardPools {
		if addr == p.PoolAddress {
			return &p
		}
	}

	return nil
}
