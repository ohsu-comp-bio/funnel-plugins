// Main "authorize" application that can load and register plugins from the
// plugin-binaries directory.
//
// Adapted from 'RPC-based plugins in Go' (https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go)
// by Eli Bendersky (https://eli.thegreenplace.net/) 🚀

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"example.com/auth"
	"example.com/plugin"
)

// Register the http.NoBody type to avoid the following error:
// "Error calling Plugin.ProcessContents: gob: type not registered for interface: http.noBody"
func init() {
	gob.Register(http.NoBody)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)

	go func() {
		fmt.Println("Listening on http://localhost:8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("Server exited: %v", err)
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
}

var rolePattern = regexp.MustCompile(":(\\w+):`([^`]*)`")

// authorize turns the text of post.Contents into HTML and returns it; it uses
// the plugin manager to invoke loaded plugins on the contents and the roles
// within it.
func authorize(pm *plugin.Manager, request *http.Request) auth.Auth {
	return pm.ApplyContentsHooks(request)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	// Load plugins from the plugin-binaries directory.
	var pm plugin.Manager
	if err := (&pm).LoadPlugins("./plugin-binaries/"); err != nil {
		log.Fatal("loading plugins:", err)
	}
	defer pm.Close()

	result := authorize(&pm, r)

	if result.Token == "" {
		fmt.Fprintf(w, "Error: User not found ❌\n")
	} else {
		fmt.Fprintf(w, "User: %s, Token: %s ✅", result.User, result.Token)
	}
}
