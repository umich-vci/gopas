# RemoteMachinesAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteMachines** | Pointer to **string** | List of remote machines, separated by semicolons. | [optional] 
**AccessRestrictedToRemoteMachines** | Pointer to **bool** | Whether or not to restrict access only to specified remote machines. | [optional] 

## Methods

### NewRemoteMachinesAccess

`func NewRemoteMachinesAccess() *RemoteMachinesAccess`

NewRemoteMachinesAccess instantiates a new RemoteMachinesAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteMachinesAccessWithDefaults

`func NewRemoteMachinesAccessWithDefaults() *RemoteMachinesAccess`

NewRemoteMachinesAccessWithDefaults instantiates a new RemoteMachinesAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteMachines

`func (o *RemoteMachinesAccess) GetRemoteMachines() string`

GetRemoteMachines returns the RemoteMachines field if non-nil, zero value otherwise.

### GetRemoteMachinesOk

`func (o *RemoteMachinesAccess) GetRemoteMachinesOk() (*string, bool)`

GetRemoteMachinesOk returns a tuple with the RemoteMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMachines

`func (o *RemoteMachinesAccess) SetRemoteMachines(v string)`

SetRemoteMachines sets RemoteMachines field to given value.

### HasRemoteMachines

`func (o *RemoteMachinesAccess) HasRemoteMachines() bool`

HasRemoteMachines returns a boolean if a field has been set.

### GetAccessRestrictedToRemoteMachines

`func (o *RemoteMachinesAccess) GetAccessRestrictedToRemoteMachines() bool`

GetAccessRestrictedToRemoteMachines returns the AccessRestrictedToRemoteMachines field if non-nil, zero value otherwise.

### GetAccessRestrictedToRemoteMachinesOk

`func (o *RemoteMachinesAccess) GetAccessRestrictedToRemoteMachinesOk() (*bool, bool)`

GetAccessRestrictedToRemoteMachinesOk returns a tuple with the AccessRestrictedToRemoteMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRestrictedToRemoteMachines

`func (o *RemoteMachinesAccess) SetAccessRestrictedToRemoteMachines(v bool)`

SetAccessRestrictedToRemoteMachines sets AccessRestrictedToRemoteMachines field to given value.

### HasAccessRestrictedToRemoteMachines

`func (o *RemoteMachinesAccess) HasAccessRestrictedToRemoteMachines() bool`

HasAccessRestrictedToRemoteMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


