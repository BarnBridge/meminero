package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/barnbridge/smartbackend/state"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manually add a block to the todo queue",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := state.NewManager(nil)
		if err != nil {
			log.Fatal(err)
		}

		block := viper.GetInt64("block")
		if block > 0 {
			err := r.AddTaskToQueue(block)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		from := viper.GetInt64("from")
		to := viper.GetInt64("to")
		if from > 0 && to > 0 {
			for i := from; i <= to; i++ {
				err := r.AddTaskToQueue(i)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(queueCmd)

	addRedisFlags(queueCmd)

	queueCmd.Flags().Int64("block", -1, "Add a single block in the todo queue")
	queueCmd.Flags().Int64("from", -1, "Add a series of blocks into the todo queue starting from the provided number (only use in combination with --to)")
	queueCmd.Flags().Int64("to", -1, "Add a series of blocks into the todo queue ending with the provided number, inclusive (only use in combination with --from)")
}
