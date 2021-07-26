package events

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) decodeProviderEvent(log gethtypes.Log) error {
	p := ethtypes.SmartYieldCompoundProvider

	if p.IsSmartYieldCompoundProviderTransferFeesEvent(&log) {
		e, err := p.SmartYieldCompoundProviderTransferFeesEvent(log)
		if err != nil {
			return errors.Wrap(err, "could not decode TransferFees event")
		}

		s.processed.ProviderTransferFees = append(s.processed.ProviderTransferFees, e)
	}

	return nil
}
