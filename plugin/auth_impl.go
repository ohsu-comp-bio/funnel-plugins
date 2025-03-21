// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"
	"net/http"

	"example.com/shared"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of Authorize that retrieves a "Secret" value for a user
type Authorize struct{}

func (Authorize) Get(user string) ([]byte, error) {
	if user == "" {
		return nil, fmt.Errorf("user is required (e.g. ./authorize <user>)")
	}
	// Currently hardcoding the endpoint of the token service
	// TODO: This should be made configurable similar to the test server (see tests/test-server/main.go)
	resp, err := http.Get("http://localhost:8080/token?user=" + user)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"authorize": &shared.AuthorizePlugin{Impl: &Authorize{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
