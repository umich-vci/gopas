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

// ComponentMonitoringSummaryVaultData struct for ComponentMonitoringSummaryVaultData
type ComponentMonitoringSummaryVaultData struct {
	IP         *string `json:"IP,omitempty"`
	Role       *string `json:"Role,omitempty"`
	IsLoggedOn *bool   `json:"IsLoggedOn,omitempty"`
}

// NewComponentMonitoringSummaryVaultData instantiates a new ComponentMonitoringSummaryVaultData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentMonitoringSummaryVaultData() *ComponentMonitoringSummaryVaultData {
	this := ComponentMonitoringSummaryVaultData{}
	return &this
}

// NewComponentMonitoringSummaryVaultDataWithDefaults instantiates a new ComponentMonitoringSummaryVaultData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentMonitoringSummaryVaultDataWithDefaults() *ComponentMonitoringSummaryVaultData {
	this := ComponentMonitoringSummaryVaultData{}
	return &this
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *ComponentMonitoringSummaryVaultData) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringSummaryVaultData) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *ComponentMonitoringSummaryVaultData) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *ComponentMonitoringSummaryVaultData) SetIP(v string) {
	o.IP = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ComponentMonitoringSummaryVaultData) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringSummaryVaultData) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ComponentMonitoringSummaryVaultData) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ComponentMonitoringSummaryVaultData) SetRole(v string) {
	o.Role = &v
}

// GetIsLoggedOn returns the IsLoggedOn field value if set, zero value otherwise.
func (o *ComponentMonitoringSummaryVaultData) GetIsLoggedOn() bool {
	if o == nil || o.IsLoggedOn == nil {
		var ret bool
		return ret
	}
	return *o.IsLoggedOn
}

// GetIsLoggedOnOk returns a tuple with the IsLoggedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentMonitoringSummaryVaultData) GetIsLoggedOnOk() (*bool, bool) {
	if o == nil || o.IsLoggedOn == nil {
		return nil, false
	}
	return o.IsLoggedOn, true
}

// HasIsLoggedOn returns a boolean if a field has been set.
func (o *ComponentMonitoringSummaryVaultData) HasIsLoggedOn() bool {
	if o != nil && o.IsLoggedOn != nil {
		return true
	}

	return false
}

// SetIsLoggedOn gets a reference to the given bool and assigns it to the IsLoggedOn field.
func (o *ComponentMonitoringSummaryVaultData) SetIsLoggedOn(v bool) {
	o.IsLoggedOn = &v
}

func (o ComponentMonitoringSummaryVaultData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IP != nil {
		toSerialize["IP"] = o.IP
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.IsLoggedOn != nil {
		toSerialize["IsLoggedOn"] = o.IsLoggedOn
	}
	return json.Marshal(toSerialize)
}

type NullableComponentMonitoringSummaryVaultData struct {
	value *ComponentMonitoringSummaryVaultData
	isSet bool
}

func (v NullableComponentMonitoringSummaryVaultData) Get() *ComponentMonitoringSummaryVaultData {
	return v.value
}

func (v *NullableComponentMonitoringSummaryVaultData) Set(val *ComponentMonitoringSummaryVaultData) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentMonitoringSummaryVaultData) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentMonitoringSummaryVaultData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentMonitoringSummaryVaultData(val *ComponentMonitoringSummaryVaultData) *NullableComponentMonitoringSummaryVaultData {
	return &NullableComponentMonitoringSummaryVaultData{value: val, isSet: true}
}

func (v NullableComponentMonitoringSummaryVaultData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentMonitoringSummaryVaultData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
