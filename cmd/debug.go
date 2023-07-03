package cmd

import (
	"context"
	"fmt"
	"time"

	"alex-dna-tech/milvus-cli/internal/client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "air auto exec function",
	Run:   debug,
}

func debug(cmd *cobra.Command, args []string) {
	target := viper.GetString("client." + clientAlias + ".url")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := client.NewMilvusClient(target)
	cobra.CheckErr(err)
	defer c.Close()
	resp, err := c.GetVersion(ctx)
	cobra.CheckErr(err)
	fmt.Printf("Response-> %#v\nError-> %#v\n", resp, err)
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
