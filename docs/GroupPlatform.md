# GroupPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether a platform is active or inactive. An inactive platform cannot be assigned to accounts but will continue to manage accounts that were already assigned to it | [optional] 
**ID** | Pointer to **int64** | Unique ID of the platform in the system | [optional] 
**PlatformID** | Pointer to **string** | Unique textual representation of the platform in the system | [optional] 
**Name** | Pointer to **string** | Display name of the platform | [optional] 

## Methods

### NewGroupPlatform

`func NewGroupPlatform() *GroupPlatform`

NewGroupPlatform instantiates a new GroupPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPlatformWithDefaults

`func NewGroupPlatformWithDefaults() *GroupPlatform`

NewGroupPlatformWithDefaults instantiates a new GroupPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GroupPlatform) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GroupPlatform) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GroupPlatform) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GroupPlatform) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetID

`func (o *GroupPlatform) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *GroupPlatform) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *GroupPlatform) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *GroupPlatform) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPlatformID

`func (o *GroupPlatform) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *GroupPlatform) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *GroupPlatform) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *GroupPlatform) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetName

`func (o *GroupPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupPlatform) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupPlatform) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


