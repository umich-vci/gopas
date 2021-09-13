# GetPlatformsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**PlatformType** | Pointer to **int32** |  | [optional] 
**Search** | Pointer to **string** | The search will be by Platform ID or Platform Name. | [optional] 
**SystemType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPlatformsRequest

`func NewGetPlatformsRequest() *GetPlatformsRequest`

NewGetPlatformsRequest instantiates a new GetPlatformsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlatformsRequestWithDefaults

`func NewGetPlatformsRequestWithDefaults() *GetPlatformsRequest`

NewGetPlatformsRequestWithDefaults instantiates a new GetPlatformsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetPlatformsRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetPlatformsRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetPlatformsRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetPlatformsRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPlatformType

`func (o *GetPlatformsRequest) GetPlatformType() int32`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *GetPlatformsRequest) GetPlatformTypeOk() (*int32, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *GetPlatformsRequest) SetPlatformType(v int32)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *GetPlatformsRequest) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetSearch

`func (o *GetPlatformsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *GetPlatformsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *GetPlatformsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *GetPlatformsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSystemType

`func (o *GetPlatformsRequest) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *GetPlatformsRequest) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *GetPlatformsRequest) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *GetPlatformsRequest) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


