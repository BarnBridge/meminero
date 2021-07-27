package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
	Governance            governance            `mapstructure:"governance"`
	Barn                  barn                  `mapstructure:"barn"`
	Erc20Transfers        erc20Transfers        `mapstructure:"erc20transfers"`
	YieldFarming          yieldFarming          `mapstructure:"yieldFarming"`
	SmartYield            smartYield            `mapstructure:"smartYield"`
}

type accountERC20Transfers struct {
	Enabled bool
}

type governance struct {
	Enabled       bool
	Address       string
	Notifications bool
}

type barn struct {
	Enabled       bool
	Address       string
	Notifications bool
}

type erc20Transfers struct {
	Enabled bool
}

type yieldFarming struct {
	Enabled bool
	Address string
}

type smartYield struct {
	Enabled       bool
	Notifications bool

	Rewards struct {
		Factories string
	}
}
