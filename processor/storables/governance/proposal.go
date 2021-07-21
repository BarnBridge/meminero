package governance

import (
	"context"
	"encoding/hex"
	"sync"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/notifications"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (g *GovStorable) handleProposals(ctx context.Context, logs []gethtypes.Log) error {
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
		g.logger.WithField("handler", "proposals").Debug("no events found")
		return nil
	}

	err := g.getProposalsDetailsFromChain(ctx, createdProposals)
	if err != nil {
		return err
	}

	return nil
}

func (g *GovStorable) getProposalsDetailsFromChain(ctx context.Context, createdEvents []ethtypes.GovernanceProposalCreatedEvent) error {
	wg, _ := errgroup.WithContext(ctx)
	var mu sync.Mutex
	for _, p := range createdEvents {
		p := p
		a := ethtypes.Governance.ABI
		wg.Go(func() error {
			subwg, _ := errgroup.WithContext(ctx)

			var proposal Proposal
			var proposalAction ProposalActions
			subwg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "proposals", []interface{}{p.ProposalId}, &proposal))
			subwg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "getActions", []interface{}{p.ProposalId}, &proposalAction))
			err := subwg.Wait()
			if err != nil {
				return err
			}

			mu.Lock()
			g.Processed.proposals = append(g.Processed.proposals, proposal)
			g.Processed.proposalsActions = append(g.Processed.proposalsActions, proposalAction)
			mu.Unlock()

			return nil
		})

	}

	err := wg.Wait()
	if err != nil {
		return errors.Wrap(err, "could not get proposals info")
	}
	return nil
}

func (g *GovStorable) storeProposals(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.proposals) == 0 {
		g.logger.WithField("handler", "proposals").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job
	for i, p := range g.Processed.proposals {

		var targets, values, signatures, calldatas types.JSONStringArray
		a := g.Processed.proposalsActions[i]
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
			g.block.BlockCreationTime,
			g.block.Number,
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
			IncludedInBlockNumber: g.block.Number,
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
