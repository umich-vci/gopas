# UpdateSafeMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SafeUrlId** | Pointer to **string** | The name of the Safe to be used when calling Safe APIs. | [optional] 
**SafeName** | Pointer to **string** | The unique name of the Safe. | [optional] 
**SafeNumber** | Pointer to **int64** | The unique numerical ID of the Safe. | [optional] 
**MemberId** | Pointer to **int64** | The Vault user ID, Domain user ID, or group ID of the Safe member. | [optional] 
**MemberName** | Pointer to **string** | The Vault user name, Domain user name or group name of the Safe member. | [optional] 
**MemberType** | Pointer to **string** | The member type.  Valid values: user\\group | [optional] 
**MembershipExpirationDate** | Pointer to **int64** | The member&#39;s expiration date for this Safe.  For members that do not have an expiration date, this value will be null. | [optional] 
**IsExpiredMembershipEnable** | Pointer to **bool** | Whether or not the membership for the Safe is expired.For expired members, the value will be True. | [optional] 
**IsPredefinedUser** | Pointer to **bool** | Whether the member is a predefined user or group of the Vault. | [optional] 
**IsReadOnly** | Pointer to **bool** | Whether or not the current user can update the permissions of the member | [optional] 
**Permissions** | Pointer to **map[string]bool** | &lt;p&gt;The permissions that the user or group has in this Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;useAccounts&lt;/B&gt; - Use accounts but not view passwords.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;retrieveAccounts&lt;/B&gt; - Retrieve and view accounts in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;listAccounts&lt;/B&gt; - View Accounts list.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;addAccounts&lt;/B&gt; - Add accounts in the Safe. Users who have this permission automatically have UpdateAccountProperties as well.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;updateAccountContent&lt;/B&gt; - Update existing account content.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;updateAccountProperties&lt;/B&gt; - Update existing account properties.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;initiateCPMAccountManagementOperations&lt;/B&gt; - Initiate password management operations through CPM, such as changing, verifying, and reconciling passwords. When this parameter is set to False, the SpecifyNextAccountContent parameter is also automatically set to False.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;specifyNextAccountContent&lt;/B&gt; - Specify the password that is used when the CPM changes the password value. This parameter can only be specified when the InitiateCPMAccountManagementOperations parameter is set to True. When InitiateCPMAccountManagementOperations is set to False this parameter is automatically set to False.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;renameAccounts&lt;/B&gt; - Rename existing accounts in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;deleteAccounts&lt;/B&gt; - Delete existing passwords in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;unlockAccounts&lt;/B&gt; - Unlock accounts that are locked by other users.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;manageSafe&lt;/B&gt; - Perform administrative tasksin the Safe, including: Update Safe properties, Recover the Safe, Delete the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;manageSafeMembers&lt;/B&gt; - Add and remove Safe members, and update their authorizations in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;backupSafe&lt;/B&gt; - Create a backup of a Safe and its contents, and store in another location.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;viewAuditLog&lt;/B&gt; - View account and user activity in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;viewSafeMembers&lt;/B&gt; - View Safe member&#39;s permissions.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;requestsAuthorizationLevel1&lt;/B&gt; - Request Authorization Level 1.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;requestsAuthorizationLevel2&lt;/B&gt; - Request Authorization Level 2.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;accessWithoutConfirmation&lt;/B&gt; - Access the Safe without confirmation from authorized users. This overrides the Safe properties that specify that Safe members require confirmation to access the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;createFolders&lt;/B&gt; - Create folders in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;deleteFolders&lt;/B&gt; - Delete folders from the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;moveAccountsAndFolders&lt;/B&gt; - Move accounts and folders in the Safe to different folders and subfolders.&lt;/p&gt; | [optional] 

## Methods

### NewUpdateSafeMemberResponse

`func NewUpdateSafeMemberResponse() *UpdateSafeMemberResponse`

NewUpdateSafeMemberResponse instantiates a new UpdateSafeMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSafeMemberResponseWithDefaults

`func NewUpdateSafeMemberResponseWithDefaults() *UpdateSafeMemberResponse`

NewUpdateSafeMemberResponseWithDefaults instantiates a new UpdateSafeMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafeUrlId

`func (o *UpdateSafeMemberResponse) GetSafeUrlId() string`

GetSafeUrlId returns the SafeUrlId field if non-nil, zero value otherwise.

### GetSafeUrlIdOk

`func (o *UpdateSafeMemberResponse) GetSafeUrlIdOk() (*string, bool)`

GetSafeUrlIdOk returns a tuple with the SafeUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeUrlId

