package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editorCmd represents the metrics command
var editorCmd = &cobra.Command{
	Use:   "editor",
	Short: "Editor related commands",
	Long:  "The editor is first looked up in the config. If not present, the environment variable EDITOR will be used.",
	RunE: func(cmd *cobra.Command, args []string) error {
		editor, path, err := getEditor()
		if err != nil {
			fmt.Printf("Currently no editor is configured. Either set one with 'editor set [path to editor]' or by setting the environment variable EDITOR\n")
			return nil
		}

		fmt.Printf("%s", editor)
		if path != editor {
			fmt.Printf(" (%s)", path)
		}
		fmt.Printf("\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().Bool("raw", false, "Display raw result from the API as JSON")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
