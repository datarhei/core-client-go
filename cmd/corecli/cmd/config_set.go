package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// setConfigCmd represents the list command
var setConfigCmd = &cobra.Command{
	Use:   "set [(-f|--from-file) path]",
	Short: "Set a new config",
	Long:  "Set a new config.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		source, _ := cmd.Flags().GetString("from-file")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		s := os.Stdin

		if source != "-" {
			file, err := os.Open(source)
			if err != nil {
				return err
			}

			s = file
			defer s.Close()
		}

		data, err := io.ReadAll(s)
		if err != nil {
			return err
		}

		var config interface{}

		if err := json.Unmarshal(data, &config); err != nil {
			return err
		}

		if err := client.ConfigSet(config); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(setConfigCmd)

	setConfigCmd.Flags().StringP("from-file", "f", "-", "Where to read the file from, '-' for stdin")
}
