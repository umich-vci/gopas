# GetUserSSHKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]PublicSSHKeyModel**](PublicSSHKeyModel.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserSSHKeysResponse

`func NewGetUserSSHKeysResponse() *GetUserSSHKeysResponse`

NewGetUserSSHKeysResponse instantiates a new GetUserSSHKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserSSHKeysResponseWithDefaults

`func NewGetUserSSHKeysResponseWithDefaults() *GetUserSSHKeysResponse`

NewGetUserSSHKeysResponseWithDefaults instantiates a new GetUserSSHKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetUserSSHKeysResponse) GetValue() []PublicSSHKeyModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetUserSSHKeysResponse) GetValueOk() (*[]PublicSSHKeyModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetUserSSHKeysResponse) SetValue(v []PublicSSHKeyModel)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetUserSSHKeysResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCount

`func (o *GetUserSSHKeysResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetUserSSHKeysResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetUserSSHKeysResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetUserSSHKeysResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextLink

`func (o *GetUserSSHKeysResponse) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *GetUserSSHKeysResponse) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *GetUserSSHKeysResponse) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *GetUserSSHKeysResponse) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


