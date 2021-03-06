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

// General struct for General
type General struct {
	// The platform unique ID
	Id *string `json:"id,omitempty"`
	// The platform name
	Name *string `json:"name,omitempty"`
	// The type of system to which the platform is applied
	SystemType *string `json:"systemType,omitempty"`
	// Indicates whether a platform is active or inactive. valid values: true\\false
	Active *bool `json:"active,omitempty"`
	// The platform description
	Description *string `json:"description,omitempty"`
	// The ID of the default platform that this platform is based on (duplicated from)
	PlatformBaseID *string `json:"platformBaseID,omitempty"`
	// Indicates if the platform is a group platform. valid values: Group\\Regular
	PlatformType *string `json:"platformType,omitempty"`
}

// NewGeneral instantiates a new General object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneral() *General {
	this := General{}
	return &this
}

// NewGeneralWithDefaults instantiates a new General object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralWithDefaults() *General {
	this := General{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *General) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *General) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *General) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *General) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *General) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *General) SetName(v string) {
	o.Name = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *General) GetSystemType() string {
	if o == nil || o.SystemType == nil {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetSystemTypeOk() (*string, bool) {
	if o == nil || o.SystemType == nil {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *General) HasSystemType() bool {
	if o != nil && o.SystemType != nil {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *General) SetSystemType(v string) {
	o.SystemType = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *General) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *General) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *General) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *General) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *General) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *General) SetDescription(v string) {
	o.Description = &v
}

// GetPlatformBaseID returns the PlatformBaseID field value if set, zero value otherwise.
func (o *General) GetPlatformBaseID() string {
	if o == nil || o.PlatformBaseID == nil {
		var ret string
		return ret
	}
	return *o.PlatformBaseID
}

// GetPlatformBaseIDOk returns a tuple with the PlatformBaseID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetPlatformBaseIDOk() (*string, bool) {
	if o == nil || o.PlatformBaseID == nil {
		return nil, false
	}
	return o.PlatformBaseID, true
}

// HasPlatformBaseID returns a boolean if a field has been set.
func (o *General) HasPlatformBaseID() bool {
	if o != nil && o.PlatformBaseID != nil {
		return true
	}

	return false
}

// SetPlatformBaseID gets a reference to the given string and assigns it to the PlatformBaseID field.
func (o *General) SetPlatformBaseID(v string) {
	o.PlatformBaseID = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *General) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *General) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *General) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *General) SetPlatformType(v string) {
	o.PlatformType = &v
}

func (o General) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SystemType != nil {
		toSerialize["systemType"] = o.SystemType
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PlatformBaseID != nil {
		toSerialize["platformBaseID"] = o.PlatformBaseID
	}
	if o.PlatformType != nil {
		toSerialize["platformType"] = o.PlatformType
	}
	return json.Marshal(toSerialize)
}

type NullableGeneral struct {
	value *General
	isSet bool
}

func (v NullableGeneral) Get() *General {
	return v.value
}

func (v *NullableGeneral) Set(val *General) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneral(val *General) *NullableGeneral {
	return &NullableGeneral{value: val, isSet: true}
}

func (v NullableGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
