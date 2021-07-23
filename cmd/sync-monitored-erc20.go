package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var syncERC20 = &cobra.Command{
	Use:   "sync-erc20",
	Short: "Sync monitored erc20",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		dbConn := database.Connection()

		data, err := ioutil.ReadFile(viper.GetString("erc20file"))
		if err != nil {
			log.Fatal(err)
		}

		var monitored struct {
			Tokens []string `json:"tokens"`
		}

		err = json.Unmarshal(data, &monitored)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("removing current monitored erc20 from database")

		_, err = dbConn.Exec(context.Background(), `delete from public.monitored_erc20;`)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("done removing tokens")
		log.WithField("count", len(monitored.Tokens)).Info("adding monitored erc20 from file to database")

		for _, a := range monitored.Tokens {
			_, err := dbConn.Exec(context.Background(), "insert into public.monitored_erc20 (address) values($1)", utils.NormalizeAddress(a))
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	RootCmd.AddCommand(syncERC20)

	addDBFlags(syncERC20)

	syncERC20.Flags().String("erc20file", ".", "Path to list of monitored erc20 tokens in json format")
}
