# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountID** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountID

`func (o *Account) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Account) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Account) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.

### HasAccountID

`func (o *Account) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### GetProperties

`func (o *Account) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Account) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Account) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Account) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


