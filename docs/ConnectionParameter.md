# ConnectionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**ShouldSave** | Pointer to **bool** |  | [optional] 

## Methods

### NewConnectionParameter

`func NewConnectionParameter() *ConnectionParameter`

NewConnectionParameter instantiates a new ConnectionParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionParameterWithDefaults

`func NewConnectionParameterWithDefaults() *ConnectionParameter`

NewConnectionParameterWithDefaults instantiates a new ConnectionParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ConnectionParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConnectionParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConnectionParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConnectionParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetShouldSave

`func (o *ConnectionParameter) GetShouldSave() bool`

GetShouldSave returns the ShouldSave field if non-nil, zero value otherwise.

### GetShouldSaveOk

`func (o *ConnectionParameter) GetShouldSaveOk() (*bool, bool)`

GetShouldSaveOk returns a tuple with the ShouldSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSave

`func (o *ConnectionParameter) SetShouldSave(v bool)`

SetShouldSave sets ShouldSave field to given value.

### HasShouldSave

`func (o *ConnectionParameter) HasShouldSave() bool`

HasShouldSave returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


