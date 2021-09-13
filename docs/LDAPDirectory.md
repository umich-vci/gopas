# LDAPDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryType** | **string** | The name of the directory profile file that represents the profile the Vault should use when working with the specified LDAP directory. | 
**BindUsername** | Pointer to **string** | The user that will authenticate to the directory server. | [optional] 
**BindPassword** | Pointer to **string** | The password for the user specified in the Bind User field. | [optional] 
**Port** | Pointer to **int32** | The port that will be used to access the specified server.  The standard port for SSL LDAP connections is 636, and for non-SSL LDAP connections is 389. | [optional] 
**SSLConnect** | Pointer to **bool** | Read-only - Whether or not to connect to the external directory with SSL. | [optional] 
**LDAPDirectoryName** | Pointer to **string** | Read-only - The name of the LDAP directory where users and groups are listed.  Note: After external users and groups from this directory have been created in the Vault, this parameter must not be changed. | [optional] 
**LDAPDirectoryQueryOrder** | Pointer to **int32** | Read-only - The order in which the Vault will search directories for users before creating a corresponding User Account or Group in the Vault. | [optional] 
**LDAPDirectoryDescription** | Pointer to **string** | Read-only - A short description of the LDAP directory. | [optional] 
**VaultObjectNamesPrefix** | Pointer to **string** | Read-only - The text that will be used as a prefix for external users and groups in the Vault created from the specified directory.  This parameter is oprional,  The parameter will be ignored if AddDomainToUserName parameter is set to Yes. | [optional] 
**PasswordObjectPath** | Pointer to **string** | Read-only - The location of the bind password in the VaultInternal Safe. | [optional] 
**LDAPDirectoryGroupBaseContext** | Pointer to **string** | Read-only - The base context that will be used for external directory queries, for groups only. This parameter is optional. | [optional] 
**ReferralsChasingHopLimit** | Pointer to **int32** | Read-only - The number of recursive LDAP referrals that will be chased. The default value is -1 (unlimited). | [optional] 
**AppendFriendlyDomainNameToGroup** | Pointer to **bool** | Read-only - Whether or not to add active directory domain names to names of groups that are provisioned in the Vault. | [optional] 
**RequireReferredDirectoryDefinition** | Pointer to **bool** | Read-only - Whether or not LDAP referrals will only be supported when an external directory parameter file has been defined in the Vault for the referred directory.  The default value is false | [optional] 
**ReferralsDNSLookup** | Pointer to **bool** | Read-only - Enables referrals to be specified as domain names. The default value is false.    If this parameter is not enabled, in SSL or High-Availabilty implementations,   a directory file must be created for each domain that will be supported by the Vault,   and the domain name specified in each parameter file must be mapped in the DomainDNSName parameter in the referred directory.   In addition, each referral directory must be defined in the Windows\\System32\\Etc\\Hosts file. | [optional] 
**DisableUserEnumeration** | Pointer to **bool** | Read-only - If set to true, prevents enumerating users from directory for Safe ownership lockups. | [optional] 
**AdditionalQueryFilterOptimize** | Pointer to **bool** | Read-only - Causes an additional query filter to be done in scope of specific user. | [optional] 
**ClientBrowsing** | Pointer to **bool** | Read-only - Whether or not to use this directory for PKI certificate browsing. | [optional] 
**ExternalObjectCreation** | Pointer to **bool** | Read-only - Whether or not to use this directory for user provisioning. | [optional] 
**Authentication** | Pointer to **bool** | Read-only - Whether or not to use this directory for authentication. | [optional] 
**UseLDAPCertificatesOnly** | Pointer to **bool** | Read-only - Determines whether the user certificate DN can be set manually, or taken from the Directory. | [optional] 
**DisablePaging** | Pointer to **bool** | Read-only - Determines whether or not to use page mode search while searching in the LDAP directory. This parameter is optional. | [optional] 
**ProvisionDisabledUsers** | Pointer to **bool** | Read-only - Whether or not LDAP disabled users will be created in the Vault. | [optional] 
**LDAPDirectoryUsage** | Pointer to **[]int32** | Read-only - Whether you can create external objects from this external directory, browse it, authenticate users, or do all three.  Possible values: ClientBrowsing, ExternalObjectsCreation, Authentication.  The user is able to set one value or multiple values separated by comma. | [optional] 
**DCList** | [**[]LDAPDomainController**](LDAPDomainController.md) | A list of host servers for External Directory. | 
**DomainName** | **string** | The address of the domain. | 
**DomainBaseContext** | **string** | The base context of the External Directory. | 

