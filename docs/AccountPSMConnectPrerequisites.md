# AccountPSMConnectPrerequisites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Required reason to request the access | [optional] 
**TicketingSystemName** | Pointer to **string** | The name of the Ticketing System used in the request | [optional] 
**TicketId** | Pointer to **string** | The ticket ID of the ticketing system | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**ConnectionComponent** | **string** | The name of the connection component to connect with as defined in the configuration | 
**ConnectionParams** | Pointer to [**map[string]ConnectionParameter**](ConnectionParameter.md) | A list of parameters required to perform the connection. Defined in each connection component configuration  Each item contains: key,   { value - The content of the parameter,   ShouldSave - Whether or not this value will be saved with the account for future attempts to connect to the remote machine.} | [optional] 

## Methods

### NewAccountPSMConnectPrerequisites

`func NewAccountPSMConnectPrerequisites(connectionComponent string, ) *AccountPSMConnectPrerequisites`

NewAccountPSMConnectPrerequisites instantiates a new AccountPSMConnectPrerequisites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPSMConnectPrerequisitesWithDefaults

`func NewAccountPSMConnectPrerequisitesWithDefaults() *AccountPSMConnectPrerequisites`

NewAccountPSMConnectPrerequisitesWithDefaults instantiates a new AccountPSMConnectPrerequisites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AccountPSMConnectPrerequisites) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountPSMConnectPrerequisites) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountPSMConnectPrerequisites) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccountPSMConnectPrerequisites) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTicketingSystemName

`func (o *AccountPSMConnectPrerequisites) GetTicketingSystemName() string`

GetTicketingSystemName returns the TicketingSystemName field if non-nil, zero value otherwise.

### GetTicketingSystemNameOk

`func (o *AccountPSMConnectPrerequisites) GetTicketingSystemNameOk() (*string, bool)`

GetTicketingSystemNameOk returns a tuple with the TicketingSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystemName

`func (o *AccountPSMConnectPrerequisites) SetTicketingSystemName(v string)`

SetTicketingSystemName sets TicketingSystemName field to given value.

### HasTicketingSystemName

`func (o *AccountPSMConnectPrerequisites) HasTicketingSystemName() bool`

HasTicketingSystemName returns a boolean if a field has been set.

### GetTicketId

`func (o *AccountPSMConnectPrerequisites) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *AccountPSMConnectPrerequisites) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *AccountPSMConnectPrerequisites) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *AccountPSMConnectPrerequisites) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetConnectionType

`func (o *AccountPSMConnectPrerequisites) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AccountPSMConnectPrerequisites) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AccountPSMConnectPrerequisites) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *AccountPSMConnectPrerequisites) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectionComponent

`func (o *AccountPSMConnectPrerequisites) GetConnectionComponent() string`

GetConnectionComponent returns the ConnectionComponent field if non-nil, zero value otherwise.

### GetConnectionComponentOk

`func (o *AccountPSMConnectPrerequisites) GetConnectionComponentOk() (*string, bool)`

GetConnectionComponentOk returns a tuple with the ConnectionComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionComponent

`func (o *AccountPSMConnectPrerequisites) SetConnectionComponent(v string)`

SetConnectionComponent sets ConnectionComponent field to given value.


### GetConnectionParams

`func (o *AccountPSMConnectPrerequisites) GetConnectionParams() map[string]ConnectionParameter`

GetConnectionParams returns the ConnectionParams field if non-nil, zero value otherwise.

### GetConnectionParamsOk

`func (o *AccountPSMConnectPrerequisites) GetConnectionParamsOk() (*map[string]ConnectionParameter, bool)`

GetConnectionParamsOk returns a tuple with the ConnectionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionParams

`func (o *AccountPSMConnectPrerequisites) SetConnectionParams(v map[string]ConnectionParameter)`

SetConnectionParams sets ConnectionParams field to given value.

### HasConnectionParams

`func (o *AccountPSMConnectPrerequisites) HasConnectionParams() bool`

HasConnectionParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


