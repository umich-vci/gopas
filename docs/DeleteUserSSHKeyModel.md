# DeleteUserSSHKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyFingerprint** | **string** | The key fingerprint to delete | 
**KeyStoreTypeFilter** | Pointer to **int32** | The key type filter | [optional] 

## Methods

### NewDeleteUserSSHKeyModel

`func NewDeleteUserSSHKeyModel(keyFingerprint string, ) *DeleteUserSSHKeyModel`

NewDeleteUserSSHKeyModel instantiates a new DeleteUserSSHKeyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserSSHKeyModelWithDefaults

`func NewDeleteUserSSHKeyModelWithDefaults() *DeleteUserSSHKeyModel`

NewDeleteUserSSHKeyModelWithDefaults instantiates a new DeleteUserSSHKeyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyFingerprint

`func (o *DeleteUserSSHKeyModel) GetKeyFingerprint() string`

GetKeyFingerprint returns the KeyFingerprint field if non-nil, zero value otherwise.

### GetKeyFingerprintOk

`func (o *DeleteUserSSHKeyModel) GetKeyFingerprintOk() (*string, bool)`

GetKeyFingerprintOk returns a tuple with the KeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFingerprint

`func (o *DeleteUserSSHKeyModel) SetKeyFingerprint(v string)`

SetKeyFingerprint sets KeyFingerprint field to given value.


### GetKeyStoreTypeFilter

`func (o *DeleteUserSSHKeyModel) GetKeyStoreTypeFilter() int32`

GetKeyStoreTypeFilter returns the KeyStoreTypeFilter field if non-nil, zero value otherwise.

### GetKeyStoreTypeFilterOk

`func (o *DeleteUserSSHKeyModel) GetKeyStoreTypeFilterOk() (*int32, bool)`

GetKeyStoreTypeFilterOk returns a tuple with the KeyStoreTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyStoreTypeFilter

`func (o *DeleteUserSSHKeyModel) SetKeyStoreTypeFilter(v int32)`

SetKeyStoreTypeFilter sets KeyStoreTypeFilter field to given value.

### HasKeyStoreTypeFilter

`func (o *DeleteUserSSHKeyModel) HasKeyStoreTypeFilter() bool`

HasKeyStoreTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


