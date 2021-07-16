package governance

type ActionType string

const (
	CREATED  ActionType = "CREATED"
	QUEUED   ActionType = "QUEUED"
	EXECUTED ActionType = "EXECUTED"
	CANCELED ActionType = "CANCELED"
)

