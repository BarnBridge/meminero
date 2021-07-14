package types

type Token struct {
	Address  string
	Symbol   string
	Decimals int64
	AggregatorAddress string
	PriceProviderType string
}