## Methods

### NewLDAPDirectory

`func NewLDAPDirectory(directoryType string, dCList []LDAPDomainController, domainName string, domainBaseContext string, ) *LDAPDirectory`

NewLDAPDirectory instantiates a new LDAPDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDirectoryWithDefaults

`func NewLDAPDirectoryWithDefaults() *LDAPDirectory`

NewLDAPDirectoryWithDefaults instantiates a new LDAPDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryType

`func (o *LDAPDirectory) GetDirectoryType() string`

GetDirectoryType returns the DirectoryType field if non-nil, zero value otherwise.

### GetDirectoryTypeOk

`func (o *LDAPDirectory) GetDirectoryTypeOk() (*string, bool)`

GetDirectoryTypeOk returns a tuple with the DirectoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryType

`func (o *LDAPDirectory) SetDirectoryType(v string)`

SetDirectoryType sets DirectoryType field to given value.


### GetBindUsername

`func (o *LDAPDirectory) GetBindUsername() string`

GetBindUsername returns the BindUsername field if non-nil, zero value otherwise.

### GetBindUsernameOk

`func (o *LDAPDirectory) GetBindUsernameOk() (*string, bool)`

GetBindUsernameOk returns a tuple with the BindUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindUsername

`func (o *LDAPDirectory) SetBindUsername(v string)`

SetBindUsername sets BindUsername field to given value.

### HasBindUsername

`func (o *LDAPDirectory) HasBindUsername() bool`

HasBindUsername returns a boolean if a field has been set.

### GetBindPassword

`func (o *LDAPDirectory) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *LDAPDirectory) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *LDAPDirectory) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *LDAPDirectory) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetPort

`func (o *LDAPDirectory) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPDirectory) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPDirectory) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPDirectory) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSSLConnect

`func (o *LDAPDirectory) GetSSLConnect() bool`

GetSSLConnect returns the SSLConnect field if non-nil, zero value otherwise.

### GetSSLConnectOk

`func (o *LDAPDirectory) GetSSLConnectOk() (*bool, bool)`

GetSSLConnectOk returns a tuple with the SSLConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLConnect

`func (o *LDAPDirectory) SetSSLConnect(v bool)`

SetSSLConnect sets SSLConnect field to given value.

### HasSSLConnect

`func (o *LDAPDirectory) HasSSLConnect() bool`

HasSSLConnect returns a boolean if a field has been set.

### GetLDAPDirectoryName

`func (o *LDAPDirectory) GetLDAPDirectoryName() string`

GetLDAPDirectoryName returns the LDAPDirectoryName field if non-nil, zero value otherwise.

### GetLDAPDirectoryNameOk

`func (o *LDAPDirectory) GetLDAPDirectoryNameOk() (*string, bool)`

GetLDAPDirectoryNameOk returns a tuple with the LDAPDirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPDirectoryName

`func (o *LDAPDirectory) SetLDAPDirectoryName(v string)`

SetLDAPDirectoryName sets LDAPDirectoryName field to given value.

### HasLDAPDirectoryName

`func (o *LDAPDirectory) HasLDAPDirectoryName() bool`

HasLDAPDirectoryName returns a boolean if a field has been set.

### GetLDAPDirectoryQueryOrder

`func (o *LDAPDirectory) GetLDAPDirectoryQueryOrder() int32`

GetLDAPDirectoryQueryOrder returns the LDAPDirectoryQueryOrder field if non-nil, zero value otherwise.

