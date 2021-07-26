package governance

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/notifications"
	"github.com/barnbridge/meminero/utils"
)

func (s *GovStorable) handleAbrogationProposal(ctx context.Context, logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsGovernanceAbrogationProposalStartedEvent(&log) {

			cp, err := ethtypes.Governance.GovernanceAbrogationProposalStartedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal started event")
			}
			s.Processed.abrogationProposals = append(s.Processed.abrogationProposals, cp)
		}
	}

	if len(s.Processed.abrogationProposals) == 0 {
		s.logger.WithField("handler", "abrogation proposal").Debug("no events found")
		return nil
	}

	err := s.getAPDescriptionsFromChain(ctx, s.Processed.abrogationProposals)
	if err != nil {
		return err
	}

	return nil
}

func (s *GovStorable) getAPDescriptionsFromChain(ctx context.Context, aps []ethtypes.GovernanceAbrogationProposalStartedEvent) error {
	a := ethtypes.Governance.ABI

	type Response struct {
		Creator      common.Address
		CreateTime   *big.Int
		Description  string
		ForVotes     *big.Int
		AgainstVotes *big.Int
	}

	for _, ap := range aps {
		var resp Response

		err := eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "abrogationProposals", []interface{}{ap.ProposalId}, &resp)()
		if err != nil {
			return errors.Wrap(err, "could not call governance.abrogationProposals")
		}

		if s.Processed.abrogationProposalsDescription == nil {
			s.Processed.abrogationProposalsDescription = make(map[string]string)
		}

		s.Processed.abrogationProposalsDescription[ap.ProposalId.String()] = resp.Description
	}

	return nil
}

func (s *GovStorable) storeAbrogrationProposals(ctx context.Context, tx pgx.Tx) error {
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
