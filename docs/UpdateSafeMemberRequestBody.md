# UpdateSafeMemberRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembershipExpirationDate** | Pointer to **int64** |  | [optional] 
**Permissions** | **map[string]bool** |  | 

## Methods

### NewUpdateSafeMemberRequestBody

`func NewUpdateSafeMemberRequestBody(permissions map[string]bool, ) *UpdateSafeMemberRequestBody`

NewUpdateSafeMemberRequestBody instantiates a new UpdateSafeMemberRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSafeMemberRequestBodyWithDefaults

`func NewUpdateSafeMemberRequestBodyWithDefaults() *UpdateSafeMemberRequestBody`

NewUpdateSafeMemberRequestBodyWithDefaults instantiates a new UpdateSafeMemberRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembershipExpirationDate

`func (o *UpdateSafeMemberRequestBody) GetMembershipExpirationDate() int64`

GetMembershipExpirationDate returns the MembershipExpirationDate field if non-nil, zero value otherwise.

### GetMembershipExpirationDateOk

`func (o *UpdateSafeMemberRequestBody) GetMembershipExpirationDateOk() (*int64, bool)`

GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipExpirationDate

`func (o *UpdateSafeMemberRequestBody) SetMembershipExpirationDate(v int64)`

SetMembershipExpirationDate sets MembershipExpirationDate field to given value.

### HasMembershipExpirationDate

`func (o *UpdateSafeMemberRequestBody) HasMembershipExpirationDate() bool`

HasMembershipExpirationDate returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateSafeMemberRequestBody) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateSafeMemberRequestBody) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateSafeMemberRequestBody) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


