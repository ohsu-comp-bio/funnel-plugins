/*
Task Execution Service

## Executive Summary The Task Execution Service (TES) API is a standardized schema and API for describing and executing batch execution tasks. A task defines a set of input files, a set of containers and commands to run, a set of output files and some other logging and metadata.  TES servers accept task documents and execute them asynchronously on available compute resources. A TES server could be built on top of a traditional HPC queuing system, such as Grid Engine, Slurm or cloud style compute systems such as AWS Batch or Kubernetes. ## Introduction This document describes the TES API and provides details on the specific endpoints, request formats, and responses. It is intended to provide key information for developers of TES-compatible services as well as clients that will call these TES services. Use cases include:    - Deploying existing workflow engines on new infrastructure. Workflow engines   such as CWL-Tes and Cromwell have extentions for using TES. This will allow   a system engineer to deploy them onto a new infrastructure using a job scheduling   system not previously supported by the engine.    - Developing a custom workflow management system. This API provides a common   interface to asynchronous batch processing capabilities. A developer can write   new tools against this interface and expect them to work using a variety of   backend solutions that all support the same specification.   ## Standards The TES API specification is written in OpenAPI and embodies a RESTful service philosophy. It uses JSON in requests and responses and standard HTTP/HTTPS for information transport. HTTPS should be used rather than plain HTTP except for testing or internal-only purposes. ### Authentication and Authorization Is is envisaged that most TES API instances will require users to authenticate to use the endpoints. However, the decision if authentication is required should be taken by TES API implementers.  If authentication is required, we recommend that TES implementations use an OAuth2  bearer token, although they can choose other mechanisms if appropriate.  Checking that a user is authorized to submit TES requests is a responsibility of TES implementations. ### CORS If TES API implementation is to be used by another website or domain it must implement Cross Origin Resource Sharing (CORS). Please refer to https://w3id.org/ga4gh/product-approval-support/cors for more information about GA4GH’s recommendations and how to implement CORS. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tes

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TesCreateTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TesCreateTaskResponse{}

// TesCreateTaskResponse CreateTaskResponse describes a response from the CreateTask endpoint. It will include the task ID that can be used to look up the status of the job.
type TesCreateTaskResponse struct {
	// Task identifier assigned by the server.
	Id string `json:"id"`
}

type _TesCreateTaskResponse TesCreateTaskResponse

// NewTesCreateTaskResponse instantiates a new TesCreateTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTesCreateTaskResponse(id string) *TesCreateTaskResponse {
	this := TesCreateTaskResponse{}
	this.Id = id
	return &this
}

// NewTesCreateTaskResponseWithDefaults instantiates a new TesCreateTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTesCreateTaskResponseWithDefaults() *TesCreateTaskResponse {
	this := TesCreateTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TesCreateTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TesCreateTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TesCreateTaskResponse) SetId(v string) {
	o.Id = v
}

func (o TesCreateTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TesCreateTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *TesCreateTaskResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTesCreateTaskResponse := _TesCreateTaskResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTesCreateTaskResponse)

	if err != nil {
		return err
	}

	*o = TesCreateTaskResponse(varTesCreateTaskResponse)

	return err
}

type NullableTesCreateTaskResponse struct {
	value *TesCreateTaskResponse
	isSet bool
}

func (v NullableTesCreateTaskResponse) Get() *TesCreateTaskResponse {
	return v.value
}

func (v *NullableTesCreateTaskResponse) Set(val *TesCreateTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTesCreateTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTesCreateTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTesCreateTaskResponse(val *TesCreateTaskResponse) *NullableTesCreateTaskResponse {
	return &NullableTesCreateTaskResponse{value: val, isSet: true}
}

func (v NullableTesCreateTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTesCreateTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


