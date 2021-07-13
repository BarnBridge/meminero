package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

		database, err := state.NewPGX()
		if err != nil {
			log.Fatal(err)
		}

		state, err := state.NewManager(database)
		if err != nil {
			log.Fatal(err)
		}

		g, err := glue.New(database, state)
		if err != nil {
			log.Fatal(err)
		}

		err = g.ScrapeSingleBlock(block)
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
	addAccountERC20TransfersFlags(scrapeSingleCmd)

	scrapeSingleCmd.Flags().Int64("block", -1, "The block to scrape")
}
