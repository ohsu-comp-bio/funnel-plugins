package main

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"testing"

	"example.com/auth"
)

func setupTestFile() string {
	// Create a temporary file for testing
	file, _ := os.CreateTemp("", "example-users.csv")
	file.WriteString("testuser1,token1\n")
	file.WriteString("testuser2,token2\n")
	file.WriteString("testuser3,token3\n")
	file.Close()

	return file.Name()
}

func readExampleTask(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func TestProcessContents_ValidUser(t *testing.T) {
	// Setup test file
	testFile := setupTestFile()
	defer os.Remove(testFile)

	// Replace file path for the test
	os.Setenv("CSV_FILE", testFile)

	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")

	// Create a sample HTTP request
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer Alyssa P. Hacker")
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	resp, err := authenticator.ProcessContents(req)

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
	_, err := authenticator.ProcessContents(req)

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
	_, err := authenticator.ProcessContents(req)

	if err == nil || !strings.Contains(err.Error(), "No Authorization header") {
		t.Errorf("Expected error about missing header, got %v", err)
	}
}

func TestProcessContents_UserNotFound(t *testing.T) {
	// Setup test file
	testFile := setupTestFile()
	defer os.Remove(testFile)

	// Replace file path for the test
	os.Setenv("CSV_FILE", testFile)

	// Read example task
	taskData := readExampleTask("../example-tasks/hello-world.json")
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer unknownuser")
	req.Header.Set("Content-Type", "application/json")

	authenticator := ExampleAuthorizer{}
	_, err := authenticator.ProcessContents(req)

	if err == nil || !strings.Contains(err.Error(), "User not found") {
		t.Errorf("Expected error about user not found, got %v", err)
	}
}
