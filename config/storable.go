package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
	Governance            governance            `mapstructure:"governance"`
	YieldFarming          yieldFarming          `mapstructure:"yieldFarming"`
}

type accountERC20Transfers struct {
	Enabled bool
}

type governance struct {
	Enabled       bool
	Address       string
	Notifications bool
}

type yieldFarming struct {
	Enabled bool
	Address string
}
