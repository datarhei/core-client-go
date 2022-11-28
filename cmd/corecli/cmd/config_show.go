package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// showConfigCmd represents the list command
var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current config",
	Long:  "Show the current config.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		target, _ := cmd.Flags().GetString("to-file")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		_, config, err := client.Config()
		if err != nil {
			return err
		}

		if target != "-" {
			f, err := formatJSON(config.Config, false)
			if err != nil {
				return err
			}

			file, err := os.Create(target)
			if err != nil {
				return err
			}

			defer file.Close()

			if _, err := file.Write([]byte(f)); err != nil {
				return err
			}
		} else {
			if err := writeJSON(os.Stdout, config.Config, true); err != nil {
				return err
			}
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"Created", "Loaded", "Updated"})

		t.AppendRow(table.Row{
			config.CreatedAt.Format("2006-01-02 15:04:05"),
			config.LoadedAt.Format("2006-01-02 15:04:05"),
			config.UpdatedAt.Format("2006-01-02 15:04:05"),
		})

		t.SetStyle(table.StyleLight)

		fmt.Fprintln(os.Stderr, t.Render())

		t = table.NewWriter()

		t.AppendHeader(table.Row{"Overrides"})

		for _, o := range config.Overrides {
			t.AppendRow(table.Row{o})
		}

		t.SetStyle(table.StyleLight)

		fmt.Fprintln(os.Stderr, t.Render())

		return nil
	},
}

func init() {
	configCmd.AddCommand(showConfigCmd)

	showConfigCmd.Flags().StringP("to-file", "t", "-", "Where to write the config to, '-' for stdout")
}
