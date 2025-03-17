package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Greeter is the interface that we're exposing as a plugin.
type Authorizer interface {
	Authorize() string
}

// Here is an implementation that talks over RPC
type AuthorizerRPC struct{ client *rpc.Client }

func (a *AuthorizerRPC) Authorize() string {
	var resp string
	err := a.client.Call("Plugin.Authorize", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type AuthorizerRPCServer struct {
	// This is the real implementation
	Impl Authorizer
}

func (s *AuthorizerRPCServer) Authorize(args interface{}, resp *string) error {
	*resp = s.Impl.Authorize()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods:
//  1. Server must return an RPC server for this plugin
//
// type. We construct an AuthorizerRPCServer for this.
//
//  2. Client must return an implementation of our interface that communicates
//
// over an RPC client. We return AuthorizerRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type AuthorizerPlugin struct {
	// Impl Injection
	Impl Authorizer
}

func (p *AuthorizerPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &AuthorizerRPCServer{Impl: p.Impl}, nil
}

func (AuthorizerPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &AuthorizerRPC{client: c}, nil
}
