package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/barnbridge/smartbackend/config"
	"github.com/barnbridge/smartbackend/glue"
	"github.com/barnbridge/smartbackend/integrity"
	"github.com/barnbridge/smartbackend/state"
	"github.com/barnbridge/smartbackend/state/queuekeeper"

	"github.com/spf13/cobra"
)

var scrapeQueueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Start the scraper as a long running process",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		database, err := state.NewPostgres()
		if err != nil {
			log.Fatal(err)
		}

		state, err := state.NewManager()
		if err != nil {
			log.Fatal(err)
		}

		tracker, err := initBestBlockTracker(config.Store.ETH.Config)
		if err != nil {
			log.Fatal(err)
		}

		if config.Store.Feature.Integrity.Enabled {
			integrityChecker := integrity.NewChecker(database, tracker, state)
			go integrityChecker.Run(ctx)
		}

		if config.Store.Feature.QueueKeeper.Enabled {
			keeper, err := queuekeeper.New(tracker, state)
			if err != nil {
				log.Fatal(err)
			}
			go keeper.Run(ctx)
		}

		g, err := glue.New(database, state)
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
}