### GetLDAPDirectoryQueryOrderOk

`func (o *LDAPDirectory) GetLDAPDirectoryQueryOrderOk() (*int32, bool)`

GetLDAPDirectoryQueryOrderOk returns a tuple with the LDAPDirectoryQueryOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPDirectoryQueryOrder

`func (o *LDAPDirectory) SetLDAPDirectoryQueryOrder(v int32)`

SetLDAPDirectoryQueryOrder sets LDAPDirectoryQueryOrder field to given value.

### HasLDAPDirectoryQueryOrder

`func (o *LDAPDirectory) HasLDAPDirectoryQueryOrder() bool`

HasLDAPDirectoryQueryOrder returns a boolean if a field has been set.

### GetLDAPDirectoryDescription

`func (o *LDAPDirectory) GetLDAPDirectoryDescription() string`

GetLDAPDirectoryDescription returns the LDAPDirectoryDescription field if non-nil, zero value otherwise.

### GetLDAPDirectoryDescriptionOk

`func (o *LDAPDirectory) GetLDAPDirectoryDescriptionOk() (*string, bool)`

GetLDAPDirectoryDescriptionOk returns a tuple with the LDAPDirectoryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPDirectoryDescription

`func (o *LDAPDirectory) SetLDAPDirectoryDescription(v string)`

SetLDAPDirectoryDescription sets LDAPDirectoryDescription field to given value.

### HasLDAPDirectoryDescription

`func (o *LDAPDirectory) HasLDAPDirectoryDescription() bool`

HasLDAPDirectoryDescription returns a boolean if a field has been set.

### GetVaultObjectNamesPrefix

`func (o *LDAPDirectory) GetVaultObjectNamesPrefix() string`

GetVaultObjectNamesPrefix returns the VaultObjectNamesPrefix field if non-nil, zero value otherwise.

### GetVaultObjectNamesPrefixOk

`func (o *LDAPDirectory) GetVaultObjectNamesPrefixOk() (*string, bool)`

GetVaultObjectNamesPrefixOk returns a tuple with the VaultObjectNamesPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultObjectNamesPrefix

`func (o *LDAPDirectory) SetVaultObjectNamesPrefix(v string)`

SetVaultObjectNamesPrefix sets VaultObjectNamesPrefix field to given value.

### HasVaultObjectNamesPrefix

`func (o *LDAPDirectory) HasVaultObjectNamesPrefix() bool`

HasVaultObjectNamesPrefix returns a boolean if a field has been set.

### GetPasswordObjectPath

`func (o *LDAPDirectory) GetPasswordObjectPath() string`

GetPasswordObjectPath returns the PasswordObjectPath field if non-nil, zero value otherwise.

### GetPasswordObjectPathOk

`func (o *LDAPDirectory) GetPasswordObjectPathOk() (*string, bool)`

GetPasswordObjectPathOk returns a tuple with the PasswordObjectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordObjectPath

`func (o *LDAPDirectory) SetPasswordObjectPath(v string)`

SetPasswordObjectPath sets PasswordObjectPath field to given value.

### HasPasswordObjectPath

`func (o *LDAPDirectory) HasPasswordObjectPath() bool`

HasPasswordObjectPath returns a boolean if a field has been set.

### GetLDAPDirectoryGroupBaseContext

`func (o *LDAPDirectory) GetLDAPDirectoryGroupBaseContext() string`

GetLDAPDirectoryGroupBaseContext returns the LDAPDirectoryGroupBaseContext field if non-nil, zero value otherwise.

### GetLDAPDirectoryGroupBaseContextOk

`func (o *LDAPDirectory) GetLDAPDirectoryGroupBaseContextOk() (*string, bool)`

GetLDAPDirectoryGroupBaseContextOk returns a tuple with the LDAPDirectoryGroupBaseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPDirectoryGroupBaseContext

`func (o *LDAPDirectory) SetLDAPDirectoryGroupBaseContext(v string)`

