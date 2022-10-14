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
var updateProcessCmd = &cobra.Command{
	Use:   "update [processid]",
	Short: "Update the process with the given ID",
	Long:  "Update the process with the given ID. The process with the given ID will be stopped and deleted. The new process doesn't neet to have the same ID.",
	Args:  cobra.ExactArgs(1),
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

		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		if err := client.ProcessUpdate(id, config); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(updateProcessCmd)

	updateProcessCmd.Flags().String("from-file", "-", "Load process config from file or stdin")
}
