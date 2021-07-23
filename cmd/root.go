package cmd

import (
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/meminero/config"
)

var log = logrus.WithField("module", "main")

var (
	configFile string
	version    bool

	RootCmd = &cobra.Command{
		Use:   "meminero",
		Short: "Ethereum data and indexer",
		Long:  "Scrape ethereum data from any web3-compatible node and index it into a postgres database",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				log.Fatal(err)
			}

			if configFile != "" {
				// get the filepath
				abs, err := filepath.Abs(configFile)
				if err != nil {
					log.Error("Error reading filepath: ", err.Error())
				}

				// get the config name
				base := filepath.Base(abs)

				// get the path
				path := filepath.Dir(abs)

				//
				viper.SetConfigName(strings.Split(base, ".")[0])
				viper.AddConfigPath(path)
			}

			viper.AddConfigPath(".")

			// Find and read the config file; Handle errors reading the config file
			if err := viper.ReadInConfig(); err != nil {
				log.Info("Could not load config file. Falling back to args. Error: ", err)
			}

			initLogging()
			buildDBConnectionString()

			config.Load()
		},

		Run: func(cmd *cobra.Command, args []string) {
			// fall back on default help if no args/flags are passed
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func init() {
	cobra.OnInitialize(func() {
		viper.Set("version", RootCmd.Version)
	})
	viper.SetEnvPrefix("meminero")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// persistent flags
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "/path/to/config.yml")
	RootCmd.PersistentFlags().Bool("v", false, "Set all logging modules to debug (shorthand for `--logging=*=debug`)")
	RootCmd.PersistentFlags().Bool("vv", false, "Set all logging modules to trace (shorthand for `--logging=*=trace`)")
	RootCmd.PersistentFlags().String("logging", "*=info", "Display debug messages")

	// local flags;
	RootCmd.Flags().BoolVar(&version, "version", false, "Display the current version of this CLI")
}
