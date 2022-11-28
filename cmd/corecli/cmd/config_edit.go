package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// editConfigCmd represents the list command
var editConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit core config",
	Long:  "Edit the config of the core",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		_, config, err := client.Config()
		if err != nil {
			return err
		}

		var data []byte

		data, err = json.MarshalIndent(config.Config, "", "   ")
		if err != nil {
			return err
		}

		editedData, modified, err := editData(data)
		if err != nil {
			return err
		}

		if !modified {
			// They are the same, nothing has been changed. No need to store the metadata
			fmt.Printf("No changes. Config will not be updated.\n")
			return nil
		}

		var editedConfig interface{}

		if err := json.Unmarshal(editedData, &editedConfig); err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, editedConfig, true); err != nil {
			return err
		}

		return client.ConfigSet(editedConfig)
	},
}

func init() {
	configCmd.AddCommand(editConfigCmd)
}
