package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/barnbridge/smartbackend/abi"
	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/db"
	"github.com/barnbridge/smartbackend/eth"
	"github.com/barnbridge/smartbackend/glue"
	"github.com/barnbridge/smartbackend/integrity"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/state/queuekeeper"
)

var scrapeQueueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Start the scraper as a long running process",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		abi.Init()

		err := eth.Init()
		if err != nil {
			log.Fatal(err)
		}

		db, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		state, err := state.NewManager(db.Connection())
		if err != nil {
			log.Fatal(err)
		}

		tracker, err := initBestBlockTracker(config.Store.ETH.Config)
		if err != nil {
			log.Fatal(err)
		}

		if config.Store.Feature.Integrity.Enabled {
			integrityChecker := integrity.NewChecker(db.Connection(), tracker, state)
			go integrityChecker.Run(ctx)
		}

		if config.Store.Feature.QueueKeeper.Enabled {
			keeper, err := queuekeeper.New(tracker, state)
			if err != nil {
				log.Fatal(err)
			}
			go keeper.Run(ctx)
		}

		g, err := glue.New(db.Connection(), state)
		if err != nil {
			log.Fatal(err)
		}

		go g.Run(ctx)

		<-ctx.Done()

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	scrapeCmd.AddCommand(scrapeQueueCmd)

	addDBFlags(scrapeQueueCmd)
	addRedisFlags(scrapeQueueCmd)
	addFeatureFlags(scrapeQueueCmd)
	addETHFlags(scrapeQueueCmd)
	addGenerateETHTypesFlags(scrapeQueueCmd)

	addStorableAccountERC20TransfersFlags(scrapeQueueCmd)
}
