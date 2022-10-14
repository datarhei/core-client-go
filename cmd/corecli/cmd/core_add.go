package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCoreCmd represents the add command
var addCoreCmd = &cobra.Command{
	Use:   "add [name] [host] (-u|--username) [username] (-p|--password) [password] (-o|--overwrite)",
	Short: "Add a core",
	Long:  `Add a core to the list of known cores and automatically selects it.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		overwrite, _ := cmd.Flags().GetBool("overwrite")

		list := viper.GetStringMapString("cores.list")

		name := args[0]
		host := args[1]

		if !overwrite {
			if _, ok := list[name]; ok {
				return fmt.Errorf("core with name '%s' already exists, use -o to overwrite", name)
			}
		}

		u, err := url.Parse(host)
		if err != nil {
			return err
		}

		if len(username) != 0 {
			if len(password) == 0 {
				u.User = url.User(username)
			} else {
				u.User = url.UserPassword(username, password)
			}
		}

		list[name] = u.String()

		viper.Set("cores.selected", args[0])
		viper.Set("cores.list", list)

		viper.WriteConfig()

		return nil
	},
}

func init() {
	coreCmd.AddCommand(addCoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCoreCmd.Flags().StringP("username", "u", "", "username for the core")
	addCoreCmd.Flags().StringP("password", "p", "", "password for the core")
	addCoreCmd.MarkFlagsRequiredTogether("username", "password")
	addCoreCmd.Flags().BoolP("overwrite", "o", false, "overwrite stored core if it exists")
}
