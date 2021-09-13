# DiscoveredAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserName** | **string** | The name of the account user | 
**Address** | **string** | The name or address of the machine where the account is located | 
**DiscoveryDateTime** | Pointer to **int64** | The date the account was discovered | [optional] 
**AccountEnabled** | Pointer to **bool** | The state of the account, defined in the discovery source.  Note: Domain accounts are discovered in the Active Directory. Local accounts are discovered on machines.  If this parameter is not set, it is considered null | [optional] 
**OsGroups** | Pointer to **string** | The name of the group the account belongs such as Administrators or Operators | [optional] 
**PlatformType** | **string** | The platform where the discovered account is located.  One of the following: Windows Server Local, Windows Desktop Local, Windows Domain, Unix, Unix SSH Key, AWS, AWS Access Keys, Azure Password Management | 
**Domain** | Pointer to **string** | The domain of the account | [optional] 
**LastLogonDateTime** | Pointer to **int64** | The date this account was last logged into, defined in the discovery source | [optional] 
**LastPasswordSetDateTime** | Pointer to **int64** | The date this password was last set, defined in the discovery source | [optional] 
**PasswordNeverExpires** | Pointer to **bool** | Whether or not this password expires defined in the discovery source.  If this parameter is not set, it is considered null.  This attribute cannot be set if passwordExpirationDateTime is specified | [optional] 
**OsVersion** | Pointer to **string** | The version of the OS where the account was discovered | [optional] 
**Privileged** | Pointer to **bool** | Whether the discovered account is privileged or non-privileged.  If this parameter is not set, it is considered null | [optional] 
**PrivilegedCriteria** | Pointer to **string** | The criteria that determines whether or not the discovered account is privileged. For example, the user or group name | [optional] 
**UserDisplayName** | Pointer to **string** | The user&#39;s display name | [optional] 
**Description** | Pointer to **string** | A description of the account, defined in the discovery source | [optional] 
**PasswordExpirationDateTime** | Pointer to **int64** | The expiration date of the account, defined in the discovery source | [optional] 
**OsFamily** | Pointer to **string** | The type of machine where the account was discovered.  If this parameter is not set, it is considered null and will not be returned in the result | [optional] 
**AdditionalPropertiesField** | Pointer to **map[string]interface{}** | List of name&#x3D;value pairs for additional properties added to the account.  This is an open list that is not validated.  The List of properties should be valid File Properties in the vault | [optional] 
**OrganizationalUnit** | Pointer to **string** | The Organizational Unit where the account is defined | [optional] 
**PlatformTypeAccountProperties** | Pointer to **map[string]interface{}** | Object containing key-value pairs to associate with the account, as defined by the account platform type schema.   Only properties that appear in the platform type schema are allowed | [optional] 
**Dependencies** | Pointer to [**[]DiscoveredDependency**](DiscoveredDependency.md) | List of Windows dependencies | [optional] 

## Methods

### NewDiscoveredAccount

`func NewDiscoveredAccount(userName string, address string, platformType string, ) *DiscoveredAccount`

NewDiscoveredAccount instantiates a new DiscoveredAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredAccountWithDefaults

`func NewDiscoveredAccountWithDefaults() *DiscoveredAccount`

NewDiscoveredAccountWithDefaults instantiates a new DiscoveredAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveredAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveredAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserName

`func (o *DiscoveredAccount) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DiscoveredAccount) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DiscoveredAccount) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetAddress

`func (o *DiscoveredAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveredAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveredAccount) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDiscoveryDateTime

`func (o *DiscoveredAccount) GetDiscoveryDateTime() int64`

GetDiscoveryDateTime returns the DiscoveryDateTime field if non-nil, zero value otherwise.

### GetDiscoveryDateTimeOk

`func (o *DiscoveredAccount) GetDiscoveryDateTimeOk() (*int64, bool)`

GetDiscoveryDateTimeOk returns a tuple with the DiscoveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryDateTime

`func (o *DiscoveredAccount) SetDiscoveryDateTime(v int64)`

SetDiscoveryDateTime sets DiscoveryDateTime field to given value.

### HasDiscoveryDateTime

`func (o *DiscoveredAccount) HasDiscoveryDateTime() bool`

HasDiscoveryDateTime returns a boolean if a field has been set.

### GetAccountEnabled

`func (o *DiscoveredAccount) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *DiscoveredAccount) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *DiscoveredAccount) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *DiscoveredAccount) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetOsGroups

`func (o *DiscoveredAccount) GetOsGroups() string`

GetOsGroups returns the OsGroups field if non-nil, zero value otherwise.

### GetOsGroupsOk

`func (o *DiscoveredAccount) GetOsGroupsOk() (*string, bool)`

GetOsGroupsOk returns a tuple with the OsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGroups

`func (o *DiscoveredAccount) SetOsGroups(v string)`

SetOsGroups sets OsGroups field to given value.

### HasOsGroups

`func (o *DiscoveredAccount) HasOsGroups() bool`

HasOsGroups returns a boolean if a field has been set.

### GetPlatformType

`func (o *DiscoveredAccount) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *DiscoveredAccount) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *DiscoveredAccount) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.


### GetDomain

