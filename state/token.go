package state

import (
	"context"

	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (m *Manager) loadAllTokens(ctx context.Context) error {
	rows, err := m.db.Query(ctx, `select address,symbol,decimals,aggregator_address,price_provider_type from tokens`)
	if err != nil {
		return errors.Wrap(err, "could not query database for monitored accounts")
	}

	m.Tokens = make(map[string]types.Token)
	for rows.Next() {
		var t types.Token
		err := rows.Scan(&t.Address, &t.Symbol, &t.Decimals, &t.AggregatorAddress, &t.PriceProviderType)
		if err != nil {
			return errors.Wrap(err, "could no scan monitored accounts from database")
		}
		t.Address = utils.NormalizeAddress(t.Address)
		t.AggregatorAddress = utils.NormalizeAddress(t.AggregatorAddress)
		m.Tokens[t.Address] = t
	}

	return nil
}

func (m *Manager) CheckTokenExists(addr string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Tokens[utils.NormalizeAddress(addr)]; exists {
		return true
	}

	return false
}

func (m *Manager) StoreToken(ctx context.Context, token types.Token) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, err := m.db.Exec(ctx, `insert into tokens (address,symbol,decimals,aggregator_address,price_provider_type) values ($1,$2,$3,$4,$5)`, utils.NormalizeAddress(token.Address), token.Symbol, token.Decimals, token.AggregatorAddress, token.PriceProviderType)
	if err != nil {
		return err
	}

	m.Tokens[utils.NormalizeAddress(token.Address)] = token

	return nil
}
