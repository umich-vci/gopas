# AddSafeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafeName** | Pointer to **string** | The name of the Safe. | [optional] 
**Description** | Pointer to **string** | The description of the Safe. | [optional] 
**Location** | Pointer to **string** | The location of the Safe in the Vault. | [optional] 
**NumberOfDaysRetention** | Pointer to **int32** | The number of days that password versions are saved in the Safe. | [optional] 
**NumberOfVersionsRetention** | Pointer to **int32** | The number of retained versions of every password that is stored in the Safe. | [optional] 
**OLACEnabled** | Pointer to **bool** | Whether or not to enable Object Level Access Control for the Safe. | [optional] 
**ManagingCPM** | Pointer to **string** | The name of the CPM user who will manage the Safe. | [optional] 

## Methods

### NewAddSafeRequest

`func NewAddSafeRequest() *AddSafeRequest`

NewAddSafeRequest instantiates a new AddSafeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSafeRequestWithDefaults

`func NewAddSafeRequestWithDefaults() *AddSafeRequest`

NewAddSafeRequestWithDefaults instantiates a new AddSafeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeName

`func (o *AddSafeRequest) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *AddSafeRequest) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *AddSafeRequest) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *AddSafeRequest) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSafeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSafeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSafeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSafeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *AddSafeRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddSafeRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddSafeRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddSafeRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNumberOfDaysRetention

`func (o *AddSafeRequest) GetNumberOfDaysRetention() int32`

GetNumberOfDaysRetention returns the NumberOfDaysRetention field if non-nil, zero value otherwise.

### GetNumberOfDaysRetentionOk

`func (o *AddSafeRequest) GetNumberOfDaysRetentionOk() (*int32, bool)`

GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysRetention

`func (o *AddSafeRequest) SetNumberOfDaysRetention(v int32)`

SetNumberOfDaysRetention sets NumberOfDaysRetention field to given value.

### HasNumberOfDaysRetention

`func (o *AddSafeRequest) HasNumberOfDaysRetention() bool`

HasNumberOfDaysRetention returns a boolean if a field has been set.

### GetNumberOfVersionsRetention

`func (o *AddSafeRequest) GetNumberOfVersionsRetention() int32`

GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field if non-nil, zero value otherwise.

### GetNumberOfVersionsRetentionOk

`func (o *AddSafeRequest) GetNumberOfVersionsRetentionOk() (*int32, bool)`

GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsRetention

`func (o *AddSafeRequest) SetNumberOfVersionsRetention(v int32)`

SetNumberOfVersionsRetention sets NumberOfVersionsRetention field to given value.

### HasNumberOfVersionsRetention

`func (o *AddSafeRequest) HasNumberOfVersionsRetention() bool`

HasNumberOfVersionsRetention returns a boolean if a field has been set.

### GetOLACEnabled

`func (o *AddSafeRequest) GetOLACEnabled() bool`

GetOLACEnabled returns the OLACEnabled field if non-nil, zero value otherwise.

### GetOLACEnabledOk

`func (o *AddSafeRequest) GetOLACEnabledOk() (*bool, bool)`

GetOLACEnabledOk returns a tuple with the OLACEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOLACEnabled

`func (o *AddSafeRequest) SetOLACEnabled(v bool)`

SetOLACEnabled sets OLACEnabled field to given value.

### HasOLACEnabled

`func (o *AddSafeRequest) HasOLACEnabled() bool`

HasOLACEnabled returns a boolean if a field has been set.

### GetManagingCPM

`func (o *AddSafeRequest) GetManagingCPM() string`

GetManagingCPM returns the ManagingCPM field if non-nil, zero value otherwise.

### GetManagingCPMOk

`func (o *AddSafeRequest) GetManagingCPMOk() (*string, bool)`

GetManagingCPMOk returns a tuple with the ManagingCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingCPM

`func (o *AddSafeRequest) SetManagingCPM(v string)`

SetManagingCPM sets ManagingCPM field to given value.

### HasManagingCPM

`func (o *AddSafeRequest) HasManagingCPM() bool`

HasManagingCPM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


