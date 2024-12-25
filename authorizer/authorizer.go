package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
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

func (ExampleAuthorizer) Authorize(headers map[string][]string, body []byte) (auth.Auth, error) {
	// TOOD: Currently we're just using the first Authorization header in the request
	// How might we support multiple Authorization headers?
	if len(headers["Authorization"]) == 0 {
		return auth.Auth{}, fmt.Errorf("No Authorization header found")
	}
	authHeader := headers["Authorization"][0]

	// Assuming the Authorization header is in the format "Bearer <username>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return auth.Auth{}, fmt.Errorf("Invalid Authorization header '%v'", authHeader)
	}

	user, err := ExampleAuthorizer{}.getUser(parts[1])
	if err != nil {
		return auth.Auth{}, err
	}

	// Deserialize the task from the request body and check for Task-specific auth info
	if len(body) == 0 {
		return user, fmt.Errorf("No task found in the request body")
	}

	var task tes.TesTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		return auth.Auth{}, fmt.Errorf("Error unmarshalling request body into TES Task: %v", err)
	}

	// TODO: Add Task-specific authorization logic here

	log.Info(context.Background(), "User authenticated", "user", user.User)

	return user, nil
}

func (ExampleAuthorizer) getUser(user string) (auth.Auth, error) {
	// Check if the user is authorized
	// Read the "internal" User Database
	// Here we're just using a CSV file to represent the list of authorized users
	// A real-world example would use a database or an external service (e.g. OAuth)
	userFile := os.Getenv("EXAMPLE_USERS")
	if userFile == "" {
		log.Info(context.Background(), "EXAMPLE_USERS not set, using default example-users.csv")
		userFile = "authorizer/example-users.csv"
	}

	file, err := os.Open(userFile)
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

	return auth.Auth{}, fmt.Errorf("User %s not found", user)
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
