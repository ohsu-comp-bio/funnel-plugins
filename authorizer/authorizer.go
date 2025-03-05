package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/auth"
	"example.com/plugin"
	"example.com/tes"
	"github.com/golang/gddo/log"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

type ExampleAuthorizer struct{}

func (ExampleAuthorizer) Hooks() []string {
	return []string{"contents"}
}

func (ExampleAuthorizer) Authorize(authHeader http.Header, task tes.Task) (auth.Auth, error) {
	// TOOD: Currently we're just using the first Authorization header in the request
	// How might we support multiple Authorization headers?
	if authHeader == nil {
		return auth.Auth{}, fmt.Errorf("No Authorization header found")
	}

	user := authHeader.Get("Authorization")
	if user == "" {
		return auth.Auth{}, fmt.Errorf("No user found in Authorization header %s", user)
	}

	if !strings.HasPrefix(user, "Bearer ") {
		return auth.Auth{}, fmt.Errorf("Invalid Authorization header: expected 'Bearer <token>', got %s", user)
	}

	user = strings.TrimPrefix(user, "Bearer ")

	creds, err := ExampleAuthorizer{}.getUser(user)
	if err != nil {
		log.Info(context.Background(), "401: User Unauthorized", "user", user)
		return auth.Auth{}, err
	}

	log.Info(context.Background(), "200: User Authorized", "user", creds.User)

	return creds, nil
}

func (ExampleAuthorizer) getUser(user string) (auth.Auth, error) {
	// Query the external service for user authorization
	url := fmt.Sprintf("http://localhost:8080/s3?user=%s", user)
	resp, err := http.Get(url)
	if err != nil {
		return auth.Auth{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return auth.Auth{}, fmt.Errorf("User %s not found", user)
	}

	var creds auth.Auth
	if err := json.NewDecoder(resp.Body).Decode(&creds); err != nil {
		return auth.Auth{}, err
	}

	return creds, nil
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level: hclog.Warn,
	})

	plugins := map[string]goplugin.Plugin{
		"authorize": &plugin.AuthorizePlugin{
			Impl: ExampleAuthorizer{},
		},
	}

	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: plugin.Handshake,
		Plugins:         plugins,
		Logger:          logger,
	})
}
