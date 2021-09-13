# PlatformListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platforms** | Pointer to [**[]PlatformModel**](PlatformModel.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewPlatformListModel

`func NewPlatformListModel() *PlatformListModel`

NewPlatformListModel instantiates a new PlatformListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformListModelWithDefaults

`func NewPlatformListModelWithDefaults() *PlatformListModel`

NewPlatformListModelWithDefaults instantiates a new PlatformListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatforms

`func (o *PlatformListModel) GetPlatforms() []PlatformModel`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PlatformListModel) GetPlatformsOk() (*[]PlatformModel, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PlatformListModel) SetPlatforms(v []PlatformModel)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PlatformListModel) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetTotal

`func (o *PlatformListModel) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PlatformListModel) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PlatformListModel) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PlatformListModel) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


