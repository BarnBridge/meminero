package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/smartbackend/db"
	"github.com/barnbridge/smartbackend/eth"

	"github.com/barnbridge/smartbackend/glue"
	"github.com/barnbridge/smartbackend/state"
)

var scrapeSingleCmd = &cobra.Command{
	Use:   "single",
	Short: "Scrape a single blocks",
	Run: func(cmd *cobra.Command, args []string) {
		block := viper.GetInt64("block")

		if block == -1 {
			log.Fatal("No block was specified")
		}

		err := eth.Init()
		if err != nil {
			log.Fatal(err)
		}

		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		state, err := state.NewManager(database.Connection())
		if err != nil {
			log.Fatal(err)
		}

		g, err := glue.New(database.Connection(), state)
		if err != nil {
			log.Fatal(err)
		}

		err = g.ScrapeSingleBlock(context.Background(), block)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	scrapeCmd.AddCommand(scrapeSingleCmd)

	addDBFlags(scrapeSingleCmd)
	addRedisFlags(scrapeSingleCmd)
	addFeatureFlags(scrapeSingleCmd)
	addETHFlags(scrapeSingleCmd)
	addGenerateETHTypesFlags(scrapeSingleCmd)

	scrapeSingleCmd.Flags().Int64("block", -1, "The block to scrape")
}
