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

// checks if the TesInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TesInput{}

// TesInput Input describes Task input files.
type TesInput struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// REQUIRED, unless \"content\" is set.  URL in long term storage, for example:  - s3://my-object-store/file1  - gs://my-bucket/file2  - file:///path/to/my/file  - /path/to/my/file
	Url *string `json:"url,omitempty"`
	// Path of the file inside the container. Must be an absolute path.
	Path string `json:"path"`
	Type *TesFileType `json:"type,omitempty"`
	// File content literal.  Implementations should support a minimum of 128 KiB in this field and may define their own maximum.  UTF-8 encoded  If content is not empty, \"url\" must be ignored.
	Content *string `json:"content,omitempty"`
	// Indicate that a file resource could be accessed using a streaming interface, ie a FUSE mounted s3 object. This flag indicates that using a streaming mount, as opposed to downloading the whole file to the local scratch space, may be faster despite the latency and overhead. This does not mean that the backend will use a streaming interface, as it may not be provided by the vendor, but if the capacity is avalible it can be used without degrading the performance of the underlying program.
	Streamable *bool `json:"streamable,omitempty"`
}

type _TesInput TesInput

// NewTesInput instantiates a new TesInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTesInput(path string) *TesInput {
	this := TesInput{}
	this.Path = path
	var type_ TesFileType = FILE
	this.Type = &type_
	return &this
}

// NewTesInputWithDefaults instantiates a new TesInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTesInputWithDefaults() *TesInput {
	this := TesInput{}
	var type_ TesFileType = FILE
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TesInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TesInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TesInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TesInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TesInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TesInput) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TesInput) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TesInput) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TesInput) SetUrl(v string) {
	o.Url = &v
}

// GetPath returns the Path field value
func (o *TesInput) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *TesInput) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *TesInput) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TesInput) GetType() TesFileType {
	if o == nil || IsNil(o.Type) {
		var ret TesFileType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetTypeOk() (*TesFileType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TesInput) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TesFileType and assigns it to the Type field.
func (o *TesInput) SetType(v TesFileType) {
	o.Type = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TesInput) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TesInput) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *TesInput) SetContent(v string) {
	o.Content = &v
}

// GetStreamable returns the Streamable field value if set, zero value otherwise.
func (o *TesInput) GetStreamable() bool {
	if o == nil || IsNil(o.Streamable) {
		var ret bool
		return ret
	}
	return *o.Streamable
}

// GetStreamableOk returns a tuple with the Streamable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesInput) GetStreamableOk() (*bool, bool) {
	if o == nil || IsNil(o.Streamable) {
		return nil, false
	}
	return o.Streamable, true
}

// HasStreamable returns a boolean if a field has been set.
func (o *TesInput) HasStreamable() bool {
	if o != nil && !IsNil(o.Streamable) {
		return true
	}

	return false
}

// SetStreamable gets a reference to the given bool and assigns it to the Streamable field.
func (o *TesInput) SetStreamable(v bool) {
	o.Streamable = &v
}

func (o TesInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TesInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Streamable) {
		toSerialize["streamable"] = o.Streamable
	}
	return toSerialize, nil
}

func (o *TesInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varTesInput := _TesInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTesInput)

	if err != nil {
		return err
	}

	*o = TesInput(varTesInput)

	return err
}

type NullableTesInput struct {
	value *TesInput
	isSet bool
}

func (v NullableTesInput) Get() *TesInput {
	return v.value
}

func (v *NullableTesInput) Set(val *TesInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTesInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTesInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTesInput(val *TesInput) *NullableTesInput {
	return &NullableTesInput{value: val, isSet: true}
}

func (v NullableTesInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTesInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


