# AutomaticSecretManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticManagementEnabled** | Pointer to **bool** | Whether the account secret is automatically managed by the CPM. Default: true. | [optional] 
**ManualManagementReason** | Pointer to **string** | Reason for disabling automatic secret management | [optional] 
**Status** | Pointer to **string** | Account management status. | [optional] 
**LastModifiedTime** | Pointer to **int64** | Last modified date of the account. | [optional] 
**LastReconciledTime** | Pointer to **int64** | Last reconciled date of the account. | [optional] 
**LastVerifiedTime** | Pointer to **int64** | Last verified date of the account. | [optional] 

## Methods

### NewAutomaticSecretManagement

`func NewAutomaticSecretManagement() *AutomaticSecretManagement`

NewAutomaticSecretManagement instantiates a new AutomaticSecretManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticSecretManagementWithDefaults

`func NewAutomaticSecretManagementWithDefaults() *AutomaticSecretManagement`

NewAutomaticSecretManagementWithDefaults instantiates a new AutomaticSecretManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticManagementEnabled

`func (o *AutomaticSecretManagement) GetAutomaticManagementEnabled() bool`

GetAutomaticManagementEnabled returns the AutomaticManagementEnabled field if non-nil, zero value otherwise.

### GetAutomaticManagementEnabledOk

`func (o *AutomaticSecretManagement) GetAutomaticManagementEnabledOk() (*bool, bool)`

GetAutomaticManagementEnabledOk returns a tuple with the AutomaticManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticManagementEnabled

`func (o *AutomaticSecretManagement) SetAutomaticManagementEnabled(v bool)`

SetAutomaticManagementEnabled sets AutomaticManagementEnabled field to given value.

### HasAutomaticManagementEnabled

`func (o *AutomaticSecretManagement) HasAutomaticManagementEnabled() bool`

HasAutomaticManagementEnabled returns a boolean if a field has been set.

### GetManualManagementReason

`func (o *AutomaticSecretManagement) GetManualManagementReason() string`

GetManualManagementReason returns the ManualManagementReason field if non-nil, zero value otherwise.

### GetManualManagementReasonOk

`func (o *AutomaticSecretManagement) GetManualManagementReasonOk() (*string, bool)`

GetManualManagementReasonOk returns a tuple with the ManualManagementReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualManagementReason

`func (o *AutomaticSecretManagement) SetManualManagementReason(v string)`

SetManualManagementReason sets ManualManagementReason field to given value.

### HasManualManagementReason

`func (o *AutomaticSecretManagement) HasManualManagementReason() bool`

HasManualManagementReason returns a boolean if a field has been set.

### GetStatus

`func (o *AutomaticSecretManagement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomaticSecretManagement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomaticSecretManagement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomaticSecretManagement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *AutomaticSecretManagement) GetLastModifiedTime() int64`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AutomaticSecretManagement) GetLastModifiedTimeOk() (*int64, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AutomaticSecretManagement) SetLastModifiedTime(v int64)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *AutomaticSecretManagement) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetLastReconciledTime

`func (o *AutomaticSecretManagement) GetLastReconciledTime() int64`

GetLastReconciledTime returns the LastReconciledTime field if non-nil, zero value otherwise.

### GetLastReconciledTimeOk

`func (o *AutomaticSecretManagement) GetLastReconciledTimeOk() (*int64, bool)`

GetLastReconciledTimeOk returns a tuple with the LastReconciledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconciledTime

`func (o *AutomaticSecretManagement) SetLastReconciledTime(v int64)`

SetLastReconciledTime sets LastReconciledTime field to given value.

### HasLastReconciledTime

`func (o *AutomaticSecretManagement) HasLastReconciledTime() bool`

HasLastReconciledTime returns a boolean if a field has been set.

### GetLastVerifiedTime

`func (o *AutomaticSecretManagement) GetLastVerifiedTime() int64`

GetLastVerifiedTime returns the LastVerifiedTime field if non-nil, zero value otherwise.

### GetLastVerifiedTimeOk

`func (o *AutomaticSecretManagement) GetLastVerifiedTimeOk() (*int64, bool)`

GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedTime

`func (o *AutomaticSecretManagement) SetLastVerifiedTime(v int64)`

SetLastVerifiedTime sets LastVerifiedTime field to given value.

### HasLastVerifiedTime

`func (o *AutomaticSecretManagement) HasLastVerifiedTime() bool`

HasLastVerifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


