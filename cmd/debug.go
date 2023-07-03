package cmd

import (
	"context"
	"fmt"
	"time"

	"alex-dna-tech/milvus-cli/internal/client"

	pb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"

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

	c, err := client.New(target)
	cobra.CheckErr(err)
	resp, err := c.ShowCollections(ctx, &pb.ShowCollectionsRequest{})
	cobra.CheckErr(err)
	fmt.Printf("Response-> %#v\nError-> %#v\n", resp.CollectionNames, err)
	err = c.Close()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
