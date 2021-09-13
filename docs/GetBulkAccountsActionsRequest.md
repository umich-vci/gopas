# GetBulkAccountsActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of returned bulk account actions. If not specified, the default value is 10. The maximum number that can be specified is 1000. | [optional] 
**Filter** | Pointer to **string** | Search for bulk account actions filtered by spesific status. e.g. Filter&#x3D;status eq InProgress | [optional] 

## Methods

### NewGetBulkAccountsActionsRequest

`func NewGetBulkAccountsActionsRequest() *GetBulkAccountsActionsRequest`

NewGetBulkAccountsActionsRequest instantiates a new GetBulkAccountsActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBulkAccountsActionsRequestWithDefaults

`func NewGetBulkAccountsActionsRequestWithDefaults() *GetBulkAccountsActionsRequest`

NewGetBulkAccountsActionsRequestWithDefaults instantiates a new GetBulkAccountsActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetBulkAccountsActionsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBulkAccountsActionsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBulkAccountsActionsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetBulkAccountsActionsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *GetBulkAccountsActionsRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetBulkAccountsActionsRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetBulkAccountsActionsRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetBulkAccountsActionsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


