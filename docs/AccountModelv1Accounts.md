# AccountModelv1Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **string** | The account ID of the account | [optional] 
**InternalProperties** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) |  | [optional] 
**Properties** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) |  | [optional] 

## Methods

### NewAccountModelv1Accounts

`func NewAccountModelv1Accounts() *AccountModelv1Accounts`

NewAccountModelv1Accounts instantiates a new AccountModelv1Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountModelv1AccountsWithDefaults

`func NewAccountModelv1AccountsWithDefaults() *AccountModelv1Accounts`

NewAccountModelv1AccountsWithDefaults instantiates a new AccountModelv1Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *AccountModelv1Accounts) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *AccountModelv1Accounts) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *AccountModelv1Accounts) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *AccountModelv1Accounts) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetInternalProperties

`func (o *AccountModelv1Accounts) GetInternalProperties() []KeyValuePair`

GetInternalProperties returns the InternalProperties field if non-nil, zero value otherwise.

### GetInternalPropertiesOk

`func (o *AccountModelv1Accounts) GetInternalPropertiesOk() (*[]KeyValuePair, bool)`

GetInternalPropertiesOk returns a tuple with the InternalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalProperties

`func (o *AccountModelv1Accounts) SetInternalProperties(v []KeyValuePair)`

SetInternalProperties sets InternalProperties field to given value.

### HasInternalProperties

`func (o *AccountModelv1Accounts) HasInternalProperties() bool`

HasInternalProperties returns a boolean if a field has been set.

### GetProperties

`func (o *AccountModelv1Accounts) GetProperties() []KeyValuePair`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AccountModelv1Accounts) GetPropertiesOk() (*[]KeyValuePair, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AccountModelv1Accounts) SetProperties(v []KeyValuePair)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AccountModelv1Accounts) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


