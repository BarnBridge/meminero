package sync

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/utils"
)

var syncAccounts = &cobra.Command{
	Use:   "accounts",
	Short: "Sync monitored accounts",
	Run: func(cmd *cobra.Command, args []string) {
		var monitored struct {
			Accounts []string `json:"accounts"`
		}

		err := readAndUnmarshalInto(&monitored)
		if err != nil {
			log.Fatal(err)
		}

		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		log.Info("removing current monitored accounts from database")

		_, err = database.Connection().Exec(context.Background(), `delete from public.monitored_accounts;`)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("done removing accounts")
		log.WithField("count", len(monitored.Accounts)).Info("adding monitored accounts from file to database")

		for _, a := range monitored.Accounts {
			_, err := database.Connection().Exec(context.Background(), "insert into public.monitored_accounts (address) values($1)", utils.NormalizeAddress(a))
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	syncCmd.AddCommand(syncAccounts)
}