SetLDAPDirectoryGroupBaseContext sets LDAPDirectoryGroupBaseContext field to given value.

### HasLDAPDirectoryGroupBaseContext

`func (o *LDAPDirectory) HasLDAPDirectoryGroupBaseContext() bool`

HasLDAPDirectoryGroupBaseContext returns a boolean if a field has been set.

### GetReferralsChasingHopLimit

`func (o *LDAPDirectory) GetReferralsChasingHopLimit() int32`

GetReferralsChasingHopLimit returns the ReferralsChasingHopLimit field if non-nil, zero value otherwise.

### GetReferralsChasingHopLimitOk

`func (o *LDAPDirectory) GetReferralsChasingHopLimitOk() (*int32, bool)`

GetReferralsChasingHopLimitOk returns a tuple with the ReferralsChasingHopLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralsChasingHopLimit

`func (o *LDAPDirectory) SetReferralsChasingHopLimit(v int32)`

SetReferralsChasingHopLimit sets ReferralsChasingHopLimit field to given value.

### HasReferralsChasingHopLimit

`func (o *LDAPDirectory) HasReferralsChasingHopLimit() bool`

HasReferralsChasingHopLimit returns a boolean if a field has been set.

### GetAppendFriendlyDomainNameToGroup

`func (o *LDAPDirectory) GetAppendFriendlyDomainNameToGroup() bool`

GetAppendFriendlyDomainNameToGroup returns the AppendFriendlyDomainNameToGroup field if non-nil, zero value otherwise.

### GetAppendFriendlyDomainNameToGroupOk

`func (o *LDAPDirectory) GetAppendFriendlyDomainNameToGroupOk() (*bool, bool)`

GetAppendFriendlyDomainNameToGroupOk returns a tuple with the AppendFriendlyDomainNameToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendFriendlyDomainNameToGroup

`func (o *LDAPDirectory) SetAppendFriendlyDomainNameToGroup(v bool)`

SetAppendFriendlyDomainNameToGroup sets AppendFriendlyDomainNameToGroup field to given value.

### HasAppendFriendlyDomainNameToGroup

`func (o *LDAPDirectory) HasAppendFriendlyDomainNameToGroup() bool`

HasAppendFriendlyDomainNameToGroup returns a boolean if a field has been set.

### GetRequireReferredDirectoryDefinition

`func (o *LDAPDirectory) GetRequireReferredDirectoryDefinition() bool`

GetRequireReferredDirectoryDefinition returns the RequireReferredDirectoryDefinition field if non-nil, zero value otherwise.

### GetRequireReferredDirectoryDefinitionOk

`func (o *LDAPDirectory) GetRequireReferredDirectoryDefinitionOk() (*bool, bool)`

GetRequireReferredDirectoryDefinitionOk returns a tuple with the RequireReferredDirectoryDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReferredDirectoryDefinition

`func (o *LDAPDirectory) SetRequireReferredDirectoryDefinition(v bool)`

SetRequireReferredDirectoryDefinition sets RequireReferredDirectoryDefinition field to given value.

### HasRequireReferredDirectoryDefinition

`func (o *LDAPDirectory) HasRequireReferredDirectoryDefinition() bool`

HasRequireReferredDirectoryDefinition returns a boolean if a field has been set.

### GetReferralsDNSLookup

`func (o *LDAPDirectory) GetReferralsDNSLookup() bool`

GetReferralsDNSLookup returns the ReferralsDNSLookup field if non-nil, zero value otherwise.

### GetReferralsDNSLookupOk

`func (o *LDAPDirectory) GetReferralsDNSLookupOk() (*bool, bool)`

GetReferralsDNSLookupOk returns a tuple with the ReferralsDNSLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralsDNSLookup

`func (o *LDAPDirectory) SetReferralsDNSLookup(v bool)`

SetReferralsDNSLookup sets ReferralsDNSLookup field to given value.

### HasReferralsDNSLookup

`func (o *LDAPDirectory) HasReferralsDNSLookup() bool`

