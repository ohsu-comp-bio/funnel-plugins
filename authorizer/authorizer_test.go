package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"example.com/auth"
)

func readExampleTask(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func TestProcessContents_ValidUser(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")

	// Create a sample HTTP request
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer Alyssa P. Hacker")
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	body, _ := io.ReadAll(req.Body)
	resp, err := authenticator.Authorize(req.Header, body)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := auth.Auth{
		User:  "Alyssa P. Hacker",
		Token: "<Alyssa's Secret>",
	}

	if resp.User != expected.User || resp.Token != expected.Token {
		t.Errorf("Expected (%s, %s), got (%s, %s)",
			expected.User, expected.Token,
			resp.User, resp.Token)
	}
}

func TestProcessContents_InvalidHeader(t *testing.T) {
	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "InvalidHeader")
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	body, _ := io.ReadAll(req.Body)
	_, err := authenticator.Authorize(req.Header, body)

	if err == nil || !strings.Contains(err.Error(), "Invalid Authorization header") {
		t.Errorf("Expected error about invalid header, got %v", err)
	}
}

func TestProcessContents_NoHeader(t *testing.T) {
	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	body, _ := io.ReadAll(req.Body)
	_, err := authenticator.Authorize(req.Header, body)

	if err == nil || !strings.Contains(err.Error(), "No Authorization header found") {
		t.Errorf("Expected error about missing header, got %v", err)
	}
}

func TestProcessContents_UserNotFound(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer unknownuser")
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	body, _ := io.ReadAll(req.Body)
	_, err := authenticator.Authorize(req.Header, body)

	if err == nil || !strings.Contains(err.Error(), "User unknownuser not found") {
		t.Errorf("Expected error about user not found, got %v", err)
	}
}
