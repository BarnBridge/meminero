package governance

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Proposal struct {
	Id           *big.Int
	Proposer     common.Address
	Description  string
	Title        string
	CreateTime   *big.Int
	Eta          *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
	ProposalParameters
}
type ProposalParameters struct {
	WarmUpDuration      *big.Int
	ActiveDuration      *big.Int
	QueueDuration       *big.Int
	GracePeriodDuration *big.Int
	AcceptanceThreshold *big.Int
	MinQuorum           *big.Int
}

type ProposalActions struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}