package governance

import (
	"context"
	"sync"
	"time"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/notifications"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (g *GovStorable) handleAbrogationProposal(ctx context.Context, logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceAbrogationProposalStartedEvent(&log) {

			cp, err := ethtypes.Governance.GovernanceAbrogationProposalStartedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal started event")
			}
			g.Processed.abrogationProposals = append(g.Processed.abrogationProposals, cp)
		}
	}

	if len(g.Processed.abrogationProposals) == 0 {
		g.logger.WithField("handler", "abrogation proposal").Debug("no events found")
		return nil
	}

	err := g.getAPDescriptionsFromChain(ctx, g.Processed.abrogationProposals)
	if err != nil {
		return err
	}

	return nil
}

func (g *GovStorable) getAPDescriptionsFromChain(ctx context.Context, aps []ethtypes.GovernanceAbrogationProposalStartedEvent) error {
	wg, _ := errgroup.WithContext(ctx)
	var mu sync.Mutex
	for _, ap := range aps {
		ap := ap
		a := ethtypes.Governance.ABI
		wg.Go(func() error {
			var description string
			subwg, _ := errgroup.WithContext(ctx)
			subwg.Go(eth.CallContractFunction(*a, utils.NormalizeAddress(config.Store.Storable.Governance.Address), "abrogationProposals", []interface{}{ap.ProposalId}, &description))
			err := wg.Wait()
			if err != nil {
				return errors.Wrap(err, "")
			}
			mu.Lock()
			g.Processed.abrogationProposalsDescription[ap.ProposalId.String()] = description
			mu.Unlock()

			return nil
		})

	}
	return nil
}

func (g *GovStorable) storeAbrogrationProposals(ctx context.Context, tx pgx.Tx) error {
	if len(g.Processed.abrogationProposals) == 0 {
		g.logger.WithField("handler", "abrogation proposal").Debug("no events found")
		return nil
	}

	var rows [][]interface{}
	var jobs []*notifications.Job
	for _, ap := range g.Processed.abrogationProposals {
		rows = append(rows, []interface{}{
			ap.ProposalId.Int64(),
			ap.Caller.String(),
			g.block.BlockCreationTime,
			g.Processed.abrogationProposalsDescription[ap.ProposalId.String()],
			ap.Raw.TxHash.String(),
			ap.Raw.TxIndex,
			ap.Raw.Index,
			ap.Raw.BlockNumber,
		})

		jd := notifications.AbrogationProposalCreatedJobData{
			Id:                    ap.ProposalId.Int64(),
			Proposer:              ap.Caller.String(),
			CreateTime:            g.block.BlockCreationTime,
			IncludedInBlockNumber: g.block.Number,
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
