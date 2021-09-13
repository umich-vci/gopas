# BulkAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryModificationTime** | Pointer to **int64** |  | [optional] 
**GroupName** | Pointer to **string** | The name of the group to associate the account with. | [optional] 
**GroupPlatformId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**UploadIndex** | Pointer to **int32** | The numeric identifier for the account. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the account. | [optional] 
**Address** | Pointer to **string** | The name or address of the machine where the account will be used.  Valid values: DNS/IP/URL where the account is managed | [optional] 
**UserName** | Pointer to **string** | Account user&#39;s name. | [optional] 
**PlatformId** | **string** | The platform assigned to this account.  Valid values: Valid platform IDs, example: WinServerLocal | 
**SafeName** | **string** | The Safe where the account will be created. | 
**SecretType** | Pointer to **string** | Description The type of password.  Valid values password, key | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**PlatformAccountProperties** | Pointer to **map[string]string** | Object containing key-value pairs to associate with the account, as defined by the account platform.  These properties are validated against the mandatory and optional properties of the specified platform&#39;s definition.  Optional properties that do not exist on the account will not be returned here.  Internal properties are not returned.  Valid values example: {\&quot;Location\&quot;: \&quot;IT\&quot;, \&quot;OwnerName\&quot;:\&quot;MSSPAdmin\&quot;} | [optional] 
**SecretManagement** | Pointer to [**AutomaticSecretManagement**](AutomaticSecretManagement.md) |  | [optional] 
**RemoteMachinesAccess** | Pointer to [**RemoteMachinesAccess**](RemoteMachinesAccess.md) |  | [optional] 
**CreatedTime** | Pointer to **int64** | Date and time account was created | [optional] 

## Methods

### NewBulkAccountModel

`func NewBulkAccountModel(platformId string, safeName string, ) *BulkAccountModel`

NewBulkAccountModel instantiates a new BulkAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountModelWithDefaults

`func NewBulkAccountModelWithDefaults() *BulkAccountModel`

NewBulkAccountModelWithDefaults instantiates a new BulkAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryModificationTime

`func (o *BulkAccountModel) GetCategoryModificationTime() int64`

GetCategoryModificationTime returns the CategoryModificationTime field if non-nil, zero value otherwise.

### GetCategoryModificationTimeOk

`func (o *BulkAccountModel) GetCategoryModificationTimeOk() (*int64, bool)`

GetCategoryModificationTimeOk returns a tuple with the CategoryModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryModificationTime

`func (o *BulkAccountModel) SetCategoryModificationTime(v int64)`

SetCategoryModificationTime sets CategoryModificationTime field to given value.

### HasCategoryModificationTime

`func (o *BulkAccountModel) HasCategoryModificationTime() bool`

HasCategoryModificationTime returns a boolean if a field has been set.

### GetGroupName

`func (o *BulkAccountModel) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *BulkAccountModel) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *BulkAccountModel) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *BulkAccountModel) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetGroupPlatformId

`func (o *BulkAccountModel) GetGroupPlatformId() string`

GetGroupPlatformId returns the GroupPlatformId field if non-nil, zero value otherwise.

### GetGroupPlatformIdOk

`func (o *BulkAccountModel) GetGroupPlatformIdOk() (*string, bool)`

GetGroupPlatformIdOk returns a tuple with the GroupPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPlatformId

`func (o *BulkAccountModel) SetGroupPlatformId(v string)`

SetGroupPlatformId sets GroupPlatformId field to given value.

### HasGroupPlatformId

`func (o *BulkAccountModel) HasGroupPlatformId() bool`

HasGroupPlatformId returns a boolean if a field has been set.

### GetError

`func (o *BulkAccountModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkAccountModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkAccountModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkAccountModel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetUploadIndex

`func (o *BulkAccountModel) GetUploadIndex() int32`

GetUploadIndex returns the UploadIndex field if non-nil, zero value otherwise.

### GetUploadIndexOk

`func (o *BulkAccountModel) GetUploadIndexOk() (*int32, bool)`

GetUploadIndexOk returns a tuple with the UploadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIndex

`func (o *BulkAccountModel) SetUploadIndex(v int32)`

SetUploadIndex sets UploadIndex field to given value.

### HasUploadIndex

`func (o *BulkAccountModel) HasUploadIndex() bool`

HasUploadIndex returns a boolean if a field has been set.

### GetId

`func (o *BulkAccountModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkAccountModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkAccountModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkAccountModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BulkAccountModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkAccountModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkAccountModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkAccountModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *BulkAccountModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkAccountModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkAccountModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BulkAccountModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUserName

`func (o *BulkAccountModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *BulkAccountModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *BulkAccountModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *BulkAccountModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPlatformId

`func (o *BulkAccountModel) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *BulkAccountModel) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *BulkAccountModel) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetSafeName

`func (o *BulkAccountModel) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *BulkAccountModel) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *BulkAccountModel) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.


### GetSecretType

`func (o *BulkAccountModel) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *BulkAccountModel) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *BulkAccountModel) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *BulkAccountModel) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecret

`func (o *BulkAccountModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BulkAccountModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BulkAccountModel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BulkAccountModel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPlatformAccountProperties

`func (o *BulkAccountModel) GetPlatformAccountProperties() map[string]string`

GetPlatformAccountProperties returns the PlatformAccountProperties field if non-nil, zero value otherwise.

### GetPlatformAccountPropertiesOk

`func (o *BulkAccountModel) GetPlatformAccountPropertiesOk() (*map[string]string, bool)`

GetPlatformAccountPropertiesOk returns a tuple with the PlatformAccountProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAccountProperties

`func (o *BulkAccountModel) SetPlatformAccountProperties(v map[string]string)`

SetPlatformAccountProperties sets PlatformAccountProperties field to given value.

### HasPlatformAccountProperties

`func (o *BulkAccountModel) HasPlatformAccountProperties() bool`

HasPlatformAccountProperties returns a boolean if a field has been set.

### GetSecretManagement

`func (o *BulkAccountModel) GetSecretManagement() AutomaticSecretManagement`

GetSecretManagement returns the SecretManagement field if non-nil, zero value otherwise.

### GetSecretManagementOk

`func (o *BulkAccountModel) GetSecretManagementOk() (*AutomaticSecretManagement, bool)`

GetSecretManagementOk returns a tuple with the SecretManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagement

`func (o *BulkAccountModel) SetSecretManagement(v AutomaticSecretManagement)`

SetSecretManagement sets SecretManagement field to given value.

### HasSecretManagement

`func (o *BulkAccountModel) HasSecretManagement() bool`

HasSecretManagement returns a boolean if a field has been set.

### GetRemoteMachinesAccess

`func (o *BulkAccountModel) GetRemoteMachinesAccess() RemoteMachinesAccess`

GetRemoteMachinesAccess returns the RemoteMachinesAccess field if non-nil, zero value otherwise.

### GetRemoteMachinesAccessOk

`func (o *BulkAccountModel) GetRemoteMachinesAccessOk() (*RemoteMachinesAccess, bool)`

GetRemoteMachinesAccessOk returns a tuple with the RemoteMachinesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMachinesAccess

`func (o *BulkAccountModel) SetRemoteMachinesAccess(v RemoteMachinesAccess)`

SetRemoteMachinesAccess sets RemoteMachinesAccess field to given value.

### HasRemoteMachinesAccess

`func (o *BulkAccountModel) HasRemoteMachinesAccess() bool`

HasRemoteMachinesAccess returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BulkAccountModel) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BulkAccountModel) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BulkAccountModel) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BulkAccountModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


