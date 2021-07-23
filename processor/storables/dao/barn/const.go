package barn

const (
	DepositEvent                = "Deposit"
	WithdrawEvent               = "Withdraw"
	LockEvent                   = "Lock"
	DelegateEvent               = "Delegate"
	DelegatePowerIncreasedEvent = "DelegatedPowerIncreased"
	DelegatePowerDecreasedEvent = "DelegatedPowerDecreased"
)

type ActionType string

const (
	DEPOSIT  ActionType = "DEPOSIT"
	WITHDRAW ActionType = "WITHDRAW"
)

const (
	DELEGATE_START ActionType = "START"
	DELEGATE_STOP  ActionType = "STOP"
)

const (
	DELEGATE_INCREASE ActionType = "INCREASE"
	DELEGATE_DECREASE ActionType = "DECREASE"
)

const ZeroAddress = "0x0000000000000000000000000000000000000000"
