package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type store struct {
	Database database `mapstructure:"db"`
	Redis    redis    `mapstructure:"redis"`
	API      api      `mapstructure:"api"`
	Feature  features `mapstructure:"feature"`
	ETH      eth      `mapstructure:"eth"`
}

var Store store

func Load() {
	err := viper.Unmarshal(&Store)
	if err != nil {
		logrus.Fatal(err)
	}
}
