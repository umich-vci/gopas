# CredentialsManagementPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to [**TaskRole**](TaskRole.md) |  | [optional] 
**Change** | Pointer to [**TaskRole**](TaskRole.md) |  | [optional] 
**Reconcile** | Pointer to [**ReconcileTask**](ReconcileTask.md) |  | [optional] 
**SecretUpdateConfiguration** | Pointer to [**SecretUpdateConfiguration**](SecretUpdateConfiguration.md) |  | [optional] 

## Methods

### NewCredentialsManagementPolicy

`func NewCredentialsManagementPolicy() *CredentialsManagementPolicy`

NewCredentialsManagementPolicy instantiates a new CredentialsManagementPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsManagementPolicyWithDefaults

`func NewCredentialsManagementPolicyWithDefaults() *CredentialsManagementPolicy`

NewCredentialsManagementPolicyWithDefaults instantiates a new CredentialsManagementPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerification

`func (o *CredentialsManagementPolicy) GetVerification() TaskRole`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *CredentialsManagementPolicy) GetVerificationOk() (*TaskRole, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *CredentialsManagementPolicy) SetVerification(v TaskRole)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *CredentialsManagementPolicy) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetChange

`func (o *CredentialsManagementPolicy) GetChange() TaskRole`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *CredentialsManagementPolicy) GetChangeOk() (*TaskRole, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *CredentialsManagementPolicy) SetChange(v TaskRole)`

SetChange sets Change field to given value.

### HasChange

`func (o *CredentialsManagementPolicy) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetReconcile

`func (o *CredentialsManagementPolicy) GetReconcile() ReconcileTask`

GetReconcile returns the Reconcile field if non-nil, zero value otherwise.

### GetReconcileOk

`func (o *CredentialsManagementPolicy) GetReconcileOk() (*ReconcileTask, bool)`

GetReconcileOk returns a tuple with the Reconcile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconcile

`func (o *CredentialsManagementPolicy) SetReconcile(v ReconcileTask)`

SetReconcile sets Reconcile field to given value.

### HasReconcile

`func (o *CredentialsManagementPolicy) HasReconcile() bool`

HasReconcile returns a boolean if a field has been set.

### GetSecretUpdateConfiguration

`func (o *CredentialsManagementPolicy) GetSecretUpdateConfiguration() SecretUpdateConfiguration`

GetSecretUpdateConfiguration returns the SecretUpdateConfiguration field if non-nil, zero value otherwise.

### GetSecretUpdateConfigurationOk

`func (o *CredentialsManagementPolicy) GetSecretUpdateConfigurationOk() (*SecretUpdateConfiguration, bool)`

GetSecretUpdateConfigurationOk returns a tuple with the SecretUpdateConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretUpdateConfiguration

`func (o *CredentialsManagementPolicy) SetSecretUpdateConfiguration(v SecretUpdateConfiguration)`

SetSecretUpdateConfiguration sets SecretUpdateConfiguration field to given value.

### HasSecretUpdateConfiguration

`func (o *CredentialsManagementPolicy) HasSecretUpdateConfiguration() bool`

HasSecretUpdateConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


