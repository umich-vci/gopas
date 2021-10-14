# GetSafeMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseCache** | Pointer to **bool** | Whether to retrieve from session or not. | [optional] 

## Methods

### NewGetSafeMemberRequest

`func NewGetSafeMemberRequest() *GetSafeMemberRequest`

NewGetSafeMemberRequest instantiates a new GetSafeMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSafeMemberRequestWithDefaults

`func NewGetSafeMemberRequestWithDefaults() *GetSafeMemberRequest`

NewGetSafeMemberRequestWithDefaults instantiates a new GetSafeMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseCache

`func (o *GetSafeMemberRequest) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *GetSafeMemberRequest) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *GetSafeMemberRequest) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *GetSafeMemberRequest) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


