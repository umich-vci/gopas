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

// A4c struct for A4c
type A4c struct {
	PSMConnectorID string `json:"PSMConnectorID"`
	Enabled        bool   `json:"Enabled"`
}

// NewA4c instantiates a new A4c object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewA4c(pSMConnectorID string, enabled bool) *A4c {
	this := A4c{}
	this.PSMConnectorID = pSMConnectorID
	this.Enabled = enabled
	return &this
}

// NewA4cWithDefaults instantiates a new A4c object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewA4cWithDefaults() *A4c {
	this := A4c{}
	return &this
}

// GetPSMConnectorID returns the PSMConnectorID field value
func (o *A4c) GetPSMConnectorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PSMConnectorID
}

// GetPSMConnectorIDOk returns a tuple with the PSMConnectorID field value
// and a boolean to check if the value has been set.
func (o *A4c) GetPSMConnectorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PSMConnectorID, true
}

// SetPSMConnectorID sets field value
func (o *A4c) SetPSMConnectorID(v string) {
	o.PSMConnectorID = v
}

// GetEnabled returns the Enabled field value
func (o *A4c) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *A4c) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *A4c) SetEnabled(v bool) {
	o.Enabled = v
}

func (o A4c) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PSMConnectorID"] = o.PSMConnectorID
	}
	if true {
		toSerialize["Enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableA4c struct {
	value *A4c
	isSet bool
}

func (v NullableA4c) Get() *A4c {
	return v.value
}

func (v *NullableA4c) Set(val *A4c) {
	v.value = val
	v.isSet = true
}

func (v NullableA4c) IsSet() bool {
	return v.isSet
}

func (v *NullableA4c) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableA4c(val *A4c) *NullableA4c {
	return &NullableA4c{value: val, isSet: true}
}

func (v NullableA4c) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableA4c) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}