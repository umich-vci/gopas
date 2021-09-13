# TaskRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformAutomatic** | Pointer to **bool** | Indicates whether the action will be initiated periodically. | [optional] 
**RequirePasswordEveryXDays** | Pointer to **int32** | The number of days between each periodic action. | [optional] 
**AutoOnAdd** | Pointer to **bool** | Indicates whether the action will be performed after account are added. | [optional] 
**IsRequirePasswordEveryXDaysAnException** | Pointer to **bool** | Indicates whether the action will be performed after account are added. | [optional] 
**AllowManual** | Pointer to **bool** | Indicates whether the action can be initiated manually. | [optional] 

## Methods

### NewTaskRole

`func NewTaskRole() *TaskRole`

NewTaskRole instantiates a new TaskRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRoleWithDefaults

`func NewTaskRoleWithDefaults() *TaskRole`

NewTaskRoleWithDefaults instantiates a new TaskRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformAutomatic

`func (o *TaskRole) GetPerformAutomatic() bool`

GetPerformAutomatic returns the PerformAutomatic field if non-nil, zero value otherwise.

### GetPerformAutomaticOk

`func (o *TaskRole) GetPerformAutomaticOk() (*bool, bool)`

GetPerformAutomaticOk returns a tuple with the PerformAutomatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformAutomatic

`func (o *TaskRole) SetPerformAutomatic(v bool)`

SetPerformAutomatic sets PerformAutomatic field to given value.

### HasPerformAutomatic

`func (o *TaskRole) HasPerformAutomatic() bool`

HasPerformAutomatic returns a boolean if a field has been set.

### GetRequirePasswordEveryXDays

`func (o *TaskRole) GetRequirePasswordEveryXDays() int32`

GetRequirePasswordEveryXDays returns the RequirePasswordEveryXDays field if non-nil, zero value otherwise.

### GetRequirePasswordEveryXDaysOk

`func (o *TaskRole) GetRequirePasswordEveryXDaysOk() (*int32, bool)`

GetRequirePasswordEveryXDaysOk returns a tuple with the RequirePasswordEveryXDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePasswordEveryXDays

`func (o *TaskRole) SetRequirePasswordEveryXDays(v int32)`

SetRequirePasswordEveryXDays sets RequirePasswordEveryXDays field to given value.

### HasRequirePasswordEveryXDays

`func (o *TaskRole) HasRequirePasswordEveryXDays() bool`

HasRequirePasswordEveryXDays returns a boolean if a field has been set.

### GetAutoOnAdd

`func (o *TaskRole) GetAutoOnAdd() bool`

GetAutoOnAdd returns the AutoOnAdd field if non-nil, zero value otherwise.

### GetAutoOnAddOk

`func (o *TaskRole) GetAutoOnAddOk() (*bool, bool)`

GetAutoOnAddOk returns a tuple with the AutoOnAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOnAdd

`func (o *TaskRole) SetAutoOnAdd(v bool)`

SetAutoOnAdd sets AutoOnAdd field to given value.

### HasAutoOnAdd

`func (o *TaskRole) HasAutoOnAdd() bool`

HasAutoOnAdd returns a boolean if a field has been set.

### GetIsRequirePasswordEveryXDaysAnException

`func (o *TaskRole) GetIsRequirePasswordEveryXDaysAnException() bool`

GetIsRequirePasswordEveryXDaysAnException returns the IsRequirePasswordEveryXDaysAnException field if non-nil, zero value otherwise.

### GetIsRequirePasswordEveryXDaysAnExceptionOk

`func (o *TaskRole) GetIsRequirePasswordEveryXDaysAnExceptionOk() (*bool, bool)`

GetIsRequirePasswordEveryXDaysAnExceptionOk returns a tuple with the IsRequirePasswordEveryXDaysAnException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequirePasswordEveryXDaysAnException

`func (o *TaskRole) SetIsRequirePasswordEveryXDaysAnException(v bool)`

SetIsRequirePasswordEveryXDaysAnException sets IsRequirePasswordEveryXDaysAnException field to given value.

### HasIsRequirePasswordEveryXDaysAnException

`func (o *TaskRole) HasIsRequirePasswordEveryXDaysAnException() bool`

HasIsRequirePasswordEveryXDaysAnException returns a boolean if a field has been set.

### GetAllowManual

`func (o *TaskRole) GetAllowManual() bool`

GetAllowManual returns the AllowManual field if non-nil, zero value otherwise.

### GetAllowManualOk

`func (o *TaskRole) GetAllowManualOk() (*bool, bool)`

GetAllowManualOk returns a tuple with the AllowManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowManual

`func (o *TaskRole) SetAllowManual(v bool)`

SetAllowManual sets AllowManual field to given value.

### HasAllowManual

`func (o *TaskRole) HasAllowManual() bool`

HasAllowManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


