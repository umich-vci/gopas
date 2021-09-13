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

// Incident struct for Incident
type Incident struct {
	Id        *string `json:"Id,omitempty"`
	Url       *string `json:"Url,omitempty"`
	Score     *int32  `json:"Score,omitempty"`
	Name      *string `json:"Name,omitempty"`
	StartDate *int64  `json:"StartDate,omitempty"`
}

// NewIncident instantiates a new Incident object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncident() *Incident {
	this := Incident{}
	return &this
}

// NewIncidentWithDefaults instantiates a new Incident object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentWithDefaults() *Incident {
	this := Incident{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Incident) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Incident) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Incident) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Incident) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Incident) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Incident) SetUrl(v string) {
	o.Url = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Incident) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Incident) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *Incident) SetScore(v int32) {
	o.Score = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Incident) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Incident) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Incident) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Incident) GetStartDate() int64 {
	if o == nil || o.StartDate == nil {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetStartDateOk() (*int64, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Incident) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *Incident) SetStartDate(v int64) {
	o.StartDate = &v
}

func (o Incident) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.Score != nil {
		toSerialize["Score"] = o.Score
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["StartDate"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableIncident struct {
	value *Incident
	isSet bool
}

func (v NullableIncident) Get() *Incident {
	return v.value
}

func (v *NullableIncident) Set(val *Incident) {
	v.value = val
	v.isSet = true
}

func (v NullableIncident) IsSet() bool {
	return v.isSet
}

func (v *NullableIncident) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncident(val *Incident) *NullableIncident {
	return &NullableIncident{value: val, isSet: true}
}

func (v NullableIncident) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncident) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}