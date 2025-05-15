package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/plugins/proto"
	"github.com/ohsu-comp-bio/funnel/plugins/shared"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/hashicorp/go-plugin"
	"github.com/ohsu-comp-bio/funnel/tes"
)

// Here is a real implementation of Authorize that retrieves a "Secret" value for a user
type Authorize struct{}

func (a Authorize) PluginAction(params map[string]string, headers map[string]*proto.StringList, config *config.Config, task *tes.Task, taskType proto.Type) (*proto.JobResponse, error) {
	shared.Logger.Debug("Params: ", params)
	shared.Logger.Debug("Headers: ", headers)
	shared.Logger.Debug("Config: ", config)
	shared.Logger.Debug("Task: ", task)

	host, ok := params["Host"]
	if !ok || host == "" {
		return &proto.JobResponse{
				Code:    400,
				Message: "host is required in params (e.g. params['host'])"},
			fmt.Errorf("host is required in params (e.g. params['host'])")
	}
	user := headers["authorization"].Values[0]
	if !ok || user == "" {
		return &proto.JobResponse{
				Code:    400,
				Message: "user is required in the Auth header"},
			fmt.Errorf("user is required in the Auth header")
	}
	shared.Logger.Info("GET", "user", user, "host", host)
	Body := &proto.Job{
		Config:  config,
		Task:    task,
		Headers: headers,
		Params:  params,
		Type:    taskType,
	}

	var requestBody []byte
	var err error
	marshalOptions := protojson.MarshalOptions{}
	requestBody, err = marshalOptions.Marshal(Body)
	if err != nil {
		return &proto.JobResponse{
				Code:    400,
				Message: fmt.Sprintf("error marshaling JSON body: %#v", Body)},
			fmt.Errorf("error marshaling JSON body: %w. Should be of type proto.Job", err)
	}

	var req *http.Request
	var method string
	switch taskType {
	case proto.Type_CREATE:
		method = "PUT"
	case proto.Type_GET:
		method = "GET"
	case proto.Type_CANCEL:
		method = "DELETE"
	default:
		return &proto.JobResponse{
				Code:    400,
				Message: fmt.Sprintf("unsupported task type: %v", taskType)},
			fmt.Errorf("unsupported task type: %v", taskType)
	}

	req, err = http.NewRequest(method, host+user, bytes.NewBuffer(requestBody))
	if err != nil {
		return &proto.JobResponse{
				Code:    500,
				Message: fmt.Sprintf("error creating request: %w", err)},
			fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		for _, val := range v.Values {
			req.Header.Add(k, val)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &proto.JobResponse{
				Code:    500,
				Message: fmt.Sprintf("error receiving request: %w", err)},
			fmt.Errorf("error receiving request: %w", err)
	}

	defer resp.Body.Close()
	shared.Logger.Info("Response", "status", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &proto.JobResponse{
				Code:    500,
				Message: fmt.Sprintf("%w", err)},
			fmt.Errorf("%w", err)
	}
	receivedData := &proto.JobResponse{}
	unmarshalOptions := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = unmarshalOptions.Unmarshal(body, receivedData)
	if err != nil {
		return &proto.JobResponse{
				Code:    500,
				Message: fmt.Sprintf("%w", err)},
			fmt.Errorf("%w", err)
	}

	shared.Logger.Info("+++++++++++++++++++++++++++++RESP: ", receivedData)
	return receivedData, nil
}

func main() {
	log.Println("Server: registering gob types")
	gob.Register(&config.TimeoutConfig_Duration{})
	gob.Register(&config.TimeoutConfig_Disabled{})

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"authorize": &shared.AuthorizePlugin{Impl: &Authorize{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
