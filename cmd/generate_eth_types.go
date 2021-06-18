//go:generate go run ../main.go generate-eth-types --ethtypes.abi-folder ../ethtypes/_source --ethtypes.package-path ../ethtypes

package cmd

import (
	"github.com/lacasian/ethwheels/ethgen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genETHDecoders = &cobra.Command{
	Use:   "generate-eth-types",
	Short: "Generate ETH data types",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("starting to generate ABI bindings")

		abisDir := viper.GetString("ethtypes.abi-folder")
		packagePath := viper.GetString("ethtypes.package-path")

		err := ethgen.NewFromABIs(abisDir, packagePath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(genETHDecoders)

	addGenerateETHTypesFlags(genETHDecoders)
}
