package accountERC20Transfers

import (
	web3types "github.com/alethio/web3-go/types"
	"github.com/barnbridge/smartbackend/ethtypes"
)
func (s *Storable) decodeLogs(logs []web3types.Log,erc20Decoder *ethtypes.ERC20Decoder) error {
	for _, log := range logs {
		t, err := s.decodeTransfer(log,erc20Decoder)
		if err != nil {
			return err
		} else if t != nil {

			s.processed.transfers = append(s.processed.transfers, *t)
		}
	}
	return nil
}

func (s *Storable) decodeTransfer(log web3types.Log,erc20Decoder *ethtypes.ERC20Decoder) (*ethtypes.ERC20TransferEvent, error) {

	erc20Transfer,err := erc20Decoder.ERC20TransferEventW3(log)
	if err != nil {
		return nil,err
	}

	return &erc20Transfer,nil
}