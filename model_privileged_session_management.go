/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopas

import (
	"encoding/json"
)

// PrivilegedSessionManagement struct for PrivilegedSessionManagement
type PrivilegedSessionManagement struct {
	// List of PSM Connectors linked to the platform
	PSMConnectors *[]A2j `json:"PSMConnectors,omitempty"`
	// PSM server ID linked to the platform
	PSMServerId string `json:"PSMServerId"`
	// Name of the PSM server linked to the platform
	PSMServerName *string `json:"PSMServerName,omitempty"`
}

// NewPrivilegedSessionManagement instantiates a new PrivilegedSessionManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedSessionManagement(pSMServerId string) *PrivilegedSessionManagement {
	this := PrivilegedSessionManagement{}
	this.PSMServerId = pSMServerId
	return &this
}

// NewPrivilegedSessionManagementWithDefaults instantiates a new PrivilegedSessionManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedSessionManagementWithDefaults() *PrivilegedSessionManagement {
	this := PrivilegedSessionManagement{}
	return &this
}

// GetPSMConnectors returns the PSMConnectors field value if set, zero value otherwise.
func (o *PrivilegedSessionManagement) GetPSMConnectors() []A2j {
	if o == nil || o.PSMConnectors == nil {
		var ret []A2j
		return ret
	}
	return *o.PSMConnectors
}

// GetPSMConnectorsOk returns a tuple with the PSMConnectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedSessionManagement) GetPSMConnectorsOk() (*[]A2j, bool) {
	if o == nil || o.PSMConnectors == nil {
		return nil, false
	}
	return o.PSMConnectors, true
}

// HasPSMConnectors returns a boolean if a field has been set.
func (o *PrivilegedSessionManagement) HasPSMConnectors() bool {
	if o != nil && o.PSMConnectors != nil {
		return true
	}

	return false
}

// SetPSMConnectors gets a reference to the given []A2j and assigns it to the PSMConnectors field.
func (o *PrivilegedSessionManagement) SetPSMConnectors(v []A2j) {
	o.PSMConnectors = &v
}

// GetPSMServerId returns the PSMServerId field value
func (o *PrivilegedSessionManagement) GetPSMServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PSMServerId
}

// GetPSMServerIdOk returns a tuple with the PSMServerId field value
// and a boolean to check if the value has been set.
func (o *PrivilegedSessionManagement) GetPSMServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PSMServerId, true
}

// SetPSMServerId sets field value
func (o *PrivilegedSessionManagement) SetPSMServerId(v string) {
	o.PSMServerId = v
}

// GetPSMServerName returns the PSMServerName field value if set, zero value otherwise.
func (o *PrivilegedSessionManagement) GetPSMServerName() string {
	if o == nil || o.PSMServerName == nil {
		var ret string
		return ret
	}
	return *o.PSMServerName
}

// GetPSMServerNameOk returns a tuple with the PSMServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedSessionManagement) GetPSMServerNameOk() (*string, bool) {
	if o == nil || o.PSMServerName == nil {
		return nil, false
	}
	return o.PSMServerName, true
}

// HasPSMServerName returns a boolean if a field has been set.
func (o *PrivilegedSessionManagement) HasPSMServerName() bool {
	if o != nil && o.PSMServerName != nil {
		return true
	}

	return false
}

// SetPSMServerName gets a reference to the given string and assigns it to the PSMServerName field.
func (o *PrivilegedSessionManagement) SetPSMServerName(v string) {
	o.PSMServerName = &v
}

func (o PrivilegedSessionManagement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PSMConnectors != nil {
		toSerialize["PSMConnectors"] = o.PSMConnectors
	}
	if true {
		toSerialize["PSMServerId"] = o.PSMServerId
	}
	if o.PSMServerName != nil {
		toSerialize["PSMServerName"] = o.PSMServerName
	}
	return json.Marshal(toSerialize)
}

type NullablePrivilegedSessionManagement struct {
	value *PrivilegedSessionManagement
	isSet bool
}

func (v NullablePrivilegedSessionManagement) Get() *PrivilegedSessionManagement {
	return v.value
}

func (v *NullablePrivilegedSessionManagement) Set(val *PrivilegedSessionManagement) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedSessionManagement) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedSessionManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedSessionManagement(val *PrivilegedSessionManagement) *NullablePrivilegedSessionManagement {
	return &NullablePrivilegedSessionManagement{value: val, isSet: true}
}

func (v NullablePrivilegedSessionManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedSessionManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
