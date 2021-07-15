package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
}

type accountERC20Transfers struct {
	Enabled bool
}

