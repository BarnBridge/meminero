package config

type storable struct {
	AccountERC20Transfers accountERC20Transfers `mapstructure:"accountERC20Transfers"`
	Governance            governance            `mapstructure:"governance"`
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

type smartExposure struct {
	Enabled               bool
	EPoolPeripheryAddress string
	ETokenFactoryAddress  string
}
