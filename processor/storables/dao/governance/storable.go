package governance

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/barnbridge/meminero/ethtypes"
	"github.com/barnbridge/meminero/types"
)

type Storable struct {
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

const storableID = "dao.governance"

func New(block *types.Block) *Storable {
	return &Storable{
		block:  block,
		logger: logrus.WithField("module", fmt.Sprintf("storable(%s)", storableID)),
	}
}

func (s *Storable) ID() string {
	return storableID
}

func (s *Storable) Result() interface{} {
	return s.Processed
}
