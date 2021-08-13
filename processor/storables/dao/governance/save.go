package governance

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/notifications"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeProposals(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals")
	}

	err = s.storeAbrogationProposals(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogration proposals")
	}

	err = s.storeEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals events")
	}

	err = s.storeProposalVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposal's votes")
	}

	err = s.storeProposalCanceledVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposal's  canceled votes")
	}

	err = s.storeProposalAbrogationVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal's votes")
	}

	err = s.storeAbrogationProposalCanceledVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal's  canceled votes")
	}

	return nil
}

func (s *Storable) storeProposals(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.proposals) == 0 {
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job
	for i, p := range s.Processed.proposals {
		var targets, values, signatures, calldatas types.JSONStringArray

		a := s.Processed.proposalsActions[i]
		for i := 0; i < len(a.Targets); i++ {
			targets = append(targets, a.Targets[i].String())
			values = append(values, a.Values[i].String())
			signatures = append(signatures, a.Signatures[i])
			calldatas = append(calldatas, hex.EncodeToString(a.Calldatas[i]))
		}

		rows = append(rows, []interface{}{
			p.Id.Int64(),
			utils.NormalizeAddress(p.Proposer.String()),
			p.Description,
			p.Title,
			p.CreateTime.Int64(),
			targets,
			values,
			signatures,
			calldatas,
			p.Parameters.WarmUpDuration.Int64(),
			p.Parameters.ActiveDuration.Int64(),
			p.Parameters.QueueDuration.Int64(),
			p.Parameters.GracePeriodDuration.Int64(),
			p.Parameters.AcceptanceThreshold.Int64(),
			p.Parameters.MinQuorum.Int64(),
			s.block.Number,
			s.block.BlockCreationTime,
		})

		jd := notifications.ProposalCreatedJobData{
			Id:                    p.Id.Int64(),
			Proposer:              utils.NormalizeAddress(p.Proposer.String()),
			Title:                 p.Title,
			CreateTime:            p.CreateTime.Int64(),
			WarmUpDuration:        p.Parameters.WarmUpDuration.Int64(),
			ActiveDuration:        p.Parameters.ActiveDuration.Int64(),
			QueueDuration:         p.Parameters.QueueDuration.Int64(),
			GraceDuration:         p.Parameters.GracePeriodDuration.Int64(),
			IncludedInBlockNumber: s.block.Number,
		}

		j, err := notifications.NewProposalCreatedJob(&jd)
		if err != nil {
			return errors.Wrap(err, "could not create notification job")
		}

		jobs = append(jobs, j)
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "proposals"},
		[]string{"proposal_id", "proposer", "description", "title", "create_time", "targets", "values", "signatures", "calldatas", "warm_up_duration", "active_duration", "queue_duration", "grace_period_duration", "acceptance_threshold", "min_quorum", "included_in_block", "block_timestamp"},
		pgx.CopyFromRows(rows),
	)

	if config.Store.Storable.Governance.Notifications {
		err := notifications.ExecuteJobsWithTx(ctx, tx, jobs...)
		if err != nil && err != context.DeadlineExceeded {
			return errors.Wrap(err, "could not execute notification jobs")
		}
	}

	return err
}

func (s *Storable) storeAbrogationProposals(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.abrogationProposals) == 0 {
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job
	for _, ap := range s.Processed.abrogationProposals {
		rows = append(rows, []interface{}{
			ap.ProposalId.Int64(),
			utils.NormalizeAddress(ap.Caller.String()),
			s.block.BlockCreationTime,
			s.Processed.abrogationProposalsDescription[ap.ProposalId.String()],
			utils.NormalizeAddress(ap.Raw.TxHash.String()),
			ap.Raw.TxIndex,
			ap.Raw.Index,
			ap.Raw.BlockNumber,
		})

		jd := notifications.AbrogationProposalCreatedJobData{
			Id:                    ap.ProposalId.Int64(),
			Proposer:              utils.NormalizeAddress(ap.Caller.String()),
			CreateTime:            s.block.BlockCreationTime,
			IncludedInBlockNumber: s.block.Number,
		}
		j, err := notifications.NewAbrogationProposalCreatedJob(&jd)
		if err != nil {
			return errors.Wrap(err, "could not create notification job")
		}

		jobs = append(jobs, j)
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "abrogation_proposals"},
		[]string{"proposal_id", "creator", "create_time", "description", "tx_hash", "tx_index", "log_index", "included_in_block"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogration_proposals")
	}

	if config.Store.Storable.Governance.Notifications {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		err = notifications.ExecuteJobsWithTx(ctx, tx, jobs...)
		if err != nil && err != context.DeadlineExceeded {
			return errors.Wrap(err, "could not execute notification jobs")
		}
	}

	return nil
}

func (s *Storable) storeEvents(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.proposalEvents) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, e := range s.Processed.proposalEvents {
		var eventData types.JSONObject
		if e.Eta != nil {
			eventData = make(types.JSONObject)
			eventData["eta"] = e.Eta.Int64()
		}

		rows = append(rows, []interface{}{
			e.ProposalID.Int64(),
			utils.NormalizeAddress(e.Caller.String()),
			e.EventType,
			eventData,
			s.block.BlockCreationTime,
			s.block.Number,
			e.TransactionHash,
			e.TransactionIndex,
			e.LogIndex,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "proposal_events"},
		[]string{"proposal_id", "caller", "event_type", "event_data", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return errors.Wrap(err, "could not store proposal events")
	}

	return nil
}

func (s *Storable) storeProposalVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.votes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.votes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			v.Support,
			v.PowerDecimal(0),
			s.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "votes"},
		[]string{"proposal_id", "user_id", "support", "power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromSlice(len(s.Processed.votes), func(i int) ([]interface{}, error) {
			return rows[i], nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "could not store proposal votes")
	}

	return nil
}

func (s *Storable) storeProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.canceledVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.canceledVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			s.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}
	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "votes_canceled"},
		[]string{"proposal_id", "user_id", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store proposal canceled votes")
	}

	return nil
}

func (s *Storable) storeProposalAbrogationVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.abrogationVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.abrogationVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			v.Support,
			v.PowerDecimal(0),
			s.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "abrogation_votes"},
		[]string{"proposal_id", "user_id", "support", "power", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal  votes")
	}

	return nil
}

func (s *Storable) storeAbrogationProposalCanceledVotes(ctx context.Context, tx pgx.Tx) error {
	if len(s.Processed.abrogationCanceledVotes) == 0 {
		return nil
	}

	var rows [][]interface{}
	for _, v := range s.Processed.abrogationCanceledVotes {
		rows = append(rows, []interface{}{
			v.ProposalId.Int64(),
			utils.NormalizeAddress(v.User.String()),
			s.block.BlockCreationTime,
			v.Raw.BlockNumber,
			utils.NormalizeAddress(v.Raw.TxHash.String()),
			v.Raw.TxIndex,
			v.Raw.Index,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"governance", "abrogation_votes_canceled"},
		[]string{"proposal_id", "user_id", "block_timestamp", "included_in_block", "tx_hash", "tx_index", "log_index"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal canceled votes")
	}

	return nil
}
