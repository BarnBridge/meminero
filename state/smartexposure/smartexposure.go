package smartexposure

import (
	"github.com/barnbridge/meminero/utils"
)

type SmartExposure struct {
	Pools    map[string]*SEPool
	Tranches map[string]*SETranche
}

func New() *SmartExposure {
	return &SmartExposure{
		Pools:    make(map[string]*SEPool),
		Tranches: make(map[string]*SETranche),
	}
}

func (se *SmartExposure) SEPools() map[string]*SEPool {
	return se.Pools
}

func (se *SmartExposure) SEPoolByETokenAddress(address string) *SEPool {
	tranche := se.Tranches[utils.NormalizeAddress(address)]
	if tranche.ETokenAddress == "" {
		return nil
	}

	return se.Pools[tranche.EPoolAddress]
}

func (se *SmartExposure) SEPoolByAddress(address string) *SEPool {
	return se.Pools[utils.NormalizeAddress(address)]
}

func (se *SmartExposure) SETranches() map[string]*SETranche {
	return se.Tranches
}

func (se *SmartExposure) SETrancheByETokenAddress(address string) *SETranche {
	return se.Tranches[address]
}

func (se *SmartExposure) AddNewTrancheToState(tranche SETranche) {
	se.Tranches[tranche.ETokenAddress] = &tranche
}
