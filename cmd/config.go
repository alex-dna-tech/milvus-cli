/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if input == "" {
				return errors.New("Empty URL input")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "URL host:port",
			Validate: validate,
			Default:  "localhost:19530",
		}

		url, err := prompt.Run()
		if err != nil {
			fmt.Printf("URL prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", url)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringP("url", "u", "", "URL host:port")
	configCmd.Flags().StringP("alias", "a", "", "Connect alias")
}
