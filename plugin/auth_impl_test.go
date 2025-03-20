package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Define a struct that matches the expected JSON response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Config  Credentials `json:"config"` // Key-value pairs for configuration
}

type Credentials struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func TestAuthorizedUser(t *testing.T) {
	auth := Authorize{}
	raw, err := auth.Get("example")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Parse actual JSON response
	var actual Response
	if err := json.Unmarshal(raw, &actual); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	// Define expected response
	expected := Response{
		Code:    200,
		Message: "User found!",
		Config: Credentials{
			Key:    "key1",
			Secret: "secret1",
		},
	}

	// Compare using reflect.DeepEqual
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestUnauthorizedUser(t *testing.T) {
	auth := Authorize{}
	raw, err := auth.Get("error")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Parse actual JSON response
	var actual Response
	if err := json.Unmarshal(raw, &actual); err != nil {
		t.Fatalf("failed to parse JSON response: %v", err)
	}

	// Define expected response
	expected := Response{
		Code:    401,
		Message: "User not found",
		Config:  Credentials{},
	}

	// Compare using reflect.DeepEqual
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
