package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

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

		listenOn := fmt.Sprintf(":%d", config.Store.Metrics.Port)
		sm := http.NewServeMux()
		sm.Handle("/metrics", promhttp.Handler())
		metricsSrv := &http.Server{Addr: listenOn, Handler: sm}
		go func() {
			log.Infof("serving metrics on %s", listenOn)
			err := metricsSrv.ListenAndServe()
			if err != nil && ctx.Err() == nil {
				log.Fatal(err)
			}
		}()

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

		// TODO i think listening to ctx done like this does not leave time for threads to exit cleanly
		<-ctx.Done()

		// cleanup
		_ = metricsSrv.Close()

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	scrapeCmd.AddCommand(scrapeQueueCmd)
}
