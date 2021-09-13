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

// GetDiscoveredAccountsResponse struct for GetDiscoveredAccountsResponse
type GetDiscoveredAccountsResponse struct {
	Value *[]GetDiscoveredAccount `json:"value,omitempty"`
	// Total number of results across all pages
	Count *int32 `json:"count,omitempty"`
	// An opaque URL to the next page of results.   Should be present only if requested page size (limit) is not specified and there are more results than a single page as defined by the server.  The last page shouldn't have 'nextLink' in the response.   If the limit in the request is too high, an error is returned
	NextLink *string `json:"nextLink,omitempty"`
}

// NewGetDiscoveredAccountsResponse instantiates a new GetDiscoveredAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoveredAccountsResponse() *GetDiscoveredAccountsResponse {
	this := GetDiscoveredAccountsResponse{}
	return &this
}

// NewGetDiscoveredAccountsResponseWithDefaults instantiates a new GetDiscoveredAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoveredAccountsResponseWithDefaults() *GetDiscoveredAccountsResponse {
	this := GetDiscoveredAccountsResponse{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsResponse) GetValue() []GetDiscoveredAccount {
	if o == nil || o.Value == nil {
		var ret []GetDiscoveredAccount
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsResponse) GetValueOk() (*[]GetDiscoveredAccount, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []GetDiscoveredAccount and assigns it to the Value field.
func (o *GetDiscoveredAccountsResponse) SetValue(v []GetDiscoveredAccount) {
	o.Value = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetDiscoveredAccountsResponse) SetCount(v int32) {
	o.Count = &v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *GetDiscoveredAccountsResponse) GetNextLink() string {
	if o == nil || o.NextLink == nil {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveredAccountsResponse) GetNextLinkOk() (*string, bool) {
	if o == nil || o.NextLink == nil {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *GetDiscoveredAccountsResponse) HasNextLink() bool {
	if o != nil && o.NextLink != nil {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *GetDiscoveredAccountsResponse) SetNextLink(v string) {
	o.NextLink = &v
}

func (o GetDiscoveredAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.NextLink != nil {
		toSerialize["nextLink"] = o.NextLink
	}
	return json.Marshal(toSerialize)
}

type NullableGetDiscoveredAccountsResponse struct {
	value *GetDiscoveredAccountsResponse
	isSet bool
}

func (v NullableGetDiscoveredAccountsResponse) Get() *GetDiscoveredAccountsResponse {
	return v.value
}

func (v *NullableGetDiscoveredAccountsResponse) Set(val *GetDiscoveredAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveredAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveredAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveredAccountsResponse(val *GetDiscoveredAccountsResponse) *NullableGetDiscoveredAccountsResponse {
	return &NullableGetDiscoveredAccountsResponse{value: val, isSet: true}
}

func (v NullableGetDiscoveredAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveredAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