HasReferralsDNSLookup returns a boolean if a field has been set.

### GetDisableUserEnumeration

`func (o *LDAPDirectory) GetDisableUserEnumeration() bool`

GetDisableUserEnumeration returns the DisableUserEnumeration field if non-nil, zero value otherwise.

### GetDisableUserEnumerationOk

`func (o *LDAPDirectory) GetDisableUserEnumerationOk() (*bool, bool)`

GetDisableUserEnumerationOk returns a tuple with the DisableUserEnumeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUserEnumeration

`func (o *LDAPDirectory) SetDisableUserEnumeration(v bool)`

SetDisableUserEnumeration sets DisableUserEnumeration field to given value.

### HasDisableUserEnumeration

`func (o *LDAPDirectory) HasDisableUserEnumeration() bool`

HasDisableUserEnumeration returns a boolean if a field has been set.

### GetAdditionalQueryFilterOptimize

`func (o *LDAPDirectory) GetAdditionalQueryFilterOptimize() bool`

GetAdditionalQueryFilterOptimize returns the AdditionalQueryFilterOptimize field if non-nil, zero value otherwise.

### GetAdditionalQueryFilterOptimizeOk

`func (o *LDAPDirectory) GetAdditionalQueryFilterOptimizeOk() (*bool, bool)`

GetAdditionalQueryFilterOptimizeOk returns a tuple with the AdditionalQueryFilterOptimize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQueryFilterOptimize

`func (o *LDAPDirectory) SetAdditionalQueryFilterOptimize(v bool)`

SetAdditionalQueryFilterOptimize sets AdditionalQueryFilterOptimize field to given value.

### HasAdditionalQueryFilterOptimize

`func (o *LDAPDirectory) HasAdditionalQueryFilterOptimize() bool`

HasAdditionalQueryFilterOptimize returns a boolean if a field has been set.

### GetClientBrowsing

`func (o *LDAPDirectory) GetClientBrowsing() bool`

GetClientBrowsing returns the ClientBrowsing field if non-nil, zero value otherwise.

### GetClientBrowsingOk

`func (o *LDAPDirectory) GetClientBrowsingOk() (*bool, bool)`

GetClientBrowsingOk returns a tuple with the ClientBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBrowsing

`func (o *LDAPDirectory) SetClientBrowsing(v bool)`

SetClientBrowsing sets ClientBrowsing field to given value.

### HasClientBrowsing

`func (o *LDAPDirectory) HasClientBrowsing() bool`

HasClientBrowsing returns a boolean if a field has been set.

### GetExternalObjectCreation

`func (o *LDAPDirectory) GetExternalObjectCreation() bool`

GetExternalObjectCreation returns the ExternalObjectCreation field if non-nil, zero value otherwise.

### GetExternalObjectCreationOk

`func (o *LDAPDirectory) GetExternalObjectCreationOk() (*bool, bool)`

GetExternalObjectCreationOk returns a tuple with the ExternalObjectCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalObjectCreation

`func (o *LDAPDirectory) SetExternalObjectCreation(v bool)`

SetExternalObjectCreation sets ExternalObjectCreation field to given value.

### HasExternalObjectCreation

`func (o *LDAPDirectory) HasExternalObjectCreation() bool`

HasExternalObjectCreation returns a boolean if a field has been set.

### GetAuthentication

`func (o *LDAPDirectory) GetAuthentication() bool`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *LDAPDirectory) GetAuthenticationOk() (*bool, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *LDAPDirectory) SetAuthentication(v bool)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *LDAPDirectory) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetUseLDAPCertificatesOnly

`func (o *LDAPDirectory) GetUseLDAPCertificatesOnly() bool`

GetUseLDAPCertificatesOnly returns the UseLDAPCertificatesOnly field if non-nil, zero value otherwise.

### GetUseLDAPCertificatesOnlyOk

`func (o *LDAPDirectory) GetUseLDAPCertificatesOnlyOk() (*bool, bool)`

