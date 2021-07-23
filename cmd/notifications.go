package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/notifications"
)

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "generate and store notifications",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		d, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		n, err := notifications.NewWorker(d.Connection())

		n.Run(ctx)

		// TODO i think listening to ctx done like this does not leave time for threads to exit cleanly
		<-ctx.Done()

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	RootCmd.AddCommand(notificationsCmd)

	addDBFlags(notificationsCmd)
}
