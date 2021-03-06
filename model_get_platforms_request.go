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

// GetPlatformsRequest struct for GetPlatformsRequest
type GetPlatformsRequest struct {
	Active       *bool  `json:"active,omitempty"`
	PlatformType *int32 `json:"platformType,omitempty"`
	// The search will be by Platform ID or Platform Name.
	Search     *string `json:"search,omitempty"`
	SystemType *string `json:"systemType,omitempty"`
}

// NewGetPlatformsRequest instantiates a new GetPlatformsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlatformsRequest() *GetPlatformsRequest {
	this := GetPlatformsRequest{}
	return &this
}

// NewGetPlatformsRequestWithDefaults instantiates a new GetPlatformsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlatformsRequestWithDefaults() *GetPlatformsRequest {
	this := GetPlatformsRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetPlatformsRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlatformsRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetPlatformsRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetPlatformsRequest) SetActive(v bool) {
	o.Active = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *GetPlatformsRequest) GetPlatformType() int32 {
	if o == nil || o.PlatformType == nil {
		var ret int32
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlatformsRequest) GetPlatformTypeOk() (*int32, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *GetPlatformsRequest) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given int32 and assigns it to the PlatformType field.
func (o *GetPlatformsRequest) SetPlatformType(v int32) {
	o.PlatformType = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *GetPlatformsRequest) GetSearch() string {
	if o == nil || o.Search == nil {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlatformsRequest) GetSearchOk() (*string, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *GetPlatformsRequest) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *GetPlatformsRequest) SetSearch(v string) {
	o.Search = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *GetPlatformsRequest) GetSystemType() string {
	if o == nil || o.SystemType == nil {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlatformsRequest) GetSystemTypeOk() (*string, bool) {
	if o == nil || o.SystemType == nil {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *GetPlatformsRequest) HasSystemType() bool {
	if o != nil && o.SystemType != nil {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *GetPlatformsRequest) SetSystemType(v string) {
	o.SystemType = &v
}

func (o GetPlatformsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.PlatformType != nil {
		toSerialize["platformType"] = o.PlatformType
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.SystemType != nil {
		toSerialize["systemType"] = o.SystemType
	}
	return json.Marshal(toSerialize)
}

type NullableGetPlatformsRequest struct {
	value *GetPlatformsRequest
	isSet bool
}

func (v NullableGetPlatformsRequest) Get() *GetPlatformsRequest {
	return v.value
}

func (v *NullableGetPlatformsRequest) Set(val *GetPlatformsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlatformsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlatformsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlatformsRequest(val *GetPlatformsRequest) *NullableGetPlatformsRequest {
	return &NullableGetPlatformsRequest{value: val, isSet: true}
}

func (v NullableGetPlatformsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlatformsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
