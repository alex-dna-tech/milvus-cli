/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [flags]",
	Short: "Set connection parameters",
	RunE:  config,
}

func init() {
	configCmd.Flags().StringP("url", "u", "", "URL host:port")
	configCmd.Flags().StringP("alias", "a", "", "Alias")

	rootCmd.AddCommand(configCmd)
}

func config(cmd *cobra.Command, args []string) error {
	var (
		url, alias string

		err error
	)

	// Get url flag or prompt
	url, err = FlagOrPrompt(cmd, "url", "", "URL (host:port)", urlValidate)
	if err != nil {
		return err
	}

	// Get alias flag or prompt
	alias, err = FlagOrPrompt(cmd, "alias", "default", "URL (host:port)", urlValidate)
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

func FlagOrPrompt(cmd *cobra.Command, pFlag, pDefault, pLabel string,
	pValidate func(string) error) (string, error) {
	p, err := cmd.Flags().GetString(pFlag)
	if err != nil {
		fmt.Printf("%s flag GetString error: %v\n", capitalize(pFlag), err)
		return "", err
	}
	p = strings.TrimSpace(p)
	if p == "" {
		pPrompt := promptui.Prompt{
			Label:    pLabel,
			Validate: pValidate,
		}

		if pDefault != "" {
			pPrompt.Default = pDefault
		}
		p, err = pPrompt.Run()
		if err != nil {
			fmt.Printf("%s prompt Run error: %v\n", capitalize(pFlag), err)
			return "", err
		}
	} else {
		err = pValidate(p)
		if err != nil {
			fmt.Printf("%s validation error: %v\n", capitalize(pFlag), err)
			return "", err
		}
	}
	return p, nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
