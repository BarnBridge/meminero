package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/smartbackend/db"
	"github.com/barnbridge/smartbackend/state"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database to an empty state by truncating all the tables",
	Run: func(cmd *cobra.Command, args []string) {
		if !viper.GetBool("force") {
			fmt.Print("This will reset the database. Are you sure? [y/N]: ")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			if strings.TrimSuffix(strings.ToLower(text), "\n") != "y" {
				fmt.Println("Nobody was harmed.")
				return
			}
		}

		stateManager, err := state.NewManager(nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Deleting todo queue from redis ... ")

		err = stateManager.Reset()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[done]")

		fmt.Print("Truncating database ... ")

		database, err := db.New()
		if err != nil {
			log.Fatal(err)
		}

		_, err = database.Connection().Exec(context.Background(), `
			drop schema public cascade;
			create schema public;
		
			drop schema if exists yield_farming cascade;
			drop schema if exists governance cascade;
			drop schema if exists smart_yield cascade;
			drop schema if exists smart_exposure cascade;
		`)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[done]")

		fmt.Println("Database was reset. Have fun!")
	},
}

func init() {
	RootCmd.AddCommand(resetCmd)

	addDBFlags(resetCmd)
	addRedisFlags(resetCmd)

	resetCmd.Flags().Bool("force", false, "Skip interactive shell")
}
