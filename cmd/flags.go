package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func addDBFlags(cmd *cobra.Command) {
	cmd.Flags().String("db.connection-string", "", "Postgres connection string.")
	cmd.Flags().String("db.host", "localhost", "Database host")
	cmd.Flags().String("db.port", "5432", "Database port")
	cmd.Flags().String("db.sslmode", "disable", "Database sslmode")
	cmd.Flags().String("db.dbname", "name", "Database name")
	cmd.Flags().String("db.user", "", "Database user (also allowed via PG_USER env)")
	cmd.Flags().String("db.password", "", "Database password (also allowed via PG_PASSWORD env)")
	cmd.Flags().Bool("db.automigrate", true, "Enable/disable the automatic migrations feature")
	cmd.Flags().String("db.migrations-path", "db/migrations", "Path to migrations directory")
}

func addRedisFlags(cmd *cobra.Command) {
	cmd.Flags().String("redis.server", "localhost:6379", "Redis server URL")
	cmd.Flags().String("redis.list", "todo", "The name of the list to be used for task management")
	cmd.Flags().String("redis.password", "", "Redis password")
}

func addMetricsFlags(cmd *cobra.Command) {
	cmd.Flags().Int64("metrics.port", 9909, "Port on which to serve Prometheus metrics")
}

func addFeatureFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("feature.integrity.enabled", true, "Enable/disable the integrity checker")
	cmd.Flags().Bool("feature.queuekeeper.enabled", true, "Enable/disable the queue keeper (watch new heads and store into the queue)")
	cmd.Flags().Int64("feature.queuekeeper.lag", 10, "The amount of blocks to lag behind the tip of the chain")
	cmd.Flags().Bool("feature.replace-blocks", false, "Enable this if the scraper should replace existing blocks instead of skipping them")
}

func addETHFlags(cmd *cobra.Command) {
	cmd.Flags().String("eth.client.http", "", "HTTP endpoint of JSON-RPC enabled Ethereum node")
	cmd.Flags().String("eth.client.ws", "", "WS endpoint of JSON-RPC enabled Ethereum node (provide this only if you want to use websocket subscription for tracking best block)")
	cmd.Flags().Duration("eth.client.poll-interval", 15*time.Second, "Interval to be used for polling the Ethereum node for best block")
}

func addGenerateETHTypesFlags(cmd *cobra.Command) {
	cmd.Flags().String("ethtypes.abi-folder", "ethtypes/_source", "Folder containing ABI JSONs")
	cmd.Flags().String("ethtypes.package-path", "ethtypes", "Path where to generate packages. Final folder represents package name")
}

func addStorableAccountERC20TransfersFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.accountERC20Transfers.enabled", true, "Enable/disable erc20 transfers scraping")
}

func addStorableGovernanceFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.governance.enabled", true, "Enable/disable governance scraping")
	cmd.PersistentFlags().Bool("storable.governance.notifications", true, "Enable/disable governance notifications")
	cmd.PersistentFlags().String("storable.governance.address", "0x4cAE362D7F227e3d306f70ce4878E245563F3069", "Address of governance contract")
}

func addStorableMonitoredERC20TransfersFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.erc20Transfers.enabled", true, "Enable/disable erc20Transfers scraping")
}

func addStorableYieldFarmingFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.yieldFarming.enabled", true, "Enable/disable yieldFarming scraping")
	cmd.PersistentFlags().String("storable.yieldFarming.address", "", "Address of governance contract")
}
