# TargetPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether a platform is active or inactive. An inactive platform cannot be assigned to accounts but will continue to manage accounts that were already assigned to it | [optional] 
**SystemType** | Pointer to **string** | The type of system to target belong to | [optional] 
**AllowedSafes** | Pointer to **string** | Regular expression of safes in which accounts from this platform can be managed | [optional] 
**PrivilegedAccessWorkflows** | Pointer to [**PrivilegedAccessWorkflowsPolicy**](PrivilegedAccessWorkflowsPolicy.md) |  | [optional] 
**CredentialsManagementPolicy** | Pointer to [**CredentialsManagementPolicy**](CredentialsManagementPolicy.md) |  | [optional] 
**PrivilegedSessionManagement** | Pointer to [**PrivilegedSessionManagementBase**](PrivilegedSessionManagementBase.md) |  | [optional] 
**ID** | Pointer to **int64** | Unique ID of the platform in the system | [optional] 
**PlatformID** | Pointer to **string** | Unique textual representation of the platform in the system | [optional] 
**Name** | Pointer to **string** | Display name of the platform | [optional] 

## Methods

### NewTargetPlatform

`func NewTargetPlatform() *TargetPlatform`

NewTargetPlatform instantiates a new TargetPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetPlatformWithDefaults

`func NewTargetPlatformWithDefaults() *TargetPlatform`

NewTargetPlatformWithDefaults instantiates a new TargetPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *TargetPlatform) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TargetPlatform) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TargetPlatform) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TargetPlatform) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSystemType

`func (o *TargetPlatform) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *TargetPlatform) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *TargetPlatform) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *TargetPlatform) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetAllowedSafes

`func (o *TargetPlatform) GetAllowedSafes() string`

GetAllowedSafes returns the AllowedSafes field if non-nil, zero value otherwise.

### GetAllowedSafesOk

`func (o *TargetPlatform) GetAllowedSafesOk() (*string, bool)`

GetAllowedSafesOk returns a tuple with the AllowedSafes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSafes

`func (o *TargetPlatform) SetAllowedSafes(v string)`

SetAllowedSafes sets AllowedSafes field to given value.

### HasAllowedSafes

`func (o *TargetPlatform) HasAllowedSafes() bool`

HasAllowedSafes returns a boolean if a field has been set.

### GetPrivilegedAccessWorkflows

`func (o *TargetPlatform) GetPrivilegedAccessWorkflows() PrivilegedAccessWorkflowsPolicy`

GetPrivilegedAccessWorkflows returns the PrivilegedAccessWorkflows field if non-nil, zero value otherwise.

### GetPrivilegedAccessWorkflowsOk

`func (o *TargetPlatform) GetPrivilegedAccessWorkflowsOk() (*PrivilegedAccessWorkflowsPolicy, bool)`

GetPrivilegedAccessWorkflowsOk returns a tuple with the PrivilegedAccessWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedAccessWorkflows

`func (o *TargetPlatform) SetPrivilegedAccessWorkflows(v PrivilegedAccessWorkflowsPolicy)`

SetPrivilegedAccessWorkflows sets PrivilegedAccessWorkflows field to given value.

### HasPrivilegedAccessWorkflows

`func (o *TargetPlatform) HasPrivilegedAccessWorkflows() bool`

HasPrivilegedAccessWorkflows returns a boolean if a field has been set.

### GetCredentialsManagementPolicy

`func (o *TargetPlatform) GetCredentialsManagementPolicy() CredentialsManagementPolicy`

GetCredentialsManagementPolicy returns the CredentialsManagementPolicy field if non-nil, zero value otherwise.

### GetCredentialsManagementPolicyOk

`func (o *TargetPlatform) GetCredentialsManagementPolicyOk() (*CredentialsManagementPolicy, bool)`

GetCredentialsManagementPolicyOk returns a tuple with the CredentialsManagementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsManagementPolicy

`func (o *TargetPlatform) SetCredentialsManagementPolicy(v CredentialsManagementPolicy)`

SetCredentialsManagementPolicy sets CredentialsManagementPolicy field to given value.

### HasCredentialsManagementPolicy

`func (o *TargetPlatform) HasCredentialsManagementPolicy() bool`

HasCredentialsManagementPolicy returns a boolean if a field has been set.

### GetPrivilegedSessionManagement

`func (o *TargetPlatform) GetPrivilegedSessionManagement() PrivilegedSessionManagementBase`

GetPrivilegedSessionManagement returns the PrivilegedSessionManagement field if non-nil, zero value otherwise.

### GetPrivilegedSessionManagementOk

`func (o *TargetPlatform) GetPrivilegedSessionManagementOk() (*PrivilegedSessionManagementBase, bool)`

GetPrivilegedSessionManagementOk returns a tuple with the PrivilegedSessionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedSessionManagement

`func (o *TargetPlatform) SetPrivilegedSessionManagement(v PrivilegedSessionManagementBase)`

SetPrivilegedSessionManagement sets PrivilegedSessionManagement field to given value.

### HasPrivilegedSessionManagement

`func (o *TargetPlatform) HasPrivilegedSessionManagement() bool`

HasPrivilegedSessionManagement returns a boolean if a field has been set.

### GetID

`func (o *TargetPlatform) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *TargetPlatform) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *TargetPlatform) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *TargetPlatform) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPlatformID

`func (o *TargetPlatform) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *TargetPlatform) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *TargetPlatform) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *TargetPlatform) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetName

`func (o *TargetPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetPlatform) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


