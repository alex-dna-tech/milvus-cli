// Package cmd creates cobra Commands
package cmd

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

var (
	cfgFile, serverAlias string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "milvus-cli",
	Short: "MilvusDB CLI Interface",
	Long: `Milvus Command-Line Interface (CLI) is a command-line tool that supports
database connection, data operations, and import and export of data.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.milvus-cli.yml)")
	rootCmd.PersistentFlags().StringVarP(&serverAlias, "server", "s", "default", "server connection alias from config")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".milvus-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yml")
		viper.SetConfigName(".milvus-cli")
	}

	if !viper.IsSet("client.default.url") {
		viper.SetDefault("client.default.url", "localhost:19530")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintf(os.Stderr, "Using: config file - \"%s\", server alias - \"%s\"\n", viper.ConfigFileUsed(), serverAlias)
	}

	// Overwrite if environment variable is set
	viper.AutomaticEnv()

	err := viper.SafeWriteConfig()
	if err != nil {
		return
	}
}

func getServerURL() string {
	return viper.GetString("client." + serverAlias + ".url")
}

// getStringVal use flag value or prompt and validate it
func getStringVal(cmd *cobra.Command, pFlag, pDefault, pLabel string,
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

// getBoolVal use flag value or prompt if value false
func getBoolVal(cmd *cobra.Command, bFlagName string, bLabel string) (bool, error) {
	b, err := cmd.Flags().GetBool(bFlagName)
	if err != nil {
		fmt.Printf("%s flag GetString error: %v\n", capitalize(bFlagName), err)
		return false, err
	}
	if !b {
		bSelect := promptui.Select{
			Label: bLabel,
			Items: []string{"No", "Yes"},
		}

		i, _, err := bSelect.Run()
		if err != nil {
			fmt.Printf("%s select Run error: %v\n", capitalize(bFlagName), err)
			return false, err
		}

		if i == 0 {
			b = false
		} else {
			b = true
		}
	}

	return b, nil
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func allowedDataTypes() (allowed []string) {
	for _, v := range schemapb.DataType_name {
		allowed = append(allowed, v)
	}
	return
}
