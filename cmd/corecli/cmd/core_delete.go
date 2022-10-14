package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCoreCmd represents the delete command
var deleteCoreCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Remove a core",
	Long:  `Remove a core from the list of known cores.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		selected := viper.GetString("cores.selected")
		list := viper.GetStringMapString("cores.list")

		if name == selected {
			selected = ""
		}

		delete(list, name)

		viper.Set("cores.selected", selected)
		viper.Set("cores.list", list)

		viper.WriteConfig()

		fmt.Printf("Deleted core '%s'\n", name)

		return nil
	},
}

func init() {
	coreCmd.AddCommand(deleteCoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
