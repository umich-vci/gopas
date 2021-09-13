# GetDiscoveredAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Search** | Pointer to **string** | Searches according to REST standard (search&#x3D;\&quot;search word\&quot;).    Search is supported for userName and address. | [optional] 
**SearchType** | Pointer to **string** | The type of search to perform.     The keyword can either be contained within the account property values, or at the beginning of the value  specified in the search parameter.    When using a keyword at the beginning of a value, performance is enhanced.    Valid values: contains (default) or startswith | [optional] 
**Offset** | Pointer to **int32** | The offset of the first returned accounts into the list of results. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of returned accounts.     If not specified, the server limits the results to 100.    The maximum number that can be specified is 1000. | [optional] 
**Filter** | Pointer to **string** | Search for accounts filtered by:    • platformType - Get discovered accounts of a specific platform, using the platform type name.        Usage: filter&#x3D;platformType eq Windows Server Local            Type: string            Valid values:                o Windows Server Local                    o Windows Desktop Local                    o Windows Domain                    o Unix                    o Unix SSH Key                    o AWS                    o AWS Access Keys                    o Azure Password Management        • privileged - Get only privileged or non-privileged discovered accounts.        Usage: filter&#x3D;privileged eq true            Type: boolean            Valid values: true/false        • accountEnabled - Get only enabled or disabled discovered accounts.           Usage: filter&#x3D;accountEnabled eq true            Type: boolean            Valid values: true/false                  Note: To use more than one filter, you can use the AND operator.         filter&#x3D;platformType eq Windows Server Local AND privileged eq true | [optional] 

## Methods

### NewGetDiscoveredAccountsRequest

`func NewGetDiscoveredAccountsRequest() *GetDiscoveredAccountsRequest`

NewGetDiscoveredAccountsRequest instantiates a new GetDiscoveredAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveredAccountsRequestWithDefaults

`func NewGetDiscoveredAccountsRequestWithDefaults() *GetDiscoveredAccountsRequest`

NewGetDiscoveredAccountsRequestWithDefaults instantiates a new GetDiscoveredAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearch

`func (o *GetDiscoveredAccountsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *GetDiscoveredAccountsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *GetDiscoveredAccountsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *GetDiscoveredAccountsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSearchType

`func (o *GetDiscoveredAccountsRequest) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *GetDiscoveredAccountsRequest) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *GetDiscoveredAccountsRequest) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *GetDiscoveredAccountsRequest) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetOffset

`func (o *GetDiscoveredAccountsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetDiscoveredAccountsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetDiscoveredAccountsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetDiscoveredAccountsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetDiscoveredAccountsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetDiscoveredAccountsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetDiscoveredAccountsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetDiscoveredAccountsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *GetDiscoveredAccountsRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetDiscoveredAccountsRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetDiscoveredAccountsRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetDiscoveredAccountsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


