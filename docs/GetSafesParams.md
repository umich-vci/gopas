# GetSafesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAccounts** | Pointer to **bool** | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be false. | [optional] 
**ExtendedDetails** | Pointer to **bool** | Whether or not to return all Safe data or only SafeName as part of the response. If not sent, the value will be true. | [optional] 

## Methods

### NewGetSafesParams

`func NewGetSafesParams() *GetSafesParams`

NewGetSafesParams instantiates a new GetSafesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSafesParamsWithDefaults

`func NewGetSafesParamsWithDefaults() *GetSafesParams`

NewGetSafesParamsWithDefaults instantiates a new GetSafesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAccounts

`func (o *GetSafesParams) GetIncludeAccounts() bool`

GetIncludeAccounts returns the IncludeAccounts field if non-nil, zero value otherwise.

### GetIncludeAccountsOk

`func (o *GetSafesParams) GetIncludeAccountsOk() (*bool, bool)`

GetIncludeAccountsOk returns a tuple with the IncludeAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAccounts

`func (o *GetSafesParams) SetIncludeAccounts(v bool)`

SetIncludeAccounts sets IncludeAccounts field to given value.

### HasIncludeAccounts

`func (o *GetSafesParams) HasIncludeAccounts() bool`

HasIncludeAccounts returns a boolean if a field has been set.

### GetExtendedDetails

`func (o *GetSafesParams) GetExtendedDetails() bool`

GetExtendedDetails returns the ExtendedDetails field if non-nil, zero value otherwise.

### GetExtendedDetailsOk

`func (o *GetSafesParams) GetExtendedDetailsOk() (*bool, bool)`

GetExtendedDetailsOk returns a tuple with the ExtendedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedDetails

`func (o *GetSafesParams) SetExtendedDetails(v bool)`

SetExtendedDetails sets ExtendedDetails field to given value.

### HasExtendedDetails

`func (o *GetSafesParams) HasExtendedDetails() bool`

HasExtendedDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


