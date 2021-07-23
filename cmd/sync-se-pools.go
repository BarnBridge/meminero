package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/barnbridge/smartbackend/db"
	"github.com/barnbridge/smartbackend/types"
	"github.com/barnbridge/smartbackend/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var syncSEPoolsCmd = &cobra.Command{
	Use:   "sync-se-pools",
	Short: "Sync SmartExposure pools in the database with the ones in the json file",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		dbConn := database.Connection()

		data, err := ioutil.ReadFile(viper.GetString("sePoolsFile"))
		if err != nil {
			log.Fatal(err)
		}

		var pools struct {
			SmartExposure []types.SEPool `json:"smartExposure"`
		}

		err = json.Unmarshal(data, &pools)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("removing current pools from database")

		_, err = dbConn.Exec(context.Background(), `delete from smart_exposure.pools;`)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("done removing pools")

		log.WithField("count", len(pools.SmartExposure)).Info("adding SmartExposure pools from file to database")

		for _, p := range pools.SmartExposure {
			_, err = dbConn.Exec(context.Background(), "insert into smart_exposure.pools (pool_address, pool_name, token_a_address, token_a_symbol, token_a_decimals, token_b_address, token_b_symbol, token_b_decimals,start_at_block) values ($1, $2, $3, $4, $5, $6, $7, $8,$9);",
				utils.NormalizeAddress(p.EPoolAddress),
				strings.ToUpper(p.ProtocolId),
				utils.NormalizeAddress(p.ATokenAddress),
				p.ATokenSymbol,
				p.ATokenDecimals,
				utils.NormalizeAddress(p.BTokenAddress),
				p.BTokenSymbol,
				p.BTokenDecimals,
				p.StartAtBlock,
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("done")
	},
}

func init() {
	RootCmd.AddCommand(syncSEPoolsCmd)

	addDBFlags(syncSEPoolsCmd)

	syncSEPoolsCmd.Flags().String("sePoolsFile", "./pools.kovan.json", "Path to list of pools in json format")
}
