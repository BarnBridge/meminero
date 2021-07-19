package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type store struct {
	Database database `mapstructure:"db"`
	Redis    redis    `mapstructure:"redis"`
	Metrics  metrics  `mapstructure:"metrics"`
	API      api      `mapstructure:"api"`
	Feature  features `mapstructure:"feature"`
	ETH      eth      `mapstructure:"eth"`
	EthTypes ethtypes `mapstructure:"ethtypes"`
	Storable storable `mapstructure:"storable"`
}

var Store store

func Load() {
	err := viper.Unmarshal(&Store)
	if err != nil {
		logrus.Fatal(err)
	}
}
