package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func GetTokenPrice(ctx context.Context, tx pgx.Tx, tokenAddress string, blockNumber int64) (decimal.Decimal, error) {
	var price decimal.Decimal

	err := tx.QueryRow(ctx, `select price_usd from tokens_prices where token_address = $1 and included_in_block <= $2 order by included_in_block desc limit 1`, tokenAddress, blockNumber).Scan(&price)
	if err != nil {
		return price, err
	}
	return price, nil
}
