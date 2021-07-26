package cmd

import (
	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/syncer"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync various information from files to database",
	Run: func(cmd *cobra.Command, args []string) {
		err := syncer.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(syncCmd)

	addDBFlags(syncCmd)
	addSyncerFlags(syncCmd)
}
