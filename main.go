// Main "authorize" application that can load and register plugins from the
// plugin-binaries directory.
//
// Adapted from 'RPC-based plugins in Go' (https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go)
// by Eli Bendersky (https://eli.thegreenplace.net/) 🚀

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/auth"
	"example.com/plugin"
)

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

// authorize turns the text of post.Contents into HTML and returns it; it uses
// the plugin manager to invoke loaded plugins on the contents within it.
func authorize(pm *plugin.Manager, headers map[string][]string, body []byte) (auth.Auth, error) {
	return pm.ApplyContentsHooks(headers, body)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	// Load plugins from the plugin-binaries directory.
	var pm plugin.Manager
	if err := (&pm).LoadPlugins("./plugin-binaries/"); err != nil {
		log.Fatal("loading plugins:", err)
	}
	defer pm.Close()

	headers := r.Header

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	resp, err := authorize(&pm, headers, body)

	if err != nil {
		fmt.Fprintf(w, "Error: %v ❌\n", err)
	} else {
		fmt.Fprintf(w, "Response: %v ✅\n", resp)
	}
}
