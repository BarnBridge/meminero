package governance

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/ethtypes"
)

func (s *Storable) getProposalsDetailsFromChain(ctx context.Context, createdEvents []ethtypes.GovernanceProposalCreatedEvent) error {
	a := ethtypes.Governance.ABI

	wg, _ := errgroup.WithContext(ctx)

	for _, p := range createdEvents {
		p := p

		var proposal Proposal
		var proposalAction ProposalActions
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "proposals", []interface{}{p.ProposalId}, &proposal))
		wg.Go(eth.CallContractFunction(*a, config.Store.Storable.Governance.Address, "getActions", []interface{}{p.ProposalId}, &proposalAction))

		err := wg.Wait()
		if err != nil {
			return err
		}

		s.Processed.proposals = append(s.Processed.proposals, proposal)
		s.Processed.proposalsActions = append(s.Processed.proposalsActions, proposalAction)
	}

	return nil
}

func (s *Storable) getAPDescriptionsFromChain(ctx context.Context, aps []ethtypes.GovernanceAbrogationProposalStartedEvent) error {
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
