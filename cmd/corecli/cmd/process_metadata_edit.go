package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/datarhei/core-client-go/v16/api"
	"github.com/spf13/cobra"
)

// editMetadataProcessCmd represents the list command
var editMetadataProcessCmd = &cobra.Command{
	Use:   "edit [processid] [key]",
	Short: "Edit metadata",
	Long:  "Edit a specific metadata key",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		key := args[1]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		found := true

		m, err := client.ProcessMetadata(id, key)
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

		if err := writeJSON(os.Stdout, em, true); err != nil {
			return err
		}

		return client.ProcessMetadataSet(id, key, em)
	},
}

func init() {
	metadataProcessCmd.AddCommand(editMetadataProcessCmd)
}
