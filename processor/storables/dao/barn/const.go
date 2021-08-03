package barn

type ActionType string

const (
	Deposit  ActionType = "DEPOSIT"
	Withdraw ActionType = "WITHDRAW"
)

const (
	DelegateStart ActionType = "START"
	DelegateStop  ActionType = "STOP"
)

const (
	DelegateIncrease ActionType = "INCREASE"
	DelegateDecrease ActionType = "DECREASE"
)
