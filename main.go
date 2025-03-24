package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"example.com/shared"
)

func run(user string, dir string) (string, error) {
	m := &shared.Manager{}
	defer m.Close()

	authorize, err := m.Client(dir)
	if err != nil {
		return "", fmt.Errorf("failed to get client: %w", err)
	}

	resp, err := authorize.Get(user, "http://localhost:8080/token?user=")
	if err != nil {
		return "", fmt.Errorf("failed to authorize: %w", err)
	}

	// Pretty print directly from raw JSON response
	var out bytes.Buffer
	if err := json.Indent(&out, resp, "", "  "); err != nil {
		return "", fmt.Errorf("error formatting JSON: %w", err)
	}

	return out.String(), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <user>\n", os.Args[0])
		os.Exit(1)
	}

	out, err := run(os.Args[1], "build/plugins")
	if err != nil {
		fmt.Println("Error calling plugin:", err)
		os.Exit(1)
	}

	fmt.Println(out)
	os.Exit(0)
}
