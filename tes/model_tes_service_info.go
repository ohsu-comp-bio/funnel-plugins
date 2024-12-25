/*
Task Execution Service

## Executive Summary The Task Execution Service (TES) API is a standardized schema and API for describing and executing batch execution tasks. A task defines a set of input files, a set of containers and commands to run, a set of output files and some other logging and metadata.  TES servers accept task documents and execute them asynchronously on available compute resources. A TES server could be built on top of a traditional HPC queuing system, such as Grid Engine, Slurm or cloud style compute systems such as AWS Batch or Kubernetes. ## Introduction This document describes the TES API and provides details on the specific endpoints, request formats, and responses. It is intended to provide key information for developers of TES-compatible services as well as clients that will call these TES services. Use cases include:    - Deploying existing workflow engines on new infrastructure. Workflow engines   such as CWL-Tes and Cromwell have extentions for using TES. This will allow   a system engineer to deploy them onto a new infrastructure using a job scheduling   system not previously supported by the engine.    - Developing a custom workflow management system. This API provides a common   interface to asynchronous batch processing capabilities. A developer can write   new tools against this interface and expect them to work using a variety of   backend solutions that all support the same specification.   ## Standards The TES API specification is written in OpenAPI and embodies a RESTful service philosophy. It uses JSON in requests and responses and standard HTTP/HTTPS for information transport. HTTPS should be used rather than plain HTTP except for testing or internal-only purposes. ### Authentication and Authorization Is is envisaged that most TES API instances will require users to authenticate to use the endpoints. However, the decision if authentication is required should be taken by TES API implementers.  If authentication is required, we recommend that TES implementations use an OAuth2  bearer token, although they can choose other mechanisms if appropriate.  Checking that a user is authorized to submit TES requests is a responsibility of TES implementations. ### CORS If TES API implementation is to be used by another website or domain it must implement Cross Origin Resource Sharing (CORS). Please refer to https://w3id.org/ga4gh/product-approval-support/cors for more information about GA4GH’s recommendations and how to implement CORS. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tes

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TesServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TesServiceInfo{}

// TesServiceInfo struct for TesServiceInfo
type TesServiceInfo struct {
	// Unique ID of this service. Reverse domain name notation is recommended, though not required. The identifier should attempt to be globally unique so it can be used in downstream aggregator services e.g. Service Registry.
	Id string `json:"id"`
	// Name of this service. Should be human readable.
	Name string `json:"name"`
	Type TesServiceType `json:"type"`
	// Description of the service. Should be human readable and provide information about the service.
	Description *string `json:"description,omitempty"`
	Organization ServiceOrganization `json:"organization"`
	// URL of the contact for the provider of this service, e.g. a link to a contact form (RFC 3986 format), or an email (RFC 2368 format).
	ContactUrl *string `json:"contactUrl,omitempty"`
	// URL of the documentation of this service (RFC 3986 format). This should help someone learn how to use your service, including any specifics required to access data, e.g. authentication.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
	// Timestamp describing when the service was first deployed and available (RFC 3339 format)
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp describing when the service was last updated (RFC 3339 format)
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Environment the service is running in. Use this to distinguish between production, development and testing/staging deployments. Suggested values are prod, test, dev, staging. However this is advised and not enforced.
	Environment *string `json:"environment,omitempty"`
	// Version of the service being described. Semantic versioning is recommended, but other identifiers, such as dates or commit hashes, are also allowed. The version should be changed whenever the service is updated.
	Version string `json:"version"`
	// Lists some, but not necessarily all, storage locations supported by the service.
	Storage []string `json:"storage,omitempty"`
	// Lists all tesResources.backend_parameters keys supported by the service
	TesResourcesBackendParameters []string `json:"tesResources_backend_parameters,omitempty"`
}

type _TesServiceInfo TesServiceInfo

// NewTesServiceInfo instantiates a new TesServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTesServiceInfo(id string, name string, type_ TesServiceType, organization ServiceOrganization, version string) *TesServiceInfo {
	this := TesServiceInfo{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Organization = organization
	this.Version = version
	return &this
}

// NewTesServiceInfoWithDefaults instantiates a new TesServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTesServiceInfoWithDefaults() *TesServiceInfo {
	this := TesServiceInfo{}
	return &this
}

// GetId returns the Id field value
func (o *TesServiceInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TesServiceInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TesServiceInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TesServiceInfo) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TesServiceInfo) GetType() TesServiceType {
	if o == nil {
		var ret TesServiceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetTypeOk() (*TesServiceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TesServiceInfo) SetType(v TesServiceType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TesServiceInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TesServiceInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TesServiceInfo) SetDescription(v string) {
	o.Description = &v
}

// GetOrganization returns the Organization field value
func (o *TesServiceInfo) GetOrganization() ServiceOrganization {
	if o == nil {
		var ret ServiceOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetOrganizationOk() (*ServiceOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *TesServiceInfo) SetOrganization(v ServiceOrganization) {
	o.Organization = v
}

// GetContactUrl returns the ContactUrl field value if set, zero value otherwise.
func (o *TesServiceInfo) GetContactUrl() string {
	if o == nil || IsNil(o.ContactUrl) {
		var ret string
		return ret
	}
	return *o.ContactUrl
}

// GetContactUrlOk returns a tuple with the ContactUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetContactUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContactUrl) {
		return nil, false
	}
	return o.ContactUrl, true
}

// HasContactUrl returns a boolean if a field has been set.
func (o *TesServiceInfo) HasContactUrl() bool {
	if o != nil && !IsNil(o.ContactUrl) {
		return true
	}

	return false
}

// SetContactUrl gets a reference to the given string and assigns it to the ContactUrl field.
func (o *TesServiceInfo) SetContactUrl(v string) {
	o.ContactUrl = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *TesServiceInfo) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *TesServiceInfo) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *TesServiceInfo) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TesServiceInfo) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TesServiceInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TesServiceInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TesServiceInfo) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TesServiceInfo) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TesServiceInfo) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *TesServiceInfo) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *TesServiceInfo) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *TesServiceInfo) SetEnvironment(v string) {
	o.Environment = &v
}

// GetVersion returns the Version field value
func (o *TesServiceInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TesServiceInfo) SetVersion(v string) {
	o.Version = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *TesServiceInfo) GetStorage() []string {
	if o == nil || IsNil(o.Storage) {
		var ret []string
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetStorageOk() ([]string, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *TesServiceInfo) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []string and assigns it to the Storage field.
func (o *TesServiceInfo) SetStorage(v []string) {
	o.Storage = v
}

// GetTesResourcesBackendParameters returns the TesResourcesBackendParameters field value if set, zero value otherwise.
func (o *TesServiceInfo) GetTesResourcesBackendParameters() []string {
	if o == nil || IsNil(o.TesResourcesBackendParameters) {
		var ret []string
		return ret
	}
	return o.TesResourcesBackendParameters
}

// GetTesResourcesBackendParametersOk returns a tuple with the TesResourcesBackendParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TesServiceInfo) GetTesResourcesBackendParametersOk() ([]string, bool) {
	if o == nil || IsNil(o.TesResourcesBackendParameters) {
		return nil, false
	}
	return o.TesResourcesBackendParameters, true
}

// HasTesResourcesBackendParameters returns a boolean if a field has been set.
func (o *TesServiceInfo) HasTesResourcesBackendParameters() bool {
	if o != nil && !IsNil(o.TesResourcesBackendParameters) {
		return true
	}

	return false
}

// SetTesResourcesBackendParameters gets a reference to the given []string and assigns it to the TesResourcesBackendParameters field.
func (o *TesServiceInfo) SetTesResourcesBackendParameters(v []string) {
	o.TesResourcesBackendParameters = v
}

func (o TesServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TesServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["organization"] = o.Organization
	if !IsNil(o.ContactUrl) {
		toSerialize["contactUrl"] = o.ContactUrl
	}
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentationUrl"] = o.DocumentationUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.TesResourcesBackendParameters) {
		toSerialize["tesResources_backend_parameters"] = o.TesResourcesBackendParameters
	}
	return toSerialize, nil
}

func (o *TesServiceInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"organization",
		"version",
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

	varTesServiceInfo := _TesServiceInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTesServiceInfo)

	if err != nil {
		return err
	}

	*o = TesServiceInfo(varTesServiceInfo)

	return err
}

type NullableTesServiceInfo struct {
	value *TesServiceInfo
	isSet bool
}

func (v NullableTesServiceInfo) Get() *TesServiceInfo {
	return v.value
}

func (v *NullableTesServiceInfo) Set(val *TesServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTesServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTesServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTesServiceInfo(val *TesServiceInfo) *NullableTesServiceInfo {
	return &NullableTesServiceInfo{value: val, isSet: true}
}

func (v NullableTesServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTesServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

