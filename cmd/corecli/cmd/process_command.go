package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// stateProcessCmd represents the show command
var commandProcessCmd = &cobra.Command{
	Use:   "command [processid]",
	Short: "Show the ffmpeg command of the process with the given ID",
	Long:  "Show the ffmpeg command of the process with the given ID",
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

		fmt.Printf("%s\n", strings.Join(state.Command, " "))

		return nil
	},
}

func init() {
	processCmd.AddCommand(commandProcessCmd)
}
