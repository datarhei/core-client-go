package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// unsetEditorCmd represents the metrics command
var unsetEditorCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset the editor in the config",
	Long:  "Unset the editor in the config. The value of the environment variable EDITOR will be used.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set("editor", "")
		viper.WriteConfig()

		fmt.Printf("The editor has been unset. The value of the environment variable EDITOR will be used.\n")

		return nil
	},
}

func init() {
	editorCmd.AddCommand(unsetEditorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().Bool("raw", false, "Display raw result from the API as JSON")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
