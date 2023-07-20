package cmd

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config [flags]",
	Aliases: []string{"cfg"},
	Short:   "Set connection parameters",
	RunE:    config,
}

func init() {
	configCmd.Flags().StringP("url", "u", "", "URL host:port")
	configCmd.Flags().StringP("alias", "a", "", "Server connection alias")

	rootCmd.AddCommand(configCmd)
}

func config(cmd *cobra.Command, args []string) error {
	var (
		url, alias string
		err        error
	)

	// Get url flag or prompt
	url, err = getStringVal(cmd, "url", "", "URL (host:port)", urlValidate)
	if err != nil {
		return err
	}

	// Get alias flag or prompt
	alias, err = getStringVal(cmd, "alias", "default", "Alias", aliasValidate)
	if err != nil {
		return err
	}

	viper.Set("client."+alias+".url", url)
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// Validate URL
func urlValidate(input string) error {
	if input == "" {
		return errors.New("empty URL input")
	}

	s := strings.Split(input, ":")
	if len(s) != 2 {
		return errors.New("input must be host:port")
	}
	if s[0] == "" {
		return errors.New("host must be not empty")
	}
	//TODO: check host to be valid hostname on IP
	if s[1] == "" {
		return errors.New("port must be not empty")
	}
	p, err := strconv.Atoi(s[1])
	if err != nil {
		return errors.New("port must be a number")
	}
	if p < 0 || p > 65536 {
		return errors.New("port must be in range 0-65535")
	}
	return nil
}

// Validate alias
func aliasValidate(input string) error {
	if input == "" {
		return errors.New("empty alias input")
	}
	return nil
}
