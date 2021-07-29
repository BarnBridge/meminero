package syncer

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

type Tokens []types.Token

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
			insert into tokens (address, symbol, decimals, prices)
			values ($1, $2, $3, $4) 
			on conflict (address) 
			do update set symbol = $2, decimals = $3, prices = $4
		`,
			utils.NormalizeAddress(token.Address),
			token.Symbol,
			token.Decimals,
			token.Prices,
		)
		if err != nil {
			return errors.Wrap(err, "could not insert token")
		}
	}

	return nil
}
