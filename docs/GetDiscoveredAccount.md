# GetDiscoveredAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the discovered account | [optional] 
**Name** | Pointer to **string** | The name of the account in the PasswordManager_Pending safe | [optional] 
**UserName** | Pointer to **string** | The name of the discovered account user | [optional] 
**Address** | Pointer to **string** | The name or address of the machine\\domain where the account was discovered | [optional] 
**DiscoveryDateTime** | Pointer to **int64** | The date the account was discovered | [optional] 
**AccountEnabled** | Pointer to **bool** | The state of the account, defined in the discovery source.  Note: The state of domain accounts is based on the Active Directory. The state of local accounts is based on the local machine.  If this parameter is not set, it is considered null | [optional] 
**OsGroups** | Pointer to **string** | The group names that the account belongs to, such as Administrators or Operators | [optional] 
**PlatformType** | Pointer to **string** | The platform where the discovered account is located | [optional] 
**Domain** | Pointer to **string** | The domain of the account | [optional] 
**LastLogonDateTime** | Pointer to **int64** | The date this account was last logged into, defined in the discovery source | [optional] 
**LastPasswordSetDateTime** | Pointer to **int64** | The date this password was last set, defined in the discovery source | [optional] 
**PasswordNeverExpires** | Pointer to **bool** | Whether or not this password expires, defined in the discovery source.  If this parameter is not set, it is considered null | [optional] 
**OsVersion** | Pointer to **string** | The version of the OS where the account was discovered | [optional] 
**Privileged** | Pointer to **bool** | Whether the discovered account is privileged or non-privileged.  If this parameter is not set, it is considered null | [optional] 
**PrivilegedCriteria** | Pointer to **string** | The criteria that determines whether or not the discovered account is privileged. For example, the user or group name | [optional] 
**UserDisplayName** | Pointer to **string** | The user&#39;s display name | [optional] 
**Description** | Pointer to **string** | A description of the account, defined in the discovery source | [optional] 
**PasswordExpirationDateTime** | Pointer to **int64** | The expiration date of the account, defined in the discovery source | [optional] 
**OsFamily** | Pointer to **string** | The type of machine where the account was discovered.  If this parameter is not set, it is considered null and will not be returned in the result | [optional] 
**OrganizationalUnit** | Pointer to **string** | The organizational unit where the account is defined | [optional] 
**AdditionalPropertiesField** | Pointer to **map[string]string** | List of name&#x3D;value pairs for additional properties of the account.  The list of properties is valid file properties in the Vault | [optional] 
**PlatformTypeAccountProperties** | Pointer to **map[string]string** | The object that contains the key-value pairs to associate with the account, as defined by the account platform type schema.   Only properties that appear in the platform type schema are allowed | [optional] 
**NumberOfDependencies** | Pointer to **int32** | The number of dependencies for the discovered account | [optional] 
**Dependencies** | Pointer to [**[]GetDiscoveredDependency**](GetDiscoveredDependency.md) | The list of dependency details for the discovered account | [optional] 

## Methods

### NewGetDiscoveredAccount

`func NewGetDiscoveredAccount() *GetDiscoveredAccount`

NewGetDiscoveredAccount instantiates a new GetDiscoveredAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveredAccountWithDefaults

`func NewGetDiscoveredAccountWithDefaults() *GetDiscoveredAccount`

NewGetDiscoveredAccountWithDefaults instantiates a new GetDiscoveredAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDiscoveredAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDiscoveredAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDiscoveredAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDiscoveredAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoveredAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveredAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveredAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveredAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserName

`func (o *GetDiscoveredAccount) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetDiscoveredAccount) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetDiscoveredAccount) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetDiscoveredAccount) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAddress

`func (o *GetDiscoveredAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDiscoveredAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDiscoveredAccount) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDiscoveredAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDiscoveryDateTime

`func (o *GetDiscoveredAccount) GetDiscoveryDateTime() int64`

GetDiscoveryDateTime returns the DiscoveryDateTime field if non-nil, zero value otherwise.

### GetDiscoveryDateTimeOk

`func (o *GetDiscoveredAccount) GetDiscoveryDateTimeOk() (*int64, bool)`

