package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/barnbridge/smartbackend/api"
	"github.com/barnbridge/smartbackend/state"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the API only",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		d, err := state.NewPostgres()
		if err != nil {
			log.Fatal(err)
		}

		a := api.New(d)
		go a.Run()

		<-ctx.Done()

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)

	addDBFlags(apiCmd)
	addRedisFlags(apiCmd)
	addAPIFlags(apiCmd)
}
