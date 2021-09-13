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

// SetNextCredentialsProperties struct for SetNextCredentialsProperties
type SetNextCredentialsProperties struct {
	// The new account credentials that will be allocated to the account in the Vault.
	NewCredentials string `json:"NewCredentials"`
	// Whether or not the password will be changed immediately in the Vault.
	ChangeImmediately *bool `json:"ChangeImmediately,omitempty"`
}

// NewSetNextCredentialsProperties instantiates a new SetNextCredentialsProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetNextCredentialsProperties(newCredentials string) *SetNextCredentialsProperties {
	this := SetNextCredentialsProperties{}
	this.NewCredentials = newCredentials
	return &this
}

// NewSetNextCredentialsPropertiesWithDefaults instantiates a new SetNextCredentialsProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetNextCredentialsPropertiesWithDefaults() *SetNextCredentialsProperties {
	this := SetNextCredentialsProperties{}
	return &this
}

// GetNewCredentials returns the NewCredentials field value
func (o *SetNextCredentialsProperties) GetNewCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewCredentials
}

// GetNewCredentialsOk returns a tuple with the NewCredentials field value
// and a boolean to check if the value has been set.
func (o *SetNextCredentialsProperties) GetNewCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewCredentials, true
}

// SetNewCredentials sets field value
func (o *SetNextCredentialsProperties) SetNewCredentials(v string) {
	o.NewCredentials = v
}

// GetChangeImmediately returns the ChangeImmediately field value if set, zero value otherwise.
func (o *SetNextCredentialsProperties) GetChangeImmediately() bool {
	if o == nil || o.ChangeImmediately == nil {
		var ret bool
		return ret
	}
	return *o.ChangeImmediately
}

// GetChangeImmediatelyOk returns a tuple with the ChangeImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetNextCredentialsProperties) GetChangeImmediatelyOk() (*bool, bool) {
	if o == nil || o.ChangeImmediately == nil {
		return nil, false
	}
	return o.ChangeImmediately, true
}

// HasChangeImmediately returns a boolean if a field has been set.
func (o *SetNextCredentialsProperties) HasChangeImmediately() bool {
	if o != nil && o.ChangeImmediately != nil {
		return true
	}

	return false
}

// SetChangeImmediately gets a reference to the given bool and assigns it to the ChangeImmediately field.
func (o *SetNextCredentialsProperties) SetChangeImmediately(v bool) {
	o.ChangeImmediately = &v
}

func (o SetNextCredentialsProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["NewCredentials"] = o.NewCredentials
	}
	if o.ChangeImmediately != nil {
		toSerialize["ChangeImmediately"] = o.ChangeImmediately
	}
	return json.Marshal(toSerialize)
}

type NullableSetNextCredentialsProperties struct {
	value *SetNextCredentialsProperties
	isSet bool
}

func (v NullableSetNextCredentialsProperties) Get() *SetNextCredentialsProperties {
	return v.value
}

func (v *NullableSetNextCredentialsProperties) Set(val *SetNextCredentialsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableSetNextCredentialsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableSetNextCredentialsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetNextCredentialsProperties(val *SetNextCredentialsProperties) *NullableSetNextCredentialsProperties {
	return &NullableSetNextCredentialsProperties{value: val, isSet: true}
}

func (v NullableSetNextCredentialsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetNextCredentialsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}