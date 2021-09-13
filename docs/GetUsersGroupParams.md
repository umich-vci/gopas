# GetUsersGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMembers** | Pointer to **bool** | Whether or not to return members for each User Group as part of the response. If not sent, the value will be False. | [optional] 

## Methods

### NewGetUsersGroupParams

`func NewGetUsersGroupParams() *GetUsersGroupParams`

NewGetUsersGroupParams instantiates a new GetUsersGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersGroupParamsWithDefaults

`func NewGetUsersGroupParamsWithDefaults() *GetUsersGroupParams`

NewGetUsersGroupParamsWithDefaults instantiates a new GetUsersGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeMembers

`func (o *GetUsersGroupParams) GetIncludeMembers() bool`

GetIncludeMembers returns the IncludeMembers field if non-nil, zero value otherwise.

### GetIncludeMembersOk

`func (o *GetUsersGroupParams) GetIncludeMembersOk() (*bool, bool)`

GetIncludeMembersOk returns a tuple with the IncludeMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMembers

`func (o *GetUsersGroupParams) SetIncludeMembers(v bool)`

SetIncludeMembers sets IncludeMembers field to given value.

### HasIncludeMembers

`func (o *GetUsersGroupParams) HasIncludeMembers() bool`

HasIncludeMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


