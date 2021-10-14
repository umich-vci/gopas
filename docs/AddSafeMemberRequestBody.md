# AddSafeMemberRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberName** | **string** | The Vault user name, Domain user name or group name of the Safe member. | 
**SearchIn** | Pointer to **string** | The Vault or Domain where the user or group was found. | [optional] 
**MembershipExpirationDate** | Pointer to **int64** |  | [optional] 
**Permissions** | **map[string]bool** |  | 

## Methods

### NewAddSafeMemberRequestBody

`func NewAddSafeMemberRequestBody(memberName string, permissions map[string]bool, ) *AddSafeMemberRequestBody`

NewAddSafeMemberRequestBody instantiates a new AddSafeMemberRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSafeMemberRequestBodyWithDefaults

`func NewAddSafeMemberRequestBodyWithDefaults() *AddSafeMemberRequestBody`

NewAddSafeMemberRequestBodyWithDefaults instantiates a new AddSafeMemberRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberName

`func (o *AddSafeMemberRequestBody) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *AddSafeMemberRequestBody) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *AddSafeMemberRequestBody) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.


### GetSearchIn

`func (o *AddSafeMemberRequestBody) GetSearchIn() string`

GetSearchIn returns the SearchIn field if non-nil, zero value otherwise.

### GetSearchInOk

`func (o *AddSafeMemberRequestBody) GetSearchInOk() (*string, bool)`

GetSearchInOk returns a tuple with the SearchIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIn

`func (o *AddSafeMemberRequestBody) SetSearchIn(v string)`

SetSearchIn sets SearchIn field to given value.

### HasSearchIn

`func (o *AddSafeMemberRequestBody) HasSearchIn() bool`

HasSearchIn returns a boolean if a field has been set.

### GetMembershipExpirationDate

`func (o *AddSafeMemberRequestBody) GetMembershipExpirationDate() int64`

GetMembershipExpirationDate returns the MembershipExpirationDate field if non-nil, zero value otherwise.

### GetMembershipExpirationDateOk

`func (o *AddSafeMemberRequestBody) GetMembershipExpirationDateOk() (*int64, bool)`

GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipExpirationDate

`func (o *AddSafeMemberRequestBody) SetMembershipExpirationDate(v int64)`

SetMembershipExpirationDate sets MembershipExpirationDate field to given value.

### HasMembershipExpirationDate

`func (o *AddSafeMemberRequestBody) HasMembershipExpirationDate() bool`

HasMembershipExpirationDate returns a boolean if a field has been set.

### GetPermissions

`func (o *AddSafeMemberRequestBody) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddSafeMemberRequestBody) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddSafeMemberRequestBody) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


