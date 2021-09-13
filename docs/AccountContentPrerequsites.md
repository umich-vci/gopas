# AccountContentPrerequsites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason that is required to retrieve the password/SSH key | [optional] 
**TicketId** | Pointer to **string** | The ticket ID of the ticketing system | [optional] 
**TicketingSystem** | Pointer to **string** | The name of the ticketing system | [optional] 
**IsUse** | Pointer to **bool** | Internal parameter (for PSMP only) | [optional] 
**ActionType** | Pointer to **string** | The action this password/SSH key will be used for. (Show/Copy/Connect for password; Retrieve for SSH key) | [optional] 
**Machine** | Pointer to **string** | The address of the remote machine the user wants to connect to using the password/SSH key | [optional] 
**Version** | Pointer to **int64** | The version number of the required password/SSH key. Must be a positive number. If the value is left empty or the value passed does not exist, then the current password/SSH key version is returned | [optional] 

## Methods

### NewAccountContentPrerequsites

`func NewAccountContentPrerequsites() *AccountContentPrerequsites`

NewAccountContentPrerequsites instantiates a new AccountContentPrerequsites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountContentPrerequsitesWithDefaults

`func NewAccountContentPrerequsitesWithDefaults() *AccountContentPrerequsites`

NewAccountContentPrerequsitesWithDefaults instantiates a new AccountContentPrerequsites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AccountContentPrerequsites) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountContentPrerequsites) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountContentPrerequsites) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccountContentPrerequsites) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTicketId

`func (o *AccountContentPrerequsites) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *AccountContentPrerequsites) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *AccountContentPrerequsites) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *AccountContentPrerequsites) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetTicketingSystem

`func (o *AccountContentPrerequsites) GetTicketingSystem() string`

GetTicketingSystem returns the TicketingSystem field if non-nil, zero value otherwise.

### GetTicketingSystemOk

`func (o *AccountContentPrerequsites) GetTicketingSystemOk() (*string, bool)`

GetTicketingSystemOk returns a tuple with the TicketingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystem

`func (o *AccountContentPrerequsites) SetTicketingSystem(v string)`

SetTicketingSystem sets TicketingSystem field to given value.

### HasTicketingSystem

`func (o *AccountContentPrerequsites) HasTicketingSystem() bool`

HasTicketingSystem returns a boolean if a field has been set.

### GetIsUse

`func (o *AccountContentPrerequsites) GetIsUse() bool`

GetIsUse returns the IsUse field if non-nil, zero value otherwise.

### GetIsUseOk

`func (o *AccountContentPrerequsites) GetIsUseOk() (*bool, bool)`

GetIsUseOk returns a tuple with the IsUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUse

`func (o *AccountContentPrerequsites) SetIsUse(v bool)`

SetIsUse sets IsUse field to given value.

### HasIsUse

`func (o *AccountContentPrerequsites) HasIsUse() bool`

HasIsUse returns a boolean if a field has been set.

### GetActionType

`func (o *AccountContentPrerequsites) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AccountContentPrerequsites) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AccountContentPrerequsites) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *AccountContentPrerequsites) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetMachine

`func (o *AccountContentPrerequsites) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *AccountContentPrerequsites) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *AccountContentPrerequsites) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *AccountContentPrerequsites) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetVersion

`func (o *AccountContentPrerequsites) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AccountContentPrerequsites) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AccountContentPrerequsites) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AccountContentPrerequsites) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