`func (o *UpdateSafeMemberResponse) SetSafeUrlId(v string)`

SetSafeUrlId sets SafeUrlId field to given value.

### HasSafeUrlId

`func (o *UpdateSafeMemberResponse) HasSafeUrlId() bool`

HasSafeUrlId returns a boolean if a field has been set.

### GetSafeName

`func (o *UpdateSafeMemberResponse) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *UpdateSafeMemberResponse) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *UpdateSafeMemberResponse) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *UpdateSafeMemberResponse) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetSafeNumber

`func (o *UpdateSafeMemberResponse) GetSafeNumber() int64`

GetSafeNumber returns the SafeNumber field if non-nil, zero value otherwise.

### GetSafeNumberOk

`func (o *UpdateSafeMemberResponse) GetSafeNumberOk() (*int64, bool)`

GetSafeNumberOk returns a tuple with the SafeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeNumber

`func (o *UpdateSafeMemberResponse) SetSafeNumber(v int64)`

SetSafeNumber sets SafeNumber field to given value.

### HasSafeNumber

`func (o *UpdateSafeMemberResponse) HasSafeNumber() bool`

HasSafeNumber returns a boolean if a field has been set.

### GetMemberId

`func (o *UpdateSafeMemberResponse) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UpdateSafeMemberResponse) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UpdateSafeMemberResponse) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UpdateSafeMemberResponse) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMemberName

`func (o *UpdateSafeMemberResponse) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *UpdateSafeMemberResponse) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *UpdateSafeMemberResponse) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.

### HasMemberName

`func (o *UpdateSafeMemberResponse) HasMemberName() bool`

HasMemberName returns a boolean if a field has been set.

### GetMemberType

`func (o *UpdateSafeMemberResponse) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *UpdateSafeMemberResponse) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *UpdateSafeMemberResponse) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *UpdateSafeMemberResponse) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetMembershipExpirationDate

`func (o *UpdateSafeMemberResponse) GetMembershipExpirationDate() int64`

GetMembershipExpirationDate returns the MembershipExpirationDate field if non-nil, zero value otherwise.

### GetMembershipExpirationDateOk

`func (o *UpdateSafeMemberResponse) GetMembershipExpirationDateOk() (*int64, bool)`

GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipExpirationDate

`func (o *UpdateSafeMemberResponse) SetMembershipExpirationDate(v int64)`

SetMembershipExpirationDate sets MembershipExpirationDate field to given value.

### HasMembershipExpirationDate

`func (o *UpdateSafeMemberResponse) HasMembershipExpirationDate() bool`

HasMembershipExpirationDate returns a boolean if a field has been set.

### GetIsExpiredMembershipEnable

`func (o *UpdateSafeMemberResponse) GetIsExpiredMembershipEnable() bool`

GetIsExpiredMembershipEnable returns the IsExpiredMembershipEnable field if non-nil, zero value otherwise.

### GetIsExpiredMembershipEnableOk

`func (o *UpdateSafeMemberResponse) GetIsExpiredMembershipEnableOk() (*bool, bool)`

GetIsExpiredMembershipEnableOk returns a tuple with the IsExpiredMembershipEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpiredMembershipEnable

`func (o *UpdateSafeMemberResponse) SetIsExpiredMembershipEnable(v bool)`

SetIsExpiredMembershipEnable sets IsExpiredMembershipEnable field to given value.

### HasIsExpiredMembershipEnable

`func (o *UpdateSafeMemberResponse) HasIsExpiredMembershipEnable() bool`

HasIsExpiredMembershipEnable returns a boolean if a field has been set.

### GetIsPredefinedUser

`func (o *UpdateSafeMemberResponse) GetIsPredefinedUser() bool`

GetIsPredefinedUser returns the IsPredefinedUser field if non-nil, zero value otherwise.

### GetIsPredefinedUserOk

`func (o *UpdateSafeMemberResponse) GetIsPredefinedUserOk() (*bool, bool)`

GetIsPredefinedUserOk returns a tuple with the IsPredefinedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPredefinedUser

`func (o *UpdateSafeMemberResponse) SetIsPredefinedUser(v bool)`

SetIsPredefinedUser sets IsPredefinedUser field to given value.

### HasIsPredefinedUser

`func (o *UpdateSafeMemberResponse) HasIsPredefinedUser() bool`

HasIsPredefinedUser returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *UpdateSafeMemberResponse) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *UpdateSafeMemberResponse) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *UpdateSafeMemberResponse) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *UpdateSafeMemberResponse) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateSafeMemberResponse) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateSafeMemberResponse) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateSafeMemberResponse) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateSafeMemberResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


