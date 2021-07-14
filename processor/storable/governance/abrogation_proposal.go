package governance

import (
	"context"
	"fmt"
	"sync"

	web3types "github.com/alethio/web3-go/types"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (g *GovStorable) handleAbrogationProposal(logs []web3types.Log, tx pgx.Tx,governanceDecoder *ethtypes.GovernanceDecoder) error {
	var aps []ethtypes.GovernanceAbrogationProposalStartedEvent

	for _, log := range logs {
		if governanceDecoder.IsGovernanceAbrogationProposalStartedEvent(log) {
			cp ,err := governanceDecoder.GovernanceAbrogationProposalStartedEvent(log)
			if err != nil {
				return errors.Wrap(err,"could not decode abrogation proposal started event")
			}
			aps = append(aps, cp)
		}
	}

	if len(aps) == 0 {
		g.logger.WithField("handler", "abrogation proposal").Debug("no events found")
		return nil
	}

	var wg = &errgroup.Group{}
	var mu = &sync.Mutex{}
	apsDescription := make(map[string]string)
	g.getAPDescriptions(wg,mu,aps,apsDescription)
	err := wg.Wait()
	if err != nil {
		return err
	}


	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"abrogation_proposals"},
		[]string{"proposal_id", "creator","create_time","description","tx_hash","tx_index","log_index","included_in_block"},
		pgx.CopyFromSlice(len(aps), func(i int) ([]interface{}, error) {
			return []interface{}{aps[i].ProposalId.String(),aps[i].Caller.String(), g.Preprocessed.BlockTimestamp,apsDescription[aps[i].ProposalId.String()],aps[i].Raw.BlockHash.String(),aps[i].Raw.TxIndex,aps[i].Raw.TxIndex,g.Preprocessed.BlockNumber}, nil
		}),
	)
	if err != nil {
		return errors.Wrap(err,"could not store abrogration_proposals")
	}
	//var jobs []*notifications.Job
/*	for _, cp := range aps {
		_, err = stmt.Exec(cp.ProposalID.Int64(), cp.Caller.String(), cp.CreateTime, cp.Description, cp.TransactionHash, cp.TransactionIndex, cp.LogIndex, cp.LoggedBy, g.Preprocessed.BlockNumber)
		if err != nil {
			return errors.Wrap(err, "could not execute statement")
		}

		jd := notifications.AbrogationProposalCreatedJobData{
			Id:                    cp.ProposalID.Int64(),
			Proposer:              cp.Caller.String(),
			CreateTime:            cp.CreateTime,
			IncludedInBlockNumber: g.Preprocessed.BlockNumber,
		}
		j, err := notifications.NewAbrogationProposalCreatedJob(&jd)
		if err != nil {
			return errors.Wrap(err, "could not create notification job")
		}

		jobs = append(jobs, j)
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

func (g *GovStorable) getAPDescriptions(wg *errgroup.Group,  mu *sync.Mutex,aps []ethtypes.GovernanceAbrogationProposalStartedEvent, results map[string]string) {
	for _, ap :=range aps {
		ap := ap
		wg.Go(func() error {
			input, err := utils.ABIGenerateInput(g.govAbi, "abrogationProposals", ap.ProposalId)
			if err != nil {
				return err
			}
			data, err := utils.CallAtBlock(g.ethRPC,g.config.GovernanceAddress , input, g.Preprocessed.BlockNumber)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not call %g.%g", g.config.GovernanceAddress, "abrogationProposals"))
			}

			decoded, err := utils.DecodeFunctionOutput(g.govAbi, "abrogationProposals", data)
			if err != nil {
				return  errors.Wrap(err, fmt.Sprintf("could not decode output from %g.%g", g.config.GovernanceAddress, "abrogationProposals"))
			}
			mu.Lock()
			results[ap.ProposalId.String()] = decoded["description"].(string)
			mu.Unlock()
			return nil
		})
	}
}