# PrivateSSHKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **int32** | The key exported format | [optional] 
**PrivateKey** | Pointer to **string** | The private key | [optional] 
**KeyAlg** | Pointer to **int32** | Key algorithm | [optional] 

## Methods

### NewPrivateSSHKeyModel

`func NewPrivateSSHKeyModel() *PrivateSSHKeyModel`

NewPrivateSSHKeyModel instantiates a new PrivateSSHKeyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSSHKeyModelWithDefaults

`func NewPrivateSSHKeyModelWithDefaults() *PrivateSSHKeyModel`

NewPrivateSSHKeyModelWithDefaults instantiates a new PrivateSSHKeyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *PrivateSSHKeyModel) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PrivateSSHKeyModel) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PrivateSSHKeyModel) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PrivateSSHKeyModel) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPrivateKey

`func (o *PrivateSSHKeyModel) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PrivateSSHKeyModel) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PrivateSSHKeyModel) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *PrivateSSHKeyModel) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetKeyAlg

`func (o *PrivateSSHKeyModel) GetKeyAlg() int32`

GetKeyAlg returns the KeyAlg field if non-nil, zero value otherwise.

### GetKeyAlgOk

`func (o *PrivateSSHKeyModel) GetKeyAlgOk() (*int32, bool)`

GetKeyAlgOk returns a tuple with the KeyAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlg

`func (o *PrivateSSHKeyModel) SetKeyAlg(v int32)`

SetKeyAlg sets KeyAlg field to given value.

### HasKeyAlg

`func (o *PrivateSSHKeyModel) HasKeyAlg() bool`

HasKeyAlg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


