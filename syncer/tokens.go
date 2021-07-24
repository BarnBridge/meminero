package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Token struct {
	Address           string `json:"address"`
	Symbol            string `json:"symbol"`
	Decimals          int64  `json:"decimals"`
	AggregatorAddress string `json:"aggregatorAddress"`
	PriceProviderType string `json:"priceProviderType"`
}

type Tokens []Token

func (t Tokens) Sync(tx pgx.Tx) error {
	if len(t) == 0 {
		return nil
	}

	start := time.Now()
	log.WithField("count", len(t)).Info("syncing tokens")
	defer func() {
		log.WithField("duration", time.Since(start)).Info("done syncing tokens")
	}()

	for _, token := range t {
		_, err := tx.Exec(context.Background(), `
			insert into tokens (address, symbol, decimals, aggregator_address, price_provider_type)
			values ($1, $2, $3, $4, $5) 
			on conflict (address) 
			do update set symbol = $2, decimals = $3, aggregator_address = $4, price_provider_type = $5
		`, token.Address, token.Symbol, token.Decimals, token.AggregatorAddress, token.PriceProviderType)
		if err != nil {
			return errors.Wrap(err, "could not insert token")
		}
	}

	return nil
}
