package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// showMetadataCmd represents the list command
var showMetadataCmd = &cobra.Command{
	Use:   "show [key]?",
	Short: "Show metadat",
	Long:  "Show all metadata or only a specific key",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := ""
		if len(args) == 1 {
			key = args[0]
		}

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		m, err := client.Metadata(key)
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, m, true); err != nil {
			return err
		}

		return nil

	},
}

func init() {
	metadataCmd.AddCommand(showMetadataCmd)
}
