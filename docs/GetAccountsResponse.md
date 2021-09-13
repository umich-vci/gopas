# GetAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]AccountModel**](AccountModel.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAccountsResponse

`func NewGetAccountsResponse() *GetAccountsResponse`

NewGetAccountsResponse instantiates a new GetAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsResponseWithDefaults

`func NewGetAccountsResponseWithDefaults() *GetAccountsResponse`

NewGetAccountsResponseWithDefaults instantiates a new GetAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetAccountsResponse) GetValue() []AccountModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetAccountsResponse) GetValueOk() (*[]AccountModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetAccountsResponse) SetValue(v []AccountModel)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetAccountsResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCount

`func (o *GetAccountsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetAccountsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetAccountsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetAccountsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextLink

`func (o *GetAccountsResponse) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *GetAccountsResponse) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *GetAccountsResponse) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *GetAccountsResponse) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


