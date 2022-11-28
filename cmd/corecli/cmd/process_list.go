package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	coreclient "github.com/datarhei/core-client-go/v16"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// listProcessCmd represents the list command
var listProcessCmd = &cobra.Command{
	Use:   "list",
	Short: "List all processes",
	Long:  "List all processes of the selected core",
	RunE: func(cmd *cobra.Command, args []string) error {
		asRaw, _ := cmd.Flags().GetBool("raw")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		list, err := client.ProcessList(coreclient.ProcessListOptions{
			Filter: []string{"state"},
		})
		if err != nil {
			return err
		}

		if asRaw {
			if err := writeJSON(os.Stdout, list, true); err != nil {
				return err
			}

			return nil
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Reference", "Order", "State", "Memory", "CPU", "Runtime"})

		for _, p := range list {
			runtime := p.State.Runtime
			if p.State.State != "running" {
				runtime = 0
			}

			order := strings.ToUpper(p.State.Order)
			switch order {
			case "START":
				order = text.FgGreen.Sprint(order)
			case "STOP":
				order = text.Colors{text.FgWhite, text.Faint}.Sprint(order)
			}

			state := strings.ToUpper(p.State.State)
			switch state {
			case "RUNNING":
				state = text.FgGreen.Sprint(state)
			case "FINISHED":
				state = text.Colors{text.FgWhite, text.Faint}.Sprint(state)
			case "FAILED":
				state = text.FgRed.Sprint(state)
			case "STARTING":
				state = text.FgCyan.Sprint(state)
			case "FINISHING":
				state = text.FgCyan.Sprint(state)
			case "KILLED":
				state = text.Colors{text.FgRed, text.Faint}.Sprint(state)
			}

			t.AppendRow(table.Row{
				p.ID,
				p.Reference,
				order,
				state,
				formatByteCountBinary(p.State.Memory),
				fmt.Sprintf("%.1f%%", p.State.CPU),
				(time.Duration(runtime) * time.Second).String(),
			})
		}

		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 3, Align: text.AlignRight},
			{Number: 4, Align: text.AlignRight},
			{Number: 5, Align: text.AlignRight},
			{Number: 6, Align: text.AlignRight},
			{Number: 7, Align: text.AlignRight},
		})

		t.SortBy([]table.SortBy{
			{Number: 2, Mode: table.Asc},
			{Number: 1, Mode: table.Asc},
			{Number: 4, Mode: table.Asc},
			{Number: 6, Mode: table.Dsc},
		})

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		return nil
	},
}

func init() {
	processCmd.AddCommand(listProcessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
