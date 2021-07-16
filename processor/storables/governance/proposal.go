package governance

import (
	"context"
	"encoding/hex"

	abiStore "github.com/barnbridge/smartbackend/abi"
	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (g *GovStorable) handleProposals(ctx context.Context,logs []gethtypes.Log) error {
	var createdProposals []ethtypes.GovernanceProposalCreatedEvent
	for _, log := range logs {
		if 	ethtypes.Governance.IsGovernanceProposalCreatedEvent(&log) {
			p ,err :=ethtypes.Governance.GovernanceProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err,"could not decode proposal created event")
			}

			createdProposals = append(createdProposals, p)
		}
	}

	if len(createdProposals) == 0 {
		g.logger.WithField("handler", "proposals").Debug("no events found")
		return nil
	}

	err := g.getProposalsDetailsFromChain(ctx,createdProposals)
	if err != nil {
		return err
	}


	//var jobs []*notifications.Job
	/*
	stmt, err := tx.Prepare(pq.CopyIn("governance_proposals", "proposal_id", "proposer", "description", "title", "create_time", "targets", "values", "signatures", "calldatas", "warm_up_duration", "active_duration", "queue_duration", "grace_period_duration", "acceptance_threshold", "min_quorum", "included_in_block", "block_timestamp"))
	if err != nil {
		return errors.Wrap(err, "could not prepare statement")
	}

	for i, p := range proposals {
		a := actions[i]
		var targets, values, signatures, calldatas types.JSONStringArray

		for i := 0; i < len(a.Targets); i++ {
			targets = append(targets, a.Targets[i].String())
			values = append(values, a.Values[i].String())
			signatures = append(signatures, a.Signatures[i])
			calldatas = append(calldatas, hex.EncodeToString(a.Calldatas[i]))
		}

		_, err = stmt.Exec(p.Id.Int64(), p.Proposer.String(), p.Description, p.Title, p.CreateTime.Int64(), targets, values, signatures, calldatas, p.WarmUpDuration.Int64(), p.ActiveDuration.Int64(), p.QueueDuration.Int64(), p.GracePeriodDuration.Int64(), p.AcceptanceThreshold.Int64(), p.MinQuorum.Int64(), g.Preprocessed.BlockNumber, g.Preprocessed.BlockTimestamp)
		if err != nil {
			return errors.Wrap(err, "could not execute statement")
		}

		jd := notifications.ProposalCreatedJobData{
			Id:                    p.Id.Int64(),
			Proposer:              p.Proposer.String(),
			Title:                 p.Title,
			CreateTime:            p.CreateTime.Int64(),
			WarmUpDuration:        p.WarmUpDuration.Int64(),
			ActiveDuration:        p.ActiveDuration.Int64(),
			QueueDuration:         p.QueueDuration.Int64(),
			GraceDuration:         p.GracePeriodDuration.Int64(),
			IncludedInBlockNumber: g.Preprocessed.BlockNumber,
		}
		j, err := notifications.NewProposalCreatedJob(&jd)
		if err != nil {
			return errors.Wrap(err, "could not create notification job")
		}

		jobs = append(jobs, j)
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return errors.Wrap(err, "could not close statement")
	}

	if g.config.Notifications {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		err = notifications.ExecuteJobsWithTx(ctx, tx, jobs...)
		if err != nil && err != context.DeadlineExceeded {
			return errors.Wrap(err, "could not execute notification jobs")
		}
	}*/

	return nil
}

func (g *GovStorable) getProposalsDetailsFromChain(ctx context.Context, createdEvents []ethtypes.GovernanceProposalCreatedEvent) error {
	for _, p := range createdEvents {
		p := p
		a,err := abiStore.Get("governance")
		if err != nil {
			return errors.Wrap(err,"could not find governance abi")
		}
		wg, _ := errgroup.WithContext(ctx)
		var proposal Proposal
		var proposalAction ProposalActions
		wg.Go(eth.CallContractFunction(*a,utils.NormalizeAddress(config.Store.Storable.Governance.GovernanceAddress),"proposals", []interface{}{p.ProposalId},&proposal))
		wg.Go(eth.CallContractFunction(*a,utils.NormalizeAddress(config.Store.Storable.Governance.GovernanceAddress),"getActions", []interface{}{p.ProposalId},&proposalAction))
		err = wg.Wait()
		if err != nil {
			return errors.Wrap(err,"could not get proposals info")
		}
		g.Processed.proposals = append(g.Processed.proposals, proposal)
		g.Processed.proposalsActions = append(g.Processed.proposalsActions, proposalAction)

	}
	return nil
}

func (g *GovStorable) storeProposals(ctx context.Context,tx pgx.Tx) error {
	var rows [][]interface{}
	for i,p := range g.Processed.proposals {

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
			p.WarmUpDuration.Int64(),
			p.ActiveDuration.Int64(),
			p.QueueDuration.Int64(),
			p.GracePeriodDuration.Int64(),
			p.AcceptanceThreshold.Int64(),
			p.MinQuorum.Int64(),
			g.block.BlockCreationTime,
			g.block.Number,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"proposals"},
		[]string{"proposal_id", "proposer","description", "title", "create_time", "targets", "values", "signatures", "calldatas", "warm_up_duration", "active_duration", "queue_duration", "grace_period_duration", "acceptance_threshold", "min_quorum", "included_in_block", "block_timestamp"},
		pgx.CopyFromSlice(len(g.Processed.proposals), func(i int) ([]interface{}, error) {
			return []interface{}{rows}, nil
		}),
	)

	return err
}