package config

import (
	"github.com/lacasian/ethwheels/bestblock"
)

type database struct {
	Host        string
	Port        string
	SSLMode     string
	DBName      string
	User        string
	Password    string
	Automigrate bool

	ConnectionString string `mapstructure:"connection-string"`
}

type redis struct {
	Server   string
	List     string
	Password string
}

type api struct {
	Port        string
	DevCors     bool   `mapstructure:"dev-cors"`
	DevCorsHost string `mapstructure:"dev-cors-host"`
}

type features struct {
	Uncles struct {
		Enabled bool
	}
	Integrity struct {
		Enabled bool
	}
	QueueKeeper struct {
		Enabled bool
		Lag     int64
	}
	ReplaceBlocks bool `mapstructure:"replace-blocks"`
}

type eth struct {
	bestblock.Config `mapstructure:"client"`
}

type accountERC20Transfers struct {
	Enabled bool
}