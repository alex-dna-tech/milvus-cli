package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe milvus elements",
	Run:   describe,
}

func init() {
	rootCmd.AddCommand(describeCmd)
}

func describe(cmd *cobra.Command, args []string) {
	fmt.Println("describe called")
}
