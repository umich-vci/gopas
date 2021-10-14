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

// SafeListItem struct for SafeListItem
type SafeListItem struct {
	// The unique numerical ID of the Safe.
	SafeNumber *int64 `json:"safeNumber,omitempty"`
	// The location of the Safe in the Vault.
	Location *string          `json:"location,omitempty"`
	Creator  *SafeCreator     `json:"creator,omitempty"`
	Accounts *[]AccountInSafe `json:"accounts,omitempty"`
	// Whether or not to enable Object Level Access Control for the new Safe.
	OlacEnabled *bool `json:"olacEnabled,omitempty"`
	// The number of retained versions of every password that is stored in the Safe.
	NumberOfVersionsRetention *int32 `json:"numberOfVersionsRetention,omitempty"`
	// The number of days that password versions are saved in the Safe.
	NumberOfDaysRetention *int32 `json:"numberOfDaysRetention,omitempty"`
	// Whether or not to automatically purge files after the end of the Object History Retention Period defined in the Safe properties.  Reports Safes and PSM Recording Safes are created automatically with AutoPurgeEnabled set to Yes.These Safes cannot be managed by the CPM.
	AutoPurgeEnabled *bool `json:"autoPurgeEnabled,omitempty"`
	// Unix creation time of the Safe.
	CreationTime *int64 `json:"creationTime,omitempty"`
	// Unix time when the Safe was last updated.
	LastModificationTime *int64 `json:"lastModificationTime,omitempty"`
	// The unique ID of the Safe to be used when calling Safe APIs.
	SafeUrlId *string `json:"safeUrlId,omitempty"`
	// The name of the Safe.
	SafeName *string `json:"safeName,omitempty"`
	// The description of the Safe.
	Description *string `json:"description,omitempty"`
	// The name of the CPM user who will manage the new Safe.
	ManagingCPM *string `json:"managingCPM,omitempty"`
	// Whether or not the membership for the Safe is expired.For expired member, the value will be True.
	IsExpiredMember *bool `json:"isExpiredMember,omitempty"`
}

// NewSafeListItem instantiates a new SafeListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafeListItem() *SafeListItem {
	this := SafeListItem{}
	return &this
}

// NewSafeListItemWithDefaults instantiates a new SafeListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafeListItemWithDefaults() *SafeListItem {
	this := SafeListItem{}
	return &this
}

// GetSafeNumber returns the SafeNumber field value if set, zero value otherwise.
func (o *SafeListItem) GetSafeNumber() int64 {
	if o == nil || o.SafeNumber == nil {
		var ret int64
		return ret
	}
	return *o.SafeNumber
}

// GetSafeNumberOk returns a tuple with the SafeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetSafeNumberOk() (*int64, bool) {
	if o == nil || o.SafeNumber == nil {
		return nil, false
	}
	return o.SafeNumber, true
}

// HasSafeNumber returns a boolean if a field has been set.
func (o *SafeListItem) HasSafeNumber() bool {
	if o != nil && o.SafeNumber != nil {
		return true
	}

	return false
}

