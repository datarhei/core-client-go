package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// probeProcessCmd represents the show command
var probeProcessCmd = &cobra.Command{
	Use:   "probe [processid]",
	Short: "Probe the process with the given ID",
	Long:  "Probe the process with the given ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		probe, err := client.ProcessProbe(id)
		if err != nil {
			return err
		}

		f, err := formatJSON(probe, true)
		if err != nil {
			return err
		}

		fmt.Println(f)

		return nil
	},
}

func init() {
	processCmd.AddCommand(probeProcessCmd)
}
