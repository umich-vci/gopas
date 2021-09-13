# GroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | **string** | The name of the member or LDAP group to add to the Vault group. | 
**MemberType** | Pointer to **string** | The type of the member in order to differ between members that come from the domain or from the Vault.    Valid values: domain/vault | [optional] 
**DomainName** | **string** | The dns address of the domain. | 

## Methods

### NewGroupMember

`func NewGroupMember(memberId string, domainName string, ) *GroupMember`

NewGroupMember instantiates a new GroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberWithDefaults

`func NewGroupMemberWithDefaults() *GroupMember`

NewGroupMemberWithDefaults instantiates a new GroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *GroupMember) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *GroupMember) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *GroupMember) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetMemberType

`func (o *GroupMember) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *GroupMember) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *GroupMember) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *GroupMember) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetDomainName

`func (o *GroupMember) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GroupMember) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GroupMember) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


