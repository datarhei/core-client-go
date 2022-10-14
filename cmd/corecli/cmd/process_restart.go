package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// restartProcessCmd represents the show command
var restartProcessCmd = &cobra.Command{
	Use:   "restart [processid]",
	Short: "Restart the process with the given ID",
	Long:  "Restart the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		if err := client.ProcessCommand(id, "restart"); err != nil {
			return err
		}

		fmt.Printf("%s restart\n", id)

		return nil
	},
}

func init() {
	processCmd.AddCommand(restartProcessCmd)
}
