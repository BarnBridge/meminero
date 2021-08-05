package governance

import (
	"context"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/utils"
)

func (s *Storable) Execute(ctx context.Context) error {
	var govLogs []gethtypes.Log
	for _, data := range s.block.Txs {
		for _, log := range data.LogEntries {
			if utils.NormalizeAddress(log.Address.String()) == utils.NormalizeAddress(config.Store.Storable.Governance.Address) {
				govLogs = append(govLogs, log)
			}
		}
	}

	if len(govLogs) == 0 {
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

func (s *Storable) handleProposals(ctx context.Context, logs []gethtypes.Log) error {
	var createdProposals []ethtypes.GovernanceProposalCreatedEvent
	for _, log := range logs {
		if ethtypes.Governance.IsProposalCreatedEvent(&log) {
			p, err := ethtypes.Governance.ProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal created event")
			}

			createdProposals = append(createdProposals, p)
		}
	}

	if len(createdProposals) == 0 {
		return nil
	}

	err := s.getProposalsDetailsFromChain(ctx, createdProposals)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) handleAbrogationProposal(ctx context.Context, logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsAbrogationProposalStartedEvent(&log) {

			cp, err := ethtypes.Governance.AbrogationProposalStartedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal started event")
			}
			s.Processed.abrogationProposals = append(s.Processed.abrogationProposals, cp)
		}
	}

	if len(s.Processed.abrogationProposals) == 0 {
		return nil
	}

	err := s.getAPDescriptionsFromChain(ctx, s.Processed.abrogationProposals)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storable) handleEvents(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsProposalCreatedEvent(&log) {
			e, err := ethtypes.Governance.ProposalCreatedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal created event")
			}
			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				EventType:  CREATED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalQueuedEvent(&log) {
			e, err := ethtypes.Governance.ProposalQueuedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal queued event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				Eta:        e.Eta,
				EventType:  QUEUED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalExecutedEvent(&log) {
			e, err := ethtypes.Governance.ProposalExecutedEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal executed event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				EventType:  EXECUTED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

		if ethtypes.Governance.IsProposalCanceledEvent(&log) {
			e, err := ethtypes.Governance.ProposalCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal canceled event")
			}

			s.Processed.proposalEvents = append(s.Processed.proposalEvents, ProposalEvent{
				ProposalID: e.ProposalId,
				Caller:     e.Caller,
				EventType:  CANCELED,
				BaseLog: BaseLog{
					TransactionHash:  utils.NormalizeAddress(e.Raw.TxHash.String()),
					TransactionIndex: int64(e.Raw.TxIndex),
					LogIndex:         int64(e.Raw.Index),
				},
			})
			continue
		}

	}
	return nil
}

func (s *Storable) handleVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsVoteEvent(&log) {
			vote, err := ethtypes.Governance.VoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote event")
			}

			s.Processed.votes = append(s.Processed.votes, vote)
		}

		if ethtypes.Governance.IsVoteCanceledEvent(&log) {
			vote, err := ethtypes.Governance.VoteCanceledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode proposal vote canceled event")
			}

			s.Processed.canceledVotes = append(s.Processed.canceledVotes, vote)
		}
	}

	return nil
}

func (s *Storable) handleAbrogationProposalVotes(logs []gethtypes.Log) error {
	for _, log := range logs {
		if ethtypes.Governance.IsAbrogationProposalVoteEvent(&log) {
			vote, err := ethtypes.Governance.AbrogationProposalVoteEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}

			s.Processed.abrogationVotes = append(s.Processed.abrogationVotes, vote)
		}

		if ethtypes.Governance.IsAbrogationProposalVoteCancelledEvent(&log) {
			vote, err := ethtypes.Governance.AbrogationProposalVoteCancelledEvent(log)
			if err != nil {
				return errors.Wrap(err, "could not decode abrogation proposal event")
			}

			s.Processed.abrogationCanceledVotes = append(s.Processed.abrogationCanceledVotes, vote)
		}
	}

	return nil
}
