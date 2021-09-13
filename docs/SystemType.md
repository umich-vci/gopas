# SystemType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformsCount** | Pointer to **int32** | Number of platforms that match the filter in the system type | [optional] 
**Name** | Pointer to **string** | Display name of the system types | [optional] 

## Methods

### NewSystemType

`func NewSystemType() *SystemType`

NewSystemType instantiates a new SystemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemTypeWithDefaults

`func NewSystemTypeWithDefaults() *SystemType`

NewSystemTypeWithDefaults instantiates a new SystemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformsCount

`func (o *SystemType) GetPlatformsCount() int32`

GetPlatformsCount returns the PlatformsCount field if non-nil, zero value otherwise.

### GetPlatformsCountOk

`func (o *SystemType) GetPlatformsCountOk() (*int32, bool)`

GetPlatformsCountOk returns a tuple with the PlatformsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformsCount

`func (o *SystemType) SetPlatformsCount(v int32)`

SetPlatformsCount sets PlatformsCount field to given value.

### HasPlatformsCount

`func (o *SystemType) HasPlatformsCount() bool`

HasPlatformsCount returns a boolean if a field has been set.

### GetName

`func (o *SystemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemType) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


