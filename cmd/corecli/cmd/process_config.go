package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// configProcessCmd represents the show command
var configProcessCmd = &cobra.Command{
	Use:   "config [processid]",
	Short: "Show the config of the process with the given ID",
	Long:  "Show the config of the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		config, err := client.ProcessConfig(id)
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, config, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(configProcessCmd)
}
