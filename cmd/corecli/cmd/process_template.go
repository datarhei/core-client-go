package cmd

import (
	"os"

	"github.com/datarhei/core-client-go/v16/api"

	"github.com/spf13/cobra"
)

var templateProcessCmd = &cobra.Command{
	Use:   "template",
	Short: "Print a template for a process config",
	Long:  "Print a template for a process config.",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := api.ProcessConfig{
			ID:        "",
			Type:      "ffmpeg",
			Reference: "",
			Input: []api.ProcessConfigIO{
				{
					ID:      "",
					Address: "",
					Options: []string{},
				},
			},
			Output: []api.ProcessConfigIO{
				{
					ID:      "",
					Address: "",
					Options: []string{},
					Cleanup: []api.ProcessConfigIOCleanup{
						{
							Pattern:       "",
							MaxFiles:      0,
							MaxFileAge:    0,
							PurgeOnDelete: false,
						},
					},
				},
			},
			Options:        []string{},
			Reconnect:      false,
			ReconnectDelay: 0,
			Autostart:      false,
			StaleTimeout:   0,
			Limits:         api.ProcessConfigLimits{},
		}

		if err := writeJSON(os.Stdout, config, false); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	processCmd.AddCommand(templateProcessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
