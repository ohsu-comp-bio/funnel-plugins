package shared

import (
	"context"

	"example.com/proto"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client proto.AuthorizeClient }

func (m *GRPCClient) Get(user string, host string) ([]byte, error) {
	resp, err := m.client.Get(context.Background(), &proto.GetRequest{
		User: user,
		Host: host,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Authorize
}

func (m *GRPCServer) Get(
	ctx context.Context,
	req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(req.User, req.Host)
	return &proto.GetResponse{Value: v}, err
}
