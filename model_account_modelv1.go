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

// AccountModelv1 struct for AccountModelv1
type AccountModelv1 struct {
	// Number of accounts found by the search.
	Count    *int32                    `json:"Count,omitempty"`
	Accounts *[]map[string]interface{} `json:"accounts,omitempty"`
}

// NewAccountModelv1 instantiates a new AccountModelv1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountModelv1() *AccountModelv1 {
	this := AccountModelv1{}
	return &this
}

// NewAccountModelv1WithDefaults instantiates a new AccountModelv1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountModelv1WithDefaults() *AccountModelv1 {
	this := AccountModelv1{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AccountModelv1) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModelv1) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AccountModelv1) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AccountModelv1) SetCount(v int32) {
	o.Count = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AccountModelv1) GetAccounts() []map[string]interface{} {
	if o == nil || o.Accounts == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModelv1) GetAccountsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AccountModelv1) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []map[string]interface{} and assigns it to the Accounts field.
func (o *AccountModelv1) SetAccounts(v []map[string]interface{}) {
	o.Accounts = &v
}

func (o AccountModelv1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableAccountModelv1 struct {
	value *AccountModelv1
	isSet bool
}

func (v NullableAccountModelv1) Get() *AccountModelv1 {
	return v.value
}

func (v *NullableAccountModelv1) Set(val *AccountModelv1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountModelv1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountModelv1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountModelv1(val *AccountModelv1) *NullableAccountModelv1 {
	return &NullableAccountModelv1{value: val, isSet: true}
}

func (v NullableAccountModelv1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountModelv1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}