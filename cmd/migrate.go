package cmd

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"

	"github.com/barnbridge/smartbackend/db"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manually run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		d, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		err = d.Migrate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	addDBFlags(migrateCmd)
}
