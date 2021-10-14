# SafeListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafeNumber** | Pointer to **int64** | The unique numerical ID of the Safe. | [optional] 
**Location** | Pointer to **string** | The location of the Safe in the Vault. | [optional] 
**Creator** | Pointer to [**SafeCreator**](SafeCreator.md) |  | [optional] 
**Accounts** | Pointer to [**[]AccountInSafe**](AccountInSafe.md) |  | [optional] 
**OlacEnabled** | Pointer to **bool** | Whether or not to enable Object Level Access Control for the new Safe. | [optional] 
**NumberOfVersionsRetention** | Pointer to **int32** | The number of retained versions of every password that is stored in the Safe. | [optional] 
**NumberOfDaysRetention** | Pointer to **int32** | The number of days that password versions are saved in the Safe. | [optional] 
**AutoPurgeEnabled** | Pointer to **bool** | Whether or not to automatically purge files after the end of the Object History Retention Period defined in the Safe properties.  Reports Safes and PSM Recording Safes are created automatically with AutoPurgeEnabled set to Yes.These Safes cannot be managed by the CPM. | [optional] 
**CreationTime** | Pointer to **int64** | Unix creation time of the Safe. | [optional] 
**LastModificationTime** | Pointer to **int64** | Unix time when the Safe was last updated. | [optional] 
**SafeUrlId** | Pointer to **string** | The unique ID of the Safe to be used when calling Safe APIs. | [optional] 
**SafeName** | Pointer to **string** | The name of the Safe. | [optional] 
**Description** | Pointer to **string** | The description of the Safe. | [optional] 
**ManagingCPM** | Pointer to **string** | The name of the CPM user who will manage the new Safe. | [optional] 
**IsExpiredMember** | Pointer to **bool** | Whether or not the membership for the Safe is expired.For expired member, the value will be True. | [optional] 

## Methods

### NewSafeListItem

`func NewSafeListItem() *SafeListItem`

NewSafeListItem instantiates a new SafeListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeListItemWithDefaults

`func NewSafeListItemWithDefaults() *SafeListItem`

NewSafeListItemWithDefaults instantiates a new SafeListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeNumber

`func (o *SafeListItem) GetSafeNumber() int64`

GetSafeNumber returns the SafeNumber field if non-nil, zero value otherwise.

### GetSafeNumberOk

`func (o *SafeListItem) GetSafeNumberOk() (*int64, bool)`

GetSafeNumberOk returns a tuple with the SafeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeNumber

`func (o *SafeListItem) SetSafeNumber(v int64)`

SetSafeNumber sets SafeNumber field to given value.

### HasSafeNumber

`func (o *SafeListItem) HasSafeNumber() bool`

HasSafeNumber returns a boolean if a field has been set.

### GetLocation

`func (o *SafeListItem) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SafeListItem) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SafeListItem) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SafeListItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCreator

`func (o *SafeListItem) GetCreator() SafeCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SafeListItem) GetCreatorOk() (*SafeCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SafeListItem) SetCreator(v SafeCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SafeListItem) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetAccounts

`func (o *SafeListItem) GetAccounts() []AccountInSafe`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SafeListItem) GetAccountsOk() (*[]AccountInSafe, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SafeListItem) SetAccounts(v []AccountInSafe)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SafeListItem) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetOlacEnabled

`func (o *SafeListItem) GetOlacEnabled() bool`

GetOlacEnabled returns the OlacEnabled field if non-nil, zero value otherwise.

### GetOlacEnabledOk

`func (o *SafeListItem) GetOlacEnabledOk() (*bool, bool)`

GetOlacEnabledOk returns a tuple with the OlacEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlacEnabled

`func (o *SafeListItem) SetOlacEnabled(v bool)`

SetOlacEnabled sets OlacEnabled field to given value.

### HasOlacEnabled

`func (o *SafeListItem) HasOlacEnabled() bool`

HasOlacEnabled returns a boolean if a field has been set.

### GetNumberOfVersionsRetention

`func (o *SafeListItem) GetNumberOfVersionsRetention() int32`

GetNumberOfVersionsRetention returns the NumberOfVersionsRetention field if non-nil, zero value otherwise.

### GetNumberOfVersionsRetentionOk

`func (o *SafeListItem) GetNumberOfVersionsRetentionOk() (*int32, bool)`

GetNumberOfVersionsRetentionOk returns a tuple with the NumberOfVersionsRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfVersionsRetention

`func (o *SafeListItem) SetNumberOfVersionsRetention(v int32)`

SetNumberOfVersionsRetention sets NumberOfVersionsRetention field to given value.

### HasNumberOfVersionsRetention

`func (o *SafeListItem) HasNumberOfVersionsRetention() bool`

HasNumberOfVersionsRetention returns a boolean if a field has been set.

### GetNumberOfDaysRetention

`func (o *SafeListItem) GetNumberOfDaysRetention() int32`

GetNumberOfDaysRetention returns the NumberOfDaysRetention field if non-nil, zero value otherwise.

### GetNumberOfDaysRetentionOk

`func (o *SafeListItem) GetNumberOfDaysRetentionOk() (*int32, bool)`

GetNumberOfDaysRetentionOk returns a tuple with the NumberOfDaysRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysRetention

`func (o *SafeListItem) SetNumberOfDaysRetention(v int32)`

SetNumberOfDaysRetention sets NumberOfDaysRetention field to given value.

### HasNumberOfDaysRetention

`func (o *SafeListItem) HasNumberOfDaysRetention() bool`

HasNumberOfDaysRetention returns a boolean if a field has been set.

### GetAutoPurgeEnabled

`func (o *SafeListItem) GetAutoPurgeEnabled() bool`

GetAutoPurgeEnabled returns the AutoPurgeEnabled field if non-nil, zero value otherwise.

### GetAutoPurgeEnabledOk

`func (o *SafeListItem) GetAutoPurgeEnabledOk() (*bool, bool)`

GetAutoPurgeEnabledOk returns a tuple with the AutoPurgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPurgeEnabled

`func (o *SafeListItem) SetAutoPurgeEnabled(v bool)`

SetAutoPurgeEnabled sets AutoPurgeEnabled field to given value.

### HasAutoPurgeEnabled

`func (o *SafeListItem) HasAutoPurgeEnabled() bool`

HasAutoPurgeEnabled returns a boolean if a field has been set.

### GetCreationTime

`func (o *SafeListItem) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SafeListItem) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SafeListItem) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SafeListItem) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *SafeListItem) GetLastModificationTime() int64`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *SafeListItem) GetLastModificationTimeOk() (*int64, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *SafeListItem) SetLastModificationTime(v int64)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *SafeListItem) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetSafeUrlId

`func (o *SafeListItem) GetSafeUrlId() string`

GetSafeUrlId returns the SafeUrlId field if non-nil, zero value otherwise.

### GetSafeUrlIdOk

`func (o *SafeListItem) GetSafeUrlIdOk() (*string, bool)`

GetSafeUrlIdOk returns a tuple with the SafeUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeUrlId

`func (o *SafeListItem) SetSafeUrlId(v string)`

SetSafeUrlId sets SafeUrlId field to given value.

### HasSafeUrlId

`func (o *SafeListItem) HasSafeUrlId() bool`

HasSafeUrlId returns a boolean if a field has been set.

### GetSafeName

`func (o *SafeListItem) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *SafeListItem) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *SafeListItem) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *SafeListItem) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetDescription

`func (o *SafeListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SafeListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SafeListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SafeListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManagingCPM

`func (o *SafeListItem) GetManagingCPM() string`

GetManagingCPM returns the ManagingCPM field if non-nil, zero value otherwise.

### GetManagingCPMOk

`func (o *SafeListItem) GetManagingCPMOk() (*string, bool)`

GetManagingCPMOk returns a tuple with the ManagingCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingCPM

`func (o *SafeListItem) SetManagingCPM(v string)`

SetManagingCPM sets ManagingCPM field to given value.

### HasManagingCPM

`func (o *SafeListItem) HasManagingCPM() bool`

HasManagingCPM returns a boolean if a field has been set.

### GetIsExpiredMember

`func (o *SafeListItem) GetIsExpiredMember() bool`

GetIsExpiredMember returns the IsExpiredMember field if non-nil, zero value otherwise.

### GetIsExpiredMemberOk

`func (o *SafeListItem) GetIsExpiredMemberOk() (*bool, bool)`

GetIsExpiredMemberOk returns a tuple with the IsExpiredMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpiredMember

`func (o *SafeListItem) SetIsExpiredMember(v bool)`

SetIsExpiredMember sets IsExpiredMember field to given value.

### HasIsExpiredMember

`func (o *SafeListItem) HasIsExpiredMember() bool`

HasIsExpiredMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


