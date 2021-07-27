package scrape

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) storeEPoolTransactions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.seTransactions) == 0 {
		s.logger.WithField("module", "smart exposure epool transactions").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	for _, t := range s.processed.seTransactions {
		rows = append(rows, []interface{}{
			t.UserAddress,
			t.ETokenAddress,
			t.Amount,
			t.AmountA,
			t.AmountB,
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
		ratioA, _ := t.TokenARatio.Float64()
		ratioB, _ := t.TokenBRatio.Float64()
		rows = append(rows, []interface{}{
			t.EPoolAddress,
			t.ETokenAddress,
			t.ETokenSymbol,
			t.SFactorE,
			t.TargetRatio,
			ratioA,
			ratioB,
			s.block.Number,
		})

		s.state.SmartExposure.AddNewTrancheToState(t)
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_exposure", "tranches"},
		[]string{"pool_address", "etoken_address", "etoken_symbol", "s_factor_e", "target_ratio", "token_a_ratio", "token_b_ratio", "start_at_block"},
		pgx.CopyFromRows(rows),
	)

	return err
}
