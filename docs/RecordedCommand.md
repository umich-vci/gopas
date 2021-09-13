# RecordedCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityText** | Pointer to **string** |  | [optional] 
**ActivityType** | Pointer to **int32** |  | [optional] 
**ActivityId** | Pointer to **string** |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Offsets** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRecordedCommand

`func NewRecordedCommand() *RecordedCommand`

NewRecordedCommand instantiates a new RecordedCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordedCommandWithDefaults

`func NewRecordedCommandWithDefaults() *RecordedCommand`

NewRecordedCommandWithDefaults instantiates a new RecordedCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityText

`func (o *RecordedCommand) GetActivityText() string`

GetActivityText returns the ActivityText field if non-nil, zero value otherwise.

### GetActivityTextOk

`func (o *RecordedCommand) GetActivityTextOk() (*string, bool)`

GetActivityTextOk returns a tuple with the ActivityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityText

`func (o *RecordedCommand) SetActivityText(v string)`

SetActivityText sets ActivityText field to given value.

### HasActivityText

`func (o *RecordedCommand) HasActivityText() bool`

HasActivityText returns a boolean if a field has been set.

### GetActivityType

`func (o *RecordedCommand) GetActivityType() int32`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *RecordedCommand) GetActivityTypeOk() (*int32, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *RecordedCommand) SetActivityType(v int32)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *RecordedCommand) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetActivityId

`func (o *RecordedCommand) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *RecordedCommand) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *RecordedCommand) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *RecordedCommand) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetFormats

`func (o *RecordedCommand) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *RecordedCommand) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *RecordedCommand) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *RecordedCommand) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### GetOffsets

`func (o *RecordedCommand) GetOffsets() map[string]string`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *RecordedCommand) GetOffsetsOk() (*map[string]string, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *RecordedCommand) SetOffsets(v map[string]string)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *RecordedCommand) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


