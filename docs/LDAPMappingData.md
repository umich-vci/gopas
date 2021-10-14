# LDAPMappingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LDAPBranch** | **string** | The LDAP branch used for the directory queries. | 
**VaultGroups** | Pointer to **[]string** | A list of Vault groups that the user will be added to once mapped. | [optional] 
**MappingAuthorizations** | Pointer to **[]string** | The security attributes and authorizations that are applied when LDAP user accounts in the current mapping are created in the Vault.  Here is the list with the possible authorizations:  AddSafes,   AuditUsers,   AddUpdateUsers,   ResetUsersPasswords,   ActivateUsers,  AddNetworkAreas,  ManageServerFileCategories,  BackupAllSafes,  RestoreAllSafes. | [optional] 
**Location** | Pointer to **string** | The specific Vault location path that users in the current mapping are added under.  This value cannot be updated. | [optional] 
**AuthenticationMethod** | Pointer to **[]string** | Read-only - The authentication method that the users belong to this map will use to log onto the Vault. | [optional] 
**UserType** | Pointer to **string** | Read-only - The interfaces that users in the current mapping can use to access the Vault.  This will only be available to users with Add/Update users permissions. | [optional] 
**DisableUser** | Pointer to **bool** | Read-only - Whether users in the current mapping are temporarily inaccessible. | [optional] 
**UserActivityLogPeriod** | Pointer to **int32** | The number of days that activity records are stored for users in the current mapping before they can be deleted. | [optional] 
**UserExpiration** | Pointer to **int64** | Read-only - The date after which user accounts in the current mapping are no longer accessible, in Unix time (\&quot;0\&quot; if never). | [optional] 
**LogonFromHour** | Pointer to **int32** | Read-only - Whether users in the current mapping can log on to the Vault only starting from specific hours. | [optional] 
**LogonToHour** | Pointer to **int32** | Read-only - Whether users in the current mapping can log on to the Vault only up to specific hours. | [optional] 
**MappingID** | Pointer to **int64** | Read-only - Unique ID of the directory mapping | [optional] 
**DirectoryMappingOrder** | Pointer to **int32** | The order in which the maps are matched with users and groups   from the External Directory when they are created in the Vault.  This field is read only. | [optional] 
**MappingName** | **string** | The name of the role.  For example: Vault Admins, Safe Managers. | 
**LDAPQuery** | Pointer to **string** | LDAP users that the filter applies to them are assigned to relevant roles in the system. | [optional] 
**DomainGroups** | Pointer to **[]string** | Users who belong to these LDAP groups are assigned to relevant roles in the system.  Note: This field is required if the VaultGroups field was sent. | [optional] 

## Methods

### NewLDAPMappingData

`func NewLDAPMappingData(lDAPBranch string, mappingName string, ) *LDAPMappingData`

NewLDAPMappingData instantiates a new LDAPMappingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPMappingDataWithDefaults

`func NewLDAPMappingDataWithDefaults() *LDAPMappingData`

NewLDAPMappingDataWithDefaults instantiates a new LDAPMappingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLDAPBranch

`func (o *LDAPMappingData) GetLDAPBranch() string`

GetLDAPBranch returns the LDAPBranch field if non-nil, zero value otherwise.

### GetLDAPBranchOk

`func (o *LDAPMappingData) GetLDAPBranchOk() (*string, bool)`

GetLDAPBranchOk returns a tuple with the LDAPBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPBranch

`func (o *LDAPMappingData) SetLDAPBranch(v string)`

SetLDAPBranch sets LDAPBranch field to given value.


### GetVaultGroups

`func (o *LDAPMappingData) GetVaultGroups() []string`

GetVaultGroups returns the VaultGroups field if non-nil, zero value otherwise.

### GetVaultGroupsOk

`func (o *LDAPMappingData) GetVaultGroupsOk() (*[]string, bool)`

GetVaultGroupsOk returns a tuple with the VaultGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultGroups

`func (o *LDAPMappingData) SetVaultGroups(v []string)`

