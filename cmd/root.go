// Package cmd creates cobra Commands
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, clientAlias string

// var viperConfig *viper.Viper

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "milvus-cli",
	Short: "MilvusDB CLI Interface",
	Long: `Milvus Command-Line Interface (CLI) is a command-line tool that supports
database connection, data operations, and import and export of data.

Based on Milvus Goland SDK, it allows the execution of commands through a
terminal using interactive command-line prompts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.PersistentFlags().StringVarP(&clientAlias, "client", "c", "default", "client alias stored in config")
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
		fmt.Fprintf(os.Stderr, "Using: config file - \"%s\", server alias - \"%s\"\n", viper.ConfigFileUsed(), clientAlias)
	}

	// Overwrite if environment variable is set
	viper.AutomaticEnv()

	err := viper.SafeWriteConfig()
	if err != nil {
		return
	}
}
