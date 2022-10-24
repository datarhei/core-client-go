package cmd

import (
	"fmt"

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

		f, err := formatJSON(srt, true)
		if err != nil {
			return err
		}

		fmt.Println(f)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(srtCmd)
}
