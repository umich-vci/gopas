# GetAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | List of keywords separated with space to search in accounts. | [optional] 
**Sort** | Pointer to **[]string** | Property or properties to sort returned accounts, followed by asc (default) or desc to control sort direction. Multiple sorts are comma-separated. Maximum number of properties is 3. | [optional] 
**Offset** | Pointer to **int32** | Offset of the first returned account into the collection of results. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000. | [optional] 
**Filter** | Pointer to **string** | Search for accounts filtered by specific Safe. Filter&#x3D;safename eq [safe name] | [optional] 
**SearchType** | Pointer to **string** | Type of search to perform - if keywords are contained in the relevant account properties values or in the start of the properties values (the later enhances performance) | [optional] 

## Methods

### NewGetAccountsRequest

`func NewGetAccountsRequest() *GetAccountsRequest`

NewGetAccountsRequest instantiates a new GetAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountsRequestWithDefaults

`func NewGetAccountsRequestWithDefaults() *GetAccountsRequest`

NewGetAccountsRequestWithDefaults instantiates a new GetAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *GetAccountsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *GetAccountsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *GetAccountsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *GetAccountsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSort

`func (o *GetAccountsRequest) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetAccountsRequest) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetAccountsRequest) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetAccountsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetOffset

`func (o *GetAccountsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetAccountsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetAccountsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetAccountsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetAccountsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetAccountsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetAccountsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetAccountsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *GetAccountsRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetAccountsRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetAccountsRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetAccountsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearchType

`func (o *GetAccountsRequest) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *GetAccountsRequest) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *GetAccountsRequest) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *GetAccountsRequest) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


