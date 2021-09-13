# AddAccountGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | The name of the newly created group. | 
**GroupPlatformId** | **string** | The name of the platform of the group.  The associated platform must be set to PolicyType &#x3D; Group. | 
**Safe** | **string** | The name of the Safe where the group will be created. | 

## Methods

### NewAddAccountGroupData

`func NewAddAccountGroupData(groupName string, groupPlatformId string, safe string, ) *AddAccountGroupData`

NewAddAccountGroupData instantiates a new AddAccountGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountGroupDataWithDefaults

`func NewAddAccountGroupDataWithDefaults() *AddAccountGroupData`

NewAddAccountGroupDataWithDefaults instantiates a new AddAccountGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *AddAccountGroupData) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AddAccountGroupData) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AddAccountGroupData) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetGroupPlatformId

`func (o *AddAccountGroupData) GetGroupPlatformId() string`

GetGroupPlatformId returns the GroupPlatformId field if non-nil, zero value otherwise.

### GetGroupPlatformIdOk

`func (o *AddAccountGroupData) GetGroupPlatformIdOk() (*string, bool)`

GetGroupPlatformIdOk returns a tuple with the GroupPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPlatformId

`func (o *AddAccountGroupData) SetGroupPlatformId(v string)`

SetGroupPlatformId sets GroupPlatformId field to given value.


### GetSafe

`func (o *AddAccountGroupData) GetSafe() string`

GetSafe returns the Safe field if non-nil, zero value otherwise.

### GetSafeOk

`func (o *AddAccountGroupData) GetSafeOk() (*string, bool)`

GetSafeOk returns a tuple with the Safe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafe

`func (o *AddAccountGroupData) SetSafe(v string)`

SetSafe sets Safe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


