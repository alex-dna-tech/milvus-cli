/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [flags]",
	Short: "Set connection parameters",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			url, alias string
			err        error
		)

		// Get url flag or prompt
		url, err = cmd.Flags().GetString("url")
		if err != nil {
			fmt.Printf("URL FlagGetString error: %v\n", err)
			return
		}
		if url == "" {
			fmt.Printf("URL flag not set %v\n", err)

			urlPrompt := promptui.Prompt{
				Label:    "URL (host:port)",
				Validate: urlValidate,
			}

			url, err = urlPrompt.Run()
		} else {
			err = urlValidate(url)
			if err != nil {
				fmt.Printf("URL validation error: %v\n", err)
				return
			}
		}

		// Get alias flag or prompt
		alias, err = cmd.Flags().GetString("alias")
		if err != nil {
			fmt.Printf("Alias FlagGetString error: %v\n", err)
			return
		}
		if alias == "" {
			aliasPrompt := promptui.Prompt{
				Label:    "Alias",
				Validate: aliasValidate,
				Default:  "default",
			}

			alias, err = aliasPrompt.Run()
		} else {
			err = aliasValidate(alias)
			if err != nil {
				fmt.Printf("Alias validation error: %v\n", err)
				return
			}
		}

		viper.Set("client."+alias+".url", url)
		viper.WriteConfig()
	},
}

func init() {
	configCmd.Flags().StringP("url", "u", "", "URL host:port")
	configCmd.Flags().StringP("alias", "a", "", "Alias")

	rootCmd.AddCommand(configCmd)
}

// Validate URL
func urlValidate(input string) error {
	if input == "" {
		return errors.New("Empty URL input")
	}

	s := strings.Split(input, ":")
	if len(s) != 2 {
		return errors.New("Input must be host:port")
	}
	if s[0] == "" {
		return errors.New("Host must be not empty")
	}
	//TODO: check host to be valid hostname on IP
	if s[1] == "" {
		return errors.New("Port must be not empty")
	}
	p, err := strconv.Atoi(s[1])
	if err != nil {
		return errors.New("Port must be a number")
	}
	if p < 0 && p > 65536 {
		return errors.New("Port must be in range 0-65535")
	}
	return nil
}

// Validate alias
func aliasValidate(input string) error {
	if input == "" {
		return errors.New("Empty alias input")
	}
	return nil
}
