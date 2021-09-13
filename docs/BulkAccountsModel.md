# BulkAccountsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountsList** | Pointer to [**[]BulkAccountModel**](BulkAccountModel.md) | The object that contains a list of account object. Each account object contains the parameters for that account. | [optional] 
**Source** | Pointer to **string** | Free text that describes the souce of the bulk account upload. | [optional] 

## Methods

### NewBulkAccountsModel

`func NewBulkAccountsModel() *BulkAccountsModel`

NewBulkAccountsModel instantiates a new BulkAccountsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountsModelWithDefaults

`func NewBulkAccountsModelWithDefaults() *BulkAccountsModel`

NewBulkAccountsModelWithDefaults instantiates a new BulkAccountsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountsList

`func (o *BulkAccountsModel) GetAccountsList() []BulkAccountModel`

GetAccountsList returns the AccountsList field if non-nil, zero value otherwise.

### GetAccountsListOk

`func (o *BulkAccountsModel) GetAccountsListOk() (*[]BulkAccountModel, bool)`

GetAccountsListOk returns a tuple with the AccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsList

`func (o *BulkAccountsModel) SetAccountsList(v []BulkAccountModel)`

SetAccountsList sets AccountsList field to given value.

### HasAccountsList

`func (o *BulkAccountsModel) HasAccountsList() bool`

HasAccountsList returns a boolean if a field has been set.

### GetSource

`func (o *BulkAccountsModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BulkAccountsModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BulkAccountsModel) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BulkAccountsModel) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


