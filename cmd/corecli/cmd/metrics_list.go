package cmd

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listMetricsCmd represents the list command
var listMetricsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all known metrics",
	Long:  "List all known metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		list, err := client.MetricsList()
		if err != nil {
			return err
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"Name", "Description", "Labels"})

		for _, m := range list {
			t.AppendRow(table.Row{m.Name, m.Description, strings.Join(m.Labels, ",")})
		}

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		return nil

	},
}

func init() {
	metricsCmd.AddCommand(listMetricsCmd)
}
