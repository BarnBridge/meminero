package state

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
)

func (m *Manager) loadAllTokens() error {
	rows, err := m.db.Query(context.Background(), `select address,symbol,decimals,aggregator_address,price_provider_type from tokens`)
	if err != nil {
		return errors.Wrap(err, "could not query database for monitored accounts")
	}

	var tokens []types.Token
	for rows.Next() {
		var t types.Token
		err := rows.Scan(&t.Address, &t.Symbol, &t.Decimals, &t.AggregatorAddress, &t.PriceProviderType)
		if err != nil {
			return errors.Wrap(err, "could no scan monitored accounts from database")
		}

		tokens = append(tokens, t)
	}

	m.Tokens = tokens

	return nil
}

func (m *Manager) CheckTokenExists(addr string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, a := range m.Tokens {
		if utils.NormalizeAddress(a.Address) == utils.NormalizeAddress(addr) {
			return true
		}
	}

	return false
}

func (m *Manager) StoreToken(token types.Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, err := m.db.Exec(context.Background(), `insert into tokens (address,symbol,decimals,aggregator_address,price_provider_type) values ($1,$2,$3,$4,$5)`, token.Address, token.Symbol, token.Decimals, token.AggregatorAddress, token.PriceProviderType)
	if err != nil {
		return err
	}

	m.Tokens = append(m.Tokens, token)

	return nil
}
