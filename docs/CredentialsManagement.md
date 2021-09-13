# CredentialsManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSafes** | Pointer to **string** | A list of safes (regular expression) to which this platform can be applied | [optional] 
**AllowManualChange** | Pointer to **bool** | Indicates whether a &#39;Change now&#39; process can be initiated manually. valid values: true\\false | [optional] 
**PerformPeriodicChange** | Pointer to **bool** | Indicates whether accounts related to this platform will be changed periodically. valid values: true\\false | [optional] 
**RequirePasswordChangeEveryXDays** | Pointer to **int32** | The number of days between each periodic change. | [optional] 
**AllowManualVerification** | Pointer to **bool** | Indicates whether a &#39;Verify now&#39; process can be initiated manually. valid values: true\\false | [optional] 
**PerformPeriodicVerification** | Pointer to **bool** | Indicates whether accounts related to this platform will be verified periodically. valid values: true\\false | [optional] 
**RequirePasswordVerificationEveryXDays** | Pointer to **int32** | The number of days between each periodic verification. | [optional] 
**AllowManualReconciliation** | Pointer to **bool** | Indicates whether a &#39;Reconcile now&#39; process can be initiated manually. valid values: true\\false | [optional] 
**AutomaticReconcileWhenUnsynched** | Pointer to **bool** | Whether or not passwords will be reconciled automatically after the CPM detects a password on a remote machine  that is not synchronized with its corresponding password in the Server.valid values: true\\false | [optional] 

## Methods

### NewCredentialsManagement

`func NewCredentialsManagement() *CredentialsManagement`

NewCredentialsManagement instantiates a new CredentialsManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsManagementWithDefaults

`func NewCredentialsManagementWithDefaults() *CredentialsManagement`

NewCredentialsManagementWithDefaults instantiates a new CredentialsManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSafes

`func (o *CredentialsManagement) GetAllowedSafes() string`

GetAllowedSafes returns the AllowedSafes field if non-nil, zero value otherwise.

### GetAllowedSafesOk

`func (o *CredentialsManagement) GetAllowedSafesOk() (*string, bool)`

GetAllowedSafesOk returns a tuple with the AllowedSafes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSafes

`func (o *CredentialsManagement) SetAllowedSafes(v string)`

SetAllowedSafes sets AllowedSafes field to given value.

### HasAllowedSafes

`func (o *CredentialsManagement) HasAllowedSafes() bool`

HasAllowedSafes returns a boolean if a field has been set.

### GetAllowManualChange

`func (o *CredentialsManagement) GetAllowManualChange() bool`

GetAllowManualChange returns the AllowManualChange field if non-nil, zero value otherwise.

### GetAllowManualChangeOk

`func (o *CredentialsManagement) GetAllowManualChangeOk() (*bool, bool)`

GetAllowManualChangeOk returns a tuple with the AllowManualChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowManualChange

`func (o *CredentialsManagement) SetAllowManualChange(v bool)`

SetAllowManualChange sets AllowManualChange field to given value.

### HasAllowManualChange

`func (o *CredentialsManagement) HasAllowManualChange() bool`

HasAllowManualChange returns a boolean if a field has been set.

### GetPerformPeriodicChange

`func (o *CredentialsManagement) GetPerformPeriodicChange() bool`

GetPerformPeriodicChange returns the PerformPeriodicChange field if non-nil, zero value otherwise.

### GetPerformPeriodicChangeOk

`func (o *CredentialsManagement) GetPerformPeriodicChangeOk() (*bool, bool)`

GetPerformPeriodicChangeOk returns a tuple with the PerformPeriodicChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformPeriodicChange

`func (o *CredentialsManagement) SetPerformPeriodicChange(v bool)`

SetPerformPeriodicChange sets PerformPeriodicChange field to given value.

### HasPerformPeriodicChange

`func (o *CredentialsManagement) HasPerformPeriodicChange() bool`

HasPerformPeriodicChange returns a boolean if a field has been set.

### GetRequirePasswordChangeEveryXDays

`func (o *CredentialsManagement) GetRequirePasswordChangeEveryXDays() int32`

GetRequirePasswordChangeEveryXDays returns the RequirePasswordChangeEveryXDays field if non-nil, zero value otherwise.

### GetRequirePasswordChangeEveryXDaysOk

`func (o *CredentialsManagement) GetRequirePasswordChangeEveryXDaysOk() (*int32, bool)`

GetRequirePasswordChangeEveryXDaysOk returns a tuple with the RequirePasswordChangeEveryXDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePasswordChangeEveryXDays

`func (o *CredentialsManagement) SetRequirePasswordChangeEveryXDays(v int32)`

SetRequirePasswordChangeEveryXDays sets RequirePasswordChangeEveryXDays field to given value.

