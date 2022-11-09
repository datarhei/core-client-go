package cmd

import (
	"github.com/spf13/cobra"
)

// deleteMetadataCmd represents the list command
var deleteMetadataCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete metadata",
	Long:  "Delete a specific metadata key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		return client.MetadataSet(key, nil)
	},
}

func init() {
	metadataCmd.AddCommand(deleteMetadataCmd)
}
