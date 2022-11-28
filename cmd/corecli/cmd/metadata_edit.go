package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/datarhei/core-client-go/v16/api"
	"github.com/spf13/cobra"
)

// editMetadataCmd represents the list command
var editMetadataCmd = &cobra.Command{
	Use:   "edit [key]",
	Short: "Edit metadata",
	Long:  "Edit a specific metadata key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		found := true

		m, err := client.Metadata(key)
		if err != nil {
			apierr, ok := err.(api.Error)
			if !ok {
				return err
			}

			if apierr.Code != 404 {
				return err
			}

			found = false
		}

		var data []byte

		if found {
			data, err = json.MarshalIndent(m, "", "   ")
			if err != nil {
				return err
			}
		}

		editedData, modified, err := editData(data)
		if err != nil {
			return err
		}

		if !modified {
			// They are the same, nothing has been changed. No need to store the metadata
			fmt.Printf("No changes. Metadata will not be updated.")
			return nil
		}

		var em api.Metadata

		if err := json.Unmarshal(editedData, &em); err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, editedData, true); err != nil {
			return err
		}

		return client.MetadataSet(key, em)
	},
}

func init() {
	metadataCmd.AddCommand(editMetadataCmd)
}
