// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
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

	result, err := authorize.Get(user)
	if err != nil {
		return "", fmt.Errorf("failed to authorize: %w", err)
	}

	return string(result), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <user>\n", os.Args[0])
	}

	out, err := run(os.Args[1], "build/plugins")
	if err != nil {
		fmt.Println("Error calling plugin:", err)
		os.Exit(1)
	}

	fmt.Println(out)
	os.Exit(0)
}
