package state

import (
	"context"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (m *Manager) CheckTokenExists(log gethtypes.Log) bool {
	if len(m.Tokens) == 0 {
		return false
	}

	for _, a := range m.Tokens {
		if utils.NormalizeAddress(a.Address) == utils.NormalizeAddress(log.Address.String()){
			return true
		}
	}
	return false
}

func (m *Manager) loadAllTokens() error {
	rows, err := m.db.Query(context.Background(), `select address,symbol,decimals,aggregator_address,price_provider_type from tokens`)
	if err != nil {
		return errors.Wrap(err, "could not query database for monitored accounts")
	}

	var tokens []types.Token
	for rows.Next() {
		var t types.Token
		err := rows.Scan(&t.Address,&t.Symbol,&t.Decimals,&t.AggregatorAddress,&t.PriceProviderType)
		if err != nil {
			return errors.Wrap(err, "could no scan monitored accounts from database")
		}

		tokens = append(tokens, t)
	}

	m.Tokens = tokens

	return nil
}

func(m *Manager) StoreToken(token types.Token) error {
	_, err := m.db.Exec(context.Background(),`insert into tokens (address,symbol,decimals,aggregator_address,price_provider_type) values ($1,$2,$3,$4,$5)`, token.Address, token.Symbol, token.Decimals,token.AggregatorAddress,token.PriceProviderType)
	if err != nil {
		return err
	}
	m.Tokens = append(m.Tokens, token)
	return nil

}
