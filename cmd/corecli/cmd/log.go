package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// logCmd represents the metrics command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logging related commands",
	Long:  "Logging related commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		logs, err := client.Log()
		if err != nil {
			return err
		}

		for _, l := range logs {
			if err := writeJSON(os.Stdout, l, true); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
