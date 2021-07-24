package sync

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/barnbridge/meminero/cmd"
)

var log = logrus.WithField("module", "main")

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync various information from files to database",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	cmd.RootCmd.AddCommand(syncCmd)

	cmd.AddDBFlags(syncCmd)

	syncCmd.PersistentFlags().String("file", "", "Path to the file to be synced, in json format")
}
