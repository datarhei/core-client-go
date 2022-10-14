package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reloadProcessCmd represents the show command
var reloadProcessCmd = &cobra.Command{
	Use:   "reload [processid]",
	Short: "Reload the process with the given ID",
	Long:  "Reload the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		if err := client.ProcessCommand(id, "reload"); err != nil {
			return err
		}

		fmt.Printf("%s reload\n", id)

		return nil
	},
}

func init() {
	processCmd.AddCommand(reloadProcessCmd)
}