// SetSafeNumber gets a reference to the given int64 and assigns it to the SafeNumber field.
func (o *SafeListItem) SetSafeNumber(v int64) {
	o.SafeNumber = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SafeListItem) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SafeListItem) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SafeListItem) SetLocation(v string) {
	o.Location = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SafeListItem) GetCreator() SafeCreator {
	if o == nil || o.Creator == nil {
		var ret SafeCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetCreatorOk() (*SafeCreator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SafeListItem) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given SafeCreator and assigns it to the Creator field.
func (o *SafeListItem) SetCreator(v SafeCreator) {
	o.Creator = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *SafeListItem) GetAccounts() []AccountInSafe {
	if o == nil || o.Accounts == nil {
		var ret []AccountInSafe
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetAccountsOk() (*[]AccountInSafe, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *SafeListItem) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []AccountInSafe and assigns it to the Accounts field.
func (o *SafeListItem) SetAccounts(v []AccountInSafe) {
	o.Accounts = &v
}

// GetOlacEnabled returns the OlacEnabled field value if set, zero value otherwise.
func (o *SafeListItem) GetOlacEnabled() bool {
	if o == nil || o.OlacEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OlacEnabled
}

// GetOlacEnabledOk returns a tuple with the OlacEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetOlacEnabledOk() (*bool, bool) {
	if o == nil || o.OlacEnabled == nil {
		return nil, false
	}
	return o.OlacEnabled, true
}

// HasOlacEnabled returns a boolean if a field has been set.
func (o *SafeListItem) HasOlacEnabled() bool {
	if o != nil && o.OlacEnabled != nil {
		return true
	}

	return false
}

// SetOlacEnabled gets a reference to the given bool and assigns it to the OlacEnabled field.
func (o *SafeListItem) SetOlacEnabled(v bool) {
	o.OlacEnabled = &v
}

// GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field value if set, zero value otherwise.
func (o *SafeListItem) GetNumberOfVersionsRetention() int32 {
	if o == nil || o.NumberOfVersionsRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfVersionsRetention
}

// GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetNumberOfVersionsRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfVersionsRetention == nil {
		return nil, false
	}
	return o.NumberOfVersionsRetention, true
}

// HasNumberOfVersionsRetention returns a boolean if a field has been set.
func (o *SafeListItem) HasNumberOfVersionsRetention() bool {
	if o != nil && o.NumberOfVersionsRetention != nil {
		return true
	}

	return false
}

// SetNumberOfVersionsRetention gets a reference to the given int32 and assigns it to the NumberOfVersionsRetention field.
func (o *SafeListItem) SetNumberOfVersionsRetention(v int32) {
	o.NumberOfVersionsRetention = &v
}

// GetNumberOfDaysRetention returns the NumberOfDaysRetention field value if set, zero value otherwise.
func (o *SafeListItem) GetNumberOfDaysRetention() int32 {
	if o == nil || o.NumberOfDaysRetention == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfDaysRetention
}

// GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetNumberOfDaysRetentionOk() (*int32, bool) {
	if o == nil || o.NumberOfDaysRetention == nil {
		return nil, false
	}
	return o.NumberOfDaysRetention, true
}

// HasNumberOfDaysRetention returns a boolean if a field has been set.
func (o *SafeListItem) HasNumberOfDaysRetention() bool {
	if o != nil && o.NumberOfDaysRetention != nil {
		return true
	}

	return false
}

// SetNumberOfDaysRetention gets a reference to the given int32 and assigns it to the NumberOfDaysRetention field.
func (o *SafeListItem) SetNumberOfDaysRetention(v int32) {
	o.NumberOfDaysRetention = &v
}

// GetAutoPurgeEnabled returns the AutoPurgeEnabled field value if set, zero value otherwise.
func (o *SafeListItem) GetAutoPurgeEnabled() bool {
	if o == nil || o.AutoPurgeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoPurgeEnabled
}

// GetAutoPurgeEnabledOk returns a tuple with the AutoPurgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetAutoPurgeEnabledOk() (*bool, bool) {
	if o == nil || o.AutoPurgeEnabled == nil {
		return nil, false
	}
	return o.AutoPurgeEnabled, true
}

// HasAutoPurgeEnabled returns a boolean if a field has been set.
func (o *SafeListItem) HasAutoPurgeEnabled() bool {
	if o != nil && o.AutoPurgeEnabled != nil {
		return true
	}

	return false
}

// SetAutoPurgeEnabled gets a reference to the given bool and assigns it to the AutoPurgeEnabled field.
func (o *SafeListItem) SetAutoPurgeEnabled(v bool) {
	o.AutoPurgeEnabled = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *SafeListItem) GetCreationTime() int64 {
	if o == nil || o.CreationTime == nil {
		var ret int64
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetCreationTimeOk() (*int64, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *SafeListItem) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int64 and assigns it to the CreationTime field.
func (o *SafeListItem) SetCreationTime(v int64) {
	o.CreationTime = &v
}

// GetLastModificationTime returns the LastModificationTime field value if set, zero value otherwise.
func (o *SafeListItem) GetLastModificationTime() int64 {
	if o == nil || o.LastModificationTime == nil {
		var ret int64
		return ret
	}
	return *o.LastModificationTime
}

// GetLastModificationTimeOk returns a tuple with the LastModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetLastModificationTimeOk() (*int64, bool) {
	if o == nil || o.LastModificationTime == nil {
		return nil, false
	}
	return o.LastModificationTime, true
}

// HasLastModificationTime returns a boolean if a field has been set.
func (o *SafeListItem) HasLastModificationTime() bool {
	if o != nil && o.LastModificationTime != nil {
		return true
	}

	return false
}

// SetLastModificationTime gets a reference to the given int64 and assigns it to the LastModificationTime field.
func (o *SafeListItem) SetLastModificationTime(v int64) {
	o.LastModificationTime = &v
}

// GetSafeUrlId returns the SafeUrlId field value if set, zero value otherwise.
func (o *SafeListItem) GetSafeUrlId() string {
	if o == nil || o.SafeUrlId == nil {
		var ret string
		return ret
	}
	return *o.SafeUrlId
}

// GetSafeUrlIdOk returns a tuple with the SafeUrlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetSafeUrlIdOk() (*string, bool) {
	if o == nil || o.SafeUrlId == nil {
		return nil, false
	}
	return o.SafeUrlId, true
}

