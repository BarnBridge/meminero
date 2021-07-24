package sync

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/db"
)

var syncLabels = &cobra.Command{
	Use:   "labels",
	Short: "Sync monitored erc20",
	Run: func(cmd *cobra.Command, args []string) {
		var data struct {
			Labels []struct {
				Address string `json:"address"`
				Label   string `json:"label"`
			} `json:"labels"`
		}

		err := readAndUnmarshalInto(&data)
		if err != nil {
			log.Fatal(err)
		}

		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		log.WithField("count", len(data.Labels)).Info("syncing labels from file to database")

		for _, a := range data.Labels {
			_, err := database.Connection().Exec(
				context.Background(),
				`insert into public.labels (address, label) values($1, $2) on conflict do nothing`,
				a,
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	syncCmd.AddCommand(syncLabels)
}
