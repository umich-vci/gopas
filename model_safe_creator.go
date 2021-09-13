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

// SafeCreator struct for SafeCreator
type SafeCreator struct {
	// The ID of the user that created this Safe.
	Id *string `json:"id,omitempty"`
	// The name of the user that created this safe.
	Name *string `json:"name,omitempty"`
}

// NewSafeCreator instantiates a new SafeCreator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafeCreator() *SafeCreator {
	this := SafeCreator{}
	return &this
}

// NewSafeCreatorWithDefaults instantiates a new SafeCreator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafeCreatorWithDefaults() *SafeCreator {
	this := SafeCreator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SafeCreator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeCreator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SafeCreator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SafeCreator) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SafeCreator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeCreator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SafeCreator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SafeCreator) SetName(v string) {
	o.Name = &v
}

func (o SafeCreator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSafeCreator struct {
	value *SafeCreator
	isSet bool
}

func (v NullableSafeCreator) Get() *SafeCreator {
	return v.value
}

func (v *NullableSafeCreator) Set(val *SafeCreator) {
	v.value = val
	v.isSet = true
}

func (v NullableSafeCreator) IsSet() bool {
	return v.isSet
}

func (v *NullableSafeCreator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafeCreator(val *SafeCreator) *NullableSafeCreator {
	return &NullableSafeCreator{value: val, isSet: true}
}

func (v NullableSafeCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafeCreator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
