# SafeMemberItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberName** | **string** | The Vault user name, Domain user name or group name of the Safe member. | 
**SearchIn** | Pointer to **string** | The Vault or Domain where the user or group was found. | [optional] 
**MembershipExpirationDate** | Pointer to **int64** | The member&#39;s expiration date for this Safe.  For members that do not have an expiration date, this value will be null. | [optional] 
**Permissions** | **map[string]bool** | &lt;p&gt;The permissions that the user or group has in this Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;useAccounts&lt;/B&gt; - Use accounts but not view passwords.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;retrieveAccounts&lt;/B&gt; - Retrieve and view accounts in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;listAccounts&lt;/B&gt; - View Accounts list.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;addAccounts&lt;/B&gt; - Add accounts in the Safe. Users who have this permission automatically have UpdateAccountProperties as well.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;updateAccountContent&lt;/B&gt; - Update existing account content.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;updateAccountProperties&lt;/B&gt; - Update existing account properties.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;initiateCPMAccountManagementOperations&lt;/B&gt; - Initiate password management operations through CPM, such as changing, verifying, and reconciling passwords. When this parameter is set to False, the SpecifyNextAccountContent parameter is also automatically set to False.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;specifyNextAccountContent&lt;/B&gt; - Specify the password that is used when the CPM changes the password value. This parameter can only be specified when the InitiateCPMAccountManagementOperations parameter is set to True. When InitiateCPMAccountManagementOperations is set to False this parameter is automatically set to False.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;renameAccounts&lt;/B&gt; - Rename existing accounts in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;deleteAccounts&lt;/B&gt; - Delete existing passwords in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;unlockAccounts&lt;/B&gt; - Unlock accounts that are locked by other users.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;manageSafe&lt;/B&gt; - Perform administrative tasksin the Safe, including: Update Safe properties, Recover the Safe, Delete the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;manageSafeMembers&lt;/B&gt; - Add and remove Safe members, and update their authorizations in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;backupSafe&lt;/B&gt; - Create a backup of a Safe and its contents, and store in another location.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;viewAuditLog&lt;/B&gt; - View account and user activity in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;viewSafeMembers&lt;/B&gt; - View Safe member&#39;s permissions.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;requestsAuthorizationLevel1&lt;/B&gt; - Request Authorization Level 1.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;requestsAuthorizationLevel2&lt;/B&gt; - Request Authorization Level 2.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;accessWithoutConfirmation&lt;/B&gt; - Access the Safe without confirmation from authorized users. This overrides the Safe properties that specify that Safe members require confirmation to access the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;createFolders&lt;/B&gt; - Create folders in the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;deleteFolders&lt;/B&gt; - Delete folders from the Safe.&lt;/p&gt;  &lt;p&gt;    &lt;B&gt;moveAccountsAndFolders&lt;/B&gt; - Move accounts and folders in the Safe to different folders and subfolders.&lt;/p&gt; | 

## Methods

### NewSafeMemberItem

`func NewSafeMemberItem(memberName string, permissions map[string]bool, ) *SafeMemberItem`

NewSafeMemberItem instantiates a new SafeMemberItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeMemberItemWithDefaults

`func NewSafeMemberItemWithDefaults() *SafeMemberItem`

NewSafeMemberItemWithDefaults instantiates a new SafeMemberItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberName

`func (o *SafeMemberItem) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *SafeMemberItem) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *SafeMemberItem) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.


### GetSearchIn

`func (o *SafeMemberItem) GetSearchIn() string`

GetSearchIn returns the SearchIn field if non-nil, zero value otherwise.

### GetSearchInOk

`func (o *SafeMemberItem) GetSearchInOk() (*string, bool)`

GetSearchInOk returns a tuple with the SearchIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIn

`func (o *SafeMemberItem) SetSearchIn(v string)`

SetSearchIn sets SearchIn field to given value.

### HasSearchIn

`func (o *SafeMemberItem) HasSearchIn() bool`

HasSearchIn returns a boolean if a field has been set.

### GetMembershipExpirationDate

`func (o *SafeMemberItem) GetMembershipExpirationDate() int64`

GetMembershipExpirationDate returns the MembershipExpirationDate field if non-nil, zero value otherwise.

### GetMembershipExpirationDateOk

`func (o *SafeMemberItem) GetMembershipExpirationDateOk() (*int64, bool)`

GetMembershipExpirationDateOk returns a tuple with the MembershipExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipExpirationDate

`func (o *SafeMemberItem) SetMembershipExpirationDate(v int64)`

SetMembershipExpirationDate sets MembershipExpirationDate field to given value.

### HasMembershipExpirationDate

`func (o *SafeMemberItem) HasMembershipExpirationDate() bool`

HasMembershipExpirationDate returns a boolean if a field has been set.

### GetPermissions

`func (o *SafeMemberItem) GetPermissions() map[string]bool`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SafeMemberItem) GetPermissionsOk() (*map[string]bool, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SafeMemberItem) SetPermissions(v map[string]bool)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


