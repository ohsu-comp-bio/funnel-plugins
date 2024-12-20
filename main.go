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
	"strings"

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
func authorize(pm *plugin.Manager, auth *auth.Auth) auth.Auth {
	return pm.ApplyContentsHooks(auth.User)
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

	// Assuming the Authorization header is in the format "Bearer <username>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := parts[1]
	auth := &auth.Auth{User: user}
	result := authorize(&pm, auth)

	if result.Token == "" {
		fmt.Fprintf(w, "Error: User %s not found\n", user)
	} else {
		fmt.Fprintf(w, "User: %s, Token: %s\n", result.User, result.Token)
	}
}
