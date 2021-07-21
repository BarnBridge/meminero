package barn

import (
	"context"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/notifications"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func (s *Storable) storeDelegateActions(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.delegateActions) == 0 {
		s.logger.WithField("handler", "delegateActions").Debug("no events found")
		return nil
	}

	var rows [][]interface{}

	for _, d := range s.processed.delegateActions {
		rows = append(rows, []interface{}{
			utils.NormalizeAddress(d.From.String()),
			utils.NormalizeAddress(d.To.String()),
			d.ActionType,
			s.block.BlockCreationTime,
			s.block.Number,
			utils.NormalizeAddress(d.Raw.TxHash.String()),
			d.Raw.TxIndex,
			d.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "barn_delegate_actions"},
		[]string{"sender", "receiver", "action_type", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)

	return err
}

func (s *Storable) storeDelegateChanges(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.delegateChanges) == 0 {
		s.logger.WithField("handler", "delegateChanges").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job
	for _, d := range s.processed.delegateChanges {
		value := decimal.NewFromBigInt(d.Amount, 0)
		newDelegated := decimal.NewFromBigInt(d.ToNewDelegatedPower, 0)
		rows = append(rows, []interface{}{
			d.ActionType,
			d.Sender,
			d.Receiver,
			value,
			newDelegated,
			s.block.BlockCreationTime,
			s.block.Number,
			d.TransactionHash,
			d.TransactionIndex,
			d.LogIndex,
		})
		if d.ActionType == DELEGATE_INCREASE {
			if d.ToNewDelegatedPower.Cmp(d.Amount) == 0 {
				jd := notifications.DelegateJobData{
					StartTime:             s.block.BlockCreationTime,
					From:                  d.Sender,
					To:                    d.Receiver,
					Amount:                decimal.NewFromBigInt(d.ToNewDelegatedPower, 0),
					IncludedInBlockNumber: s.block.Number,
				}
				j, err := notifications.NewDelegateStartJob(&jd)
				if err != nil {
					return errors.Wrap(err, "could not create notification job")
				}

				jobs = append(jobs, j)
			}
		}
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "barn_delegate_changes"},
		[]string{"action_type", "sender", "receiver", "amount", "receiver_new_delegated_power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}
	if config.Store.Storable.Barn.Notifications && len(jobs) > 0 {
		err := notifications.ExecuteJobsWithTx(ctx, tx, jobs...)
		if err != nil && err != context.DeadlineExceeded {
			return errors.Wrap(err, "could not execute notification jobs")
		}
	}

	return nil
}

func (s *Storable) storeLockEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.processed.locks) == 0 {
		s.logger.WithField("handler", "locks").Debug("no events found")
		return nil
	}
	var rows [][]interface{}
	for _, l := range s.processed.locks {
		rows = append(rows, []interface{}{
			utils.NormalizeAddress(l.User.String()),
		})
	}

	return nil
}