package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	formatter "github.com/lacasian/logrus-module-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/lacasian/ethwheels/bestblock"
)

func initLogging() {
	logging := viper.GetString("logging")

	if viper.GetBool("v") {
		logging = "*=debug"
	}

	if viper.GetBool("vv") {
		logging = "*=trace"
	}

	if logging == "" {
		logging = "*=info"
	}

	viper.Set("logging", logging)

	gin.SetMode(gin.DebugMode)

	modules := formatter.NewModulesMap(logging)
	if level, exists := modules["gin"]; exists {
		if level < logrus.DebugLevel {
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		level := modules["*"]
		if level < logrus.DebugLevel {
			gin.SetMode(gin.ReleaseMode)
		}
	}

	f, err := formatter.New(modules)
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(f)

	log.Debug("Debug mode")
}

func buildDBConnectionString() {
	if viper.GetString("db.connection-string") == "" {
		user := viper.GetString("db.user")
		pass := viper.GetString("db.password")

		p := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s", viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.sslmode"), viper.GetString("db.dbname"), user, pass)
		viper.Set("db.connection-string", p)
	}
}

func initBestBlockTracker(config bestblock.Config) (*bestblock.Tracker, error) {
	bbtracker, err := bestblock.NewTracker(config)
	if err != nil {
		return nil, errors.Wrap(err, "could not init best block tracker")
	}

	go bbtracker.Run()
	err = <-bbtracker.Err()
	if err != nil {
		return nil, errors.Wrap(err, "could not run best block tracker")
	}

	go func() {
		// todo: can we handle these errors?
		for err := range bbtracker.Err() {
			log.Error(err)
		}
	}()

	return bbtracker, nil
}
