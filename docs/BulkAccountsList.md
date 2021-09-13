# BulkAccountsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] [readonly] 
**Items** | Pointer to [**[]BulkAccountModel**](BulkAccountModel.md) |  | [optional] 

## Methods

### NewBulkAccountsList

`func NewBulkAccountsList() *BulkAccountsList`

NewBulkAccountsList instantiates a new BulkAccountsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountsListWithDefaults

`func NewBulkAccountsListWithDefaults() *BulkAccountsList`

NewBulkAccountsListWithDefaults instantiates a new BulkAccountsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BulkAccountsList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkAccountsList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkAccountsList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BulkAccountsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *BulkAccountsList) GetItems() []BulkAccountModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkAccountsList) GetItemsOk() (*[]BulkAccountModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkAccountsList) SetItems(v []BulkAccountModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkAccountsList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


