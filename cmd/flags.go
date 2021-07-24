package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func AddDBFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("db.connection-string", "", "Postgres connection string.")
	cmd.PersistentFlags().String("db.host", "localhost", "Database host")
	cmd.PersistentFlags().String("db.port", "5432", "Database port")
	cmd.PersistentFlags().String("db.sslmode", "disable", "Database sslmode")
	cmd.PersistentFlags().String("db.dbname", "name", "Database name")
	cmd.PersistentFlags().String("db.user", "", "Database user (also allowed via PG_USER env)")
	cmd.PersistentFlags().String("db.password", "", "Database password (also allowed via PG_PASSWORD env)")
	cmd.PersistentFlags().Bool("db.automigrate", true, "Enable/disable the automatic migrations feature")
	cmd.PersistentFlags().String("db.migrations-path", "db/migrations", "Path to migrations directory")
}

func AddRedisFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("redis.server", "localhost:6379", "Redis server URL")
	cmd.PersistentFlags().String("redis.list", "todo", "The name of the list to be used for task management")
	cmd.PersistentFlags().String("redis.password", "", "Redis password")
}

func AddMetricsFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Int64("metrics.port", 9909, "Port on which to serve Prometheus metrics")
}

func AddFeatureFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("feature.integrity.enabled", true, "Enable/disable the integrity checker")
	cmd.PersistentFlags().Bool("feature.queuekeeper.enabled", true, "Enable/disable the queue keeper (watch new heads and store into the queue)")
	cmd.PersistentFlags().Int64("feature.queuekeeper.lag", 10, "The amount of blocks to lag behind the tip of the chain")
	cmd.PersistentFlags().Bool("feature.replace-blocks", false, "Enable this if the scraper should replace existing blocks instead of skipping them")
}

func AddETHFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("eth.client.http", "", "HTTP endpoint of JSON-RPC enabled Ethereum node")
	cmd.PersistentFlags().String("eth.client.ws", "", "WS endpoint of JSON-RPC enabled Ethereum node (provide this only if you want to use websocket subscription for tracking best block)")
	cmd.PersistentFlags().Duration("eth.client.poll-interval", 15*time.Second, "Interval to be used for polling the Ethereum node for best block")
}

func AddGenerateETHTypesFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("ethtypes.abi-folder", "ethtypes/_source", "Folder containing ABI JSONs")
	cmd.PersistentFlags().String("ethtypes.package-path", "ethtypes", "Path where to generate packages. Final folder represents package name")
}

func AddStorableAccountERC20TransfersFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.accountERC20Transfers.enabled", true, "Enable/disable erc20 transfers scraping")
}

func AddStorableGovernanceFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.governance.enabled", true, "Enable/disable governance scraping")
	cmd.PersistentFlags().Bool("storable.governance.notifications", true, "Enable/disable governance notifications")
	cmd.PersistentFlags().String("storable.governance.address", "", "Address of governance contract")
}

func AddStorableMonitoredERC20TransfersFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.erc20Transfers.enabled", true, "Enable/disable erc20Transfers scraping")
}

func AddStorableYieldFarmingFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.yieldFarming.enabled", true, "Enable/disable yieldFarming scraping")
	cmd.PersistentFlags().String("storable.yieldFarming.address", "", "Address of governance contract")
}

func AddStorableBarnFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("storable.barn.enabled", true, "Enable/disable barn scraping")
	cmd.PersistentFlags().Bool("storable.barn.notifications", true, "Enable/disable barn notifications")
	cmd.PersistentFlags().String("storable.barn.address", "", "Address of barn staking contract")
}
