# UpdateSafeRequestBody

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

### NewUpdateSafeRequestBody

`func NewUpdateSafeRequestBody() *UpdateSafeRequestBody`

NewUpdateSafeRequestBody instantiates a new UpdateSafeRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSafeRequestBodyWithDefaults

`func NewUpdateSafeRequestBodyWithDefaults() *UpdateSafeRequestBody`

NewUpdateSafeRequestBodyWithDefaults instantiates a new UpdateSafeRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeName

`func (o *UpdateSafeRequestBody) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *UpdateSafeRequestBody) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *UpdateSafeRequestBody) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *UpdateSafeRequestBody) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSafeRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSafeRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSafeRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSafeRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateSafeRequestBody) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateSafeRequestBody) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateSafeRequestBody) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateSafeRequestBody) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNumberOfDaysRetention

`func (o *UpdateSafeRequestBody) GetNumberOfDaysRetention() int32`

GetNumberOfDaysRetention returns the NumberOfDaysRetention field if non-nil, zero value otherwise.

### GetNumberOfDaysRetentionOk

`func (o *UpdateSafeRequestBody) GetNumberOfDaysRetentionOk() (*int32, bool)`

GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysRetention

`func (o *UpdateSafeRequestBody) SetNumberOfDaysRetention(v int32)`

SetNumberOfDaysRetention sets NumberOfDaysRetention field to given value.

### HasNumberOfDaysRetention

`func (o *UpdateSafeRequestBody) HasNumberOfDaysRetention() bool`

HasNumberOfDaysRetention returns a boolean if a field has been set.

### GetNumberOfVersionsRetention

`func (o *UpdateSafeRequestBody) GetNumberOfVersionsRetention() int32`

GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field if non-nil, zero value otherwise.

### GetNumberOfVersionsRetentionOk

`func (o *UpdateSafeRequestBody) GetNumberOfVersionsRetentionOk() (*int32, bool)`

GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsRetention

`func (o *UpdateSafeRequestBody) SetNumberOfVersionsRetention(v int32)`

SetNumberOfVersionsRetention sets NumberOfVersionsRetention field to given value.

### HasNumberOfVersionsRetention

`func (o *UpdateSafeRequestBody) HasNumberOfVersionsRetention() bool`

HasNumberOfVersionsRetention returns a boolean if a field has been set.

### GetOLACEnabled

`func (o *UpdateSafeRequestBody) GetOLACEnabled() bool`

GetOLACEnabled returns the OLACEnabled field if non-nil, zero value otherwise.

### GetOLACEnabledOk

`func (o *UpdateSafeRequestBody) GetOLACEnabledOk() (*bool, bool)`

GetOLACEnabledOk returns a tuple with the OLACEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOLACEnabled

`func (o *UpdateSafeRequestBody) SetOLACEnabled(v bool)`

SetOLACEnabled sets OLACEnabled field to given value.

### HasOLACEnabled

`func (o *UpdateSafeRequestBody) HasOLACEnabled() bool`

HasOLACEnabled returns a boolean if a field has been set.

### GetManagingCPM

`func (o *UpdateSafeRequestBody) GetManagingCPM() string`

GetManagingCPM returns the ManagingCPM field if non-nil, zero value otherwise.

### GetManagingCPMOk

`func (o *UpdateSafeRequestBody) GetManagingCPMOk() (*string, bool)`

GetManagingCPMOk returns a tuple with the ManagingCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingCPM

`func (o *UpdateSafeRequestBody) SetManagingCPM(v string)`

SetManagingCPM sets ManagingCPM field to given value.

### HasManagingCPM

`func (o *UpdateSafeRequestBody) HasManagingCPM() bool`

HasManagingCPM returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


