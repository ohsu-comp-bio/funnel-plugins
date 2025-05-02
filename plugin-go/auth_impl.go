package main

import (
	"fmt"
	"io"
	"net/http"

	"example.com/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/tes"
)

// Here is a real implementation of Authorize that retrieves a "Secret" value for a user
type Authorize struct{}

func (a Authorize) Get(params, headers map[string]string, config *config.Config, task *tes.Task) ([]byte, error) {
	user, ok := params["user"]
	if !ok || user == "" {
		return nil, fmt.Errorf("user is required in params (e.g. params['user'])")
	}
	host, ok := params["host"]
	if !ok || host == "" {
		return nil, fmt.Errorf("host is required in params (e.g. params['host'])")
	}

	shared.Logger.Info("Get", "user", user, "host", host)
	req, err := http.NewRequest("GET", host+user, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	shared.Logger.Info("Response", "status", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	shared.Logger.Info("Response", "body", string(body))
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
