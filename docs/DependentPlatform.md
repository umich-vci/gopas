# DependentPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfLinkedTargetPlatforms** | Pointer to **int32** |  | [optional] 
**CredentialsManagementPolicy** | Pointer to [**CredentialsManagementPolicy**](CredentialsManagementPolicy.md) |  | [optional] 
**ID** | Pointer to **int64** | Unique ID of the platform in the system | [optional] 
**PlatformID** | Pointer to **string** | Unique textual representation of the platform in the system | [optional] 
**Name** | Pointer to **string** | Display name of the platform | [optional] 

## Methods

### NewDependentPlatform

`func NewDependentPlatform() *DependentPlatform`

NewDependentPlatform instantiates a new DependentPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependentPlatformWithDefaults

`func NewDependentPlatformWithDefaults() *DependentPlatform`

NewDependentPlatformWithDefaults instantiates a new DependentPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfLinkedTargetPlatforms

`func (o *DependentPlatform) GetNumberOfLinkedTargetPlatforms() int32`

GetNumberOfLinkedTargetPlatforms returns the NumberOfLinkedTargetPlatforms field if non-nil, zero value otherwise.

### GetNumberOfLinkedTargetPlatformsOk

`func (o *DependentPlatform) GetNumberOfLinkedTargetPlatformsOk() (*int32, bool)`

GetNumberOfLinkedTargetPlatformsOk returns a tuple with the NumberOfLinkedTargetPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLinkedTargetPlatforms

`func (o *DependentPlatform) SetNumberOfLinkedTargetPlatforms(v int32)`

SetNumberOfLinkedTargetPlatforms sets NumberOfLinkedTargetPlatforms field to given value.

### HasNumberOfLinkedTargetPlatforms

`func (o *DependentPlatform) HasNumberOfLinkedTargetPlatforms() bool`

HasNumberOfLinkedTargetPlatforms returns a boolean if a field has been set.

### GetCredentialsManagementPolicy

`func (o *DependentPlatform) GetCredentialsManagementPolicy() CredentialsManagementPolicy`

GetCredentialsManagementPolicy returns the CredentialsManagementPolicy field if non-nil, zero value otherwise.

### GetCredentialsManagementPolicyOk

`func (o *DependentPlatform) GetCredentialsManagementPolicyOk() (*CredentialsManagementPolicy, bool)`

GetCredentialsManagementPolicyOk returns a tuple with the CredentialsManagementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsManagementPolicy

`func (o *DependentPlatform) SetCredentialsManagementPolicy(v CredentialsManagementPolicy)`

SetCredentialsManagementPolicy sets CredentialsManagementPolicy field to given value.

### HasCredentialsManagementPolicy

`func (o *DependentPlatform) HasCredentialsManagementPolicy() bool`

HasCredentialsManagementPolicy returns a boolean if a field has been set.

### GetID

`func (o *DependentPlatform) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *DependentPlatform) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *DependentPlatform) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *DependentPlatform) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPlatformID

`func (o *DependentPlatform) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *DependentPlatform) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *DependentPlatform) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *DependentPlatform) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetName

`func (o *DependentPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DependentPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DependentPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DependentPlatform) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


