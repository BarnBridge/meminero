package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/barnbridge/smartbackend/db"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var syncAccounts = &cobra.Command{
	Use:   "sync-accounts",
	Short: "Sync monitored accounts",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		dbConn := database.Connection()

		data, err := ioutil.ReadFile(viper.GetString("file"))
		if err != nil {
			log.Fatal(err)
		}

		var monitored struct {
			Accounts []string `json:"accounts"`
		}

		err = json.Unmarshal(data, &monitored)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("removing current monitored accounts from database")

		_, err = dbConn.Exec(context.Background(), `delete from public.monitored_accounts;`)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("done removing accounts")
		log.WithField("count", len(monitored.Accounts)).Info("adding monitored accounts from file to database")

		for _, a := range monitored.Accounts {
			_, err := dbConn.Exec(context.Background(), "insert into public.monitored_accounts (address) values($1)", utils.NormalizeAddress(a))
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	RootCmd.AddCommand(syncAccounts)

	addDBFlags(syncAccounts)

	syncAccounts.Flags().String("file", "./accounts.kovan.json", "Path to list of monitored accounts in json format")
}
