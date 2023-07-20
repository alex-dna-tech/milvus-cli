package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add", "c"},
	Short:   "Create milvus elements",
	Run:     create,
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(cmd *cobra.Command, args []string) {
	items := []string{"collection", "partition"}
	prompt := promptui.Select{
		Label: "Create",
		Items: items,
	}

	_, create, err := prompt.Run()
	cobra.CheckErr(err)

	switch create {
	case "collection":
		collectionCreateCmd.Run(collectionCreateCmd, []string{})
	case "partition":
		partitionCreateCmd.Run(partitionCreateCmd, []string{})
	}
}
