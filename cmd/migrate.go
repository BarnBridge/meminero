package cmd

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/barnbridge/smartbackend/migrations"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manually run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("postgres", viper.GetString("db.connection-string"))
		if err != nil {
			log.Fatal(err)
		}

		err = goose.Up(db, "/tmp")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	addDBFlags(migrateCmd)
}
