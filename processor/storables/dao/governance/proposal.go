package governance

import (
	"context"
	"encoding/hex"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/notifications"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
)

func (s *GovStorable) handleProposals(ctx context.Context, logs []gethtypes.Log) error {
	var createdProposals []ethtypes.GovernanceProposalCreatedEvent
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceProposalCreatedEvent(&log) {
			p, err := ethtypes.Governance.GovernanceProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal created event")
			}

			createdProposals = append(createdProposals, p)
		}
	}

	if len(createdProposals) == 0 {
		s.logger.WithField("handler", "proposals").Debug("no events found")
		return nil
	}

	err := s.getProposalsDetailsFromChain(ctx, createdProposals)
	if err != nil {
		return err
	}

	return nil
}

func (s *GovStorable) getProposalsDetailsFromChain(ctx context.Context, createdEvents []ethtypes.GovernanceProposalCreatedEvent) error {
	a := ethtypes.Governance.ABI

	wg, _ := errgroup.WithContext(ctx)

	for _, p := range createdEvents {
		p := p

		var proposal Proposal
		var proposalAction ProposalActions
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "proposals", []interface{}{p.ProposalId}, &proposal))
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "getActions", []interface{}{p.ProposalId}, &proposalAction))

		err := wg.Wait()
		if err != nil {
			return err
		}

		s.Processed.proposals = append(s.Processed.proposals, proposal)
		s.Processed.proposalsActions = append(s.Processed.proposalsActions, proposalAction)
	}

	return nil
}

func (s *GovStorable) storeProposals(ctx context.Context, tx pgx.Tx) error {
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
