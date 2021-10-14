# UpdateSafeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafeUrlId** | Pointer to **string** | The name of the Safe to be used when calling Safe APIs. | [optional] 
**SafeName** | Pointer to **string** | The name of the Safe. | [optional] 
**SafeNumber** | Pointer to **int64** | The unique numerical ID of the Safe. | [optional] 
**Description** | Pointer to **string** | The description of the Safe. | [optional] 
**Location** | Pointer to **string** | The location of the Safe in the Vault. | [optional] 
**Creator** | Pointer to [**SafeCreatorResponse**](SafeCreatorResponse.md) |  | [optional] 
**OlacEnabled** | Pointer to **bool** | Whether or not to enable Object Level Access Control for the new Safe. | [optional] 
**ManagingCPM** | Pointer to **string** | The name of the CPM user who will manage the new Safe. | [optional] 
**NumberOfVersionsRetention** | Pointer to **int32** | The number of retained versions of every password that is stored in the Safe. | [optional] 
**NumberOfDaysRetention** | Pointer to **int32** | The number of days that password versions are saved in the Safe. | [optional] 
**AutoPurgeEnabled** | Pointer to **bool** | Whether or not to automatically purge files after the end of the Object History Retention Period defined in the Safe properties.  Reports Safes and PSM Recording Safes are created automatically with AutoPurgeEnabled set to Yes.These Safes cannot be managed by the CPM. | [optional] 
**CreationTime** | Pointer to **int64** | Unix creation time of the Safe. | [optional] 
**LastModificationTime** | Pointer to **int64** | Unix time when the Safe was last updated. | [optional] 

## Methods

### NewUpdateSafeResponse

`func NewUpdateSafeResponse() *UpdateSafeResponse`

NewUpdateSafeResponse instantiates a new UpdateSafeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSafeResponseWithDefaults

`func NewUpdateSafeResponseWithDefaults() *UpdateSafeResponse`

NewUpdateSafeResponseWithDefaults instantiates a new UpdateSafeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeUrlId

`func (o *UpdateSafeResponse) GetSafeUrlId() string`

GetSafeUrlId returns the SafeUrlId field if non-nil, zero value otherwise.

### GetSafeUrlIdOk

`func (o *UpdateSafeResponse) GetSafeUrlIdOk() (*string, bool)`

GetSafeUrlIdOk returns a tuple with the SafeUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeUrlId

`func (o *UpdateSafeResponse) SetSafeUrlId(v string)`

SetSafeUrlId sets SafeUrlId field to given value.

### HasSafeUrlId

`func (o *UpdateSafeResponse) HasSafeUrlId() bool`

HasSafeUrlId returns a boolean if a field has been set.

### GetSafeName

