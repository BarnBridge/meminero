package types

type SEPool struct {
	EPoolAddress string
	ProtocolId   string

	ATokenAddress  string
	ATokenSymbol   string
	ATokenDecimals int64

	BTokenAddress  string
	BTokenSymbol   string
	BTokenDecimals int64

	StartAtBlock int64
}
