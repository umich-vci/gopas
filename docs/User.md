# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableUser** | Pointer to **bool** | Whether or not the user is enabled. | [optional] 
**ChangePassOnNextLogon** | Pointer to **bool** | Whether or not the user must change their password on the next logon. | [optional] 
**ExpiryDate** | Pointer to **int64** | The date when the user will expire. | [optional] 
**Suspended** | Pointer to **bool** | Whether or not the user is suspended due to maximum violations. | [optional] 
**LastSuccessfulLoginDate** | Pointer to **int64** | Returns the last successful login date of the user. | [optional] 
**UnAuthorizedInterfaces** | Pointer to **[]string** | The CyberArk interfaces that this user is authorized to use.  The possible values depends on the specific user type as defined in the license. | [optional] 
**AuthenticationMethod** | Pointer to **[]string** | The authentication method that the User will use to log on. | [optional] 
**PasswordNeverExpires** | Pointer to **bool** | Whether the User’s password will be retained until he decides to change it. | [optional] 
**DistinguishedName** | Pointer to **string** | The User’s distinguished name.  The usage is for PKI authentication, this will match the certificate Subject Name or DN. | [optional] 
**Description** | Pointer to **string** | Notes and comments. | [optional] 
**BusinessAddress** | Pointer to [**UserBusinessAddress**](UserBusinessAddress.md) |  | [optional] 
**Internet** | Pointer to [**UserInternet**](UserInternet.md) |  | [optional] 
**Phones** | Pointer to [**UserPhone**](UserPhone.md) |  | [optional] 
**PersonalDetails** | Pointer to [**UserPersonalDetails**](UserPersonalDetails.md) |  | [optional] 
**InitialPassword** | Pointer to **string** | The password that the user will use to log on the first time.  This password must meet the password policy requirements. | [optional] 
**Id** | Pointer to **int64** | The unique id of the user. | [optional] 
**Username** | **string** | The name of the user. | 
**Source** | Pointer to **string** | The source of the user.  Possible values: CyberArk, LDAP. | [optional] 
**UserType** | Pointer to **string** | The type of the user.  Possible types could be any user types according to the license. | [optional] 
**ComponentUser** | Pointer to **bool** | Whether the user is a known component or not.  Possible values: true (if it is component) or false otherwise.    The following user types are considered as components:      CPM      ENE      PVWA      PSM      AppProvider      OPMProvider      PIMProvider      PSMPServer      PSMPADBridge      PSMHTML5Gateway      CIFS      FTP      SFE      DCAInstance      FEWA      SEG | [optional] 
**GroupsMembership** | Pointer to [**[]GroupMembershipDetails**](GroupMembershipDetails.md) | The security attributes and authorizations, contains:  AddSafe  AuditUsers  AddUpdateUsers  ResetUsersPasswords  ActivateUsers  AddNetworkAreas  ManageDirectoryMapping  ManageServerFileCategories  BackupAllSafes  RestoreAllSafes | [optional] 
**UserDN** | Pointer to **string** | The distinguished name of the user. Only applies to LDAP users. | [optional] 
**VaultAuthorization** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** | The Vault Location. | [optional] 

## Methods

### NewUser

`func NewUser(username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableUser

`func (o *User) GetEnableUser() bool`

GetEnableUser returns the EnableUser field if non-nil, zero value otherwise.

### GetEnableUserOk

`func (o *User) GetEnableUserOk() (*bool, bool)`

GetEnableUserOk returns a tuple with the EnableUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUser

`func (o *User) SetEnableUser(v bool)`

SetEnableUser sets EnableUser field to given value.

### HasEnableUser

`func (o *User) HasEnableUser() bool`

HasEnableUser returns a boolean if a field has been set.

### GetChangePassOnNextLogon

`func (o *User) GetChangePassOnNextLogon() bool`

GetChangePassOnNextLogon returns the ChangePassOnNextLogon field if non-nil, zero value otherwise.

### GetChangePassOnNextLogonOk

`func (o *User) GetChangePassOnNextLogonOk() (*bool, bool)`

GetChangePassOnNextLogonOk returns a tuple with the ChangePassOnNextLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassOnNextLogon

`func (o *User) SetChangePassOnNextLogon(v bool)`

SetChangePassOnNextLogon sets ChangePassOnNextLogon field to given value.

### HasChangePassOnNextLogon

`func (o *User) HasChangePassOnNextLogon() bool`

HasChangePassOnNextLogon returns a boolean if a field has been set.

### GetExpiryDate

`func (o *User) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *User) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *User) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *User) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetSuspended

