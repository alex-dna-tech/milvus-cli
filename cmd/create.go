/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create milvus elements",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		// fmt.Printf("conf.client: %#v\n", viper.Get("client"))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