SetVaultGroups sets VaultGroups field to given value.

### HasVaultGroups

`func (o *LDAPMappingData) HasVaultGroups() bool`

HasVaultGroups returns a boolean if a field has been set.

### GetMappingAuthorizations

`func (o *LDAPMappingData) GetMappingAuthorizations() []string`

GetMappingAuthorizations returns the MappingAuthorizations field if non-nil, zero value otherwise.

### GetMappingAuthorizationsOk

`func (o *LDAPMappingData) GetMappingAuthorizationsOk() (*[]string, bool)`

GetMappingAuthorizationsOk returns a tuple with the MappingAuthorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingAuthorizations

`func (o *LDAPMappingData) SetMappingAuthorizations(v []string)`

SetMappingAuthorizations sets MappingAuthorizations field to given value.

### HasMappingAuthorizations

`func (o *LDAPMappingData) HasMappingAuthorizations() bool`

HasMappingAuthorizations returns a boolean if a field has been set.

### GetLocation

`func (o *LDAPMappingData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LDAPMappingData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LDAPMappingData) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LDAPMappingData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *LDAPMappingData) GetAuthenticationMethod() []string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *LDAPMappingData) GetAuthenticationMethodOk() (*[]string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *LDAPMappingData) SetAuthenticationMethod(v []string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *LDAPMappingData) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetUserType

`func (o *LDAPMappingData) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *LDAPMappingData) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *LDAPMappingData) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *LDAPMappingData) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetDisableUser

`func (o *LDAPMappingData) GetDisableUser() bool`

GetDisableUser returns the DisableUser field if non-nil, zero value otherwise.

### GetDisableUserOk

`func (o *LDAPMappingData) GetDisableUserOk() (*bool, bool)`

GetDisableUserOk returns a tuple with the DisableUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUser

`func (o *LDAPMappingData) SetDisableUser(v bool)`

SetDisableUser sets DisableUser field to given value.

### HasDisableUser

`func (o *LDAPMappingData) HasDisableUser() bool`

HasDisableUser returns a boolean if a field has been set.

### GetUserActivityLogPeriod

`func (o *LDAPMappingData) GetUserActivityLogPeriod() int32`

GetUserActivityLogPeriod returns the UserActivityLogPeriod field if non-nil, zero value otherwise.

### GetUserActivityLogPeriodOk

`func (o *LDAPMappingData) GetUserActivityLogPeriodOk() (*int32, bool)`

GetUserActivityLogPeriodOk returns a tuple with the UserActivityLogPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActivityLogPeriod

`func (o *LDAPMappingData) SetUserActivityLogPeriod(v int32)`

SetUserActivityLogPeriod sets UserActivityLogPeriod field to given value.

### HasUserActivityLogPeriod

`func (o *LDAPMappingData) HasUserActivityLogPeriod() bool`

HasUserActivityLogPeriod returns a boolean if a field has been set.

### GetUserExpiration

`func (o *LDAPMappingData) GetUserExpiration() int64`

GetUserExpiration returns the UserExpiration field if non-nil, zero value otherwise.

### GetUserExpirationOk

`func (o *LDAPMappingData) GetUserExpirationOk() (*int64, bool)`

GetUserExpirationOk returns a tuple with the UserExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExpiration

`func (o *LDAPMappingData) SetUserExpiration(v int64)`

SetUserExpiration sets UserExpiration field to given value.

### HasUserExpiration

`func (o *LDAPMappingData) HasUserExpiration() bool`

HasUserExpiration returns a boolean if a field has been set.

### GetLogonFromHour

`func (o *LDAPMappingData) GetLogonFromHour() int32`

GetLogonFromHour returns the LogonFromHour field if non-nil, zero value otherwise.

### GetLogonFromHourOk

`func (o *LDAPMappingData) GetLogonFromHourOk() (*int32, bool)`

GetLogonFromHourOk returns a tuple with the LogonFromHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonFromHour

`func (o *LDAPMappingData) SetLogonFromHour(v int32)`

