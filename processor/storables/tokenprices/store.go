package tokenprices

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) storeTokensPrice(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.prices) == 0 {
		return nil
	}

	var rows [][]interface{}

	for tokenAddress, p := range s.processed.prices {
		token := s.state.GetTokenByAddress(tokenAddress)
		price, _ := p.Float64()
		rows = append(rows, []interface{}{
			tokenAddress,
			token.Symbol,
			price,
			s.block.BlockCreationTime,
			s.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"token_prices"},
		[]string{"token_address", "token_symbol", "price_usd", "block_timestamp", "included_in_block"},
		pgx.CopyFromRows(rows),
	)

	return err
}
