package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	client "github.com/alex-dna-tech/milvus-client"
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"

	"github.com/golang/protobuf/proto"

	"github.com/spf13/cobra"
)

var collectionCreateCmd = &cobra.Command{
	Use:   "collection",
	Short: "Create Collection Schema",
	Long: `create collection -n <text> [-p <text>] [-a] [-d <text>] field1:field_type:field_desc field2:field_type/dim:field_desc...

Multiple fields params as arguments with schema <fieldName>:<dataType[/dimOfVector]>:<desc> format.

For example:
milvus-cli create collection \
	--name car \
	--primary id \
	--autoid \
	--desc "Car collection" \
	"id:Int64:primary field" \
	"vector:FloatVector/128:vector field" \
	color:Int64:color \
	brand:Int64:brand`,
	Run: collectionCreate,
}

var collectionDeleteCmd = &cobra.Command{
	Use:   "collection",
	Short: "Delete Collection Schema",
	Long:  `delete collection <name>`,
	Run:   collectionDelete,
}

var collectionDescribeCmd = &cobra.Command{
	Use:   "collection",
	Short: "Describe Collection Schema",
	Long:  `describe collection <name>`,
	Run:   collectionDescribe,
}

var collectionListCmd = &cobra.Command{
	Use:   "collection",
	Short: "List Collections",
	Long:  `list collections`,
	Run:   collectionList,
}

func init() {
	createCmd.AddCommand(collectionCreateCmd)
	collectionCreateCmd.Flags().StringP("name", "n", "", "Collection name.")
	collectionCreateCmd.Flags().StringP("primary", "p", "", "The name of the primary key field.")
	collectionCreateCmd.Flags().BoolP("autoid", "a", false, "(Optional) Flag to generate IDs automatically.")
	collectionCreateCmd.Flags().StringP("desc", "d", "", "(Optional) The description of the collection.")

	deleteCmd.AddCommand(collectionDeleteCmd)
	describeCmd.AddCommand(collectionDescribeCmd)
	listCmd.AddCommand(collectionListCmd)
}

func collectionCreate(cmd *cobra.Command, args []string) {
	name, err := getStringVal(cmd, "name", "", "Name",
		func(n string) error {
			if len(n) == 0 {
				return errors.New("name is empty")
			}
			return nil
		},
	)
	cobra.CheckErr(err)

	primary, err := getStringVal(cmd, "primary", "", "Primary field name",
		func(n string) error {
			return nil
		},
	)
	cobra.CheckErr(err)

	var autoid bool
	if primary != "" {
		autoid, err = getBoolVal(cmd, "autoid", "Auto ID")
		cobra.CheckErr(err)
	}

	desc, err := getStringVal(cmd, "desc", "", "Description",
		func(n string) error {
			return nil
		},
	)
	cobra.CheckErr(err)

	var (
		fields []*schemapb.FieldSchema
		exists bool
	)

	if len(args) > 0 {
		for _, f := range args {
			fs, err := parseField(f, primary, autoid, &exists)
			cobra.CheckErr(err)
			fields = append(fields, fs)
		}
	} else {
		// TODO: implement field prompt
	}

	schema := &schemapb.CollectionSchema{
		Name:        name,
		Description: desc,
		AutoID:      autoid,
		Fields:      fields,
	}

	schemaBytes, err := proto.Marshal(schema)
	cobra.CheckErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c, err := client.New(ctx, getServerURL())
	cobra.CheckErr(err)
	defer c.Close()

	resp, err := c.CreateCollection(ctx, &milvuspb.CreateCollectionRequest{
		CollectionName: name,
		Schema:         schemaBytes,
	})
	cobra.CheckErr(err)
	fmt.Printf("Response-> %#v\nError-> %#v\n", resp, err)

	//TODO: collection create called need to be implemented
	// fmt.Println("collection create called", name, primary, desc)
}

func parseField(field, primary string, autoid bool, exists *bool) (*schemapb.FieldSchema, error) {
	s := strings.Split(field, ":")

	if len(s) < 2 {
		return nil, fmt.Errorf("'%s' parse error: at least <field_name> and <field_type> is required", field)
	}
	s[0] = strings.TrimSpace(s[0])
	if s[0] == "" {
		return nil, fmt.Errorf("'%s' parse error: <field_name> must be not empty", field)
	}
	s[1] = strings.TrimSpace(s[1])
	if s[1] == "" {
		return nil, fmt.Errorf("'%s' parse error: <field_type> must be not empty", field)
	}

	fs := &schemapb.FieldSchema{}
	fs.Name = s[0]
	// Split datatype, dimension
	dts := strings.Split(s[1], "/")
	dt, ok := schemapb.DataType_value[dts[0]]
	if !ok {
		return nil, fmt.Errorf("unknown field type %s, allowed types: %v", s[1], allowedDataTypes())
	}
	fs.DataType = schemapb.DataType(dt)
	if strings.HasSuffix(dts[0], "Vector") {
		//TODO: implement check if dimension is integer
		fs.TypeParams = []*commonpb.KeyValuePair{{Key: "dim", Value: dts[1]}}
	}

	if primary != "" {
		if primary == s[0] {
			*exists = true
			fs.IsPrimaryKey = true
			fs.AutoID = autoid
		}
	}

	return fs, nil
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
