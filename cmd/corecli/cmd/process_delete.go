package cmd

import (
	"fmt"

	coreclient "github.com/datarhei/core-client-go/v16"

	"github.com/spf13/cobra"
)

// deleteProcessCmd represents the show command
var deleteProcessCmd = &cobra.Command{
	Use:   "delete [processid] (-r|--reference)",
	Short: "Delete the process with the given ID",
	Long:  "Delete the process with the given ID or reference.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		reference, _ := cmd.Flags().GetBool("reference")

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		if reference {
			list, err := client.ProcessList(coreclient.ProcessListOptions{
				RefPattern: id,
			})
			if err != nil {
				return err
			}

			for _, p := range list {
				if err := client.ProcessDelete(p.ID); err != nil {
					fmt.Printf("%s error %s\n", p.ID, err.Error())
				} else {
					fmt.Printf("%s delete\n", p.ID)
				}
			}

			return nil
		}

		if err := client.ProcessDelete(id); err != nil {
			return err
		}

		fmt.Printf("%s delete\n", id)

		return nil
	},
}

func init() {
	processCmd.AddCommand(deleteProcessCmd)

	deleteProcessCmd.Flags().BoolP("reference", "r", false, "Interpret the processid as reference and delete all processes with that reference")
}
