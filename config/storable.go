package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
	Governance            governance            `mapstructure:"governance"`
	Barn                  barn                  `mapstructure:"barn"`
	Erc20Transfers        erc20Transfers        `mapstructure:"erc20transfers"`
	YieldFarming          yieldFarming          `mapstructure:"yieldFarming"`
	SmartExposure         smartExposure         `mapstructure:"smartExposure"`
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

type smartExposure struct {
	Enabled               bool
	EPoolPeripheryAddress string
	ETokenFactoryAddress  string
}
