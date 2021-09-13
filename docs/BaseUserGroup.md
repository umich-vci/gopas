# BaseUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | The name of a group in the Vault. | 
**Description** | Pointer to **string** | The description of the group. | [optional] 
**Location** | Pointer to **string** | The location of the group in the Vault’s hierarchy. | [optional] 

## Methods

### NewBaseUserGroup

`func NewBaseUserGroup(groupName string, ) *BaseUserGroup`

NewBaseUserGroup instantiates a new BaseUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseUserGroupWithDefaults

`func NewBaseUserGroupWithDefaults() *BaseUserGroup`

NewBaseUserGroupWithDefaults instantiates a new BaseUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *BaseUserGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *BaseUserGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *BaseUserGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDescription

`func (o *BaseUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *BaseUserGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseUserGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseUserGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseUserGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


