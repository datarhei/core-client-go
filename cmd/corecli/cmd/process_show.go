package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// showProcessCmd represents the show command
var showProcessCmd = &cobra.Command{
	Use:   "show [processid]",
	Short: "Show the process with the given ID",
	Long:  "Show the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		asRaw, _ := cmd.Flags().GetBool("raw")
		config, _ := cmd.Flags().GetBool("cfg")
		state, _ := cmd.Flags().GetBool("state")
		report, _ := cmd.Flags().GetBool("report")
		metadata, _ := cmd.Flags().GetBool("metadata")
		command, _ := cmd.Flags().GetBool("command")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		filter := []string{}
		if config {
			filter = append(filter, "config")
		}
		if state || command {
			filter = append(filter, "state")
		}
		if report {
			filter = append(filter, "report")
		}
		if metadata {
			filter = append(filter, "metadata")
		}

		process, err := client.Process(id, filter)
		if err != nil {
			return err
		}

		if command {
			fmt.Println(strings.Join(process.State.Command, " "))

			return nil
		}

		if asRaw {
			if err := writeJSON(os.Stdout, process, true); err != nil {
				return err
			}

			return nil
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"ID", "Reference", "State", "Memory", "CPU", "Runtime"})

		runtime := process.State.Runtime
		if process.State.State != "running" {
			runtime = 0
		}

		t.AppendRow(table.Row{
			process.ID,
			process.Reference,
			strings.ToUpper(process.State.State),
			formatByteCountBinary(process.State.Memory),
			fmt.Sprintf("%.1f%%", process.State.CPU),
			(time.Duration(runtime) * time.Second).String(),
		})

		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 2, Align: text.AlignRight},
			{Number: 3, Align: text.AlignRight},
			{Number: 4, Align: text.AlignRight},
			{Number: 5, Align: text.AlignRight},
		})

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		if len(process.State.Progress.Input) == 0 && len(process.State.Progress.Output) == 0 {
			return nil
		}

		t = table.NewWriter()

		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

		t.SetTitle("Inputs / Outputs")
		t.AppendHeader(table.Row{"", "#", "ID", "Type", "URL", "Specs"}, rowConfigAutoMerge)

		for i, p := range process.State.Progress.Input {
			var specs string
			if p.Type == "audio" {
				specs = fmt.Sprintf("%s %s %dHz", strings.ToUpper(p.Codec), p.Layout, p.Sampling)
			} else {
				specs = fmt.Sprintf("%s %dx%d", strings.ToUpper(p.Codec), p.Width, p.Height)
			}

			t.AppendRow(table.Row{
				"input",
				i,
				p.ID,
				strings.ToUpper(p.Type),
				p.Address,
				specs,
			}, rowConfigAutoMerge)
		}

		for i, p := range process.State.Progress.Output {
			var specs string
			if p.Type == "audio" {
				specs = fmt.Sprintf("%s %s %dHz", strings.ToUpper(p.Codec), p.Layout, p.Sampling)
			} else {
				specs = fmt.Sprintf("%s %dx%d", strings.ToUpper(p.Codec), p.Width, p.Height)
			}

			t.AppendRow(table.Row{
				"output",
				i,
				p.ID,
				strings.ToUpper(p.Type),
				p.Address,
				specs,
			}, rowConfigAutoMerge)
		}

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		return nil
	},
}

func init() {
	processCmd.AddCommand(showProcessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showProcessCmd.Flags().BoolP("cfg", "c", false, "Include the process config")
	showProcessCmd.Flags().BoolP("state", "s", false, "Include the process state")
	showProcessCmd.Flags().BoolP("report", "r", false, "Include the process config")
	showProcessCmd.Flags().BoolP("metadata", "m", false, "Include the process config")
	showProcessCmd.Flags().Bool("command", false, "Show the process command")
}
