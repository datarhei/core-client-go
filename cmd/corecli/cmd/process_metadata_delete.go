package cmd

import (
	"github.com/spf13/cobra"
)

// deleteMetadataProcessCmd represents the list command
var deleteMetadataProcessCmd = &cobra.Command{
	Use:   "delete [processid] [key]",
	Short: "Delete metadata",
	Long:  "Delete a specific metadata key",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		key := args[1]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		return client.ProcessMetadataSet(id, key, nil)
	},
}

func init() {
	metadataProcessCmd.AddCommand(deleteMetadataProcessCmd)
}
