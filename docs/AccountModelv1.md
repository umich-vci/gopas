# AccountModelv1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of accounts found by the search. | [optional] 
**Accounts** | Pointer to [**[]AccountModelv1Accounts**](AccountModelv1Accounts.md) |  | [optional] 

## Methods

### NewAccountModelv1

`func NewAccountModelv1() *AccountModelv1`

NewAccountModelv1 instantiates a new AccountModelv1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountModelv1WithDefaults

`func NewAccountModelv1WithDefaults() *AccountModelv1`

NewAccountModelv1WithDefaults instantiates a new AccountModelv1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AccountModelv1) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccountModelv1) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccountModelv1) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AccountModelv1) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAccounts

`func (o *AccountModelv1) GetAccounts() []AccountModelv1Accounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountModelv1) GetAccountsOk() (*[]AccountModelv1Accounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountModelv1) SetAccounts(v []AccountModelv1Accounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountModelv1) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