GetDiscoveryDateTimeOk returns a tuple with the DiscoveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryDateTime

`func (o *GetDiscoveredAccount) SetDiscoveryDateTime(v int64)`

SetDiscoveryDateTime sets DiscoveryDateTime field to given value.

### HasDiscoveryDateTime

`func (o *GetDiscoveredAccount) HasDiscoveryDateTime() bool`

HasDiscoveryDateTime returns a boolean if a field has been set.

### GetAccountEnabled

`func (o *GetDiscoveredAccount) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *GetDiscoveredAccount) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *GetDiscoveredAccount) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *GetDiscoveredAccount) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### GetOsGroups

`func (o *GetDiscoveredAccount) GetOsGroups() string`

GetOsGroups returns the OsGroups field if non-nil, zero value otherwise.

### GetOsGroupsOk

`func (o *GetDiscoveredAccount) GetOsGroupsOk() (*string, bool)`

GetOsGroupsOk returns a tuple with the OsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGroups

`func (o *GetDiscoveredAccount) SetOsGroups(v string)`

SetOsGroups sets OsGroups field to given value.

### HasOsGroups

`func (o *GetDiscoveredAccount) HasOsGroups() bool`

HasOsGroups returns a boolean if a field has been set.

### GetPlatformType

`func (o *GetDiscoveredAccount) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *GetDiscoveredAccount) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *GetDiscoveredAccount) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *GetDiscoveredAccount) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetDomain

`func (o *GetDiscoveredAccount) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDiscoveredAccount) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDiscoveredAccount) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetDiscoveredAccount) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLastLogonDateTime

`func (o *GetDiscoveredAccount) GetLastLogonDateTime() int64`

GetLastLogonDateTime returns the LastLogonDateTime field if non-nil, zero value otherwise.

### GetLastLogonDateTimeOk

`func (o *GetDiscoveredAccount) GetLastLogonDateTimeOk() (*int64, bool)`

GetLastLogonDateTimeOk returns a tuple with the LastLogonDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogonDateTime

`func (o *GetDiscoveredAccount) SetLastLogonDateTime(v int64)`

SetLastLogonDateTime sets LastLogonDateTime field to given value.

### HasLastLogonDateTime

`func (o *GetDiscoveredAccount) HasLastLogonDateTime() bool`

HasLastLogonDateTime returns a boolean if a field has been set.

### GetLastPasswordSetDateTime

`func (o *GetDiscoveredAccount) GetLastPasswordSetDateTime() int64`

GetLastPasswordSetDateTime returns the LastPasswordSetDateTime field if non-nil, zero value otherwise.

### GetLastPasswordSetDateTimeOk

`func (o *GetDiscoveredAccount) GetLastPasswordSetDateTimeOk() (*int64, bool)`

GetLastPasswordSetDateTimeOk returns a tuple with the LastPasswordSetDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordSetDateTime

`func (o *GetDiscoveredAccount) SetLastPasswordSetDateTime(v int64)`

SetLastPasswordSetDateTime sets LastPasswordSetDateTime field to given value.

### HasLastPasswordSetDateTime

`func (o *GetDiscoveredAccount) HasLastPasswordSetDateTime() bool`

HasLastPasswordSetDateTime returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *GetDiscoveredAccount) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *GetDiscoveredAccount) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *GetDiscoveredAccount) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *GetDiscoveredAccount) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetDiscoveredAccount) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetDiscoveredAccount) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetDiscoveredAccount) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetDiscoveredAccount) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPrivileged

`func (o *GetDiscoveredAccount) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *GetDiscoveredAccount) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *GetDiscoveredAccount) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *GetDiscoveredAccount) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetPrivilegedCriteria

`func (o *GetDiscoveredAccount) GetPrivilegedCriteria() string`

GetPrivilegedCriteria returns the PrivilegedCriteria field if non-nil, zero value otherwise.

### GetPrivilegedCriteriaOk

`func (o *GetDiscoveredAccount) GetPrivilegedCriteriaOk() (*string, bool)`

GetPrivilegedCriteriaOk returns a tuple with the PrivilegedCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedCriteria

