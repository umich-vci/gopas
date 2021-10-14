# SafeMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafeUrlId** | Pointer to **string** | The unique ID of the Safe to be used when calling Safes APIs. | [optional] 
**SafeName** | Pointer to **string** | The name of the Safe. | [optional] 
**SafeNumber** | Pointer to **int64** | The unique numerical ID of the Safe. | [optional] 
**MemberId** | Pointer to **int64** | Vault user Id, Domain user Id or group Id of the Safe member. | [optional] 
**MemberName** | Pointer to **string** | Vault user name, Domain user name or group name of the Safe member. | [optional] 
**MemberType** | Pointer to **string** | The member type.  Valid values: user\\group | [optional] 
**MembershipExpirationDate** | Pointer to **int64** | The members expiration date for this Safe.  For members that will never expire, this value will be null. | [optional] 
**IsExpiredMembershipEnable** | Pointer to **bool** | Whether the membership on the safe is expired or not. For expired members the value will be true. | [optional] 
**IsPredefinedUser** | Pointer to **bool** | Whether the member is a vault built-in user or group. | [optional] 
**IsReadOnly** | Pointer to **bool** | Whether or not the current user can update the permissions of the member | [optional] 
**Permissions** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewSafeMemberResponse

`func NewSafeMemberResponse() *SafeMemberResponse`

NewSafeMemberResponse instantiates a new SafeMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeMemberResponseWithDefaults

`func NewSafeMemberResponseWithDefaults() *SafeMemberResponse`

NewSafeMemberResponseWithDefaults instantiates a new SafeMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeUrlId

`func (o *SafeMemberResponse) GetSafeUrlId() string`

GetSafeUrlId returns the SafeUrlId field if non-nil, zero value otherwise.

### GetSafeUrlIdOk

`func (o *SafeMemberResponse) GetSafeUrlIdOk() (*string, bool)`

GetSafeUrlIdOk returns a tuple with the SafeUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeUrlId

`func (o *SafeMemberResponse) SetSafeUrlId(v string)`

SetSafeUrlId sets SafeUrlId field to given value.

### HasSafeUrlId

`func (o *SafeMemberResponse) HasSafeUrlId() bool`

HasSafeUrlId returns a boolean if a field has been set.

### GetSafeName

`func (o *SafeMemberResponse) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *SafeMemberResponse) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *SafeMemberResponse) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *SafeMemberResponse) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetSafeNumber

`func (o *SafeMemberResponse) GetSafeNumber() int64`

GetSafeNumber returns the SafeNumber field if non-nil, zero value otherwise.

### GetSafeNumberOk

`func (o *SafeMemberResponse) GetSafeNumberOk() (*int64, bool)`

GetSafeNumberOk returns a tuple with the SafeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeNumber

`func (o *SafeMemberResponse) SetSafeNumber(v int64)`

SetSafeNumber sets SafeNumber field to given value.

### HasSafeNumber

`func (o *SafeMemberResponse) HasSafeNumber() bool`

HasSafeNumber returns a boolean if a field has been set.

### GetMemberId

`func (o *SafeMemberResponse) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *SafeMemberResponse) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *SafeMemberResponse) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *SafeMemberResponse) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMemberName

`func (o *SafeMemberResponse) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *SafeMemberResponse) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *SafeMemberResponse) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *SafeMemberResponse) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetMemberType

`func (o *SafeMemberResponse) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *SafeMemberResponse) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *SafeMemberResponse) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *SafeMemberResponse) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetMembershipExpirationDate

`func (o *SafeMemberResponse) GetMembershipExpirationDate() int64`

GetMembershipExpirationDate returns the MembershipExpirationDate field if non-nil, zero value otherwise.

### GetMembershipExpirationDateOk

`func (o *SafeMemberResponse) GetMembershipExpirationDateOk() (*int64, bool)`

GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipExpirationDate

`func (o *SafeMemberResponse) SetMembershipExpirationDate(v int64)`

SetMembershipExpirationDate sets MembershipExpirationDate field to given value.

### HasMembershipExpirationDate

`func (o *SafeMemberResponse) HasMembershipExpirationDate() bool`

HasMembershipExpirationDate returns a boolean if a field has been set.

### GetIsExpiredMembershipEnable

`func (o *SafeMemberResponse) GetIsExpiredMembershipEnable() bool`

GetIsExpiredMembershipEnable returns the IsExpiredMembershipEnable field if non-nil, zero value otherwise.

### GetIsExpiredMembershipEnableOk

`func (o *SafeMemberResponse) GetIsExpiredMembershipEnableOk() (*bool, bool)`

GetIsExpiredMembershipEnableOk returns a tuple with the IsExpiredMembershipEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpiredMembershipEnable

`func (o *SafeMemberResponse) SetIsExpiredMembershipEnable(v bool)`

SetIsExpiredMembershipEnable sets IsExpiredMembershipEnable field to given value.

### HasIsExpiredMembershipEnable

`func (o *SafeMemberResponse) HasIsExpiredMembershipEnable() bool`

HasIsExpiredMembershipEnable returns a boolean if a field has been set.

### GetIsPredefinedUser

`func (o *SafeMemberResponse) GetIsPredefinedUser() bool`

GetIsPredefinedUser returns the IsPredefinedUser field if non-nil, zero value otherwise.

### GetIsPredefinedUserOk

`func (o *SafeMemberResponse) GetIsPredefinedUserOk() (*bool, bool)`

GetIsPredefinedUserOk returns a tuple with the IsPredefinedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPredefinedUser

`func (o *SafeMemberResponse) SetIsPredefinedUser(v bool)`

SetIsPredefinedUser sets IsPredefinedUser field to given value.

### HasIsPredefinedUser

`func (o *SafeMemberResponse) HasIsPredefinedUser() bool`

HasIsPredefinedUser returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *SafeMemberResponse) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *SafeMemberResponse) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *SafeMemberResponse) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *SafeMemberResponse) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetPermissions

`func (o *SafeMemberResponse) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SafeMemberResponse) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SafeMemberResponse) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SafeMemberResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


