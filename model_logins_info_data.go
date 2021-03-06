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

// LoginsInfoData struct for LoginsInfoData
type LoginsInfoData struct {
	LastSuccessLoginTime *int64  `json:"LastSuccessLoginTime,omitempty"`
	LastSuccessLoginIP   *string `json:"LastSuccessLoginIP,omitempty"`
	FailedLogins         *int32  `json:"FailedLogins,omitempty"`
	LastFailLoginTime    *int64  `json:"LastFailLoginTime,omitempty"`
	LastFailLoginIP      *string `json:"LastFailLoginIP,omitempty"`
}

// NewLoginsInfoData instantiates a new LoginsInfoData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginsInfoData() *LoginsInfoData {
	this := LoginsInfoData{}
	return &this
}

// NewLoginsInfoDataWithDefaults instantiates a new LoginsInfoData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginsInfoDataWithDefaults() *LoginsInfoData {
	this := LoginsInfoData{}
	return &this
}

// GetLastSuccessLoginTime returns the LastSuccessLoginTime field value if set, zero value otherwise.
func (o *LoginsInfoData) GetLastSuccessLoginTime() int64 {
	if o == nil || o.LastSuccessLoginTime == nil {
		var ret int64
		return ret
	}
	return *o.LastSuccessLoginTime
}

// GetLastSuccessLoginTimeOk returns a tuple with the LastSuccessLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginsInfoData) GetLastSuccessLoginTimeOk() (*int64, bool) {
	if o == nil || o.LastSuccessLoginTime == nil {
		return nil, false
	}
	return o.LastSuccessLoginTime, true
}

// HasLastSuccessLoginTime returns a boolean if a field has been set.
func (o *LoginsInfoData) HasLastSuccessLoginTime() bool {
	if o != nil && o.LastSuccessLoginTime != nil {
		return true
	}

	return false
}

// SetLastSuccessLoginTime gets a reference to the given int64 and assigns it to the LastSuccessLoginTime field.
func (o *LoginsInfoData) SetLastSuccessLoginTime(v int64) {
	o.LastSuccessLoginTime = &v
}

// GetLastSuccessLoginIP returns the LastSuccessLoginIP field value if set, zero value otherwise.
func (o *LoginsInfoData) GetLastSuccessLoginIP() string {
	if o == nil || o.LastSuccessLoginIP == nil {
		var ret string
		return ret
	}
	return *o.LastSuccessLoginIP
}

// GetLastSuccessLoginIPOk returns a tuple with the LastSuccessLoginIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginsInfoData) GetLastSuccessLoginIPOk() (*string, bool) {
	if o == nil || o.LastSuccessLoginIP == nil {
		return nil, false
	}
	return o.LastSuccessLoginIP, true
}

// HasLastSuccessLoginIP returns a boolean if a field has been set.
func (o *LoginsInfoData) HasLastSuccessLoginIP() bool {
	if o != nil && o.LastSuccessLoginIP != nil {
		return true
	}

	return false
}

// SetLastSuccessLoginIP gets a reference to the given string and assigns it to the LastSuccessLoginIP field.
func (o *LoginsInfoData) SetLastSuccessLoginIP(v string) {
	o.LastSuccessLoginIP = &v
}

// GetFailedLogins returns the FailedLogins field value if set, zero value otherwise.
func (o *LoginsInfoData) GetFailedLogins() int32 {
	if o == nil || o.FailedLogins == nil {
		var ret int32
		return ret
	}
	return *o.FailedLogins
}

// GetFailedLoginsOk returns a tuple with the FailedLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginsInfoData) GetFailedLoginsOk() (*int32, bool) {
	if o == nil || o.FailedLogins == nil {
		return nil, false
	}
	return o.FailedLogins, true
}

// HasFailedLogins returns a boolean if a field has been set.
func (o *LoginsInfoData) HasFailedLogins() bool {
	if o != nil && o.FailedLogins != nil {
		return true
	}

	return false
}

// SetFailedLogins gets a reference to the given int32 and assigns it to the FailedLogins field.
func (o *LoginsInfoData) SetFailedLogins(v int32) {
	o.FailedLogins = &v
}

// GetLastFailLoginTime returns the LastFailLoginTime field value if set, zero value otherwise.
func (o *LoginsInfoData) GetLastFailLoginTime() int64 {
	if o == nil || o.LastFailLoginTime == nil {
		var ret int64
		return ret
	}
	return *o.LastFailLoginTime
}

// GetLastFailLoginTimeOk returns a tuple with the LastFailLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginsInfoData) GetLastFailLoginTimeOk() (*int64, bool) {
	if o == nil || o.LastFailLoginTime == nil {
		return nil, false
	}
	return o.LastFailLoginTime, true
}

// HasLastFailLoginTime returns a boolean if a field has been set.
func (o *LoginsInfoData) HasLastFailLoginTime() bool {
	if o != nil && o.LastFailLoginTime != nil {
		return true
	}

	return false
}

// SetLastFailLoginTime gets a reference to the given int64 and assigns it to the LastFailLoginTime field.
func (o *LoginsInfoData) SetLastFailLoginTime(v int64) {
	o.LastFailLoginTime = &v
}

// GetLastFailLoginIP returns the LastFailLoginIP field value if set, zero value otherwise.
func (o *LoginsInfoData) GetLastFailLoginIP() string {
	if o == nil || o.LastFailLoginIP == nil {
		var ret string
		return ret
	}
	return *o.LastFailLoginIP
}

// GetLastFailLoginIPOk returns a tuple with the LastFailLoginIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginsInfoData) GetLastFailLoginIPOk() (*string, bool) {
	if o == nil || o.LastFailLoginIP == nil {
		return nil, false
	}
	return o.LastFailLoginIP, true
}

// HasLastFailLoginIP returns a boolean if a field has been set.
func (o *LoginsInfoData) HasLastFailLoginIP() bool {
	if o != nil && o.LastFailLoginIP != nil {
		return true
	}

	return false
}

// SetLastFailLoginIP gets a reference to the given string and assigns it to the LastFailLoginIP field.
func (o *LoginsInfoData) SetLastFailLoginIP(v string) {
	o.LastFailLoginIP = &v
}

func (o LoginsInfoData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSuccessLoginTime != nil {
		toSerialize["LastSuccessLoginTime"] = o.LastSuccessLoginTime
	}
	if o.LastSuccessLoginIP != nil {
		toSerialize["LastSuccessLoginIP"] = o.LastSuccessLoginIP
	}
	if o.FailedLogins != nil {
		toSerialize["FailedLogins"] = o.FailedLogins
	}
	if o.LastFailLoginTime != nil {
		toSerialize["LastFailLoginTime"] = o.LastFailLoginTime
	}
	if o.LastFailLoginIP != nil {
		toSerialize["LastFailLoginIP"] = o.LastFailLoginIP
	}
	return json.Marshal(toSerialize)
}

type NullableLoginsInfoData struct {
	value *LoginsInfoData
	isSet bool
}

func (v NullableLoginsInfoData) Get() *LoginsInfoData {
	return v.value
}

func (v *NullableLoginsInfoData) Set(val *LoginsInfoData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginsInfoData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginsInfoData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginsInfoData(val *LoginsInfoData) *NullableLoginsInfoData {
	return &NullableLoginsInfoData{value: val, isSet: true}
}

func (v NullableLoginsInfoData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginsInfoData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
