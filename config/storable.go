package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
	Governance            governance            `mapstructure:"governance"`
	Erc20Transfers        erc20Transfers        `mapstructure:"erc20transfers"`
}

type accountERC20Transfers struct {
	Enabled bool
}

type governance struct {
	Enabled       bool
	Address       string
	Notifications bool
}

type erc20Transfers struct {
	Enabled bool
}
