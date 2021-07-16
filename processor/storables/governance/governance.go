package governance

import (
	"context"
	"time"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GovStorable struct {
	block  *types.Block

	logger              *logrus.Entry

	Processed struct {
		proposals []Proposal
		proposalsActions []ProposalActions
		abrProposals []ethtypes.GovernanceAbrogationProposalStartedEvent
		abrProposalsDescription map[string]string
		proposalEvents []ProposalEvent
	}
}

func New(block *types.Block) *GovStorable {
	return &GovStorable{
		block:   block,
		logger:  logrus.WithField("module", "storable(governance)"),
	}
}

func (g GovStorable) Execute(ctx context.Context) error {
	g.logger.Trace("executing")
	start := time.Now()
	defer func() {
		g.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var govLogs []gethtypes.Log
	for _, data := range g.block.Txs {
		for _, log := range data.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Governance.GovernanceAddress) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
		log.Debug("no events found")
		return nil
	}

	err := g.handleProposals(ctx,govLogs)
	if err != nil {
		return err
	}


	err = g.handleAbrogationProposal(ctx,govLogs)
	if err != nil {
		return err
	}

	err = g.handleEvents(govLogs)
	if err != nil {
		return err
	}

/*	err = g.handleVotes(govLogs, tx)
	if err != nil {
		return err
	}


	err = g.handleAbrogationProposalVotes(govLogs, tx)
	if err != nil {
		return err
	}*/

	return nil
}

func (g *GovStorable) Rollback(ctx context.Context,tx pgx.Tx) error {
	_, err := tx.Exec(ctx, `delete from proposals where included_in_block = $1`, g.block.Number)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `delete from abrogation_proposals where included_in_block = $1`, g.block.Number)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `delete from proposal_events where included_in_block = $1`, g.block.Number)
	if err != nil {
		return err
	}


	return err
}

func (g *GovStorable) SaveToDatabase(ctx context.Context,tx pgx.Tx) error {
	err := g.storeProposals(ctx,tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals")
	}

	err = g.storeAbrogrationProposals(ctx,tx)
	if err != nil {
		return errors.Wrap(err,"could not store abrogration proposals")
	}

	err = g.storeEvents(ctx,tx)
	if err != nil {
		return errors.Wrap(err,"could not store proposals events")
	}

	return nil
}

func (g *GovStorable) Result() interface{} {
	return g.Processed
}