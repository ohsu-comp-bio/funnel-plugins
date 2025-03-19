package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type Response struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Config  map[string]string `json:"config"` // Key-value pairs for configuration
}

type Credentials struct {
	User   string `json:"user"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type GenericS3Storage struct {
	Key    string
	Secret string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/token", tokenHandler)

	// Currently hardcoding the endpoint of the token service
	// TODO: This should be made configurable similar to the plugin (see plugin/auth_impl.go)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler for root endpoint
func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{http.StatusOK, "Hello, world! To get a token, send a GET request to /token?user=<user>", nil}
	json.NewEncoder(w).Encode(resp)
}

// Handler for retrieving user tokens
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	// Load users from the CSV file
	userDB, err := loadUsers("tests/example-users.csv")
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	user := r.URL.Query().Get("user")
	if user == "" {
		resp := Response{http.StatusBadRequest, "User is required", nil}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, found := userDB[user]

	if found {
		config := map[string]string{
			"GenericS3Storage.Key":    token.Key,
			"GenericS3Storage.Secret": token.Secret,
		}
		resp := Response{http.StatusOK, "User found!", config}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := Response{http.StatusUnauthorized, "User not found", nil}
		json.NewEncoder(w).Encode(resp)
	}
}

// Load user tokens from the CSV file
func loadUsers(filename string) (map[string]Credentials, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	userDB := make(map[string]Credentials)
	mutex := sync.RWMutex{}
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}
		mutex.Lock()
		userDB[row[0]] = Credentials{User: row[0], Key: row[1], Secret: row[2]}
		mutex.Unlock()
	}
	return userDB, nil
}
