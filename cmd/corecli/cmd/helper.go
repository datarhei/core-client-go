package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"

	coreclient "github.com/datarhei/core-client-go/v16"
	"github.com/mattn/go-isatty"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
)

func connectSelectedCore() (coreclient.RestClient, error) {
	selected := viper.GetString("cores.selected")
	list := viper.GetStringMapString("cores.list")

	core, ok := list[selected]
	if !ok {
		return nil, fmt.Errorf("selected core doesn't exist")
	}

	u, err := url.Parse(core)
	if err != nil {
		return nil, fmt.Errorf("invalid data for core: %w", err)
	}

	address := u.Scheme + "://" + u.Host + u.Path
	password, _ := u.User.Password()

	client, err := coreclient.New(coreclient.Config{
		Address:      address,
		Username:     u.User.Username(),
		Password:     password,
		AccessToken:  u.Query().Get("accessToken"),
		RefreshToken: u.Query().Get("refreshToken"),
	})
	if err != nil {
		return nil, fmt.Errorf("can't connect to core at %s: %w", address, err)
	}

	version := client.About().Version.Number
	corename := client.About().Name
	coreid := client.About().ID
	accessToken, refreshToken := client.Tokens()

	query := u.Query()

	query.Set("accessToken", accessToken)
	query.Set("refreshToken", refreshToken)
	query.Set("version", version)
	query.Set("name", corename)
	query.Set("id", coreid)

	u.RawQuery = query.Encode()

	list[selected] = u.String()

	viper.Set("cores.list", list)
	viper.WriteConfig()

	fmt.Fprintln(os.Stderr, client.String())

	return client, nil
}

func getEditor() (string, string, error) {
	editor := viper.GetString("editor")
	if len(editor) == 0 {
		editor = os.Getenv("EDITOR")
	}

	if len(editor) == 0 {
		return "", "", fmt.Errorf("no editor defined")
	}

	path, err := exec.LookPath(editor)
	if err != nil {
		if !errors.Is(err, exec.ErrDot) {
			return "", "", fmt.Errorf("%s: %w", editor, err)
		}
	}

	return editor, path, nil
}

func editData(data []byte) ([]byte, bool, error) {
	editor, _, err := getEditor()
	if err != nil {
		return nil, false, err
	}

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
		editor := exec.Command(editor, filename)
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
			errorData, err := formatJSONError(editedData, err)
			fmt.Printf("%s\n", errorData)
			fmt.Printf("Invalid JSON: %s\n", err)
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

func formatJSONError(input []byte, err error) ([]byte, error) {
	if jsonError, ok := err.(*json.SyntaxError); ok {
		line, character, offsetError := lineAndCharacter(input, int(jsonError.Offset))
		if offsetError != nil {
			return input, err
		}

		return markJSONError(input, line-1, character-1), fmt.Errorf("syntax error at line %d, character %d: %w", line, character, err)
	}

	if jsonError, ok := err.(*json.UnmarshalTypeError); ok {
		line, character, offsetError := lineAndCharacter(input, int(jsonError.Offset))
		if offsetError != nil {
			return input, err
		}

		return markJSONError(input, line-1, character-1), fmt.Errorf("expect type '%s' for '%s' at line %d, character %d: %w", jsonError.Type.String(), jsonError.Field, line, character, err)
	}

	return input, err
}

func lineAndCharacter(input []byte, offset int) (line int, character int, err error) {
	lf := byte(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("couldn't find offset %d within the input", offset)
	}

	// Humans tend to count from 1.
	line = 1
	lastLineCharacters := 0

	for i, b := range input {
		if b == lf {
			line++
			lastLineCharacters = character
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	// Fix the reported offset because it reflects the consumed bytes from
	// parsing and not the actual position of the error.
	if line == 1 {
		character -= 1
	} else {
		character -= 2
		if character < 0 {
			line -= 1
			character = lastLineCharacters
		}
	}

	return line, character, nil
}

func markJSONError(input []byte, line, character int) []byte {
	lf := byte(0x0A)
	output := bytes.Buffer{}
	lineContext := 10

	lines := bytes.Split(input, []byte{lf})

	nlines := len(lines)
	fromLine := line - lineContext
	fromCut := true
	if fromLine < 0 {
		fromLine = 0
		fromCut = false
	}
	toLine := line + lineContext
	toCut := true
	if toLine >= nlines {
		toLine = nlines - 1
		toCut = false
	}

	if fromCut {
		output.Write([]byte(fmt.Sprintf("... %d previous lines omitted ...\n", fromLine)))
	}

	for i := fromLine; i < toLine; i++ {
		l := lines[i]

		output.Write(l)
		output.WriteByte(lf)

		if i == line {
			m := make([]byte, character+1)
			for i := range m {
				m[i] = '_'
			}
			m[character] = '^'

			output.Write(m)
			output.WriteByte(lf)
		}
	}

	if toCut {
		output.Write([]byte(fmt.Sprintf("... %d following lines omitted ...\n", nlines-toLine)))
	}

	return output.Bytes()
}

func formatJSON(d interface{}, useColor bool) (string, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	data = pretty.PrettyOptions(data, &pretty.Options{
		Width:    pretty.DefaultOptions.Width,
		Prefix:   pretty.DefaultOptions.Prefix,
		Indent:   pretty.DefaultOptions.Indent,
		SortKeys: true,
	})

	if !useColor {
		return string(data), nil
	}

	data = pretty.Color(data, nil)

	return string(data), nil
}

func writeJSON(w io.Writer, d interface{}, useColor bool) error {
	color := useColor

	if color {
		if w, ok := w.(*os.File); ok {
			if !isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()) {
				color = false
			}
		} else {
			color = false
		}
	}

	data, err := formatJSON(d, color)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, data)

	return nil
}

func formatByteCountBinary(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d  B", b)
	}

	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
