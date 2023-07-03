package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create milvus elements",
	Run: func(cmd *cobra.Command, args []string) {
		target := viper.GetString("client." + clientAlias + ".url")

		client, err := client.NewMilvusClient(target)
		defer client.Close()
		cobra.CheckErr(err)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.GetVersion(ctx)
		cobra.CheckErr(err)
		fmt.Printf("#-> %#v\n%#v\n", resp, err)
		fmt.Println("create called")
		// fmt.Printf("conf.client: %#v\n", viper.Get("client"))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
