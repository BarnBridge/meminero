package governance

import (
	"context"
	"fmt"
	"time"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/ethtypes"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GovStorable struct {
	block  *types.Block
	logger *logrus.Entry

	Processed struct {
		proposals                      []Proposal
		proposalsActions               []ProposalActions
		abrogationProposals            []ethtypes.GovernanceAbrogationProposalStartedEvent
		abrogationProposalsDescription map[string]string
		proposalEvents                 []ProposalEvent
		votes                          []ethtypes.GovernanceVoteEvent
		canceledVotes                  []ethtypes.GovernanceVoteCanceledEvent
		abrogationVotes                []ethtypes.GovernanceAbrogationProposalVoteEvent
		abrogationCanceledVotes        []ethtypes.GovernanceAbrogationProposalVoteCancelledEvent
	}
}

func New(block *types.Block) *GovStorable {
	return &GovStorable{
		block:  block,
		logger: logrus.WithField("module", "storable(governance)"),
	}
}

func (g *GovStorable) Execute(ctx context.Context) error {
	g.logger.Trace("executing")
	start := time.Now()
	defer func() {
		g.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var govLogs []gethtypes.Log
	for _, data := range g.block.Txs {
		for _, log := range data.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Governance.Address) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
		g.logger.WithField("handler", "governance").Debug("no events found")
		return nil
	}

	err := g.handleProposals(ctx, govLogs)
	if err != nil {
		return err
	}

	err = g.handleAbrogationProposal(ctx, govLogs)
	if err != nil {
		return err
	}

	err = g.handleEvents(govLogs)
	if err != nil {
		return err
	}

	err = g.handleVotes(govLogs)
	if err != nil {
		return err
	}

	err = g.handleAbrogationProposalVotes(govLogs)
	if err != nil {
		return err
	}

	return nil
}

func (g *GovStorable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := [7]string{"proposals", "abrogation_proposals", "proposal_events", "votes", "votes_canceled", "abrogation_votes", "abrogation_votes_canceled"}
	for _, t := range tables {
		s := fmt.Sprintf(`delete from governance.%s where included_in_block = $1`, t)
		b.Queue(s, g.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	err = br.Close()
	return err
}

func (g *GovStorable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := g.storeProposals(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals")
	}

	err = g.storeAbrogrationProposals(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogration proposals")
	}

	err = g.storeEvents(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals events")
	}

	err = g.storeProposalVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposal's votes")
	}

	err = g.storeProposalCanceledVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposal's  canceled votes")
	}

	err = g.storeProposalAbrogationVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal's votes")
	}

	err = g.storeAbrogationProposalCanceledVotes(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store abrogation proposal's  canceled votes")
	}

	return nil
}

func (g *GovStorable) Result() interface{} {
	return g.Processed
}
