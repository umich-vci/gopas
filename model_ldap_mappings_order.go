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

// LDAPMappingsOrder struct for LDAPMappingsOrder
type LDAPMappingsOrder struct {
	// The new order of the mappings, sorted by importance.   Must include all mappings.  Each ID can only appear once.
	MappingsOrder []int64 `json:"MappingsOrder"`
}

// NewLDAPMappingsOrder instantiates a new LDAPMappingsOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPMappingsOrder(mappingsOrder []int64) *LDAPMappingsOrder {
	this := LDAPMappingsOrder{}
	this.MappingsOrder = mappingsOrder
	return &this
}

// NewLDAPMappingsOrderWithDefaults instantiates a new LDAPMappingsOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPMappingsOrderWithDefaults() *LDAPMappingsOrder {
	this := LDAPMappingsOrder{}
	return &this
}

// GetMappingsOrder returns the MappingsOrder field value
func (o *LDAPMappingsOrder) GetMappingsOrder() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.MappingsOrder
}

// GetMappingsOrderOk returns a tuple with the MappingsOrder field value
// and a boolean to check if the value has been set.
func (o *LDAPMappingsOrder) GetMappingsOrderOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MappingsOrder, true
}

// SetMappingsOrder sets field value
func (o *LDAPMappingsOrder) SetMappingsOrder(v []int64) {
	o.MappingsOrder = v
}

func (o LDAPMappingsOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MappingsOrder"] = o.MappingsOrder
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPMappingsOrder struct {
	value *LDAPMappingsOrder
	isSet bool
}

func (v NullableLDAPMappingsOrder) Get() *LDAPMappingsOrder {
	return v.value
}

func (v *NullableLDAPMappingsOrder) Set(val *LDAPMappingsOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPMappingsOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPMappingsOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPMappingsOrder(val *LDAPMappingsOrder) *NullableLDAPMappingsOrder {
	return &NullableLDAPMappingsOrder{value: val, isSet: true}
}

func (v NullableLDAPMappingsOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPMappingsOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
