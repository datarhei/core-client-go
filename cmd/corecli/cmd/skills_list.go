package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// listSkillsCmd represents the list command
var listSkillsCmd = &cobra.Command{
	Use:   "list [name]?",
	Short: "List skills",
	Long:  "List FFmpeg skills",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := ""
		if len(args) == 1 {
			name = args[0]
		}

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		skills, err := client.Skills()
		if err != nil {
			return err
		}

		var d interface{}

		switch name {
		case "ffmpeg":
			d = skills.FFmpeg
		case "filters":
			d = skills.Filters
		case "hwaccels":
			d = skills.HWAccels
		case "codecs":
			d = skills.Codecs
		case "devices":
			d = skills.Devices
		case "formats":
			d = skills.Formats
		case "protocols":
			d = skills.Protocols
		default:
			d = skills
		}

		if err := writeJSON(os.Stdout, d, true); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	skillsCmd.AddCommand(listSkillsCmd)
}
