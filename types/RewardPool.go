package types

type RewardPoolType string

const (
	PoolTypeSingle RewardPoolType = "SINGLE"
	PoolTypeMulti  RewardPoolType = "MULTI"
)

type RewardPool struct {
	PoolType             RewardPoolType
	PoolAddress          string
	PoolTokenAddress     string
	RewardTokenAddresses []string
	StartAtBlock         int64
}
