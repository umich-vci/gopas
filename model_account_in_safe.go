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

// AccountInSafe struct for AccountInSafe
type AccountInSafe struct {
	// The ID of the accounts that reside in this Safe.
	Id *string `json:"id,omitempty"`
	// The name of the accounts that reside in this Safe.
	Name *string `json:"name,omitempty"`
}

// NewAccountInSafe instantiates a new AccountInSafe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInSafe() *AccountInSafe {
	this := AccountInSafe{}
	return &this
}

// NewAccountInSafeWithDefaults instantiates a new AccountInSafe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInSafeWithDefaults() *AccountInSafe {
	this := AccountInSafe{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountInSafe) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInSafe) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountInSafe) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountInSafe) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountInSafe) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInSafe) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountInSafe) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountInSafe) SetName(v string) {
	o.Name = &v
}

func (o AccountInSafe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAccountInSafe struct {
	value *AccountInSafe
	isSet bool
}

func (v NullableAccountInSafe) Get() *AccountInSafe {
	return v.value
}

func (v *NullableAccountInSafe) Set(val *AccountInSafe) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInSafe) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInSafe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInSafe(val *AccountInSafe) *NullableAccountInSafe {
	return &NullableAccountInSafe{value: val, isSet: true}
}

func (v NullableAccountInSafe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInSafe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
