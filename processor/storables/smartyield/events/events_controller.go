package events

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) decodeControllerEvent(log gethtypes.Log) error {
	c := ethtypes.SmartYieldCompoundController

	if c.IsSmartYieldCompoundControllerHarvestEvent(&log) {
		e, err := c.SmartYieldCompoundControllerHarvestEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode Harvest event")
		}

		s.processed.ControllerHarvests = append(s.processed.ControllerHarvests, e)
	}

	return nil
}
