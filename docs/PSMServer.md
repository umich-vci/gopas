# PSMServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | PSM server Id. | [optional] 
**Name** | Pointer to **string** | PSM server Name. | [optional] 
**Address** | Pointer to **string** | PSM server address. | [optional] 

## Methods

### NewPSMServer

`func NewPSMServer() *PSMServer`

NewPSMServer instantiates a new PSMServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPSMServerWithDefaults

`func NewPSMServerWithDefaults() *PSMServer`

NewPSMServerWithDefaults instantiates a new PSMServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *PSMServer) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *PSMServer) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *PSMServer) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *PSMServer) HasID() bool`

HasID returns a boolean if a field has been set.

### GetName

`func (o *PSMServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PSMServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PSMServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PSMServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *PSMServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PSMServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PSMServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PSMServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


