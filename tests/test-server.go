package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/plugins/proto"
	"github.com/ohsu-comp-bio/funnel/plugins/shared"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	csvFile = flag.String("users-csv", "example-users.csv", "Path to the CSV file containing user tokens")
)

func main() {
	flag.Parse() // Parse the command-line flags

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/token", tokenHandler)

	fmt.Printf("Server is running on http://0.0.0.0:8080 using users from: %s\n", *csvFile)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// Handler for root endpoint
func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp := &proto.GetResponse{ // Note the pointer here
		Code:    http.StatusOK,
		Message: "Hello, world! To get a token, send a GET request to /token?user=[USER]",
	}
	encodeResponse(w, resp)
}

// Handler for retrieving user tokens
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received token request:", r)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	receivedData := &proto.GetRequest{}
	unmarshalOptions := protojson.UnmarshalOptions{
		DiscardUnknown: true, // Or false, depending on your needs
	}

	err = unmarshalOptions.Unmarshal(bodyBytes, receivedData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal GetRequest from JSON using protojson: %v", err), http.StatusBadRequest)
		return
	}

	// Now you can access the parsed config and task objects
	fmt.Printf("Received Config: %#v\n", receivedData.Config)
	fmt.Printf("Received Task: %#v\n", receivedData.Task)
	fmt.Printf("Received Headers: %#v\n", receivedData.Headers)

	// Load users from the CSV file specified by the flag
	userDB, err := loadUsers(*csvFile)
	if err != nil {
		fmt.Println("Error loading users:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("USERS: ", userDB)

	user := r.URL.Query().Get("user")

	// No user provided in the query (Bad Request: 400)
	if user == "" {
		resp := proto.GetResponse{
			Code:    http.StatusBadRequest,
			Message: "User is required",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, found := userDB[user]

	if found {
		shared.Logger.Debug("Found token for user:", user, "Token Key:", token.AmazonS3.AWSConfig.Key)
		receivedData.Config.AmazonS3.AWSConfig.Key = token.AmazonS3.AWSConfig.Key
		receivedData.Config.AmazonS3.AWSConfig.Secret = token.AmazonS3.AWSConfig.Secret
		encodeResponse(w, &proto.GetResponse{Code: http.StatusOK, Config: receivedData.Config, Task: receivedData.Task})
	} else {
		shared.Logger.Warn("User not authorized:", user)
		encodeResponse(w, &proto.GetResponse{Code: http.StatusUnauthorized, Message: "User not authorized", Config: receivedData.Config, Task: receivedData.Task})
	}
}

func encodeResponse(w http.ResponseWriter, resp *proto.GetResponse) {
	w.WriteHeader(int(resp.Code))
	marshalOptions := protojson.MarshalOptions{} // You can customize options if needed
	responseBody, err := marshalOptions.Marshal(resp)
	if err != nil {
		shared.Logger.Error("Error marshaling protojson response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Important: Set the correct Content-Type
	w.Write(responseBody)
}

// Load user tokens from the CSV file
func loadUsers(filename string) (map[string]config.Config, error) {
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

	userDB := make(map[string]config.Config)
	mutex := sync.RWMutex{}
	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}
		mutex.Lock()
		userDB[row[0]] = config.Config{
			AmazonS3: &config.AmazonS3Storage{
				AWSConfig: &config.AWSConfig{
					Key:    row[1],
					Secret: row[2],
				},
			},
		}
		mutex.Unlock()
	}
	return userDB, nil
}