`func (o *DiscoveredAccount) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DiscoveredAccount) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DiscoveredAccount) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DiscoveredAccount) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLastLogonDateTime

`func (o *DiscoveredAccount) GetLastLogonDateTime() int64`

GetLastLogonDateTime returns the LastLogonDateTime field if non-nil, zero value otherwise.

### GetLastLogonDateTimeOk

`func (o *DiscoveredAccount) GetLastLogonDateTimeOk() (*int64, bool)`

GetLastLogonDateTimeOk returns a tuple with the LastLogonDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogonDateTime

`func (o *DiscoveredAccount) SetLastLogonDateTime(v int64)`

SetLastLogonDateTime sets LastLogonDateTime field to given value.

### HasLastLogonDateTime

`func (o *DiscoveredAccount) HasLastLogonDateTime() bool`

HasLastLogonDateTime returns a boolean if a field has been set.

### GetLastPasswordSetDateTime

`func (o *DiscoveredAccount) GetLastPasswordSetDateTime() int64`

GetLastPasswordSetDateTime returns the LastPasswordSetDateTime field if non-nil, zero value otherwise.

### GetLastPasswordSetDateTimeOk

`func (o *DiscoveredAccount) GetLastPasswordSetDateTimeOk() (*int64, bool)`

GetLastPasswordSetDateTimeOk returns a tuple with the LastPasswordSetDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordSetDateTime

`func (o *DiscoveredAccount) SetLastPasswordSetDateTime(v int64)`

SetLastPasswordSetDateTime sets LastPasswordSetDateTime field to given value.

### HasLastPasswordSetDateTime

`func (o *DiscoveredAccount) HasLastPasswordSetDateTime() bool`

HasLastPasswordSetDateTime returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *DiscoveredAccount) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *DiscoveredAccount) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *DiscoveredAccount) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *DiscoveredAccount) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetOsVersion

`func (o *DiscoveredAccount) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DiscoveredAccount) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DiscoveredAccount) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DiscoveredAccount) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPrivileged

`func (o *DiscoveredAccount) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *DiscoveredAccount) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *DiscoveredAccount) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *DiscoveredAccount) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetPrivilegedCriteria

`func (o *DiscoveredAccount) GetPrivilegedCriteria() string`

GetPrivilegedCriteria returns the PrivilegedCriteria field if non-nil, zero value otherwise.

### GetPrivilegedCriteriaOk

`func (o *DiscoveredAccount) GetPrivilegedCriteriaOk() (*string, bool)`

GetPrivilegedCriteriaOk returns a tuple with the PrivilegedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedCriteria

`func (o *DiscoveredAccount) SetPrivilegedCriteria(v string)`

SetPrivilegedCriteria sets PrivilegedCriteria field to given value.

### HasPrivilegedCriteria

`func (o *DiscoveredAccount) HasPrivilegedCriteria() bool`

HasPrivilegedCriteria returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *DiscoveredAccount) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *DiscoveredAccount) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *DiscoveredAccount) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *DiscoveredAccount) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveredAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveredAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveredAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveredAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPasswordExpirationDateTime

`func (o *DiscoveredAccount) GetPasswordExpirationDateTime() int64`

GetPasswordExpirationDateTime returns the PasswordExpirationDateTime field if non-nil, zero value otherwise.

### GetPasswordExpirationDateTimeOk

`func (o *DiscoveredAccount) GetPasswordExpirationDateTimeOk() (*int64, bool)`

GetPasswordExpirationDateTimeOk returns a tuple with the PasswordExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationDateTime

`func (o *DiscoveredAccount) SetPasswordExpirationDateTime(v int64)`

SetPasswordExpirationDateTime sets PasswordExpirationDateTime field to given value.

### HasPasswordExpirationDateTime

`func (o *DiscoveredAccount) HasPasswordExpirationDateTime() bool`

HasPasswordExpirationDateTime returns a boolean if a field has been set.

### GetOsFamily

`func (o *DiscoveredAccount) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *DiscoveredAccount) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *DiscoveredAccount) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *DiscoveredAccount) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *DiscoveredAccount) GetAdditionalPropertiesField() map[string]interface{}`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *DiscoveredAccount) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *DiscoveredAccount) SetAdditionalPropertiesField(v map[string]interface{})`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *DiscoveredAccount) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *DiscoveredAccount) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *DiscoveredAccount) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *DiscoveredAccount) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *DiscoveredAccount) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetPlatformTypeAccountProperties

`func (o *DiscoveredAccount) GetPlatformTypeAccountProperties() map[string]interface{}`

GetPlatformTypeAccountProperties returns the PlatformTypeAccountProperties field if non-nil, zero value otherwise.

### GetPlatformTypeAccountPropertiesOk

`func (o *DiscoveredAccount) GetPlatformTypeAccountPropertiesOk() (*map[string]interface{}, bool)`

GetPlatformTypeAccountPropertiesOk returns a tuple with the PlatformTypeAccountProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTypeAccountProperties

`func (o *DiscoveredAccount) SetPlatformTypeAccountProperties(v map[string]interface{})`

SetPlatformTypeAccountProperties sets PlatformTypeAccountProperties field to given value.

### HasPlatformTypeAccountProperties

`func (o *DiscoveredAccount) HasPlatformTypeAccountProperties() bool`

HasPlatformTypeAccountProperties returns a boolean if a field has been set.

### GetDependencies

`func (o *DiscoveredAccount) GetDependencies() []DiscoveredDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *DiscoveredAccount) GetDependenciesOk() (*[]DiscoveredDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *DiscoveredAccount) SetDependencies(v []DiscoveredDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *DiscoveredAccount) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


