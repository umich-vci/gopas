# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Command) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Command) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Command) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Command) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOffset

`func (o *Command) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Command) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Command) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Command) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetScore

`func (o *Command) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Command) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Command) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Command) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSeverity

`func (o *Command) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Command) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Command) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Command) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


