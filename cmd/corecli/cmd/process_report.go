package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// reportProcessCmd represents the show command
var reportProcessCmd = &cobra.Command{
	Use:   "report [processid]",
	Short: "Show the report of the process with the given ID",
	Long:  "Show the report of the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		report, err := client.ProcessReport(id)
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, report, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(reportProcessCmd)
}
