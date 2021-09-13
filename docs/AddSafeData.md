# AddSafeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfDaysRetention** | Pointer to **int32** | The number of days that password versions are saved in the Safe. | [optional] 
**NumberOfVersionsRetention** | Pointer to **int32** | The number of retained versions of every password that is stored in the Safe. | [optional] 
**OLACEnabled** | Pointer to **bool** | Whether or not to enable Object Level Access Control for the Safe. | [optional] 
**ManagingCPM** | Pointer to **string** | The name of the CPM user who will manage the Safe. | [optional] 
**SafeName** | Pointer to **string** | The name of the Safe. | [optional] 
**Description** | Pointer to **string** | The description of the Safe. | [optional] 
**Location** | Pointer to **string** | The location of the Safe in the Vault. | [optional] 

## Methods

### NewAddSafeData

`func NewAddSafeData() *AddSafeData`

NewAddSafeData instantiates a new AddSafeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSafeDataWithDefaults

`func NewAddSafeDataWithDefaults() *AddSafeData`

NewAddSafeDataWithDefaults instantiates a new AddSafeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfDaysRetention

`func (o *AddSafeData) GetNumberOfDaysRetention() int32`

GetNumberOfDaysRetention returns the NumberOfDaysRetention field if non-nil, zero value otherwise.

### GetNumberOfDaysRetentionOk

`func (o *AddSafeData) GetNumberOfDaysRetentionOk() (*int32, bool)`

GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysRetention

`func (o *AddSafeData) SetNumberOfDaysRetention(v int32)`

SetNumberOfDaysRetention sets NumberOfDaysRetention field to given value.

### HasNumberOfDaysRetention

`func (o *AddSafeData) HasNumberOfDaysRetention() bool`

HasNumberOfDaysRetention returns a boolean if a field has been set.

### GetNumberOfVersionsRetention

`func (o *AddSafeData) GetNumberOfVersionsRetention() int32`

GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field if non-nil, zero value otherwise.

### GetNumberOfVersionsRetentionOk

`func (o *AddSafeData) GetNumberOfVersionsRetentionOk() (*int32, bool)`

GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsRetention

`func (o *AddSafeData) SetNumberOfVersionsRetention(v int32)`

SetNumberOfVersionsRetention sets NumberOfVersionsRetention field to given value.

### HasNumberOfVersionsRetention

`func (o *AddSafeData) HasNumberOfVersionsRetention() bool`

HasNumberOfVersionsRetention returns a boolean if a field has been set.

### GetOLACEnabled

`func (o *AddSafeData) GetOLACEnabled() bool`

GetOLACEnabled returns the OLACEnabled field if non-nil, zero value otherwise.

### GetOLACEnabledOk

`func (o *AddSafeData) GetOLACEnabledOk() (*bool, bool)`

GetOLACEnabledOk returns a tuple with the OLACEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOLACEnabled

`func (o *AddSafeData) SetOLACEnabled(v bool)`

SetOLACEnabled sets OLACEnabled field to given value.

### HasOLACEnabled

`func (o *AddSafeData) HasOLACEnabled() bool`

HasOLACEnabled returns a boolean if a field has been set.

### GetManagingCPM

`func (o *AddSafeData) GetManagingCPM() string`

GetManagingCPM returns the ManagingCPM field if non-nil, zero value otherwise.

### GetManagingCPMOk

`func (o *AddSafeData) GetManagingCPMOk() (*string, bool)`

GetManagingCPMOk returns a tuple with the ManagingCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingCPM

`func (o *AddSafeData) SetManagingCPM(v string)`

SetManagingCPM sets ManagingCPM field to given value.

### HasManagingCPM

`func (o *AddSafeData) HasManagingCPM() bool`

HasManagingCPM returns a boolean if a field has been set.

### GetSafeName

`func (o *AddSafeData) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *AddSafeData) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *AddSafeData) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *AddSafeData) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSafeData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSafeData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSafeData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSafeData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *AddSafeData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddSafeData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddSafeData) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddSafeData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


