// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"example.com/shared"
	"github.com/hashicorp/go-plugin"
)

func run() error {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Logger:          shared.Logger,
		Cmd:             exec.Command("sh", "-c", os.Getenv("FUNNEL_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC}})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("authorize")
	if err != nil {
		return err
	}

	// We should have an Authorize function now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	authorize := raw.(shared.Authorize)

	if len(os.Args) < 2 {
		return fmt.Errorf("Usage: %s <user>", os.Args[0])
	}
	result, err := authorize.Get(os.Args[1])
	if err != nil {
		return err
	}

	fmt.Println(string(result))

	return nil
}

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	if err := run(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
