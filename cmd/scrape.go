package cmd

import (
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape blocks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	RootCmd.AddCommand(scrapeCmd)

	addStorableAccountERC20TransfersFlags(scrapeCmd)
	addStorableGovernanceFlags(scrapeCmd)
	addStorableMonitoredERC20TransfersFlags(scrapeCmd)
}
