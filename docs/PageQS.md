# PageQS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**UseCache** | Pointer to **bool** |  | [optional] 

## Methods

### NewPageQS

`func NewPageQS() *PageQS`

NewPageQS instantiates a new PageQS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageQSWithDefaults

`func NewPageQSWithDefaults() *PageQS`

NewPageQSWithDefaults instantiates a new PageQS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PageQS) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageQS) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageQS) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PageQS) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PageQS) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PageQS) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PageQS) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PageQS) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetUseCache

`func (o *PageQS) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *PageQS) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *PageQS) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *PageQS) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


