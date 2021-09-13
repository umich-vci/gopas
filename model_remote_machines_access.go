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

// RemoteMachinesAccess struct for RemoteMachinesAccess
type RemoteMachinesAccess struct {
	// List of remote machines, separated by semicolons.
	RemoteMachines *string `json:"remoteMachines,omitempty"`
	// Whether or not to restrict access only to specified remote machines.
	AccessRestrictedToRemoteMachines *bool `json:"accessRestrictedToRemoteMachines,omitempty"`
}

// NewRemoteMachinesAccess instantiates a new RemoteMachinesAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteMachinesAccess() *RemoteMachinesAccess {
	this := RemoteMachinesAccess{}
	return &this
}

// NewRemoteMachinesAccessWithDefaults instantiates a new RemoteMachinesAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteMachinesAccessWithDefaults() *RemoteMachinesAccess {
	this := RemoteMachinesAccess{}
	return &this
}

// GetRemoteMachines returns the RemoteMachines field value if set, zero value otherwise.
func (o *RemoteMachinesAccess) GetRemoteMachines() string {
	if o == nil || o.RemoteMachines == nil {
		var ret string
		return ret
	}
	return *o.RemoteMachines
}

// GetRemoteMachinesOk returns a tuple with the RemoteMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteMachinesAccess) GetRemoteMachinesOk() (*string, bool) {
	if o == nil || o.RemoteMachines == nil {
		return nil, false
	}
	return o.RemoteMachines, true
}

// HasRemoteMachines returns a boolean if a field has been set.
func (o *RemoteMachinesAccess) HasRemoteMachines() bool {
	if o != nil && o.RemoteMachines != nil {
		return true
	}

	return false
}

// SetRemoteMachines gets a reference to the given string and assigns it to the RemoteMachines field.
func (o *RemoteMachinesAccess) SetRemoteMachines(v string) {
	o.RemoteMachines = &v
}

// GetAccessRestrictedToRemoteMachines returns the AccessRestrictedToRemoteMachines field value if set, zero value otherwise.
func (o *RemoteMachinesAccess) GetAccessRestrictedToRemoteMachines() bool {
	if o == nil || o.AccessRestrictedToRemoteMachines == nil {
		var ret bool
		return ret
	}
	return *o.AccessRestrictedToRemoteMachines
}

// GetAccessRestrictedToRemoteMachinesOk returns a tuple with the AccessRestrictedToRemoteMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteMachinesAccess) GetAccessRestrictedToRemoteMachinesOk() (*bool, bool) {
	if o == nil || o.AccessRestrictedToRemoteMachines == nil {
		return nil, false
	}
	return o.AccessRestrictedToRemoteMachines, true
}

// HasAccessRestrictedToRemoteMachines returns a boolean if a field has been set.
func (o *RemoteMachinesAccess) HasAccessRestrictedToRemoteMachines() bool {
	if o != nil && o.AccessRestrictedToRemoteMachines != nil {
		return true
	}

	return false
}

// SetAccessRestrictedToRemoteMachines gets a reference to the given bool and assigns it to the AccessRestrictedToRemoteMachines field.
func (o *RemoteMachinesAccess) SetAccessRestrictedToRemoteMachines(v bool) {
	o.AccessRestrictedToRemoteMachines = &v
}

func (o RemoteMachinesAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteMachines != nil {
		toSerialize["remoteMachines"] = o.RemoteMachines
	}
	if o.AccessRestrictedToRemoteMachines != nil {
		toSerialize["accessRestrictedToRemoteMachines"] = o.AccessRestrictedToRemoteMachines
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteMachinesAccess struct {
	value *RemoteMachinesAccess
	isSet bool
}

func (v NullableRemoteMachinesAccess) Get() *RemoteMachinesAccess {
	return v.value
}

func (v *NullableRemoteMachinesAccess) Set(val *RemoteMachinesAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteMachinesAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteMachinesAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteMachinesAccess(val *RemoteMachinesAccess) *NullableRemoteMachinesAccess {
	return &NullableRemoteMachinesAccess{value: val, isSet: true}
}

func (v NullableRemoteMachinesAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteMachinesAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
