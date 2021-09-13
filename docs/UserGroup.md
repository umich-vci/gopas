# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique ID of the group. | [optional] 
**GroupType** | Pointer to **string** | Whether this is a Vault group or a Directory group. | [optional] 
**Directory** | Pointer to **string** | Displays the name of the LDAP external directory to which the external group belongs.  This is relevant only for Directory group type. | [optional] 
**Dn** | Pointer to **string** | Displays the full LDAP DN of the user in the external directory to which the external user belongs.  This is relevant only for Directory group type. | [optional] 
**Members** | Pointer to [**[]MemberGroup**](MemberGroup.md) |  | [optional] 
**GroupName** | **string** | The name of a group in the Vault. | 
**Description** | Pointer to **string** | The description of the group. | [optional] 
**Location** | Pointer to **string** | The location of the group in the Vault’s hierarchy. | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup(groupName string, ) *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupType

`func (o *UserGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *UserGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *UserGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *UserGroup) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetDirectory

`func (o *UserGroup) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *UserGroup) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *UserGroup) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *UserGroup) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetDn

`func (o *UserGroup) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *UserGroup) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *UserGroup) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *UserGroup) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetMembers

`func (o *UserGroup) GetMembers() []MemberGroup`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UserGroup) GetMembersOk() (*[]MemberGroup, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UserGroup) SetMembers(v []MemberGroup)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *UserGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetGroupName

`func (o *UserGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UserGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UserGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetDescription

`func (o *UserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *UserGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UserGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UserGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UserGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


