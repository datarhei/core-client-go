package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// srtCmd represents the metrics command
var srtCmd = &cobra.Command{
	Use:   "srt",
	Short: "SRT related commands",
	Long:  "SRT related commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		srt, err := client.SRTChannels()
		if err != nil {
			return err
		}

		if err := writeJSON(os.Stdout, srt, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(srtCmd)
}
