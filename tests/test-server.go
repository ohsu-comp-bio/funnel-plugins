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
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Config  *Credentials `json:"config,omitempty"` // Key-value pairs for configuration
}

type Credentials struct {
	Key    string `json:"key,omitempty"`
	Secret string `json:"secret,omitempty"`
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
	resp := Response{
		Code:    http.StatusOK,
		Message: "Hello, world! To get a token, send a GET request to /token?user=<user>",
		Config:  nil,
	}
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

	// No user provided in the query (Bad Request: 400)
	if user == "" {
		resp := Response{
			Code:    http.StatusBadRequest,
			Message: "User is required",
			Config:  nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, found := userDB[user]

	if found {
		// User found (OK: 200)
		resp := Response{
			Code:    http.StatusOK,
			Message: "User found!",
			Config: &Credentials{
				Key:    token.Key,
				Secret: token.Secret,
			},
		}
		json.NewEncoder(w).Encode(resp)
	} else {
		// User not found (Unauthorized: 401)
		resp := Response{
			Code:    http.StatusUnauthorized,
			Message: "User not authorized",
			Config:  nil,
		}
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
		userDB[row[0]] = Credentials{
			Key:    row[1],
			Secret: row[2],
		}
		mutex.Unlock()
	}
	return userDB, nil
}
