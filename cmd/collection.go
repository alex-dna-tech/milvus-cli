package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collectionCreateCmd = &cobra.Command{
	Use:   "collection",
	Short: "Create Collection Schema",
	Long: `create collection _name_ -f (text) -p (text) [-a] [-d (text)]
For example:
milvus-cli create collection car -f id:INT64:primary_field -f vector:FLOAT_VECTOR:128 -f color:INT64:color -f brand:INT64:brand -p id -a true -d 'car_collection'`,
	Run: collectionCreate,
}

var collectionDeleteCmd = &cobra.Command{
	Use:   "collection",
	Short: "Delete Collection Schema",
	Long:  `delete collection _name_`,
	Run:   collectionDelete,
}

var collectionDescribeCmd = &cobra.Command{
	Use:   "collection",
	Short: "Describe Collection Schema",
	Long:  `describe collection _name_`,
	Run:   collectionDescribe,
}

var collectionListCmd = &cobra.Command{
	Use:   "collection",
	Short: "List Collections",
	Long:  `list collections`,
	Run:   collectionList,
}

func init() {
	collectionCreateCmd.Flags().StringP("field", "f", "", "(Multiple) The field schema in the <fieldName>:<dataType>:<dimOfVector/desc> format.")
	collectionCreateCmd.Flags().StringP("primary", "p", "", "The name of the primary key field.")
	collectionCreateCmd.Flags().BoolP("autoid", "a", false, "(Optional) Flag to generate IDs automatically.")
	collectionCreateCmd.Flags().StringP("desc", "d", "", "(Optional) The description of the collection.")

	createCmd.AddCommand(collectionCreateCmd)
	deleteCmd.AddCommand(collectionDeleteCmd)
	describeCmd.AddCommand(collectionDescribeCmd)
	listCmd.AddCommand(collectionListCmd)
}

func collectionCreate(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	//TODO: collection create called need to be implemented
	fmt.Println("collection create called")
}

func collectionDelete(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	//TODO: collection delete called need to be implemented
	fmt.Println("collection delete called")
}

func collectionDescribe(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	//TODO: collection describe called need to be implemented
	fmt.Println("collection describe called")
}

func collectionList(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	//TODO: collection list called need to be implemented
	fmt.Println("collections list called")
}
