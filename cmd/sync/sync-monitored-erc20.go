package sync

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/db"
	"github.com/barnbridge/meminero/utils"
)

var syncERC20 = &cobra.Command{
	Use:   "monitored-erc20",
	Short: "Sync monitored erc20",
	Run: func(cmd *cobra.Command, args []string) {
		var monitored struct {
			Tokens []string `json:"tokens"`
		}

		err := readAndUnmarshalInto(&monitored)
		if err != nil {
			log.Fatal(err)
		}

		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		log.Info("removing current monitored erc20 from database")

		_, err = database.Connection().Exec(context.Background(), `delete from public.monitored_erc20;`)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("done removing tokens")
		log.WithField("count", len(monitored.Tokens)).Info("adding monitored erc20 from file to database")

		for _, a := range monitored.Tokens {
			_, err := database.Connection().Exec(context.Background(), "insert into public.monitored_erc20 (address) values($1)", utils.NormalizeAddress(a))
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Work done. Goodbye!")
	},
}

func init() {
	syncCmd.AddCommand(syncERC20)
}
