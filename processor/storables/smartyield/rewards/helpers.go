package rewards

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) isFactory(addr string) bool {
	addr = utils.NormalizeAddress(addr)
	for _, v := range s.factories {
		if addr == utils.NormalizeAddress(v) {
			return true
		}
	}

	return false
}

func (s *Storable) checkTokenExists(tokenAddress string) error {
	if s.state.CheckTokenExists(tokenAddress) {
		return nil
	}

	token, err := eth.GetERC20TokenFromChain(tokenAddress)
	if err != nil {
		return errors.Wrap(err, "could not get erc20 info from chain")
	}

	err = s.state.StoreToken(context.Background(), *token)
	if err != nil {
		return errors.Wrap(err, "could not store token to state")
	}

	return nil
}
