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

// OPConfigurations struct for OPConfigurations
type OPConfigurations struct {
	Providers *[]OPConfigurationData `json:"Providers,omitempty"`
}

// NewOPConfigurations instantiates a new OPConfigurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPConfigurations() *OPConfigurations {
	this := OPConfigurations{}
	return &this
}

// NewOPConfigurationsWithDefaults instantiates a new OPConfigurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPConfigurationsWithDefaults() *OPConfigurations {
	this := OPConfigurations{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *OPConfigurations) GetProviders() []OPConfigurationData {
	if o == nil || o.Providers == nil {
		var ret []OPConfigurationData
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurations) GetProvidersOk() (*[]OPConfigurationData, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *OPConfigurations) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []OPConfigurationData and assigns it to the Providers field.
func (o *OPConfigurations) SetProviders(v []OPConfigurationData) {
	o.Providers = &v
}

func (o OPConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["Providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableOPConfigurations struct {
	value *OPConfigurations
	isSet bool
}

func (v NullableOPConfigurations) Get() *OPConfigurations {
	return v.value
}

func (v *NullableOPConfigurations) Set(val *OPConfigurations) {
	v.value = val
	v.isSet = true
}

func (v NullableOPConfigurations) IsSet() bool {
	return v.isSet
}

func (v *NullableOPConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPConfigurations(val *OPConfigurations) *NullableOPConfigurations {
	return &NullableOPConfigurations{value: val, isSet: true}
}

func (v NullableOPConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
