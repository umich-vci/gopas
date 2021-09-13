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

// PageQS struct for PageQS
type PageQS struct {
	Limit    *int64 `json:"Limit,omitempty"`
	Offset   *int64 `json:"Offset,omitempty"`
	UseCache *bool  `json:"UseCache,omitempty"`
}

// NewPageQS instantiates a new PageQS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageQS() *PageQS {
	this := PageQS{}
	return &this
}

// NewPageQSWithDefaults instantiates a new PageQS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageQSWithDefaults() *PageQS {
	this := PageQS{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PageQS) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageQS) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PageQS) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *PageQS) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PageQS) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageQS) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PageQS) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *PageQS) SetOffset(v int64) {
	o.Offset = &v
}

// GetUseCache returns the UseCache field value if set, zero value otherwise.
func (o *PageQS) GetUseCache() bool {
	if o == nil || o.UseCache == nil {
		var ret bool
		return ret
	}
	return *o.UseCache
}

// GetUseCacheOk returns a tuple with the UseCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageQS) GetUseCacheOk() (*bool, bool) {
	if o == nil || o.UseCache == nil {
		return nil, false
	}
	return o.UseCache, true
}

// HasUseCache returns a boolean if a field has been set.
func (o *PageQS) HasUseCache() bool {
	if o != nil && o.UseCache != nil {
		return true
	}

	return false
}

// SetUseCache gets a reference to the given bool and assigns it to the UseCache field.
func (o *PageQS) SetUseCache(v bool) {
	o.UseCache = &v
}

func (o PageQS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["Offset"] = o.Offset
	}
	if o.UseCache != nil {
		toSerialize["UseCache"] = o.UseCache
	}
	return json.Marshal(toSerialize)
}

type NullablePageQS struct {
	value *PageQS
	isSet bool
}

func (v NullablePageQS) Get() *PageQS {
	return v.value
}

func (v *NullablePageQS) Set(val *PageQS) {
	v.value = val
	v.isSet = true
}

func (v NullablePageQS) IsSet() bool {
	return v.isSet
}

func (v *NullablePageQS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageQS(val *PageQS) *NullablePageQS {
	return &NullablePageQS{value: val, isSet: true}
}

func (v NullablePageQS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageQS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}