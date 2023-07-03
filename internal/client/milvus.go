package client

import (
	"context"

	pb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// milvusClient struct must imlpement MilvusClient interface
var _ MilvusClient = (*milvusClient)(nil)

type milvusClient struct {
	conn    *grpc.ClientConn
	service pb.MilvusServiceClient
	callOpt []grpc.CallOption
	dialOpt []grpc.DialOption
}

// GetVersion get milvus version
func (mc *milvusClient) GetVersion(ctx context.Context) (string, error) {
	resp, err := mc.service.GetVersion(ctx, &pb.GetVersionRequest{}, mc.callOpt...)
	if err != nil {
		return "", err
	}

	return resp.Version, nil
}

// MilvusClient is the interface used to communicate with Milvus
type MilvusClient interface {
	// Close close the remaining connection resources
	Close() error

	// GetVersion get milvus version
	GetVersion(ctx context.Context) (string, error)

	// CreateCollection create collection

}

// Close close the remaining connection resources
func (mc *milvusClient) Close() error {
	return mc.conn.Close()
}

// NewMilvusClient function creating MilvusClient
func NewMilvusClient(serverAddr string, opts ...grpc.DialOption) (*milvusClient, error) {
	// Add default options if no option is given
	if opts == nil {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	return &milvusClient{conn: conn, service: pb.NewMilvusServiceClient(conn)}, nil
}

// WithCallOptions method
func (mc *milvusClient) WithCallOptions(opts ...grpc.CallOption) *milvusClient {
	mc.callOpt = append(mc.callOpt, opts...)
	return mc
}
