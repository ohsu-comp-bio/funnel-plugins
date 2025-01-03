# Go API client for openapi

## Executive Summary
The Task Execution Service (TES) API is a standardized schema and API for describing and executing batch execution tasks. A task defines a set of input files, a set of containers and commands to run, a set of output files and some other logging and metadata.

TES servers accept task documents and execute them asynchronously on available compute resources. A TES server could be built on top of a traditional HPC queuing system, such as Grid Engine, Slurm or cloud style compute systems such as AWS Batch or Kubernetes.
## Introduction
This document describes the TES API and provides details on the specific endpoints, request formats, and responses. It is intended to provide key information for developers of TES-compatible services as well as clients that will call these TES services. Use cases include:

  - Deploying existing workflow engines on new infrastructure. Workflow engines
  such as CWL-Tes and Cromwell have extentions for using TES. This will allow
  a system engineer to deploy them onto a new infrastructure using a job scheduling
  system not previously supported by the engine.

  - Developing a custom workflow management system. This API provides a common
  interface to asynchronous batch processing capabilities. A developer can write
  new tools against this interface and expect them to work using a variety of
  backend solutions that all support the same specification.


## Standards
The TES API specification is written in OpenAPI and embodies a RESTful service philosophy. It uses JSON in requests and responses and standard HTTP/HTTPS for information transport. HTTPS should be used rather than plain HTTP except for testing or internal-only purposes.
### Authentication and Authorization
Is is envisaged that most TES API instances will require users to authenticate to use the endpoints. However, the decision if authentication is required should be taken by TES API implementers.

If authentication is required, we recommend that TES implementations use an OAuth2  bearer token, although they can choose other mechanisms if appropriate.

Checking that a user is authorized to submit TES requests is a responsibility of TES implementations.
### CORS
If TES API implementation is to be used by another website or domain it must implement Cross Origin Resource Sharing (CORS). Please refer to https://w3id.org/ga4gh/product-approval-support/cors for more information about GA4GH’s recommendations and how to implement CORS.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://raw.githubusercontent.com/ga4gh/tes/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*TaskServiceAPI* | [**CancelTask**](docs/TaskServiceAPI.md#canceltask) | **Post** /tasks/{id}:cancel | CancelTask
*TaskServiceAPI* | [**CreateTask**](docs/TaskServiceAPI.md#createtask) | **Post** /tasks | CreateTask
*TaskServiceAPI* | [**GetServiceInfo**](docs/TaskServiceAPI.md#getserviceinfo) | **Get** /service-info | GetServiceInfo
*TaskServiceAPI* | [**GetTask**](docs/TaskServiceAPI.md#gettask) | **Get** /tasks/{id} | GetTask
*TaskServiceAPI* | [**ListTasks**](docs/TaskServiceAPI.md#listtasks) | **Get** /tasks | ListTasks


## Documentation For Models

 - [Service](docs/Service.md)
 - [ServiceOrganization](docs/ServiceOrganization.md)
 - [ServiceType](docs/ServiceType.md)
 - [TesCreateTaskResponse](docs/TesCreateTaskResponse.md)
 - [TesExecutor](docs/TesExecutor.md)
 - [TesExecutorLog](docs/TesExecutorLog.md)
 - [TesFileType](docs/TesFileType.md)
 - [TesInput](docs/TesInput.md)
 - [TesListTasksResponse](docs/TesListTasksResponse.md)
 - [TesOutput](docs/TesOutput.md)
 - [TesOutputFileLog](docs/TesOutputFileLog.md)
 - [TesResources](docs/TesResources.md)
 - [TesServiceInfo](docs/TesServiceInfo.md)
 - [TesServiceType](docs/TesServiceType.md)
 - [TesState](docs/TesState.md)
 - [TesTask](docs/TesTask.md)
 - [TesTaskLog](docs/TesTaskLog.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### BasicAuth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextBasicAuth, openapi.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



