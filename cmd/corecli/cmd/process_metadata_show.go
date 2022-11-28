package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// showMetadataProcessCmd represents the show command
var showMetadataProcessCmd = &cobra.Command{
	Use:   "show [processid] [key]?",
	Short: "Show the metadata of the process with the given ID",
	Long:  "Show the metadata of the process with the given ID",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		key := ""
		if len(args) == 2 {
			key = args[1]
		}

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		metadata, err := client.ProcessMetadata(id, key)
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, metadata, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	metadataProcessCmd.AddCommand(showMetadataProcessCmd)
}
