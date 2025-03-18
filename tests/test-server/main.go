package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var (
	userDB = make(map[string]string)
	mutex  = &sync.RWMutex{}
)

// Load user tokens from the CSV file
func loadUsers(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Assuming first row is headers
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}
		if len(row) < 2 {
			continue // Skip malformed rows
		}
		mutex.Lock()
		userDB[row[0]] = row[1]
		mutex.Unlock()
	}
	return nil
}

// Handler for root endpoint
func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, world! To get a token, send a GET request to /token?user=<user>"}
	json.NewEncoder(w).Encode(response)
}

// Handler for retrieving user tokens
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	// Load users from the CSV file
	if err := loadUsers("tests/example-users.csv"); err != nil {
		fmt.Println("error loading users:", err)
		return
	}

	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, `{"error": "user is required"}`, http.StatusBadRequest)
		return
	}

	mutex.RLock()
	token, exists := userDB[user]
	mutex.RUnlock()

	if exists {
		json.NewEncoder(w).Encode(map[string]string{"user": user, "token": token})
	} else {
		http.Error(w, fmt.Sprintf(`{"error": "user '%s' not found"}`, user), http.StatusNotFound)
	}
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/token", tokenHandler)

	// Currently hardcoding the endpoint of the token service
	// TODO: This should be made configurable similar to the plugin (see plugin/auth_impl.go)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
