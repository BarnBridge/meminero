package smartexposure

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func (s *Storable) storeEPoolTransactions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.seTransactions) == 0 {
		s.logger.WithField("module", "smart exposure epool transactions").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, t := range s.processed.seTransactions {
		amount := decimal.NewFromBigInt(t.Amount, 0)
		amountA := decimal.NewFromBigInt(t.AmountA, 0)
		amountB := decimal.NewFromBigInt(t.AmountB, 0)

		rows = append(rows, []interface{}{
			t.UserAddress,
			t.ETokenAddress,
			amount,
			amountA,
			amountB,
			t.TransactionType,
			s.block.BlockCreationTime,
			s.block.Number,
			t.TxHash,
			t.TxIndex,
			t.LogIndex,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "transaction_history"},
		[]string{"user_address", "etoken_address", "amount", "amount_a", "amount_b", "transaction_type", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	return err
}

func (s *Storable) storeNewTranches(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.newTranches) == 0 {
		s.logger.WithField("module", "smart exposure new tranches").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, t := range s.processed.newTranches {
		factor := decimal.NewFromBigInt(t.SFactorE, 0)
		targetRatio := decimal.NewFromBigInt(t.TargetRatio, 0)

		rows = append(rows, []interface{}{
			t.EPoolAddress,
			t.ETokenAddress,
			t.ETokenSymbol,
			factor,
			targetRatio,
			t.TokenARatio,
			t.TokenBRatio,
			s.block.Number,
		})

		s.state.AddNewTrancheToState(t)
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "tranches"},
		[]string{"pool_address", "etoken_address", "etoken_symbol", "s_factor_e", "target_ratio", "token_a_ratio", "token_b_ratio", "start_at_block"},
		pgx.CopyFromRows(rows),
	)

	return err
}
