package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculates the distance between two vector arrays",
	Run:   calc,
}

func init() {
	rootCmd.AddCommand(calcCmd)
}

func calc(cmd *cobra.Command, args []string) {
	fmt.Println("calc called")
}
