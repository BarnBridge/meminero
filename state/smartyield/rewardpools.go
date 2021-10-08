package smartyield

import (
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (sy *SmartYield) RewardPoolByAddress(addr string) *types.RewardPool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.RewardPools {
		if addr == p.PoolAddress {
			return &p
		}
	}

	return nil
}

func (sy *SmartYield) CacheRewardPool(p types.RewardPool) {
	if sy.RewardPoolByAddress(p.PoolAddress) != nil {
		return
	}

	sy.RewardPools = append(sy.RewardPools, p)
}