SetLogonFromHour sets LogonFromHour field to given value.

### HasLogonFromHour

`func (o *LDAPMappingData) HasLogonFromHour() bool`

HasLogonFromHour returns a boolean if a field has been set.

### GetLogonToHour

`func (o *LDAPMappingData) GetLogonToHour() int32`

GetLogonToHour returns the LogonToHour field if non-nil, zero value otherwise.

### GetLogonToHourOk

`func (o *LDAPMappingData) GetLogonToHourOk() (*int32, bool)`

GetLogonToHourOk returns a tuple with the LogonToHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonToHour

`func (o *LDAPMappingData) SetLogonToHour(v int32)`

SetLogonToHour sets LogonToHour field to given value.

### HasLogonToHour

`func (o *LDAPMappingData) HasLogonToHour() bool`

HasLogonToHour returns a boolean if a field has been set.

### GetMappingID

`func (o *LDAPMappingData) GetMappingID() int64`

GetMappingID returns the MappingID field if non-nil, zero value otherwise.

### GetMappingIDOk

`func (o *LDAPMappingData) GetMappingIDOk() (*int64, bool)`

GetMappingIDOk returns a tuple with the MappingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingID

`func (o *LDAPMappingData) SetMappingID(v int64)`

SetMappingID sets MappingID field to given value.

### HasMappingID

`func (o *LDAPMappingData) HasMappingID() bool`

HasMappingID returns a boolean if a field has been set.

### GetDirectoryMappingOrder

`func (o *LDAPMappingData) GetDirectoryMappingOrder() int32`

GetDirectoryMappingOrder returns the DirectoryMappingOrder field if non-nil, zero value otherwise.

### GetDirectoryMappingOrderOk

`func (o *LDAPMappingData) GetDirectoryMappingOrderOk() (*int32, bool)`

GetDirectoryMappingOrderOk returns a tuple with the DirectoryMappingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryMappingOrder

`func (o *LDAPMappingData) SetDirectoryMappingOrder(v int32)`

SetDirectoryMappingOrder sets DirectoryMappingOrder field to given value.

### HasDirectoryMappingOrder

`func (o *LDAPMappingData) HasDirectoryMappingOrder() bool`

HasDirectoryMappingOrder returns a boolean if a field has been set.

### GetMappingName

`func (o *LDAPMappingData) GetMappingName() string`

GetMappingName returns the MappingName field if non-nil, zero value otherwise.

### GetMappingNameOk

`func (o *LDAPMappingData) GetMappingNameOk() (*string, bool)`

GetMappingNameOk returns a tuple with the MappingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingName

`func (o *LDAPMappingData) SetMappingName(v string)`

SetMappingName sets MappingName field to given value.


### GetLDAPQuery

`func (o *LDAPMappingData) GetLDAPQuery() string`

GetLDAPQuery returns the LDAPQuery field if non-nil, zero value otherwise.

### GetLDAPQueryOk

`func (o *LDAPMappingData) GetLDAPQueryOk() (*string, bool)`

GetLDAPQueryOk returns a tuple with the LDAPQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPQuery

`func (o *LDAPMappingData) SetLDAPQuery(v string)`

SetLDAPQuery sets LDAPQuery field to given value.

### HasLDAPQuery

`func (o *LDAPMappingData) HasLDAPQuery() bool`

HasLDAPQuery returns a boolean if a field has been set.

### GetDomainGroups

`func (o *LDAPMappingData) GetDomainGroups() []string`

GetDomainGroups returns the DomainGroups field if non-nil, zero value otherwise.

### GetDomainGroupsOk

`func (o *LDAPMappingData) GetDomainGroupsOk() (*[]string, bool)`

GetDomainGroupsOk returns a tuple with the DomainGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroups

`func (o *LDAPMappingData) SetDomainGroups(v []string)`

SetDomainGroups sets DomainGroups field to given value.

### HasDomainGroups

`func (o *LDAPMappingData) HasDomainGroups() bool`

HasDomainGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


