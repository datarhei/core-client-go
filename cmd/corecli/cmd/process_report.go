package cmd

import (
	"fmt"

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

		f, err := formatJSON(report, true)
		if err != nil {
			return err
		}

		fmt.Println(f)

		return nil
	},
}

func init() {
	processCmd.AddCommand(reportProcessCmd)
}