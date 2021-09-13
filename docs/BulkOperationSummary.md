# BulkOperationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkOperationSummary

`func NewBulkOperationSummary() *BulkOperationSummary`

NewBulkOperationSummary instantiates a new BulkOperationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationSummaryWithDefaults

`func NewBulkOperationSummaryWithDefaults() *BulkOperationSummary`

NewBulkOperationSummaryWithDefaults instantiates a new BulkOperationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *BulkOperationSummary) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *BulkOperationSummary) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *BulkOperationSummary) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *BulkOperationSummary) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetFailed

`func (o *BulkOperationSummary) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkOperationSummary) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkOperationSummary) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BulkOperationSummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetError

`func (o *BulkOperationSummary) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkOperationSummary) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkOperationSummary) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BulkOperationSummary) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


