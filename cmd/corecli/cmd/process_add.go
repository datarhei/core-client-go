package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/datarhei/core-client-go/v16/api"
	"github.com/spf13/cobra"
)

// updateProcessCmd represents the update command
var addProcessCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a process",
	Long:  "Add a process to the core.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fromFile, _ := cmd.Flags().GetString("from-file")
		if len(fromFile) == 0 {
			return fmt.Errorf("no process configuration file provided")
		}

		reader := os.Stdin

		if fromFile != "-" {
			file, err := os.Open(fromFile)
			if err != nil {
				return err
			}

			reader = file
		}

		data, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		config := api.ProcessConfig{}

		if err := json.Unmarshal(data, &config); err != nil {
			return err
		}

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		if err := client.ProcessAdd(config); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(addProcessCmd)

	addProcessCmd.Flags().String("from-file", "-", "Load process config from file or stdin")
}
