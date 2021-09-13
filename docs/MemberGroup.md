# MemberGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 

## Methods

### NewMemberGroup

`func NewMemberGroup() *MemberGroup`

NewMemberGroup instantiates a new MemberGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberGroupWithDefaults

`func NewMemberGroupWithDefaults() *MemberGroup`

NewMemberGroupWithDefaults instantiates a new MemberGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *MemberGroup) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MemberGroup) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MemberGroup) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MemberGroup) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetId

`func (o *MemberGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MemberGroup) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


