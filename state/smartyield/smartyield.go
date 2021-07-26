package smartyield

import (
	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

type SmartYield struct {
	Pools       []smartyield.Pool
	RewardPools []smartyield.RewardPool
}

func New() *SmartYield {
	return &SmartYield{}
}

func (sy *SmartYield) PoolByAddress(addr string) *smartyield.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.PoolAddress {
			return &p
		}
	}

	return nil
}

func (sy *SmartYield) PoolByControllerAddress(addr string) *smartyield.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.ControllerAddress {
			return &p
		}
	}

	return nil
}

func (sy *SmartYield) PoolByProviderAddress(addr string) *smartyield.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.ProviderAddress {
			return &p
		}
	}

	return nil
}

func (sy *SmartYield) IsERC721OfInterest(addr string) bool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.SeniorBondAddress || addr == p.JuniorBondAddress {
			return true
		}
	}

	return false
}

func (sy *SmartYield) PoolBySeniorBondAddress(addr string) *smartyield.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.SeniorBondAddress {
			return &p
		}
	}

	return nil
}

func (sy *SmartYield) PoolByJuniorBondAddress(addr string) *smartyield.Pool {
	addr = utils.NormalizeAddress(addr)

	for _, p := range sy.Pools {
		if addr == p.JuniorBondAddress {
			return &p
		}
	}

	return nil
}
