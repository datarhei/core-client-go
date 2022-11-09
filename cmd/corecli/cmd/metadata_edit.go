package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/datarhei/core-client-go/v16/api"
	"github.com/spf13/cobra"
)

// editMetadataCmd represents the list command
var editMetadataCmd = &cobra.Command{
	Use:   "edit [key]",
	Short: "Edit metadata",
	Long:  "Edit a specific metadata key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		client, err := connectSelectedCore()
		if err != nil {
			return err
		}

		found := true

		m, err := client.Metadata(key)
		if err != nil {
			apierr, ok := err.(api.Error)
			if !ok {
				return err
			}

			if apierr.Code != 404 {
				return err
			}

			found = false
		}

		var data []byte

		if found {
			data, err = json.MarshalIndent(m, "", "   ")
			if err != nil {
				return err
			}
		}

		editedData, modified, err := editData(data)
		if err != nil {
			return err
		}

		if !modified {
			// They are the same, nothing has been changed. No need to store the metadata
			return nil
		}

		var em api.Metadata

		if err := json.Unmarshal(editedData, &em); err != nil {
			return err
		}

		f, err := formatJSON(em, true)
		if err != nil {
			return err
		}

		fmt.Println(f)

		return client.MetadataSet(key, em)
	},
}

func init() {
	metadataCmd.AddCommand(editMetadataCmd)
}

func editData(data []byte) ([]byte, bool, error) {
	file, err := os.CreateTemp("", "corecli_*")
	if err != nil {
		return nil, false, err
	}

	filename := file.Name()

	defer os.Remove(filename)

	_, err = file.Write(data)
	file.Close()

	if err != nil {
		return nil, false, err
	}

	for {
		editor := exec.Command("nano", filename)
		editor.Stdout = os.Stdout
		editor.Stderr = os.Stderr
		editor.Stdin = os.Stdin
		if err := editor.Run(); err != nil {
			return nil, false, err
		}

		editedData, err := os.ReadFile(filename)
		if err != nil {
			return nil, false, err
		}

		var x interface{}

		if err := json.Unmarshal(editedData, &x); err != nil {
			fmt.Printf("%s\n", editedData)
			fmt.Printf("Invalid JSON: %s\n", formatJSONError(editedData, err))
			fmt.Printf("Do you want to re-open the editor (Y/n)? ")

			var char rune
			if _, err := fmt.Scanf("%c", &char); err != nil {
				return nil, false, err
			}

			if char == '\n' || char == 'Y' || char == 'y' {
				continue
			}

			return nil, false, fmt.Errorf("invalid JSON: %w", err)
		}

		return editedData, !bytes.Equal(data, editedData), nil
	}
}

func formatJSONError(input []byte, err error) error {
	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, offsetError := lineAndCharacter(input, int(jsonError.Offset))
		if offsetError != nil {
			return err
		}

		return fmt.Errorf("syntax error at line %d, character %d: %w", line, character, err)
	}

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, offsetError := lineAndCharacter(input, int(jsonError.Offset))
		if offsetError != nil {
			return err
		}

		return fmt.Errorf("expect type '%s' for '%s' at line %d, character %d: %w", jsonError.Type.String(), jsonError.Field, line, character, err)
	}

	return err
}

func lineAndCharacter(input []byte, offset int) (line int, character int, err error) {
	lf := byte(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("couldn't find offset %d within the input", offset)
	}

	// Humans tend to count from 1.
	line = 1

	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	return line, character, nil
}
