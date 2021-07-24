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

	AddDBFlags(scrapeCmd)
	AddRedisFlags(scrapeCmd)
	AddMetricsFlags(scrapeCmd)
	AddFeatureFlags(scrapeCmd)
	AddETHFlags(scrapeCmd)
	AddGenerateETHTypesFlags(scrapeCmd)

	AddStorableAccountERC20TransfersFlags(scrapeCmd)
	AddStorableGovernanceFlags(scrapeCmd)
	AddStorableMonitoredERC20TransfersFlags(scrapeCmd)
	AddStorableBarnFlags(scrapeCmd)
	AddStorableYieldFarmingFlags(scrapeCmd)
}
