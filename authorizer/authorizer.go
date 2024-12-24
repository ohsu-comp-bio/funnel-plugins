package main

import (
	"bufio"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"example.com/auth"
	"example.com/plugin"
	"example.com/tes"
	"github.com/golang/gddo/log"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

// Register the http.NoBody type to avoid the following error:
// "Error calling Plugin.ProcessContents: gob: type not registered for interface: http.noBody"
func init() {
	gob.Register(http.NoBody)
}

type ExampleAuthorizer struct{}

func (ExampleAuthorizer) Hooks() []string {
	return []string{"contents"}
}

func (ExampleAuthorizer) ProcessContents(request *http.Request) (auth.Auth, error) {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return auth.Auth{}, fmt.Errorf("No Authorization header")
	}

	// Assuming the Authorization header is in the format "Bearer <username>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return auth.Auth{}, fmt.Errorf("Invalid Authorization header")
	}

	if request.Body == nil {
		return auth.Auth{}, fmt.Errorf("No request body found in the request")
	}

	// Check request body for task-specific auth info
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return auth.Auth{}, fmt.Errorf("Error reading request body: %v", err)
	}

	// Deserialize the task from the request body
	var task tes.TesTask
	err = json.Unmarshal(body, &task)
	if err != nil {
		return auth.Auth{}, fmt.Errorf("Error unmarshalling request body into TES Task: %v", err)
	}

	user, err := ExampleAuthorizer{}.getUser(parts[1], task)
	if err != nil {
		return auth.Auth{}, err
	}

	log.Info(context.Background(), "User authenticated", "user", user.User)

	return user, nil
}

func (ExampleAuthorizer) getUser(user string, task tes.TesTask) (auth.Auth, error) {
	// Check if the user is authorized
	// Read the "internal" User Database
	// Here we're just using a CSV file to represent the list of authorized users
	// A real-world example would use a database or an external service (e.g. OAuth)
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
