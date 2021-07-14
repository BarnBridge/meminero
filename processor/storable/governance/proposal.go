package governance

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	web3types "github.com/alethio/web3-go/types"
	"github.com/barnbridge/smartbackend/contracts"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"

	"github.com/barnbridge/smartbackend/utils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (g *GovStorable) handleProposals(logs []web3types.Log, tx pgx.Tx,governanceDecoder *ethtypes.GovernanceDecoder) error {
	var createdProposals []ethtypes.GovernanceProposalCreatedEvent
	for _, log := range logs {
		if 	governanceDecoder.IsGovernanceProposalCreatedEvent(log) {
			p ,err :=governanceDecoder.GovernanceProposalCreatedEventW3(log)
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

	var wg = &errgroup.Group{}
	var mu = &sync.Mutex{}
	var proposals []Proposal
	var actions []ProposalActions
	g.getProposalsDetails(wg,mu,createdProposals,proposals,actions)
	err := wg.Wait()
	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"proposals"},
		[]string{"proposal_id", "proposer","description", "title", "create_time", "targets", "values", "signatures", "calldatas", "warm_up_duration", "active_duration", "queue_duration", "grace_period_duration", "acceptance_threshold", "min_quorum", "included_in_block", "block_timestamp"},
		pgx.CopyFromSlice(len(proposals), func(i int) ([]interface{}, error) {
			return []interface{}{proposals[i].ProposalId.String(),proposals[i].Caller.String(),
				g.Preprocessed.BlockTimestamp,proposals[aps[i].ProposalId.String()],proposals[i].Raw.BlockHash.String(),
				aps[i].Raw.TxIndex,aps[i].Raw.TxIndex,g.Preprocessed.BlockNumber}, nil
		}),
	)
	if err != nil {
		return errors.Wrap(err,"could not store abrogration_proposals")
	}
	stmt, err := tx.Prepare(context.Background(),pq.CopyIn("governance_proposals", "proposal_id", "proposer", "description", "title", "create_time", "targets", "values", "signatures", "calldatas", "warm_up_duration", "active_duration", "queue_duration", "grace_period_duration", "acceptance_threshold", "min_quorum", "included_in_block", "block_timestamp"))
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

func (g *GovStorable) getProposalsDetails(wg *errgroup.Group, mu *sync.Mutex,proposals []ethtypes.GovernanceProposalCreatedEvent, results []Proposal,actions []ProposalActions) {
	for _, p := range proposals {
		p := p
		wg.Go(func() error {
			input, err := utils.ABIGenerateInput(g.govAbi, "proposals",p.ProposalId )
			if err != nil {
				return err
			}
			data, err := utils.CallAtBlock(g.ethRPC,g.config.GovernanceAddress , input, g.Preprocessed.BlockNumber)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not call %g.%g", g.config.GovernanceAddress, "proposals"))
			}

			decoded, err := utils.DecodeFunctionOutput(g.govAbi, "proposals", data)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not decode output from %g.%g", g.config.GovernanceAddress, "proposals"))
			}

			mu.Lock()
			results = append(results, Proposal{
				Id: decoded["id"].(*big.Int),
				Proposer: decoded["proposer"].(common.Address),
				Description: decoded["description"].(string),
				Title: decoded["title"].(string),
				CreateTime: decoded["createTime"].(*big.Int),
				Eta: decoded["eta"].(*big.Int),
				ForVotes: decoded["forVotes"].(*big.Int),
				AgainstVotes: decoded["againstVotes"].(*big.Int),
				Canceled: decoded["canceled"].(bool),
				Executed: decoded["executed"].(bool),
				ProposalParameters: ProposalParameters{
					WarmUpDuration: decoded["warmUpDuration"].(*big.Int),
					ActiveDuration: decoded["activeDuration"].(*big.Int),
					QueueDuration: decoded["queueDuration"].(*big.Int),
					GracePeriodDuration: decoded["gracePeriodDuration"].(*big.Int),
					AcceptanceThreshold: decoded["acceptanceThreshold"].(*big.Int),
					MinQuorum: decoded["minQuorum"].(*big.Int),
				}})
			mu.Unlock()
			return nil
		})
		wg.Go(func() error {

			input, err := utils.ABIGenerateInput(g.govAbi, "getActions",p.ProposalId )
			if err != nil {
				return err
			}
			data, err := utils.CallAtBlock(g.ethRPC,g.config.GovernanceAddress , input, g.Preprocessed.BlockNumber)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not call %g.%g", g.config.GovernanceAddress, "getActions"))
			}

			decoded, err := utils.DecodeFunctionOutput(g.govAbi, "proposals", data)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not decode output from %g.%g", g.config.GovernanceAddress, "getActions"))
			}
			mu.Lock()
			actions = append(actions, ProposalActions{
				Targets: decoded["targets"].([]common.Address),
				Values: decoded["values"].([]*big.Int),
				Signatures: decoded["signatures"].([]string),
				Calldatas: decoded["calldatas"].([][]byte),
			})
			mu.Unlock()
			return nil
		})
	}
}