`func (o *User) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *User) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *User) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *User) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetLastSuccessfulLoginDate

`func (o *User) GetLastSuccessfulLoginDate() int64`

GetLastSuccessfulLoginDate returns the LastSuccessfulLoginDate field if non-nil, zero value otherwise.

### GetLastSuccessfulLoginDateOk

`func (o *User) GetLastSuccessfulLoginDateOk() (*int64, bool)`

GetLastSuccessfulLoginDateOk returns a tuple with the LastSuccessfulLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulLoginDate

`func (o *User) SetLastSuccessfulLoginDate(v int64)`

SetLastSuccessfulLoginDate sets LastSuccessfulLoginDate field to given value.

### HasLastSuccessfulLoginDate

`func (o *User) HasLastSuccessfulLoginDate() bool`

HasLastSuccessfulLoginDate returns a boolean if a field has been set.

### GetUnAuthorizedInterfaces

`func (o *User) GetUnAuthorizedInterfaces() []string`

GetUnAuthorizedInterfaces returns the UnAuthorizedInterfaces field if non-nil, zero value otherwise.

### GetUnAuthorizedInterfacesOk

`func (o *User) GetUnAuthorizedInterfacesOk() (*[]string, bool)`

GetUnAuthorizedInterfacesOk returns a tuple with the UnAuthorizedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnAuthorizedInterfaces

`func (o *User) SetUnAuthorizedInterfaces(v []string)`

SetUnAuthorizedInterfaces sets UnAuthorizedInterfaces field to given value.

### HasUnAuthorizedInterfaces

`func (o *User) HasUnAuthorizedInterfaces() bool`

HasUnAuthorizedInterfaces returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *User) GetAuthenticationMethod() []string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *User) GetAuthenticationMethodOk() (*[]string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *User) SetAuthenticationMethod(v []string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *User) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetPasswordNeverExpires

`func (o *User) GetPasswordNeverExpires() bool`

GetPasswordNeverExpires returns the PasswordNeverExpires field if non-nil, zero value otherwise.

### GetPasswordNeverExpiresOk

`func (o *User) GetPasswordNeverExpiresOk() (*bool, bool)`

GetPasswordNeverExpiresOk returns a tuple with the PasswordNeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordNeverExpires

`func (o *User) SetPasswordNeverExpires(v bool)`

SetPasswordNeverExpires sets PasswordNeverExpires field to given value.

### HasPasswordNeverExpires

`func (o *User) HasPasswordNeverExpires() bool`

HasPasswordNeverExpires returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *User) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *User) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *User) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *User) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBusinessAddress

`func (o *User) GetBusinessAddress() UserBusinessAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *User) GetBusinessAddressOk() (*UserBusinessAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *User) SetBusinessAddress(v UserBusinessAddress)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *User) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### GetInternet

`func (o *User) GetInternet() UserInternet`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *User) GetInternetOk() (*UserInternet, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *User) SetInternet(v UserInternet)`

SetInternet sets Internet field to given value.

### HasInternet

`func (o *User) HasInternet() bool`

HasInternet returns a boolean if a field has been set.

### GetPhones

`func (o *User) GetPhones() UserPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *User) GetPhonesOk() (*UserPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *User) SetPhones(v UserPhone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *User) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetPersonalDetails

`func (o *User) GetPersonalDetails() UserPersonalDetails`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *User) GetPersonalDetailsOk() (*UserPersonalDetails, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *User) SetPersonalDetails(v UserPersonalDetails)`

SetPersonalDetails sets PersonalDetails field to given value.

### HasPersonalDetails

`func (o *User) HasPersonalDetails() bool`

HasPersonalDetails returns a boolean if a field has been set.

### GetInitialPassword

`func (o *User) GetInitialPassword() string`

GetInitialPassword returns the InitialPassword field if non-nil, zero value otherwise.

### GetInitialPasswordOk

`func (o *User) GetInitialPasswordOk() (*string, bool)`

GetInitialPasswordOk returns a tuple with the InitialPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPassword

`func (o *User) SetInitialPassword(v string)`

SetInitialPassword sets InitialPassword field to given value.

### HasInitialPassword

`func (o *User) HasInitialPassword() bool`

HasInitialPassword returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetSource

`func (o *User) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *User) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *User) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *User) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUserType

`func (o *User) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetComponentUser

`func (o *User) GetComponentUser() bool`

GetComponentUser returns the ComponentUser field if non-nil, zero value otherwise.

### GetComponentUserOk

`func (o *User) GetComponentUserOk() (*bool, bool)`

GetComponentUserOk returns a tuple with the ComponentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUser

`func (o *User) SetComponentUser(v bool)`

SetComponentUser sets ComponentUser field to given value.

### HasComponentUser

`func (o *User) HasComponentUser() bool`

HasComponentUser returns a boolean if a field has been set.

### GetGroupsMembership

`func (o *User) GetGroupsMembership() []GroupMembershipDetails`

GetGroupsMembership returns the GroupsMembership field if non-nil, zero value otherwise.

### GetGroupsMembershipOk

`func (o *User) GetGroupsMembershipOk() (*[]GroupMembershipDetails, bool)`

GetGroupsMembershipOk returns a tuple with the GroupsMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsMembership

`func (o *User) SetGroupsMembership(v []GroupMembershipDetails)`

SetGroupsMembership sets GroupsMembership field to given value.

### HasGroupsMembership

`func (o *User) HasGroupsMembership() bool`

HasGroupsMembership returns a boolean if a field has been set.

### GetUserDN

`func (o *User) GetUserDN() string`

GetUserDN returns the UserDN field if non-nil, zero value otherwise.

### GetUserDNOk

`func (o *User) GetUserDNOk() (*string, bool)`

GetUserDNOk returns a tuple with the UserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDN

`func (o *User) SetUserDN(v string)`

SetUserDN sets UserDN field to given value.

### HasUserDN

`func (o *User) HasUserDN() bool`

HasUserDN returns a boolean if a field has been set.

### GetVaultAuthorization

`func (o *User) GetVaultAuthorization() []string`

GetVaultAuthorization returns the VaultAuthorization field if non-nil, zero value otherwise.

### GetVaultAuthorizationOk

`func (o *User) GetVaultAuthorizationOk() (*[]string, bool)`

GetVaultAuthorizationOk returns a tuple with the VaultAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthorization

`func (o *User) SetVaultAuthorization(v []string)`

SetVaultAuthorization sets VaultAuthorization field to given value.

### HasVaultAuthorization

`func (o *User) HasVaultAuthorization() bool`

HasVaultAuthorization returns a boolean if a field has been set.

### GetLocation

`func (o *User) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *User) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *User) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *User) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


