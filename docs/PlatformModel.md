# PlatformModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**General**](General.md) |  | [optional] 
**Properties** | Pointer to [**Properties**](Properties.md) |  | [optional] 
**LinkedAccounts** | Pointer to [**[]Identity**](Identity.md) | A list of all linked accounts types that are relevant for the platform, each linked account has a name and a display name | [optional] 
**CredentialsManagement** | Pointer to [**CredentialsManagement**](CredentialsManagement.md) |  | [optional] 
**SessionManagement** | Pointer to [**SessionManagement**](SessionManagement.md) |  | [optional] 
**PrivilegedAccessWorkflows** | Pointer to [**PrivilegedAccessWorkflows**](PrivilegedAccessWorkflows.md) |  | [optional] 

## Methods

### NewPlatformModel

`func NewPlatformModel() *PlatformModel`

NewPlatformModel instantiates a new PlatformModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformModelWithDefaults

`func NewPlatformModelWithDefaults() *PlatformModel`

NewPlatformModelWithDefaults instantiates a new PlatformModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *PlatformModel) GetGeneral() General`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *PlatformModel) GetGeneralOk() (*General, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *PlatformModel) SetGeneral(v General)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *PlatformModel) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetProperties

`func (o *PlatformModel) GetProperties() Properties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PlatformModel) GetPropertiesOk() (*Properties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PlatformModel) SetProperties(v Properties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PlatformModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetLinkedAccounts

`func (o *PlatformModel) GetLinkedAccounts() []Identity`

GetLinkedAccounts returns the LinkedAccounts field if non-nil, zero value otherwise.

### GetLinkedAccountsOk

`func (o *PlatformModel) GetLinkedAccountsOk() (*[]Identity, bool)`

GetLinkedAccountsOk returns a tuple with the LinkedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccounts

`func (o *PlatformModel) SetLinkedAccounts(v []Identity)`

SetLinkedAccounts sets LinkedAccounts field to given value.

### HasLinkedAccounts

`func (o *PlatformModel) HasLinkedAccounts() bool`

HasLinkedAccounts returns a boolean if a field has been set.

### GetCredentialsManagement

`func (o *PlatformModel) GetCredentialsManagement() CredentialsManagement`

GetCredentialsManagement returns the CredentialsManagement field if non-nil, zero value otherwise.

### GetCredentialsManagementOk

`func (o *PlatformModel) GetCredentialsManagementOk() (*CredentialsManagement, bool)`

GetCredentialsManagementOk returns a tuple with the CredentialsManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsManagement

`func (o *PlatformModel) SetCredentialsManagement(v CredentialsManagement)`

SetCredentialsManagement sets CredentialsManagement field to given value.

### HasCredentialsManagement

`func (o *PlatformModel) HasCredentialsManagement() bool`

HasCredentialsManagement returns a boolean if a field has been set.

### GetSessionManagement

`func (o *PlatformModel) GetSessionManagement() SessionManagement`

GetSessionManagement returns the SessionManagement field if non-nil, zero value otherwise.

### GetSessionManagementOk

`func (o *PlatformModel) GetSessionManagementOk() (*SessionManagement, bool)`

GetSessionManagementOk returns a tuple with the SessionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionManagement

`func (o *PlatformModel) SetSessionManagement(v SessionManagement)`

SetSessionManagement sets SessionManagement field to given value.

### HasSessionManagement

`func (o *PlatformModel) HasSessionManagement() bool`

HasSessionManagement returns a boolean if a field has been set.

### GetPrivilegedAccessWorkflows

`func (o *PlatformModel) GetPrivilegedAccessWorkflows() PrivilegedAccessWorkflows`

GetPrivilegedAccessWorkflows returns the PrivilegedAccessWorkflows field if non-nil, zero value otherwise.

### GetPrivilegedAccessWorkflowsOk

`func (o *PlatformModel) GetPrivilegedAccessWorkflowsOk() (*PrivilegedAccessWorkflows, bool)`

GetPrivilegedAccessWorkflowsOk returns a tuple with the PrivilegedAccessWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedAccessWorkflows

`func (o *PlatformModel) SetPrivilegedAccessWorkflows(v PrivilegedAccessWorkflows)`

SetPrivilegedAccessWorkflows sets PrivilegedAccessWorkflows field to given value.

### HasPrivilegedAccessWorkflows

`func (o *PlatformModel) HasPrivilegedAccessWorkflows() bool`

HasPrivilegedAccessWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


