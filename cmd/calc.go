/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculates the distance between two vector arrays",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calc called")
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
}
