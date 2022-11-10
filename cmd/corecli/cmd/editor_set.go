package cmd

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setEditorCmd represents the metrics command
var setEditorCmd = &cobra.Command{
	Use:   "set [path to editor]",
	Short: "Set an editor",
	Long:  "Set an editor in the config. Any value in the environment variable EDITOR will be ignored.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		editor := args[0]

		path, err := exec.LookPath(editor)
		if err != nil {
			if !errors.Is(err, exec.ErrDot) {
				return fmt.Errorf("can't find editor: %w", err)
			}
		}

		fmt.Printf("%s", editor)
		if path != editor {
			fmt.Printf(" (%s)", path)
		}
		fmt.Printf("\n")

		viper.Set("editor", editor)
		viper.WriteConfig()

		return nil
	},
}

func init() {
	editorCmd.AddCommand(setEditorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().Bool("raw", false, "Display raw result from the API as JSON")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
