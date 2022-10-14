package cmd

import (
	"github.com/spf13/cobra"
)

// coreCmd represents the core command
var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Core related commands",
	Long:  "Core related commands.",
}

func init() {
	rootCmd.AddCommand(coreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
