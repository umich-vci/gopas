# Properties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to [**[]Identity**](Identity.md) | A list of all required properties defined for the platform. Each property includes a name and a display name | [optional] 
**Optional** | Pointer to [**[]Identity**](Identity.md) | A list of all optional properties defined for the platform. Each property includes a name and a display name | [optional] 

## Methods

### NewProperties

`func NewProperties() *Properties`

NewProperties instantiates a new Properties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesWithDefaults

`func NewPropertiesWithDefaults() *Properties`

NewPropertiesWithDefaults instantiates a new Properties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *Properties) GetRequired() []Identity`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Properties) GetRequiredOk() (*[]Identity, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Properties) SetRequired(v []Identity)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Properties) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOptional

`func (o *Properties) GetOptional() []Identity`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *Properties) GetOptionalOk() (*[]Identity, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *Properties) SetOptional(v []Identity)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *Properties) HasOptional() bool`

HasOptional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


