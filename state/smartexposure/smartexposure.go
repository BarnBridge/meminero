package smartexposure

import (
	"github.com/barnbridge/meminero/processor/storables/smartexposure/types"
)

type SmartExposure struct {
	Pools    map[string]types.Pool
	Tranches map[string]types.Tranche
}

func New() *SmartExposure {
	return &SmartExposure{
		Pools:    make(map[string]types.Pool),
		Tranches: make(map[string]types.Tranche),
	}
}
