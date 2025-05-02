package shared

import (
	"fmt"
	"net/rpc"

	"example.com/proto"
	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/tes"
)

// GetResponse holds the response for the RPC Get call.
type GetResponse struct {
	Value []byte
}

// RPCClient is an implementation of Authorize that talks over RPC.
type RPCClient struct {
	client *rpc.Client
}

func (m *RPCClient) Get(params, headers map[string]string, config *config.Config, task *tes.Task) ([]byte, error) {
	var resp GetResponse
	err := m.client.Call("Plugin.Get", &proto.GetRequest{
		Params:  params,
		Headers: headers,
		Config:  config,
		Task:    task,
	}, &resp)
	if err != nil {
		return nil, fmt.Errorf("RPC Get call failed: %w", err)
	}
	return resp.Value, nil
}

type RPCServer struct {
	Impl Authorize
}

func (m *RPCServer) Get(args *proto.GetRequest, resp *GetResponse) error {
	// Call the implementation's Get method with the arguments
	v, err := m.Impl.Get(args.Params, args.Headers, args.Config, args.Task)
	if err != nil {
		return fmt.Errorf("authorize implementation failed: %w", err)
	}
	resp.Value = v
	return nil
}