// HasSafeUrlId returns a boolean if a field has been set.
func (o *SafeListItem) HasSafeUrlId() bool {
	if o != nil && o.SafeUrlId != nil {
		return true
	}

	return false
}

// SetSafeUrlId gets a reference to the given string and assigns it to the SafeUrlId field.
func (o *SafeListItem) SetSafeUrlId(v string) {
	o.SafeUrlId = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *SafeListItem) GetSafeName() string {
	if o == nil || o.SafeName == nil {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetSafeNameOk() (*string, bool) {
	if o == nil || o.SafeName == nil {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *SafeListItem) HasSafeName() bool {
	if o != nil && o.SafeName != nil {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *SafeListItem) SetSafeName(v string) {
	o.SafeName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SafeListItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SafeListItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SafeListItem) SetDescription(v string) {
	o.Description = &v
}

// GetManagingCPM returns the ManagingCPM field value if set, zero value otherwise.
func (o *SafeListItem) GetManagingCPM() string {
	if o == nil || o.ManagingCPM == nil {
		var ret string
		return ret
	}
	return *o.ManagingCPM
}

// GetManagingCPMOk returns a tuple with the ManagingCPM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetManagingCPMOk() (*string, bool) {
	if o == nil || o.ManagingCPM == nil {
		return nil, false
	}
	return o.ManagingCPM, true
}

// HasManagingCPM returns a boolean if a field has been set.
func (o *SafeListItem) HasManagingCPM() bool {
	if o != nil && o.ManagingCPM != nil {
		return true
	}

	return false
}

// SetManagingCPM gets a reference to the given string and assigns it to the ManagingCPM field.
func (o *SafeListItem) SetManagingCPM(v string) {
	o.ManagingCPM = &v
}

// GetIsExpiredMember returns the IsExpiredMember field value if set, zero value otherwise.
func (o *SafeListItem) GetIsExpiredMember() bool {
	if o == nil || o.IsExpiredMember == nil {
		var ret bool
		return ret
	}
	return *o.IsExpiredMember
}

// GetIsExpiredMemberOk returns a tuple with the IsExpiredMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeListItem) GetIsExpiredMemberOk() (*bool, bool) {
	if o == nil || o.IsExpiredMember == nil {
		return nil, false
	}
	return o.IsExpiredMember, true
}

// HasIsExpiredMember returns a boolean if a field has been set.
func (o *SafeListItem) HasIsExpiredMember() bool {
	if o != nil && o.IsExpiredMember != nil {
		return true
	}

	return false
}

// SetIsExpiredMember gets a reference to the given bool and assigns it to the IsExpiredMember field.
func (o *SafeListItem) SetIsExpiredMember(v bool) {
	o.IsExpiredMember = &v
}

func (o SafeListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SafeNumber != nil {
		toSerialize["safeNumber"] = o.SafeNumber
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.OlacEnabled != nil {
		toSerialize["olacEnabled"] = o.OlacEnabled
	}
	if o.NumberOfVersionsRetention != nil {
		toSerialize["numberOfVersionsRetention"] = o.NumberOfVersionsRetention
	}
	if o.NumberOfDaysRetention != nil {
		toSerialize["numberOfDaysRetention"] = o.NumberOfDaysRetention
	}
	if o.AutoPurgeEnabled != nil {
		toSerialize["autoPurgeEnabled"] = o.AutoPurgeEnabled
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.LastModificationTime != nil {
		toSerialize["lastModificationTime"] = o.LastModificationTime
	}
	if o.SafeUrlId != nil {
		toSerialize["safeUrlId"] = o.SafeUrlId
	}
	if o.SafeName != nil {
		toSerialize["safeName"] = o.SafeName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ManagingCPM != nil {
		toSerialize["managingCPM"] = o.ManagingCPM
	}
	if o.IsExpiredMember != nil {
		toSerialize["isExpiredMember"] = o.IsExpiredMember
	}
	return json.Marshal(toSerialize)
}

type NullableSafeListItem struct {
	value *SafeListItem
	isSet bool
}

func (v NullableSafeListItem) Get() *SafeListItem {
	return v.value
}

func (v *NullableSafeListItem) Set(val *SafeListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSafeListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSafeListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafeListItem(val *SafeListItem) *NullableSafeListItem {
	return &NullableSafeListItem{value: val, isSet: true}
}

func (v NullableSafeListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafeListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
