package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "d"},
	Short:   "Delete milvus elements",
	Run:     del,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
func del(cmd *cobra.Command, args []string) {
	fmt.Println("delete called")
}
