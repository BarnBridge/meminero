package accountERC20Transfers

import (
	"github.com/barnbridge/smartbackend/ethtypes"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (s *Storable) decodeLogs(logs []gethtypes.Log, erc20Decoder *ethtypes.ERC20Decoder) error {
	for _, log := range logs {
		t, err := s.decodeTransfer(log, erc20Decoder)
		if err != nil {
			return err
		} else if t != nil {

			s.processed.transfers = append(s.processed.transfers, *t)
		}
	}
	return nil
}

func (s *Storable) decodeTransfer(log gethtypes.Log, erc20Decoder *ethtypes.ERC20Decoder) (*ethtypes.ERC20TransferEvent, error) {
	erc20Transfer, err := erc20Decoder.ERC20TransferEvent(log)
	if err != nil {
		return nil, err
	}

	return &erc20Transfer, nil
}
