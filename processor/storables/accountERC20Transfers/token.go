package accountERC20Transfers

import (
	"github.com/barnbridge/smartbackend/contracts"
	"github.com/barnbridge/smartbackend/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (s *Storable) getToken(tokenAddress string) (*types.Token, error) {
	token, err := contracts.NewERC20(common.HexToAddress(tokenAddress), s.ethConn)
	if err != nil {
		return nil, errors.Wrap(err, "could not init erc20 contract")
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		return nil, err
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}

	return &types.Token{
		Address:  tokenAddress,
		Symbol:   symbol,
		Decimals: int64(decimals),
	}, nil
}
