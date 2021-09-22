package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/eth"

	"github.com/barnbridge/meminero/glue"
	"github.com/barnbridge/meminero/state"
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

		d, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		err = d.Migrate(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		state, err := state.NewManager(d.Connection())
		if err != nil {
			log.Fatal(err)
		}

		g, err := glue.New(d.Connection(), state)
		if err != nil {
			log.Fatal(err)
		}

		savedBlock, err := g.ScrapeSingleBlock(context.Background(), block)
		if err != nil {
			log.Fatal(err)
		}
		if !savedBlock {
			log.Info("block skipped")
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	scrapeCmd.AddCommand(scrapeSingleCmd)

	scrapeSingleCmd.Flags().Int64("block", -1, "The block to scrape")
}
