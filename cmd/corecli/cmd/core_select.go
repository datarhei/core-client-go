package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// selectCoreCmd represents the select command
var selectCoreCmd = &cobra.Command{
	Use:   "select [name]",
	Short: "Select a core for use",
	Long:  `Select a core for use from the list of known cores.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		selected := viper.GetString("cores.selected")
		list := viper.GetStringMapString("cores.list")

		if name == selected {
			return nil
		}

		if _, ok := list[name]; !ok {
			return fmt.Errorf("core with name '%s' not found", name)
		}

		viper.Set("cores.selected", name)
		viper.WriteConfig()

		return nil
	},
}

func init() {
	coreCmd.AddCommand(selectCoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
