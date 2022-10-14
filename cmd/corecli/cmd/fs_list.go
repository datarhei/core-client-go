package cmd

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// listFSCmd represents the list command
var listFSCmd = &cobra.Command{
	Use:   "list [name] [pattern]? (-s|--sort) [none|name|size|lastmod] (-o|--order) [asc|desc]",
	Short: "List files",
	Long:  "List files on filesystem",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		pattern := ""
		if len(args) == 2 {
			pattern = args[1]
		}

		sort, _ := cmd.Flags().GetString("sort")
		order, _ := cmd.Flags().GetString("order")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		list, err := client.FilesystemList(name, pattern, sort, order)
		if err != nil {
			return err
		}

		t := table.NewWriter()

		t.AppendHeader(table.Row{"Name", "Size", "Last Modification"})

		for _, f := range list {
			lastMod := time.Unix(f.LastMod, 0)
			t.AppendRow(table.Row{f.Name, formatByteCountBinary(uint64(f.Size)), lastMod.Format("2006-01-02 15:04:05")})
		}

		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 2, Align: text.AlignRight},
		})

		t.SetStyle(table.StyleLight)

		fmt.Println(t.Render())

		return nil

	},
}

func init() {
	fsCmd.AddCommand(listFSCmd)

	listFSCmd.Flags().StringP("sort", "s", "none", "Sorting criteria")
	listFSCmd.Flags().StringP("order", "o", "asc", "Sorting direction")
}
