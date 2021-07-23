package barn

import (
	"math/big"

	"github.com/barnbridge/smartbackend/ethtypes"
)

type DelegateAction struct {
	ethtypes.BarnDelegateEvent
	ActionType ActionType
}

type DelegateChange struct {
	Sender              string
	Receiver            string
	Amount              *big.Int
	ToNewDelegatedPower *big.Int
	ActionType          ActionType

	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
}

type StakingAction struct {
	UserAddress  string
	Amount       *big.Int
	BalanceAfter *big.Int
	ActionType   ActionType

	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
}
