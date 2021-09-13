# AccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryModificationTime** | Pointer to **int64** |  | [optional] 
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

### NewAccountModel

`func NewAccountModel(platformId string, safeName string, ) *AccountModel`

NewAccountModel instantiates a new AccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountModelWithDefaults

`func NewAccountModelWithDefaults() *AccountModel`

NewAccountModelWithDefaults instantiates a new AccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryModificationTime

`func (o *AccountModel) GetCategoryModificationTime() int64`

GetCategoryModificationTime returns the CategoryModificationTime field if non-nil, zero value otherwise.

### GetCategoryModificationTimeOk

`func (o *AccountModel) GetCategoryModificationTimeOk() (*int64, bool)`

GetCategoryModificationTimeOk returns a tuple with the CategoryModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryModificationTime

`func (o *AccountModel) SetCategoryModificationTime(v int64)`

SetCategoryModificationTime sets CategoryModificationTime field to given value.

### HasCategoryModificationTime

`func (o *AccountModel) HasCategoryModificationTime() bool`

HasCategoryModificationTime returns a boolean if a field has been set.

### GetId

`func (o *AccountModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *AccountModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUserName

`func (o *AccountModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AccountModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AccountModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AccountModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPlatformId

`func (o *AccountModel) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *AccountModel) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *AccountModel) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetSafeName

`func (o *AccountModel) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *AccountModel) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *AccountModel) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.


### GetSecretType

`func (o *AccountModel) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *AccountModel) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *AccountModel) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *AccountModel) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecret

`func (o *AccountModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AccountModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AccountModel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AccountModel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPlatformAccountProperties

`func (o *AccountModel) GetPlatformAccountProperties() map[string]string`

GetPlatformAccountProperties returns the PlatformAccountProperties field if non-nil, zero value otherwise.

### GetPlatformAccountPropertiesOk

`func (o *AccountModel) GetPlatformAccountPropertiesOk() (*map[string]string, bool)`

GetPlatformAccountPropertiesOk returns a tuple with the PlatformAccountProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAccountProperties

`func (o *AccountModel) SetPlatformAccountProperties(v map[string]string)`

SetPlatformAccountProperties sets PlatformAccountProperties field to given value.

### HasPlatformAccountProperties

`func (o *AccountModel) HasPlatformAccountProperties() bool`

HasPlatformAccountProperties returns a boolean if a field has been set.

### GetSecretManagement

`func (o *AccountModel) GetSecretManagement() AutomaticSecretManagement`

GetSecretManagement returns the SecretManagement field if non-nil, zero value otherwise.

### GetSecretManagementOk

`func (o *AccountModel) GetSecretManagementOk() (*AutomaticSecretManagement, bool)`

GetSecretManagementOk returns a tuple with the SecretManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagement

`func (o *AccountModel) SetSecretManagement(v AutomaticSecretManagement)`

SetSecretManagement sets SecretManagement field to given value.

### HasSecretManagement

`func (o *AccountModel) HasSecretManagement() bool`

HasSecretManagement returns a boolean if a field has been set.

### GetRemoteMachinesAccess

`func (o *AccountModel) GetRemoteMachinesAccess() RemoteMachinesAccess`

GetRemoteMachinesAccess returns the RemoteMachinesAccess field if non-nil, zero value otherwise.

### GetRemoteMachinesAccessOk

`func (o *AccountModel) GetRemoteMachinesAccessOk() (*RemoteMachinesAccess, bool)`

GetRemoteMachinesAccessOk returns a tuple with the RemoteMachinesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMachinesAccess

`func (o *AccountModel) SetRemoteMachinesAccess(v RemoteMachinesAccess)`

SetRemoteMachinesAccess sets RemoteMachinesAccess field to given value.

### HasRemoteMachinesAccess

`func (o *AccountModel) HasRemoteMachinesAccess() bool`

HasRemoteMachinesAccess returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AccountModel) GetCreatedTime() int64`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountModel) GetCreatedTimeOk() (*int64, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountModel) SetCreatedTime(v int64)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AccountModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


