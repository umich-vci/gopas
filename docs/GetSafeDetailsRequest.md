# GetSafeDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAccounts** | Pointer to **bool** | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be False. | [optional] 
**UseCache** | Pointer to **bool** | Whether to retrieve from session or not. | [optional] 

## Methods

### NewGetSafeDetailsRequest

`func NewGetSafeDetailsRequest() *GetSafeDetailsRequest`

NewGetSafeDetailsRequest instantiates a new GetSafeDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSafeDetailsRequestWithDefaults

`func NewGetSafeDetailsRequestWithDefaults() *GetSafeDetailsRequest`

NewGetSafeDetailsRequestWithDefaults instantiates a new GetSafeDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAccounts

`func (o *GetSafeDetailsRequest) GetIncludeAccounts() bool`

GetIncludeAccounts returns the IncludeAccounts field if non-nil, zero value otherwise.

### GetIncludeAccountsOk

`func (o *GetSafeDetailsRequest) GetIncludeAccountsOk() (*bool, bool)`

GetIncludeAccountsOk returns a tuple with the IncludeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAccounts

`func (o *GetSafeDetailsRequest) SetIncludeAccounts(v bool)`

SetIncludeAccounts sets IncludeAccounts field to given value.

### HasIncludeAccounts

`func (o *GetSafeDetailsRequest) HasIncludeAccounts() bool`

HasIncludeAccounts returns a boolean if a field has been set.

### GetUseCache

`func (o *GetSafeDetailsRequest) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *GetSafeDetailsRequest) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *GetSafeDetailsRequest) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *GetSafeDetailsRequest) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


