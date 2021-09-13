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

// LDAPDirectoryBase struct for LDAPDirectoryBase
type LDAPDirectoryBase struct {
	// The address of the domain.
	DomainName string `json:"DomainName"`
	// The base context of the External Directory.
	DomainBaseContext string `json:"DomainBaseContext"`
}

// NewLDAPDirectoryBase instantiates a new LDAPDirectoryBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPDirectoryBase(domainName string, domainBaseContext string) *LDAPDirectoryBase {
	this := LDAPDirectoryBase{}
	this.DomainName = domainName
	this.DomainBaseContext = domainBaseContext
	return &this
}

// NewLDAPDirectoryBaseWithDefaults instantiates a new LDAPDirectoryBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPDirectoryBaseWithDefaults() *LDAPDirectoryBase {
	this := LDAPDirectoryBase{}
	return &this
}

// GetDomainName returns the DomainName field value
func (o *LDAPDirectoryBase) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *LDAPDirectoryBase) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *LDAPDirectoryBase) SetDomainName(v string) {
	o.DomainName = v
}

// GetDomainBaseContext returns the DomainBaseContext field value
func (o *LDAPDirectoryBase) GetDomainBaseContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainBaseContext
}

// GetDomainBaseContextOk returns a tuple with the DomainBaseContext field value
// and a boolean to check if the value has been set.
func (o *LDAPDirectoryBase) GetDomainBaseContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainBaseContext, true
}

// SetDomainBaseContext sets field value
func (o *LDAPDirectoryBase) SetDomainBaseContext(v string) {
	o.DomainBaseContext = v
}

func (o LDAPDirectoryBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DomainName"] = o.DomainName
	}
	if true {
		toSerialize["DomainBaseContext"] = o.DomainBaseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPDirectoryBase struct {
	value *LDAPDirectoryBase
	isSet bool
}

func (v NullableLDAPDirectoryBase) Get() *LDAPDirectoryBase {
	return v.value
}

func (v *NullableLDAPDirectoryBase) Set(val *LDAPDirectoryBase) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPDirectoryBase) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPDirectoryBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPDirectoryBase(val *LDAPDirectoryBase) *NullableLDAPDirectoryBase {
	return &NullableLDAPDirectoryBase{value: val, isSet: true}
}

func (v NullableLDAPDirectoryBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPDirectoryBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
