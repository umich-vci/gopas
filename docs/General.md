# General

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The platform unique ID | [optional] 
**Name** | Pointer to **string** | The platform name | [optional] 
**SystemType** | Pointer to **string** | The type of system to which the platform is applied | [optional] 
**Active** | Pointer to **bool** | Indicates whether a platform is active or inactive. valid values: true\\false | [optional] 
**Description** | Pointer to **string** | The platform description | [optional] 
**PlatformBaseID** | Pointer to **string** | The ID of the default platform that this platform is based on (duplicated from) | [optional] 
**PlatformType** | Pointer to **string** | Indicates if the platform is a group platform. valid values: Group\\Regular | [optional] 

## Methods

### NewGeneral

`func NewGeneral() *General`

NewGeneral instantiates a new General object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralWithDefaults

`func NewGeneralWithDefaults() *General`

NewGeneralWithDefaults instantiates a new General object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *General) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *General) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *General) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *General) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *General) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *General) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *General) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *General) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSystemType

`func (o *General) GetSystemType() string`

GetSystemType returns the SystemType field if non-nil, zero value otherwise.

### GetSystemTypeOk

`func (o *General) GetSystemTypeOk() (*string, bool)`

GetSystemTypeOk returns a tuple with the SystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemType

`func (o *General) SetSystemType(v string)`

SetSystemType sets SystemType field to given value.

### HasSystemType

`func (o *General) HasSystemType() bool`

HasSystemType returns a boolean if a field has been set.

### GetActive

`func (o *General) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *General) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *General) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *General) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *General) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *General) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *General) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *General) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatformBaseID

`func (o *General) GetPlatformBaseID() string`

GetPlatformBaseID returns the PlatformBaseID field if non-nil, zero value otherwise.

### GetPlatformBaseIDOk

`func (o *General) GetPlatformBaseIDOk() (*string, bool)`

GetPlatformBaseIDOk returns a tuple with the PlatformBaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBaseID

`func (o *General) SetPlatformBaseID(v string)`

SetPlatformBaseID sets PlatformBaseID field to given value.

### HasPlatformBaseID

`func (o *General) HasPlatformBaseID() bool`

HasPlatformBaseID returns a boolean if a field has been set.

### GetPlatformType

`func (o *General) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *General) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *General) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *General) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


