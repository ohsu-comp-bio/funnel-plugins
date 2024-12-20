// RPC scaffolding for our server and client, using net/rpc.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package plugin

import (
	"log"
	"net/rpc"

	"example.com/auth"
)

// Types for RPC args/reply messages.

type HooksArgs struct{}

type HooksReply struct {
	Hooks []string
}

type AuthArgs struct {
	Value string
	Auth  auth.Auth
}

type AuthReply struct {
	Auth auth.Auth
	Err  error
}

type RoleArgs struct {
	Role  string
	Value string
	Auth  auth.Auth
}

type RoleReply struct {
	Value string
}

// PluginServerRPC is used by plugins to map RPC calls from the clients to
// methods of the Auth interface.
type PluginServerRPC struct {
	Impl Authorizer
}

func (s *PluginServerRPC) Hooks(args HooksArgs, reply *HooksReply) error {
	reply.Hooks = s.Impl.Hooks()
	return nil
}

func (s *PluginServerRPC) ProcessContents(args AuthArgs, reply *AuthReply) error {
	reply.Auth, reply.Err = s.Impl.ProcessContents(args.Value)
	return nil
}

func (s *PluginServerRPC) ProcessRole(args RoleArgs, reply *RoleReply) error {
	reply.Value = s.Impl.ProcessRole(args.Role, args.Value)
	return nil
}

// PluginClientRPC is used by clients (main application) to translate the
// Htmlize interface of plugins to RPC calls.
type PluginClientRPC struct {
	client *rpc.Client
}

func (c *PluginClientRPC) Hooks() []string {
	var reply HooksReply
	if err := c.client.Call("Plugin.Hooks", HooksArgs{}, &reply); err != nil {
		log.Printf("Error calling Plugin.Hooks: %v", err)
		return nil
	}
	return reply.Hooks
}

func (c *PluginClientRPC) ProcessContents(val string) (auth.Auth, error) {
	var reply AuthReply
	err := c.client.Call("Plugin.ProcessContents", AuthArgs{Value: val}, &reply)

	if err != nil {
		log.Printf("Error calling Plugin.ProcessContents: %v", err)
	}

	return reply.Auth, reply.Err
}

func (c *PluginClientRPC) ProcessRole(role string, val string) string {
	var reply RoleReply
	err := c.client.Call("Plugin.ProcessRole", RoleArgs{Role: role, Value: val}, &reply)

	if err != nil {
		log.Printf("Error calling Plugin.ProcessRole: %v", err)
	}

	return reply.Value
}
