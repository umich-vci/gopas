# BulkAccountsOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The single identifier for the bulk action. | [optional] 
**Source** | Pointer to **string** | Free-text description of the source of the bulk account upload. | [optional] 
**CreationTime** | Pointer to **int64** | The time the bulk action was created. | [optional] 
**Status** | Pointer to **int32** | The possible statuses are:  inProgress: the bulk upload is still processing  completed: the bulk upload completed and all accounts were uploaded successfully  completedWithErrors: the bulk upload completed but there were accounts that failed to upload  failed: the bulk upload process failed to run | [optional] 
**Result** | Pointer to [**BulkOperationSummary**](BulkOperationSummary.md) |  | [optional] 

## Methods

### NewBulkAccountsOperationResult

`func NewBulkAccountsOperationResult() *BulkAccountsOperationResult`

NewBulkAccountsOperationResult instantiates a new BulkAccountsOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountsOperationResultWithDefaults

`func NewBulkAccountsOperationResultWithDefaults() *BulkAccountsOperationResult`

NewBulkAccountsOperationResultWithDefaults instantiates a new BulkAccountsOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkAccountsOperationResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkAccountsOperationResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkAccountsOperationResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BulkAccountsOperationResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *BulkAccountsOperationResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkAccountsOperationResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkAccountsOperationResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkAccountsOperationResult) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCreationTime

`func (o *BulkAccountsOperationResult) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BulkAccountsOperationResult) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BulkAccountsOperationResult) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BulkAccountsOperationResult) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStatus

`func (o *BulkAccountsOperationResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkAccountsOperationResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkAccountsOperationResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkAccountsOperationResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *BulkAccountsOperationResult) GetResult() BulkOperationSummary`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BulkAccountsOperationResult) GetResultOk() (*BulkOperationSummary, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BulkAccountsOperationResult) SetResult(v BulkOperationSummary)`

SetResult sets Result field to given value.

### HasResult

`func (o *BulkAccountsOperationResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


