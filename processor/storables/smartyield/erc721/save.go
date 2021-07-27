package erc721

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/processor/storables/smartyield"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.Debug("storing")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done storing")
	}()

	err := s.saveERC721Transfers(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save erc721 transfers")
	}

	err = s.saveTxHistory(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not save transaction history")
	}

	return nil
}

func (s *Storable) saveERC721Transfers(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Transfers) == 0 {
		return nil
	}

	var rows [][]interface{}

	for _, e := range s.processed.Transfers {
		var tokenType string

		if s.state.SmartYield.PoolBySeniorBondAddress(e.Raw.Address.String()) != nil {
			tokenType = "senior"
		} else {
			tokenType = "junior"
		}

		rows = append(rows, []interface{}{
			utils.NormalizeAddress(e.Raw.Address.String()),
			tokenType,
			utils.NormalizeAddress(e.From.String()),
			utils.NormalizeAddress(e.To.String()),
			e.TokenId.Int64(),
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(e.Raw.TxHash.String()),
			e.Raw.TxIndex,
			e.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "erc721_transfers"},
		[]string{
			"token_address",
			"token_type",
			"sender",
			"receiver",
			"token_id",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into erc721_transfers")
	}

	return nil
}

func (s *Storable) saveTxHistory(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.Transfers) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.processed.Transfers {
		// we don't want to store Mint & Burn events in the transaction history table because there will already be another
		// corresponding action (eg: SeniorDeposit)
		if e.From.String() == utils.ZeroAddress || e.To.String() == utils.ZeroAddress {
			continue
		}

		var tokenType string

		if s.state.SmartYield.PoolBySeniorBondAddress(e.Raw.Address.String()) != nil {
			tokenType = "senior"
		} else {
			tokenType = "junior"
		}

		var tokenActionTypeSend, tokenActionTypeReceive smartyield.TxType
		var amount decimal.Decimal
		if tokenType == "junior" {
			tokenActionTypeSend = smartyield.JbondSend
			tokenActionTypeReceive = smartyield.JbondReceive

			err := tx.QueryRow(
				ctx,
				`select tokens_in from smart_yield.junior_2step_withdraw_events where junior_bond_address = $1 and junior_bond_id = $2`,
				utils.NormalizeAddress(e.Raw.Address.String()), e.TokenId.Int64(),
			).Scan(&amount)
			if err != nil {
				return errors.Wrap(err, "could not find JuniorBond by id in the database")
			}
		} else {
			tokenActionTypeSend = smartyield.SbondSend
			tokenActionTypeReceive = smartyield.SbondReceive

			err := tx.QueryRow(
				ctx,
				`select underlying_in from smart_yield.senior_entry_events where senior_bond_address = $1 and senior_bond_id = $2`,
				utils.NormalizeAddress(e.Raw.Address.String()), e.TokenId.Int64(),
			).Scan(&amount)
			if err != nil {
				return errors.Wrap(err, "could not find SeniorBond by id in the database")
			}
		}

		rows = append(rows, s.txHistory(e.From.String(), amount, strings.ToUpper(tokenType), tokenActionTypeSend, e.Raw))
		rows = append(rows, s.txHistory(e.To.String(), amount, strings.ToUpper(tokenType), tokenActionTypeReceive, e.Raw))
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"smart_yield", "transaction_history"},
		[]string{
			"protocol_id",
			"pool_address",
			"underlying_token_address",
			"user_address",
			"amount",
			"tranche",
			"transaction_type",
			"block_timestamp",
			"included_in_block",
			"tx_hash",
			"tx_index",
			"log_index",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not copy into transaction_history")
	}

	return nil
}
