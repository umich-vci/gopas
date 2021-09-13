# GetDiscoveredAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]GetDiscoveredAccount**](GetDiscoveredAccount.md) |  | [optional] 
**Count** | Pointer to **int32** | Total number of results across all pages | [optional] 
**NextLink** | Pointer to **string** | An opaque URL to the next page of results.   Should be present only if requested page size (limit) is not specified and there are more results than a single page as defined by the server.  The last page shouldn&#39;t have &#39;nextLink&#39; in the response.   If the limit in the request is too high, an error is returned | [optional] 

## Methods

### NewGetDiscoveredAccountsResponse

`func NewGetDiscoveredAccountsResponse() *GetDiscoveredAccountsResponse`

NewGetDiscoveredAccountsResponse instantiates a new GetDiscoveredAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveredAccountsResponseWithDefaults

`func NewGetDiscoveredAccountsResponseWithDefaults() *GetDiscoveredAccountsResponse`

NewGetDiscoveredAccountsResponseWithDefaults instantiates a new GetDiscoveredAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetDiscoveredAccountsResponse) GetValue() []GetDiscoveredAccount`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetDiscoveredAccountsResponse) GetValueOk() (*[]GetDiscoveredAccount, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetDiscoveredAccountsResponse) SetValue(v []GetDiscoveredAccount)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetDiscoveredAccountsResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCount

`func (o *GetDiscoveredAccountsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDiscoveredAccountsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDiscoveredAccountsResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetDiscoveredAccountsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextLink

`func (o *GetDiscoveredAccountsResponse) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *GetDiscoveredAccountsResponse) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *GetDiscoveredAccountsResponse) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *GetDiscoveredAccountsResponse) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


