package cmd

import (
	"context"
	"fmt"
	"time"

	client "github.com/alex-dna-tech/milvus-client"

	pb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "air auto exec function",
	Run:   debug,
}

func debug(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c, err := client.New(ctx, getServerURL())
	cobra.CheckErr(err)
	defer c.Close()

	resp, err := c.ShowCollections(ctx, &pb.ShowCollectionsRequest{})
	cobra.CheckErr(err)
	fmt.Printf("Response-> %#v\nError-> %#v\n", resp.CollectionNames, err)
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
