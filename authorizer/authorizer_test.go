package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"example.com/auth"
	"example.com/tes"
)

func readExampleTask(filename string) tes.Task {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var task tes.Task
	err = json.Unmarshal(data, &task)
	if err != nil {
		return tes.Task{}
	}

	return task
}

func TestValidUser(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	task := readExampleTask("../example-tasks/hello-world.json")
	taskData, _ := json.Marshal(task)

	// Create an Authorization Header to pass to the authorizer
	req, _ := http.NewRequest("POST", "http://localhost:8080/s3", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer Alyssa P. Hacker")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}

	var authResp auth.Auth
	json.NewDecoder(resp.Body).Decode(&authResp)

	expected := auth.Auth{
		User:  "Alyssa P. Hacker",
		Token: "<Alyssa's Secret>",
	}

	if authResp.User != expected.User || authResp.Token != expected.Token {
		t.Errorf("Expected (%s, %s), got (%s, %s)",
			expected.User, expected.Token,
			authResp.User, authResp.Token)
	}
}

func TestInvalidUser(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	task := readExampleTask("../example-tasks/hello-world.json")
	taskData, _ := json.Marshal(task)

	// Create an Authorization Header to pass to the authorizer
	req, _ := http.NewRequest("POST", "http://localhost:8080/s3", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Bearer Foo")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.StatusCode)
	}
}

func TestInvalidAuthHeader(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	task := readExampleTask("../example-tasks/hello-world.json")
	taskData, _ := json.Marshal(task)

	// Create an Authorization Header to pass to the authorizer
	req, _ := http.NewRequest("POST", "http://localhost:8080/s3", bytes.NewBuffer(taskData))
	req.Header.Set("Authorization", "Basic Alyssa P. Hacker")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.StatusCode)
	}
}

func TestMissingAuthHeader(t *testing.T) {
	// Setup test file
	os.Setenv("EXAMPLE_USERS", "example-users.csv")

	// Read example task
	task := readExampleTask("../example-tasks/hello-world.json")
	taskData, _ := json.Marshal(task)

	// Create an Authorization Header to pass to the authorizer
	req, _ := http.NewRequest("POST", "http://localhost:8080/s3", bytes.NewBuffer(taskData))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.StatusCode)
	}
}
