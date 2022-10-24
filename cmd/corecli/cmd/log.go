package cmd

import (
	"fmt"

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

			f, err := formatJSON(l, true)
			if err != nil {
				return err
			}

			fmt.Println(f)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
