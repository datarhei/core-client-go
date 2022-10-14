package cmd

import (
	"fmt"
	"net/url"
	"os"

	coreclient "github.com/datarhei/core-client-go/v16"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Process related commands",
	Long:  "Process related commands",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("process called")
	//},
}

func init() {
	rootCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	processCmd.PersistentFlags().Bool("raw", false, "Display raw result from the API as JSON")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
