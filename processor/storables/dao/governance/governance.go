package governance

import (
	"context"
	"fmt"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
	"github.com/barnbridge/meminero/utils"
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

func (s *GovStorable) Execute(ctx context.Context) error {
	s.logger.Trace("executing")
	start := time.Now()
	defer func() {
		s.logger.WithField("duration", time.Since(start)).
			Trace("done")
	}()

	var govLogs []gethtypes.Log
	for _, data := range s.block.Txs {
		for _, log := range data.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Governance.Address) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
		s.logger.WithField("handler", "governance").Debug("no events found")
		return nil
	}

	err := s.handleProposals(ctx, govLogs)
	if err != nil {
		return err
	}

	err = s.handleAbrogationProposal(ctx, govLogs)
	if err != nil {
		return err
	}

	err = s.handleEvents(govLogs)
	if err != nil {
		return err
	}

	err = s.handleVotes(govLogs)
	if err != nil {
		return err
	}

	err = s.handleAbrogationProposalVotes(govLogs)
	if err != nil {
		return err
	}

	return nil
}

func (s *GovStorable) Rollback(ctx context.Context, tx pgx.Tx) error {
	start := time.Now()
	s.logger.WithField("block", s.block.Number).Debug("rolling back block")
	defer func() {
		s.logger.WithField("duration", time.Since(start)).Debug("done rolling back block")
	}()

	b := &pgx.Batch{}
	tables := []string{"proposals", "abrogation_proposals", "proposal_events", "votes", "votes_canceled", "abrogation_votes", "abrogation_votes_canceled"}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from governance.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	err = br.Close()
	return err
}

func (s *GovStorable) SaveToDatabase(ctx context.Context, tx pgx.Tx) error {
	err := s.storeProposals(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "could not store proposals")
	}

	err = s.storeAbrogrationProposals(ctx, tx)
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

func (s *GovStorable) Result() interface{} {
	return s.Processed
}
