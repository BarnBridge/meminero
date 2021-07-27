package smartexposure

import (
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
	"github.com/barnbridge/meminero/utils"
)

func (se *SmartExposure) TrancheByETokenAddress(address string) *types.Tranche {
	t := se.Tranches[utils.NormalizeAddress(address)]
	return &t
}

func (se *SmartExposure) AddNewTrancheToState(tranche types.Tranche) {
	se.Tranches[tranche.ETokenAddress] = tranche
}