`func (o *UpdateSafeResponse) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *UpdateSafeResponse) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *UpdateSafeResponse) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *UpdateSafeResponse) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetSafeNumber

`func (o *UpdateSafeResponse) GetSafeNumber() int64`

GetSafeNumber returns the SafeNumber field if non-nil, zero value otherwise.

### GetSafeNumberOk

`func (o *UpdateSafeResponse) GetSafeNumberOk() (*int64, bool)`

GetSafeNumberOk returns a tuple with the SafeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeNumber

`func (o *UpdateSafeResponse) SetSafeNumber(v int64)`

SetSafeNumber sets SafeNumber field to given value.

### HasSafeNumber

`func (o *UpdateSafeResponse) HasSafeNumber() bool`

HasSafeNumber returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSafeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSafeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSafeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSafeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateSafeResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateSafeResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateSafeResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateSafeResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCreator

`func (o *UpdateSafeResponse) GetCreator() SafeCreatorResponse`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *UpdateSafeResponse) GetCreatorOk() (*SafeCreatorResponse, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *UpdateSafeResponse) SetCreator(v SafeCreatorResponse)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *UpdateSafeResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetOlacEnabled

`func (o *UpdateSafeResponse) GetOlacEnabled() bool`

GetOlacEnabled returns the OlacEnabled field if non-nil, zero value otherwise.

### GetOlacEnabledOk

`func (o *UpdateSafeResponse) GetOlacEnabledOk() (*bool, bool)`

GetOlacEnabledOk returns a tuple with the OlacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlacEnabled

`func (o *UpdateSafeResponse) SetOlacEnabled(v bool)`

SetOlacEnabled sets OlacEnabled field to given value.

### HasOlacEnabled

`func (o *UpdateSafeResponse) HasOlacEnabled() bool`

HasOlacEnabled returns a boolean if a field has been set.

### GetManagingCPM

`func (o *UpdateSafeResponse) GetManagingCPM() string`

GetManagingCPM returns the ManagingCPM field if non-nil, zero value otherwise.

### GetManagingCPMOk

`func (o *UpdateSafeResponse) GetManagingCPMOk() (*string, bool)`

GetManagingCPMOk returns a tuple with the ManagingCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingCPM

`func (o *UpdateSafeResponse) SetManagingCPM(v string)`

SetManagingCPM sets ManagingCPM field to given value.

### HasManagingCPM

`func (o *UpdateSafeResponse) HasManagingCPM() bool`

HasManagingCPM returns a boolean if a field has been set.

### GetNumberOfVersionsRetention

`func (o *UpdateSafeResponse) GetNumberOfVersionsRetention() int32`

GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field if non-nil, zero value otherwise.

### GetNumberOfVersionsRetentionOk

`func (o *UpdateSafeResponse) GetNumberOfVersionsRetentionOk() (*int32, bool)`

GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsRetention

`func (o *UpdateSafeResponse) SetNumberOfVersionsRetention(v int32)`

SetNumberOfVersionsRetention sets NumberOfVersionsRetention field to given value.

### HasNumberOfVersionsRetention

`func (o *UpdateSafeResponse) HasNumberOfVersionsRetention() bool`

HasNumberOfVersionsRetention returns a boolean if a field has been set.

### GetNumberOfDaysRetention

`func (o *UpdateSafeResponse) GetNumberOfDaysRetention() int32`

GetNumberOfDaysRetention returns the NumberOfDaysRetention field if non-nil, zero value otherwise.

### GetNumberOfDaysRetentionOk

`func (o *UpdateSafeResponse) GetNumberOfDaysRetentionOk() (*int32, bool)`

GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysRetention

`func (o *UpdateSafeResponse) SetNumberOfDaysRetention(v int32)`

SetNumberOfDaysRetention sets NumberOfDaysRetention field to given value.

### HasNumberOfDaysRetention

`func (o *UpdateSafeResponse) HasNumberOfDaysRetention() bool`

HasNumberOfDaysRetention returns a boolean if a field has been set.

### GetAutoPurgeEnabled

`func (o *UpdateSafeResponse) GetAutoPurgeEnabled() bool`

GetAutoPurgeEnabled returns the AutoPurgeEnabled field if non-nil, zero value otherwise.

### GetAutoPurgeEnabledOk

`func (o *UpdateSafeResponse) GetAutoPurgeEnabledOk() (*bool, bool)`

GetAutoPurgeEnabledOk returns a tuple with the AutoPurgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPurgeEnabled

`func (o *UpdateSafeResponse) SetAutoPurgeEnabled(v bool)`

SetAutoPurgeEnabled sets AutoPurgeEnabled field to given value.

### HasAutoPurgeEnabled

`func (o *UpdateSafeResponse) HasAutoPurgeEnabled() bool`

HasAutoPurgeEnabled returns a boolean if a field has been set.

### GetCreationTime

`func (o *UpdateSafeResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *UpdateSafeResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *UpdateSafeResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *UpdateSafeResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *UpdateSafeResponse) GetLastModificationTime() int64`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *UpdateSafeResponse) GetLastModificationTimeOk() (*int64, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *UpdateSafeResponse) SetLastModificationTime(v int64)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *UpdateSafeResponse) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


