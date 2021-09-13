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

// SecretUpdateConfiguration struct for SecretUpdateConfiguration
type SecretUpdateConfiguration struct {
	// Defines whether or not password changed will be performed via reset mode using the reconciliation account. This is useful in cases where the password policy prevents the user from changing his own password minimal age restriction is applied.
	ChangePasswordInResetMode *bool `json:"ChangePasswordInResetMode,omitempty"`
}

// NewSecretUpdateConfiguration instantiates a new SecretUpdateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretUpdateConfiguration() *SecretUpdateConfiguration {
	this := SecretUpdateConfiguration{}
	return &this
}

// NewSecretUpdateConfigurationWithDefaults instantiates a new SecretUpdateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretUpdateConfigurationWithDefaults() *SecretUpdateConfiguration {
	this := SecretUpdateConfiguration{}
	return &this
}

// GetChangePasswordInResetMode returns the ChangePasswordInResetMode field value if set, zero value otherwise.
func (o *SecretUpdateConfiguration) GetChangePasswordInResetMode() bool {
	if o == nil || o.ChangePasswordInResetMode == nil {
		var ret bool
		return ret
	}
	return *o.ChangePasswordInResetMode
}

// GetChangePasswordInResetModeOk returns a tuple with the ChangePasswordInResetMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretUpdateConfiguration) GetChangePasswordInResetModeOk() (*bool, bool) {
	if o == nil || o.ChangePasswordInResetMode == nil {
		return nil, false
	}
	return o.ChangePasswordInResetMode, true
}

// HasChangePasswordInResetMode returns a boolean if a field has been set.
func (o *SecretUpdateConfiguration) HasChangePasswordInResetMode() bool {
	if o != nil && o.ChangePasswordInResetMode != nil {
		return true
	}

	return false
}

// SetChangePasswordInResetMode gets a reference to the given bool and assigns it to the ChangePasswordInResetMode field.
func (o *SecretUpdateConfiguration) SetChangePasswordInResetMode(v bool) {
	o.ChangePasswordInResetMode = &v
}

func (o SecretUpdateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangePasswordInResetMode != nil {
		toSerialize["ChangePasswordInResetMode"] = o.ChangePasswordInResetMode
	}
	return json.Marshal(toSerialize)
}

type NullableSecretUpdateConfiguration struct {
	value *SecretUpdateConfiguration
	isSet bool
}

func (v NullableSecretUpdateConfiguration) Get() *SecretUpdateConfiguration {
	return v.value
}

func (v *NullableSecretUpdateConfiguration) Set(val *SecretUpdateConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretUpdateConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretUpdateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretUpdateConfiguration(val *SecretUpdateConfiguration) *NullableSecretUpdateConfiguration {
	return &NullableSecretUpdateConfiguration{value: val, isSet: true}
}

func (v NullableSecretUpdateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretUpdateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
