package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// stateProcessCmd represents the show command
var stateProcessCmd = &cobra.Command{
	Use:   "state [processid]",
	Short: "Show the state of the process with the given ID",
	Long:  "Show the state of the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		state, err := client.ProcessState(id)
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, state, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(stateProcessCmd)
}
