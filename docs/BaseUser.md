# BaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique id of the user. | [optional] 
**Username** | **string** | The name of the user. | 
**Source** | Pointer to **string** | The source of the user.  Possible values: CyberArk, LDAP. | [optional] 
**UserType** | Pointer to **string** | The type of the user.  Possible types could be any user types according to the license. | [optional] 
**ComponentUser** | Pointer to **bool** | Whether the user is a known component or not.  Possible values: true (if it is component) or false otherwise.    The following user types are considered as components:      CPM      ENE      PVWA      PSM      AppProvider      OPMProvider      PIMProvider      PSMPServer      PSMPADBridge      PSMHTML5Gateway      CIFS      FTP      SFE      DCAInstance      FEWA      SEG | [optional] 
**GroupsMembership** | Pointer to [**[]GroupMembershipDetails**](GroupMembershipDetails.md) | The security attributes and authorizations, contains:  AddSafe  AuditUsers  AddUpdateUsers  ResetUsersPasswords  ActivateUsers  AddNetworkAreas  ManageDirectoryMapping  ManageServerFileCategories  BackupAllSafes  RestoreAllSafes | [optional] 
**UserDN** | Pointer to **string** | The distinguished name of the user. Only applies to LDAP users. | [optional] 
**VaultAuthorization** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** | The Vault Location. | [optional] 
**PersonalDetails** | Pointer to [**BaseUserPersonalDetails**](BaseUserPersonalDetails.md) |  | [optional] 
**EnableUser** | Pointer to **bool** | Whether or not the user is enabled. | [optional] 
**Suspended** | Pointer to **bool** | Whether or not the user is suspended due to maximum violations. | [optional] 

## Methods

### NewBaseUser

`func NewBaseUser(username string, ) *BaseUser`

NewBaseUser instantiates a new BaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseUserWithDefaults

`func NewBaseUserWithDefaults() *BaseUser`

NewBaseUserWithDefaults instantiates a new BaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BaseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *BaseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetSource

`func (o *BaseUser) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BaseUser) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BaseUser) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BaseUser) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUserType

`func (o *BaseUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *BaseUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *BaseUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *BaseUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetComponentUser

`func (o *BaseUser) GetComponentUser() bool`

GetComponentUser returns the ComponentUser field if non-nil, zero value otherwise.

### GetComponentUserOk

`func (o *BaseUser) GetComponentUserOk() (*bool, bool)`

GetComponentUserOk returns a tuple with the ComponentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUser

`func (o *BaseUser) SetComponentUser(v bool)`

SetComponentUser sets ComponentUser field to given value.

### HasComponentUser

`func (o *BaseUser) HasComponentUser() bool`

HasComponentUser returns a boolean if a field has been set.

### GetGroupsMembership

`func (o *BaseUser) GetGroupsMembership() []GroupMembershipDetails`

GetGroupsMembership returns the GroupsMembership field if non-nil, zero value otherwise.

### GetGroupsMembershipOk

`func (o *BaseUser) GetGroupsMembershipOk() (*[]GroupMembershipDetails, bool)`

GetGroupsMembershipOk returns a tuple with the GroupsMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMembership

`func (o *BaseUser) SetGroupsMembership(v []GroupMembershipDetails)`

SetGroupsMembership sets GroupsMembership field to given value.

### HasGroupsMembership

`func (o *BaseUser) HasGroupsMembership() bool`

HasGroupsMembership returns a boolean if a field has been set.

### GetUserDN

`func (o *BaseUser) GetUserDN() string`

GetUserDN returns the UserDN field if non-nil, zero value otherwise.

### GetUserDNOk

`func (o *BaseUser) GetUserDNOk() (*string, bool)`

GetUserDNOk returns a tuple with the UserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDN

`func (o *BaseUser) SetUserDN(v string)`

SetUserDN sets UserDN field to given value.

### HasUserDN

`func (o *BaseUser) HasUserDN() bool`

HasUserDN returns a boolean if a field has been set.

### GetVaultAuthorization

`func (o *BaseUser) GetVaultAuthorization() []string`

GetVaultAuthorization returns the VaultAuthorization field if non-nil, zero value otherwise.

### GetVaultAuthorizationOk

`func (o *BaseUser) GetVaultAuthorizationOk() (*[]string, bool)`

GetVaultAuthorizationOk returns a tuple with the VaultAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthorization

`func (o *BaseUser) SetVaultAuthorization(v []string)`

SetVaultAuthorization sets VaultAuthorization field to given value.

### HasVaultAuthorization

`func (o *BaseUser) HasVaultAuthorization() bool`

HasVaultAuthorization returns a boolean if a field has been set.

### GetLocation

`func (o *BaseUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *BaseUser) GetPersonalDetails() BaseUserPersonalDetails`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *BaseUser) GetPersonalDetailsOk() (*BaseUserPersonalDetails, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *BaseUser) SetPersonalDetails(v BaseUserPersonalDetails)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *BaseUser) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.

### GetEnableUser

`func (o *BaseUser) GetEnableUser() bool`

GetEnableUser returns the EnableUser field if non-nil, zero value otherwise.

### GetEnableUserOk

`func (o *BaseUser) GetEnableUserOk() (*bool, bool)`

GetEnableUserOk returns a tuple with the EnableUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUser

`func (o *BaseUser) SetEnableUser(v bool)`

SetEnableUser sets EnableUser field to given value.

### HasEnableUser

`func (o *BaseUser) HasEnableUser() bool`

HasEnableUser returns a boolean if a field has been set.

### GetSuspended

`func (o *BaseUser) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *BaseUser) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *BaseUser) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *BaseUser) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


