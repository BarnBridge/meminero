package barn

import (
	"github.com/shopspring/decimal"

	"github.com/barnbridge/meminero/ethtypes"
)

type DelegateAction struct {
	ethtypes.BarnDelegateEvent
	ActionType ActionType
}

type DelegateChange struct {
	Sender              string
	Receiver            string
	Amount              decimal.Decimal
	ToNewDelegatedPower decimal.Decimal
	ActionType          ActionType

	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
}

type StakingAction struct {
	UserAddress  string
	Amount       decimal.Decimal
	BalanceAfter decimal.Decimal
	ActionType   ActionType

	TransactionHash  string
	TransactionIndex int64
	LogIndex         int64
}
