# NewRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account to access. | 
**Reason** | Pointer to **string** | The reason why the account needs to be accessed. | [optional] 
**TicketingSystemName** | Pointer to **string** | The name of the Ticketing System specified in the request. | [optional] 
**TicketId** | Pointer to **string** | The ticket ID given by the ticketing system. | [optional] 
**MultipleAccessRequired** | Pointer to **bool** | Whether or not the request is for multiple access. | [optional] 
**FromDate** | Pointer to **int64** | If the request is for a timeframe, the time from when the user wants to access the account, in Unix time. | [optional] 
**ToDate** | Pointer to **int64** | If the request is for a timeframe, the time until when the user wants to access the account, in Unix time. | [optional] 
**AdditionalInfo** | Pointer to **map[string]string** | Additional information included in the request. A list of values that are predefined in configuration. | [optional] 
**UseConnect** | Pointer to **bool** | Whether or not the request is for connection through the PSM. | [optional] 
**ConnectionComponent** | Pointer to **string** | If the connection is through PSM, the name of the connection component to connect with, as defined in the configuration. | [optional] 
**ConnectionParams** | Pointer to [**map[string]ConnectionParameter**](ConnectionParameter.md) | A list of parameters required to perform the connection, as defined in each connection component configuration. These parameters are listed in the table below.  Each item contains: key,   { value - The content of the parameter,   ShouldSave - Whether or not this value will be saved with the account for future attempts to connect to the remote machine.} | [optional] 

## Methods

### NewNewRequestBody

`func NewNewRequestBody(accountId string, ) *NewRequestBody`

NewNewRequestBody instantiates a new NewRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRequestBodyWithDefaults

`func NewNewRequestBodyWithDefaults() *NewRequestBody`

NewNewRequestBodyWithDefaults instantiates a new NewRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NewRequestBody) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NewRequestBody) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NewRequestBody) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetReason

`func (o *NewRequestBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NewRequestBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NewRequestBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NewRequestBody) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTicketingSystemName

`func (o *NewRequestBody) GetTicketingSystemName() string`

GetTicketingSystemName returns the TicketingSystemName field if non-nil, zero value otherwise.

### GetTicketingSystemNameOk

`func (o *NewRequestBody) GetTicketingSystemNameOk() (*string, bool)`

GetTicketingSystemNameOk returns a tuple with the TicketingSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystemName

`func (o *NewRequestBody) SetTicketingSystemName(v string)`

SetTicketingSystemName sets TicketingSystemName field to given value.

### HasTicketingSystemName

`func (o *NewRequestBody) HasTicketingSystemName() bool`

HasTicketingSystemName returns a boolean if a field has been set.

### GetTicketId

`func (o *NewRequestBody) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *NewRequestBody) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *NewRequestBody) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *NewRequestBody) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetMultipleAccessRequired

`func (o *NewRequestBody) GetMultipleAccessRequired() bool`

GetMultipleAccessRequired returns the MultipleAccessRequired field if non-nil, zero value otherwise.

### GetMultipleAccessRequiredOk

`func (o *NewRequestBody) GetMultipleAccessRequiredOk() (*bool, bool)`

GetMultipleAccessRequiredOk returns a tuple with the MultipleAccessRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccessRequired

`func (o *NewRequestBody) SetMultipleAccessRequired(v bool)`

SetMultipleAccessRequired sets MultipleAccessRequired field to given value.

### HasMultipleAccessRequired

`func (o *NewRequestBody) HasMultipleAccessRequired() bool`

HasMultipleAccessRequired returns a boolean if a field has been set.

### GetFromDate

`func (o *NewRequestBody) GetFromDate() int64`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *NewRequestBody) GetFromDateOk() (*int64, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *NewRequestBody) SetFromDate(v int64)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *NewRequestBody) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *NewRequestBody) GetToDate() int64`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *NewRequestBody) GetToDateOk() (*int64, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *NewRequestBody) SetToDate(v int64)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *NewRequestBody) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *NewRequestBody) GetAdditionalInfo() map[string]string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *NewRequestBody) GetAdditionalInfoOk() (*map[string]string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *NewRequestBody) SetAdditionalInfo(v map[string]string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *NewRequestBody) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetUseConnect

`func (o *NewRequestBody) GetUseConnect() bool`

GetUseConnect returns the UseConnect field if non-nil, zero value otherwise.

### GetUseConnectOk

`func (o *NewRequestBody) GetUseConnectOk() (*bool, bool)`

GetUseConnectOk returns a tuple with the UseConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseConnect

`func (o *NewRequestBody) SetUseConnect(v bool)`

SetUseConnect sets UseConnect field to given value.

### HasUseConnect

`func (o *NewRequestBody) HasUseConnect() bool`

HasUseConnect returns a boolean if a field has been set.

### GetConnectionComponent

`func (o *NewRequestBody) GetConnectionComponent() string`

GetConnectionComponent returns the ConnectionComponent field if non-nil, zero value otherwise.

### GetConnectionComponentOk

`func (o *NewRequestBody) GetConnectionComponentOk() (*string, bool)`

GetConnectionComponentOk returns a tuple with the ConnectionComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionComponent

`func (o *NewRequestBody) SetConnectionComponent(v string)`

SetConnectionComponent sets ConnectionComponent field to given value.

### HasConnectionComponent

`func (o *NewRequestBody) HasConnectionComponent() bool`

HasConnectionComponent returns a boolean if a field has been set.

### GetConnectionParams

`func (o *NewRequestBody) GetConnectionParams() map[string]ConnectionParameter`

GetConnectionParams returns the ConnectionParams field if non-nil, zero value otherwise.

### GetConnectionParamsOk

`func (o *NewRequestBody) GetConnectionParamsOk() (*map[string]ConnectionParameter, bool)`

GetConnectionParamsOk returns a tuple with the ConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionParams

`func (o *NewRequestBody) SetConnectionParams(v map[string]ConnectionParameter)`

SetConnectionParams sets ConnectionParams field to given value.

### HasConnectionParams

`func (o *NewRequestBody) HasConnectionParams() bool`

HasConnectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


