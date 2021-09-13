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

// BulkAccountsList struct for BulkAccountsList
type BulkAccountsList struct {
	Total *int32              `json:"Total,omitempty"`
	Items *[]BulkAccountModel `json:"Items,omitempty"`
}

// NewBulkAccountsList instantiates a new BulkAccountsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAccountsList() *BulkAccountsList {
	this := BulkAccountsList{}
	return &this
}

// NewBulkAccountsListWithDefaults instantiates a new BulkAccountsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAccountsListWithDefaults() *BulkAccountsList {
	this := BulkAccountsList{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BulkAccountsList) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAccountsList) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BulkAccountsList) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *BulkAccountsList) SetTotal(v int32) {
	o.Total = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BulkAccountsList) GetItems() []BulkAccountModel {
	if o == nil || o.Items == nil {
		var ret []BulkAccountModel
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAccountsList) GetItemsOk() (*[]BulkAccountModel, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BulkAccountsList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BulkAccountModel and assigns it to the Items field.
func (o *BulkAccountsList) SetItems(v []BulkAccountModel) {
	o.Items = &v
}

func (o BulkAccountsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["Total"] = o.Total
	}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableBulkAccountsList struct {
	value *BulkAccountsList
	isSet bool
}

func (v NullableBulkAccountsList) Get() *BulkAccountsList {
	return v.value
}

func (v *NullableBulkAccountsList) Set(val *BulkAccountsList) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAccountsList) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAccountsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAccountsList(val *BulkAccountsList) *NullableBulkAccountsList {
	return &NullableBulkAccountsList{value: val, isSet: true}
}

func (v NullableBulkAccountsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAccountsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}