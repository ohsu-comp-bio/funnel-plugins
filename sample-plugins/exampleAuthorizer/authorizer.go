package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/auth"
	"example.com/plugin"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

type ExampleAuthorizer struct{}

func (ExampleAuthorizer) Hooks() []string {
	return []string{"contents"}
}

func (ExampleAuthorizer) ProcessContents(authHeader string) (auth.Auth, error) {
	// Assuming the Authorization header is in the format "Bearer <username>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return auth.Auth{}, fmt.Errorf("Invalid Authorization header")
	}

	user := parts[1]

	// Read file
	file, err := os.Open("example-users.csv")
	if err != nil {
		return auth.Auth{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if strings.Contains(line, user) {
			result := strings.Split(line, ",")

			auth := auth.Auth{
				User:  result[0],
				Token: result[1],
			}

			return auth, nil
		}
	}

	return auth.Auth{}, fmt.Errorf("User not found")
}

func (ExampleAuthorizer) ProcessRole(role string, val string) string {
	fmt.Printf("Processing role %s\n", role)
	return val
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
