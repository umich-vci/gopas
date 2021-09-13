# GroupMembershipDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupID** | Pointer to **int64** | The unique ID of the group. | [optional] 
**GroupName** | Pointer to **string** | The name of a group in the Vault. | [optional] 
**GroupType** | Pointer to **string** | Whether this is a Vault group or a Directory group. | [optional] 

## Methods

### NewGroupMembershipDetails

`func NewGroupMembershipDetails() *GroupMembershipDetails`

NewGroupMembershipDetails instantiates a new GroupMembershipDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMembershipDetailsWithDefaults

`func NewGroupMembershipDetailsWithDefaults() *GroupMembershipDetails`

NewGroupMembershipDetailsWithDefaults instantiates a new GroupMembershipDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupID

`func (o *GroupMembershipDetails) GetGroupID() int64`

GetGroupID returns the GroupID field if non-nil, zero value otherwise.

### GetGroupIDOk

`func (o *GroupMembershipDetails) GetGroupIDOk() (*int64, bool)`

GetGroupIDOk returns a tuple with the GroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupID

`func (o *GroupMembershipDetails) SetGroupID(v int64)`

SetGroupID sets GroupID field to given value.

### HasGroupID

`func (o *GroupMembershipDetails) HasGroupID() bool`

HasGroupID returns a boolean if a field has been set.

### GetGroupName

`func (o *GroupMembershipDetails) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupMembershipDetails) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupMembershipDetails) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupMembershipDetails) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupType

`func (o *GroupMembershipDetails) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GroupMembershipDetails) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GroupMembershipDetails) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *GroupMembershipDetails) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


