package tokenprices

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.prices) == 0 {
		return nil
	}

	var rows [][]interface{}

	for tokenAddress, prices := range s.processed.prices {
		for quoteAsset, price := range prices {
			token := s.state.GetTokenByAddress(tokenAddress)
			price, _ := price.Float64()
			rows = append(rows, []interface{}{
				tokenAddress,
				token.Symbol,
				quoteAsset,
				price,
				s.block.BlockCreationTime,
				s.block.Number,
			})
		}
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"token_prices"},
		[]string{"token_address", "token_symbol", "quote_asset", "price", "block_timestamp", "included_in_block"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not execute copy")
	}

	return nil
}
