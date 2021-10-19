package config

import (
	"github.com/lacasian/ethwheels/bestblock"
)

type database struct {
	Host           string
	Port           string
	SSLMode        string
	DBName         string
	User           string
	Password       string
	Automigrate    bool
	MigrationsPath string `mapstructure:"migrations-path"`

	ConnectionString string `mapstructure:"connection-string"`
}

type redis struct {
	Server   string
	List     string
	Password string
}

type metrics struct {
	Port int64
}

type api struct {
	Port        string
	DevCors     bool   `mapstructure:"dev-cors"`
	DevCorsHost string `mapstructure:"dev-cors-host"`
}

type features struct {
	Integrity struct {
		Enabled bool
	}
	QueueKeeper struct {
		Enabled bool
		Lag     int64
	}
	ReplaceBlocks bool `mapstructure:"replace-blocks"`
	ContractState struct {
		Enabled bool
	} `mapstructure:"contract-state"`
	RequeueFailedBlocks bool `mapstructure:"requeue-failed-blocks"`
}

type eth struct {
	bestblock.Config `mapstructure:"client"`
	MaxBatch         int `mapstructure:"max-batch"`
}

type ethtypes struct {
	AbiFolder   string `mapstructure:"abi-folder"`
	PackagePath string `mapstructure:"package-path"`
}

type syncer struct {
	Path     string
	Network  string
	Datasets []string
}
