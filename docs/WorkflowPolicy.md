# WorkflowPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Indicates whether the policy is active. | [optional] 
**IsAnException** | Pointer to **bool** | Indicates whether the policy is an exception to the master policy. | [optional] 

## Methods

### NewWorkflowPolicy

`func NewWorkflowPolicy() *WorkflowPolicy`

NewWorkflowPolicy instantiates a new WorkflowPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPolicyWithDefaults

`func NewWorkflowPolicyWithDefaults() *WorkflowPolicy`

NewWorkflowPolicyWithDefaults instantiates a new WorkflowPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *WorkflowPolicy) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WorkflowPolicy) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WorkflowPolicy) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WorkflowPolicy) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsAnException

`func (o *WorkflowPolicy) GetIsAnException() bool`

GetIsAnException returns the IsAnException field if non-nil, zero value otherwise.

### GetIsAnExceptionOk

`func (o *WorkflowPolicy) GetIsAnExceptionOk() (*bool, bool)`

GetIsAnExceptionOk returns a tuple with the IsAnException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnException

`func (o *WorkflowPolicy) SetIsAnException(v bool)`

SetIsAnException sets IsAnException field to given value.

### HasIsAnException

`func (o *WorkflowPolicy) HasIsAnException() bool`

HasIsAnException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


