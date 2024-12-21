// Main "htmlize" application that can load and register plugins from the
// plugin-binaries directory.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"example.com/auth"
	"example.com/plugin"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)

	fmt.Println("Listening on http://localhost:8080")
	_ = http.ListenAndServe(":8080", mux)
}

var rolePattern = regexp.MustCompile(":(\\w+):`([^`]*)`")

// authorize turns the text of post.Contents into HTML and returns it; it uses
// the plugin manager to invoke loaded plugins on the contents and the roles
// within it.
func authorize(pm *plugin.Manager, authHeader string) auth.Auth {
	return pm.ApplyContentsHooks(authHeader)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	// Load plugins from the plugin-binaries directory.
	var pm plugin.Manager
	if err := pm.LoadPlugins("./plugin-binaries/"); err != nil {
		log.Fatal("loading plugins:", err)
	}
	defer pm.Close()

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result := authorize(&pm, authHeader)

	if result.Token == "" {
		fmt.Fprintf(w, "Error: User not found\n")
	} else {
		fmt.Fprintf(w, "User: %s, Token: %s\n", result.User, result.Token)
	}
}
