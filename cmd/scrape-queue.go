package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/config"
	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/eth"
	"github.com/barnbridge/meminero/glue"
	"github.com/barnbridge/meminero/integrity"
	"github.com/barnbridge/meminero/state"
	"github.com/barnbridge/meminero/state/queuekeeper"
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

		tracker, err := initBestBlockTracker(config.Store.ETH.Config)
		if err != nil {
			log.Fatal(err)
		}

		if config.Store.Feature.Integrity.Enabled {
			integrityChecker := integrity.NewChecker(d.Connection(), tracker, state)
			go integrityChecker.Run(ctx)
		}

		if config.Store.Feature.QueueKeeper.Enabled {
			keeper, err := queuekeeper.New(tracker, state)
			if err != nil {
				log.Fatal(err)
			}
			go keeper.Run(ctx)
		}

		g, err := glue.New(d.Connection(), state)
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
