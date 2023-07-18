package cmd

import (
	"context"
	"fmt"

	"time"

	client "github.com/alex-dna-tech/milvus-client"
	pb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/spf13/cobra"
)

// versionCmd show milvus version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show milvus version",
	Run:   version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c, err := client.New(ctx, getServerURL())
	cobra.CheckErr(err)
	defer c.Close()

	resp, err := c.GetVersion(ctx, &pb.GetVersionRequest{})
	cobra.CheckErr(err)

	fmt.Println(resp.GetVersion())
}
