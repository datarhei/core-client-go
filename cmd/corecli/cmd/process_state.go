package cmd

import (
	"fmt"

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

		f, err := formatJSON(state, true)
		if err != nil {
			return err
		}

		fmt.Println(f)

		return nil
	},
}

func init() {
	processCmd.AddCommand(stateProcessCmd)
}