`func (o *GetDiscoveredAccount) SetPrivilegedCriteria(v string)`

SetPrivilegedCriteria sets PrivilegedCriteria field to given value.

### HasPrivilegedCriteria

`func (o *GetDiscoveredAccount) HasPrivilegedCriteria() bool`

HasPrivilegedCriteria returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *GetDiscoveredAccount) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *GetDiscoveredAccount) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *GetDiscoveredAccount) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *GetDiscoveredAccount) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetDiscoveredAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDiscoveredAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDiscoveredAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDiscoveredAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPasswordExpirationDateTime

`func (o *GetDiscoveredAccount) GetPasswordExpirationDateTime() int64`

GetPasswordExpirationDateTime returns the PasswordExpirationDateTime field if non-nil, zero value otherwise.

### GetPasswordExpirationDateTimeOk

`func (o *GetDiscoveredAccount) GetPasswordExpirationDateTimeOk() (*int64, bool)`

GetPasswordExpirationDateTimeOk returns a tuple with the PasswordExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationDateTime

`func (o *GetDiscoveredAccount) SetPasswordExpirationDateTime(v int64)`

SetPasswordExpirationDateTime sets PasswordExpirationDateTime field to given value.

### HasPasswordExpirationDateTime

`func (o *GetDiscoveredAccount) HasPasswordExpirationDateTime() bool`

HasPasswordExpirationDateTime returns a boolean if a field has been set.

### GetOsFamily

`func (o *GetDiscoveredAccount) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *GetDiscoveredAccount) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *GetDiscoveredAccount) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *GetDiscoveredAccount) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *GetDiscoveredAccount) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *GetDiscoveredAccount) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *GetDiscoveredAccount) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *GetDiscoveredAccount) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *GetDiscoveredAccount) GetAdditionalPropertiesField() map[string]string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *GetDiscoveredAccount) GetAdditionalPropertiesFieldOk() (*map[string]string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *GetDiscoveredAccount) SetAdditionalPropertiesField(v map[string]string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *GetDiscoveredAccount) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetPlatformTypeAccountProperties

`func (o *GetDiscoveredAccount) GetPlatformTypeAccountProperties() map[string]string`

GetPlatformTypeAccountProperties returns the PlatformTypeAccountProperties field if non-nil, zero value otherwise.

### GetPlatformTypeAccountPropertiesOk

`func (o *GetDiscoveredAccount) GetPlatformTypeAccountPropertiesOk() (*map[string]string, bool)`

GetPlatformTypeAccountPropertiesOk returns a tuple with the PlatformTypeAccountProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTypeAccountProperties

`func (o *GetDiscoveredAccount) SetPlatformTypeAccountProperties(v map[string]string)`

SetPlatformTypeAccountProperties sets PlatformTypeAccountProperties field to given value.

### HasPlatformTypeAccountProperties

`func (o *GetDiscoveredAccount) HasPlatformTypeAccountProperties() bool`

HasPlatformTypeAccountProperties returns a boolean if a field has been set.

### GetNumberOfDependencies

`func (o *GetDiscoveredAccount) GetNumberOfDependencies() int32`

GetNumberOfDependencies returns the NumberOfDependencies field if non-nil, zero value otherwise.

### GetNumberOfDependenciesOk

`func (o *GetDiscoveredAccount) GetNumberOfDependenciesOk() (*int32, bool)`

GetNumberOfDependenciesOk returns a tuple with the NumberOfDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDependencies

`func (o *GetDiscoveredAccount) SetNumberOfDependencies(v int32)`

SetNumberOfDependencies sets NumberOfDependencies field to given value.

### HasNumberOfDependencies

`func (o *GetDiscoveredAccount) HasNumberOfDependencies() bool`

HasNumberOfDependencies returns a boolean if a field has been set.

### GetDependencies

`func (o *GetDiscoveredAccount) GetDependencies() []GetDiscoveredDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *GetDiscoveredAccount) GetDependenciesOk() (*[]GetDiscoveredDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *GetDiscoveredAccount) SetDependencies(v []GetDiscoveredDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *GetDiscoveredAccount) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


