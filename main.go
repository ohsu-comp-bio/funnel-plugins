// Main "htmlize" application that can load and register plugins from the
// plugin-binaries directory.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"example.com/auth"
	"example.com/plugin"
)

func main() {
	// Load plugins from the plugin-binaries directory.
	var pm plugin.Manager
	if err := pm.LoadPlugins("./plugin-binaries/"); err != nil {
		log.Fatal("loading plugins:", err)
	}
	defer pm.Close()

	// Create a dummy post to wrap the contents with additional metadata.
	auth := &auth.Auth{}

	// Access other arguments
	if len(os.Args) > 1 {
		auth.User = os.Args[1]
	} else {
		log.Fatal("User not provided")
	}

	fmt.Printf("Before auth ➜ User: %s\n", auth.User)
	result := authorize(&pm, auth)
	fmt.Printf(" After auth ➜ User: %s, Token: %s\n", result.User, result.Token)
}

var rolePattern = regexp.MustCompile(":(\\w+):`([^`]*)`")

// authorize turns the text of post.Contents into HTML and returns it; it uses
// the plugin manager to invoke loaded plugins on the contents and the roles
// within it.
func authorize(pm *plugin.Manager, auth *auth.Auth) auth.Auth {
	return pm.ApplyContentsHooks(auth.User)
}