### HasRequirePasswordChangeEveryXDays

`func (o *CredentialsManagement) HasRequirePasswordChangeEveryXDays() bool`

HasRequirePasswordChangeEveryXDays returns a boolean if a field has been set.

### GetAllowManualVerification

`func (o *CredentialsManagement) GetAllowManualVerification() bool`

GetAllowManualVerification returns the AllowManualVerification field if non-nil, zero value otherwise.

### GetAllowManualVerificationOk

`func (o *CredentialsManagement) GetAllowManualVerificationOk() (*bool, bool)`

GetAllowManualVerificationOk returns a tuple with the AllowManualVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowManualVerification

`func (o *CredentialsManagement) SetAllowManualVerification(v bool)`

SetAllowManualVerification sets AllowManualVerification field to given value.

### HasAllowManualVerification

`func (o *CredentialsManagement) HasAllowManualVerification() bool`

HasAllowManualVerification returns a boolean if a field has been set.

### GetPerformPeriodicVerification

`func (o *CredentialsManagement) GetPerformPeriodicVerification() bool`

GetPerformPeriodicVerification returns the PerformPeriodicVerification field if non-nil, zero value otherwise.

### GetPerformPeriodicVerificationOk

`func (o *CredentialsManagement) GetPerformPeriodicVerificationOk() (*bool, bool)`

GetPerformPeriodicVerificationOk returns a tuple with the PerformPeriodicVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformPeriodicVerification

`func (o *CredentialsManagement) SetPerformPeriodicVerification(v bool)`

SetPerformPeriodicVerification sets PerformPeriodicVerification field to given value.

### HasPerformPeriodicVerification

`func (o *CredentialsManagement) HasPerformPeriodicVerification() bool`

HasPerformPeriodicVerification returns a boolean if a field has been set.

### GetRequirePasswordVerificationEveryXDays

`func (o *CredentialsManagement) GetRequirePasswordVerificationEveryXDays() int32`

GetRequirePasswordVerificationEveryXDays returns the RequirePasswordVerificationEveryXDays field if non-nil, zero value otherwise.

### GetRequirePasswordVerificationEveryXDaysOk

`func (o *CredentialsManagement) GetRequirePasswordVerificationEveryXDaysOk() (*int32, bool)`

GetRequirePasswordVerificationEveryXDaysOk returns a tuple with the RequirePasswordVerificationEveryXDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePasswordVerificationEveryXDays

`func (o *CredentialsManagement) SetRequirePasswordVerificationEveryXDays(v int32)`

SetRequirePasswordVerificationEveryXDays sets RequirePasswordVerificationEveryXDays field to given value.

### HasRequirePasswordVerificationEveryXDays

`func (o *CredentialsManagement) HasRequirePasswordVerificationEveryXDays() bool`

HasRequirePasswordVerificationEveryXDays returns a boolean if a field has been set.

### GetAllowManualReconciliation

`func (o *CredentialsManagement) GetAllowManualReconciliation() bool`

GetAllowManualReconciliation returns the AllowManualReconciliation field if non-nil, zero value otherwise.

### GetAllowManualReconciliationOk

`func (o *CredentialsManagement) GetAllowManualReconciliationOk() (*bool, bool)`

GetAllowManualReconciliationOk returns a tuple with the AllowManualReconciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowManualReconciliation

`func (o *CredentialsManagement) SetAllowManualReconciliation(v bool)`

SetAllowManualReconciliation sets AllowManualReconciliation field to given value.

### HasAllowManualReconciliation

`func (o *CredentialsManagement) HasAllowManualReconciliation() bool`

HasAllowManualReconciliation returns a boolean if a field has been set.

### GetAutomaticReconcileWhenUnsynched

`func (o *CredentialsManagement) GetAutomaticReconcileWhenUnsynched() bool`

GetAutomaticReconcileWhenUnsynched returns the AutomaticReconcileWhenUnsynched field if non-nil, zero value otherwise.

### GetAutomaticReconcileWhenUnsynchedOk

`func (o *CredentialsManagement) GetAutomaticReconcileWhenUnsynchedOk() (*bool, bool)`

GetAutomaticReconcileWhenUnsynchedOk returns a tuple with the AutomaticReconcileWhenUnsynched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticReconcileWhenUnsynched

`func (o *CredentialsManagement) SetAutomaticReconcileWhenUnsynched(v bool)`

SetAutomaticReconcileWhenUnsynched sets AutomaticReconcileWhenUnsynched field to given value.

### HasAutomaticReconcileWhenUnsynched

`func (o *CredentialsManagement) HasAutomaticReconcileWhenUnsynched() bool`

HasAutomaticReconcileWhenUnsynched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


