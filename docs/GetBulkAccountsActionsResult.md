# GetBulkAccountsActionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] [readonly] 
**BulkActions** | Pointer to [**[]BulkAccountsOperationResult**](BulkAccountsOperationResult.md) |  | [optional] 

## Methods

### NewGetBulkAccountsActionsResult

`func NewGetBulkAccountsActionsResult() *GetBulkAccountsActionsResult`

NewGetBulkAccountsActionsResult instantiates a new GetBulkAccountsActionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBulkAccountsActionsResultWithDefaults

`func NewGetBulkAccountsActionsResultWithDefaults() *GetBulkAccountsActionsResult`

NewGetBulkAccountsActionsResultWithDefaults instantiates a new GetBulkAccountsActionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetBulkAccountsActionsResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBulkAccountsActionsResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBulkAccountsActionsResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBulkAccountsActionsResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBulkActions

`func (o *GetBulkAccountsActionsResult) GetBulkActions() []BulkAccountsOperationResult`

GetBulkActions returns the BulkActions field if non-nil, zero value otherwise.

### GetBulkActionsOk

`func (o *GetBulkAccountsActionsResult) GetBulkActionsOk() (*[]BulkAccountsOperationResult, bool)`

GetBulkActionsOk returns a tuple with the BulkActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkActions

`func (o *GetBulkAccountsActionsResult) SetBulkActions(v []BulkAccountsOperationResult)`

SetBulkActions sets BulkActions field to given value.

### HasBulkActions

`func (o *GetBulkAccountsActionsResult) HasBulkActions() bool`

HasBulkActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


