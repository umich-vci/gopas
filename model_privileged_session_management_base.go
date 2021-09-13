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

// PrivilegedSessionManagementBase struct for PrivilegedSessionManagementBase
type PrivilegedSessionManagementBase struct {
	// PSM server ID linked to the platform
	PSMServerId string `json:"PSMServerId"`
	// Name of the PSM server linked to the platform
	PSMServerName *string `json:"PSMServerName,omitempty"`
}

// NewPrivilegedSessionManagementBase instantiates a new PrivilegedSessionManagementBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedSessionManagementBase(pSMServerId string) *PrivilegedSessionManagementBase {
	this := PrivilegedSessionManagementBase{}
	this.PSMServerId = pSMServerId
	return &this
}

// NewPrivilegedSessionManagementBaseWithDefaults instantiates a new PrivilegedSessionManagementBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedSessionManagementBaseWithDefaults() *PrivilegedSessionManagementBase {
	this := PrivilegedSessionManagementBase{}
	return &this
}

// GetPSMServerId returns the PSMServerId field value
func (o *PrivilegedSessionManagementBase) GetPSMServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PSMServerId
}

// GetPSMServerIdOk returns a tuple with the PSMServerId field value
// and a boolean to check if the value has been set.
func (o *PrivilegedSessionManagementBase) GetPSMServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PSMServerId, true
}

// SetPSMServerId sets field value
func (o *PrivilegedSessionManagementBase) SetPSMServerId(v string) {
	o.PSMServerId = v
}

// GetPSMServerName returns the PSMServerName field value if set, zero value otherwise.
func (o *PrivilegedSessionManagementBase) GetPSMServerName() string {
	if o == nil || o.PSMServerName == nil {
		var ret string
		return ret
	}
	return *o.PSMServerName
}

// GetPSMServerNameOk returns a tuple with the PSMServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedSessionManagementBase) GetPSMServerNameOk() (*string, bool) {
	if o == nil || o.PSMServerName == nil {
		return nil, false
	}
	return o.PSMServerName, true
}

// HasPSMServerName returns a boolean if a field has been set.
func (o *PrivilegedSessionManagementBase) HasPSMServerName() bool {
	if o != nil && o.PSMServerName != nil {
		return true
	}

	return false
}

// SetPSMServerName gets a reference to the given string and assigns it to the PSMServerName field.
func (o *PrivilegedSessionManagementBase) SetPSMServerName(v string) {
	o.PSMServerName = &v
}

func (o PrivilegedSessionManagementBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PSMServerId"] = o.PSMServerId
	}
	if o.PSMServerName != nil {
		toSerialize["PSMServerName"] = o.PSMServerName
	}
	return json.Marshal(toSerialize)
}

type NullablePrivilegedSessionManagementBase struct {
	value *PrivilegedSessionManagementBase
	isSet bool
}

func (v NullablePrivilegedSessionManagementBase) Get() *PrivilegedSessionManagementBase {
	return v.value
}

func (v *NullablePrivilegedSessionManagementBase) Set(val *PrivilegedSessionManagementBase) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedSessionManagementBase) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedSessionManagementBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedSessionManagementBase(val *PrivilegedSessionManagementBase) *NullablePrivilegedSessionManagementBase {
	return &NullablePrivilegedSessionManagementBase{value: val, isSet: true}
}

func (v NullablePrivilegedSessionManagementBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedSessionManagementBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
