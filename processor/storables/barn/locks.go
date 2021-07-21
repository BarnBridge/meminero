package barn

import (
	"github.com/barnbridge/smartbackend/ethtypes"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (s *Storable) decodeLockEvents(logs []gethtypes.Log) error {
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
