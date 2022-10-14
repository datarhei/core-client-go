package cmd

import (
	"fmt"
	"net/url"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCoreCmd represents the backup command
var listCoreCmd = &cobra.Command{
	Use:   "list",
	Short: "List all known cores",
	Long:  "List all known cores.",
	RunE: func(cmd *cobra.Command, args []string) error {
		selected := viper.GetString("cores.selected")
		list := viper.GetStringMapString("cores.list")

		if len(list) == 0 {
			fmt.Println("No known cores")
			return nil
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"Name", "Host", "Version", "Name", "ID"})

		for name, host := range list {
			if name == selected {
				name = "*" + name
			}

			version := ""
			corename := ""
			coreid := ""

			if u, err := url.Parse(host); err == nil {
				host = u.Scheme + "://" + u.Host
				version = u.Query().Get("version")
				corename = u.Query().Get("name")
				coreid = u.Query().Get("id")
			}

			t.AppendRow(table.Row{name, host, version, corename, coreid})
		}

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		return nil
	},
}

func init() {
	coreCmd.AddCommand(listCoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
