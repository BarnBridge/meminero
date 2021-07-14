package accountERC20Transfers

import (
	"database/sql"

	"github.com/barnbridge/smartbackend/contracts"
	"github.com/barnbridge/smartbackend/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (s *Storable) checkTokenExists(tokenAddress string) error {
	//var count int64
	/*	err := s.db.QueryRow(context.Background(),`select count(*) from erc20_tokens where token_address = $1`, tokenAddress).Scan(&count)
		if err != nil {
			return err
		} else if count > 0 {
			return nil
		}

		token, err := s.getToken(tokenAddress)
		if err != nil {
			return err
		}

		err = storeToken(*token, tx)
		if err != nil {
			return err
		}*/

	return nil
}

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

func storeToken(token types.Token, tx *sql.Tx) error {
	_, err := tx.Exec(`insert into erc20_tokens (token_address,symbol,decimals) values ($1,$2,$3)`, token.Address, token.Symbol, token.Decimals)
	if err != nil {
		return err
	}

	return nil

}
