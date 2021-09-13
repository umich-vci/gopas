# BulkAccountsOperationResultExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The single identifier for the bulk action. | [optional] 
**Source** | Pointer to **string** | Free-text description of the source of the bulk account upload. | [optional] 
**CreationTime** | Pointer to **int64** | The time the bulk action was created. | [optional] 
**Status** | Pointer to **int32** | The possible statuses are:  inProgress: the bulk upload is still processing  completed: the bulk upload completed and all accounts were uploaded successfully  completedWithErrors: the bulk upload completed but there were accounts that failed to upload  failed: the bulk upload process failed to run | [optional] 
**Result** | Pointer to [**BulkOperationSummary**](BulkOperationSummary.md) |  | [optional] 
**SucceededItems** | Pointer to [**BulkAccountsList**](BulkAccountsList.md) |  | [optional] 
**FailedItems** | Pointer to [**BulkAccountsList**](BulkAccountsList.md) |  | [optional] 

## Methods

### NewBulkAccountsOperationResultExtended

`func NewBulkAccountsOperationResultExtended() *BulkAccountsOperationResultExtended`

NewBulkAccountsOperationResultExtended instantiates a new BulkAccountsOperationResultExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountsOperationResultExtendedWithDefaults

`func NewBulkAccountsOperationResultExtendedWithDefaults() *BulkAccountsOperationResultExtended`

NewBulkAccountsOperationResultExtendedWithDefaults instantiates a new BulkAccountsOperationResultExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkAccountsOperationResultExtended) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkAccountsOperationResultExtended) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkAccountsOperationResultExtended) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BulkAccountsOperationResultExtended) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *BulkAccountsOperationResultExtended) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkAccountsOperationResultExtended) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkAccountsOperationResultExtended) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkAccountsOperationResultExtended) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCreationTime

`func (o *BulkAccountsOperationResultExtended) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BulkAccountsOperationResultExtended) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BulkAccountsOperationResultExtended) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BulkAccountsOperationResultExtended) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStatus

`func (o *BulkAccountsOperationResultExtended) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkAccountsOperationResultExtended) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkAccountsOperationResultExtended) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkAccountsOperationResultExtended) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *BulkAccountsOperationResultExtended) GetResult() BulkOperationSummary`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BulkAccountsOperationResultExtended) GetResultOk() (*BulkOperationSummary, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BulkAccountsOperationResultExtended) SetResult(v BulkOperationSummary)`

SetResult sets Result field to given value.

### HasResult

`func (o *BulkAccountsOperationResultExtended) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSucceededItems

`func (o *BulkAccountsOperationResultExtended) GetSucceededItems() BulkAccountsList`

GetSucceededItems returns the SucceededItems field if non-nil, zero value otherwise.

### GetSucceededItemsOk

`func (o *BulkAccountsOperationResultExtended) GetSucceededItemsOk() (*BulkAccountsList, bool)`

GetSucceededItemsOk returns a tuple with the SucceededItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededItems

`func (o *BulkAccountsOperationResultExtended) SetSucceededItems(v BulkAccountsList)`

SetSucceededItems sets SucceededItems field to given value.

### HasSucceededItems

`func (o *BulkAccountsOperationResultExtended) HasSucceededItems() bool`

HasSucceededItems returns a boolean if a field has been set.

### GetFailedItems

`func (o *BulkAccountsOperationResultExtended) GetFailedItems() BulkAccountsList`

GetFailedItems returns the FailedItems field if non-nil, zero value otherwise.

### GetFailedItemsOk

`func (o *BulkAccountsOperationResultExtended) GetFailedItemsOk() (*BulkAccountsList, bool)`

GetFailedItemsOk returns a tuple with the FailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItems

`func (o *BulkAccountsOperationResultExtended) SetFailedItems(v BulkAccountsList)`

SetFailedItems sets FailedItems field to given value.

### HasFailedItems

`func (o *BulkAccountsOperationResultExtended) HasFailedItems() bool`

HasFailedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


