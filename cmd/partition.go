package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var partitionCreateCmd = &cobra.Command{
	Use:   "partition",
	Short: "Create Partition Schema",
	Long:  `create partition [_collection/partition_] [-d (text)]`,
	Run:   partitionCreate,
}

var partitionDeleteCmd = &cobra.Command{
	Use:   "partition",
	Short: "Delete Partition Schema",
	Long:  `delete partition [_collection/partition_]`,
	Run:   partitionDelete,
}

var partitionDescribeCmd = &cobra.Command{
	Use:   "partition",
	Short: "Describe Partition Schema",
	Long:  `describe partition [_collection/partition_]`,
	Run:   partitionDescribe,
}

var partitionListCmd = &cobra.Command{
	Use:   "partition",
	Short: "List Partitions",
	Long:  `list partitions`,
	Run:   partitionList,
}

func init() {
	partitionCreateCmd.Flags().StringP("desc", "d", "", "(Optional) The description of the collection.")
	createCmd.AddCommand(partitionCreateCmd)

	deleteCmd.AddCommand(partitionDeleteCmd)
	describeCmd.AddCommand(partitionDescribeCmd)
	listCmd.AddCommand(partitionListCmd)
}

func partitionCreate(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	fmt.Println("partition create called")
}

func partitionDelete(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	fmt.Println("partition delete called")
}

func partitionDescribe(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	fmt.Println("partition describe called")
}

func partitionList(cmd *cobra.Command, args []string) {
	fmt.Println("args:", args)
	fmt.Println("partitions list called")
}
