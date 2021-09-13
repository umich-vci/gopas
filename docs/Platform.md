# Platform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformID** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlatform

`func NewPlatform() *Platform`

NewPlatform instantiates a new Platform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformWithDefaults

`func NewPlatformWithDefaults() *Platform`

NewPlatformWithDefaults instantiates a new Platform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformID

`func (o *Platform) GetPlatformID() string`

GetPlatformID returns the PlatformID field if non-nil, zero value otherwise.

### GetPlatformIDOk

`func (o *Platform) GetPlatformIDOk() (*string, bool)`

GetPlatformIDOk returns a tuple with the PlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformID

`func (o *Platform) SetPlatformID(v string)`

SetPlatformID sets PlatformID field to given value.

### HasPlatformID

`func (o *Platform) HasPlatformID() bool`

HasPlatformID returns a boolean if a field has been set.

### GetDetails

`func (o *Platform) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Platform) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Platform) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Platform) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetActive

`func (o *Platform) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Platform) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Platform) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Platform) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


