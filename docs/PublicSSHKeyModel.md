# PublicSSHKeyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | The public key data | [optional] 
**Fingerprint** | Pointer to **string** | The public key fingerprint | [optional] 
**CreationTime** | Pointer to **int64** | The creation time of the public key | [optional] 
**SshKeyStoreType** | Pointer to **int32** | The ssh key type | [optional] 

## Methods

### NewPublicSSHKeyModel

`func NewPublicSSHKeyModel() *PublicSSHKeyModel`

NewPublicSSHKeyModel instantiates a new PublicSSHKeyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSSHKeyModelWithDefaults

`func NewPublicSSHKeyModelWithDefaults() *PublicSSHKeyModel`

NewPublicSSHKeyModelWithDefaults instantiates a new PublicSSHKeyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *PublicSSHKeyModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PublicSSHKeyModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PublicSSHKeyModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PublicSSHKeyModel) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *PublicSSHKeyModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *PublicSSHKeyModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *PublicSSHKeyModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *PublicSSHKeyModel) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCreationTime

`func (o *PublicSSHKeyModel) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PublicSSHKeyModel) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PublicSSHKeyModel) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PublicSSHKeyModel) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetSshKeyStoreType

`func (o *PublicSSHKeyModel) GetSshKeyStoreType() int32`

GetSshKeyStoreType returns the SshKeyStoreType field if non-nil, zero value otherwise.

### GetSshKeyStoreTypeOk

`func (o *PublicSSHKeyModel) GetSshKeyStoreTypeOk() (*int32, bool)`

GetSshKeyStoreTypeOk returns a tuple with the SshKeyStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyStoreType

`func (o *PublicSSHKeyModel) SetSshKeyStoreType(v int32)`

SetSshKeyStoreType sets SshKeyStoreType field to given value.

### HasSshKeyStoreType

`func (o *PublicSSHKeyModel) HasSshKeyStoreType() bool`

HasSshKeyStoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


