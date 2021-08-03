package types

type PriceProvider struct {
	Address  string `json:"address"`
	Reverse  bool   `json:"reverse"`
	Decimals int64  `json:"decimals"`
}

type Price struct {
	StartAtBlock int64           `json:"startAtBlock"`
	Quote        string          `json:"quote"`
	Provider     string          `json:"provider"`
	Path         []PriceProvider `json:"path"`
}

type Token struct {
	Address  string  `json:"address"`
	Symbol   string  `json:"symbol"`
	Decimals int64   `json:"decimals"`
	Prices   []Price `json:"prices"`
}
