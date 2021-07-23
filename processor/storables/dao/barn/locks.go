package barn

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) handleLockEvents(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Barn.IsBarnLockEvent(&log) {
			lock, err := ethtypes.Barn.BarnLockEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode lock event")
			}

			s.processed.locks = append(s.processed.locks, lock)
		}
	}
	return nil
}
