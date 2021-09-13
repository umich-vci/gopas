# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Commands** | Pointer to [**[]Command**](Command.md) |  | [optional] 

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Session) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *Session) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Session) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Session) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Session) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSeverity

`func (o *Session) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Session) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Session) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Session) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCommands

`func (o *Session) GetCommands() []Command`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *Session) GetCommandsOk() (*[]Command, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *Session) SetCommands(v []Command)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *Session) HasCommands() bool`

HasCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