GetUseLDAPCertificatesOnlyOk returns a tuple with the UseLDAPCertificatesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLDAPCertificatesOnly

`func (o *LDAPDirectory) SetUseLDAPCertificatesOnly(v bool)`

SetUseLDAPCertificatesOnly sets UseLDAPCertificatesOnly field to given value.

### HasUseLDAPCertificatesOnly

`func (o *LDAPDirectory) HasUseLDAPCertificatesOnly() bool`

HasUseLDAPCertificatesOnly returns a boolean if a field has been set.

### GetDisablePaging

`func (o *LDAPDirectory) GetDisablePaging() bool`

GetDisablePaging returns the DisablePaging field if non-nil, zero value otherwise.

### GetDisablePagingOk

`func (o *LDAPDirectory) GetDisablePagingOk() (*bool, bool)`

GetDisablePagingOk returns a tuple with the DisablePaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePaging

`func (o *LDAPDirectory) SetDisablePaging(v bool)`

SetDisablePaging sets DisablePaging field to given value.

### HasDisablePaging

`func (o *LDAPDirectory) HasDisablePaging() bool`

HasDisablePaging returns a boolean if a field has been set.

### GetProvisionDisabledUsers

`func (o *LDAPDirectory) GetProvisionDisabledUsers() bool`

GetProvisionDisabledUsers returns the ProvisionDisabledUsers field if non-nil, zero value otherwise.

### GetProvisionDisabledUsersOk

`func (o *LDAPDirectory) GetProvisionDisabledUsersOk() (*bool, bool)`

GetProvisionDisabledUsersOk returns a tuple with the ProvisionDisabledUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionDisabledUsers

`func (o *LDAPDirectory) SetProvisionDisabledUsers(v bool)`

SetProvisionDisabledUsers sets ProvisionDisabledUsers field to given value.

### HasProvisionDisabledUsers

`func (o *LDAPDirectory) HasProvisionDisabledUsers() bool`

HasProvisionDisabledUsers returns a boolean if a field has been set.

### GetLDAPDirectoryUsage

`func (o *LDAPDirectory) GetLDAPDirectoryUsage() []int32`

GetLDAPDirectoryUsage returns the LDAPDirectoryUsage field if non-nil, zero value otherwise.

### GetLDAPDirectoryUsageOk

`func (o *LDAPDirectory) GetLDAPDirectoryUsageOk() (*[]int32, bool)`

GetLDAPDirectoryUsageOk returns a tuple with the LDAPDirectoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAPDirectoryUsage

`func (o *LDAPDirectory) SetLDAPDirectoryUsage(v []int32)`

SetLDAPDirectoryUsage sets LDAPDirectoryUsage field to given value.

### HasLDAPDirectoryUsage

`func (o *LDAPDirectory) HasLDAPDirectoryUsage() bool`

HasLDAPDirectoryUsage returns a boolean if a field has been set.

### GetDCList

`func (o *LDAPDirectory) GetDCList() []LDAPDomainController`

GetDCList returns the DCList field if non-nil, zero value otherwise.

### GetDCListOk

`func (o *LDAPDirectory) GetDCListOk() (*[]LDAPDomainController, bool)`

GetDCListOk returns a tuple with the DCList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDCList

`func (o *LDAPDirectory) SetDCList(v []LDAPDomainController)`

SetDCList sets DCList field to given value.


### GetDomainName

`func (o *LDAPDirectory) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *LDAPDirectory) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *LDAPDirectory) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetDomainBaseContext

`func (o *LDAPDirectory) GetDomainBaseContext() string`

GetDomainBaseContext returns the DomainBaseContext field if non-nil, zero value otherwise.

### GetDomainBaseContextOk

`func (o *LDAPDirectory) GetDomainBaseContextOk() (*string, bool)`

GetDomainBaseContextOk returns a tuple with the DomainBaseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainBaseContext

`func (o *LDAPDirectory) SetDomainBaseContext(v string)`

SetDomainBaseContext sets DomainBaseContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


