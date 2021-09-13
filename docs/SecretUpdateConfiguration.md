# SecretUpdateConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangePasswordInResetMode** | Pointer to **bool** | Defines whether or not password changed will be performed via reset mode using the reconciliation account. This is useful in cases where the password policy prevents the user from changing his own password minimal age restriction is applied. | [optional] 

## Methods

### NewSecretUpdateConfiguration

`func NewSecretUpdateConfiguration() *SecretUpdateConfiguration`

NewSecretUpdateConfiguration instantiates a new SecretUpdateConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretUpdateConfigurationWithDefaults

`func NewSecretUpdateConfigurationWithDefaults() *SecretUpdateConfiguration`

NewSecretUpdateConfigurationWithDefaults instantiates a new SecretUpdateConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangePasswordInResetMode

`func (o *SecretUpdateConfiguration) GetChangePasswordInResetMode() bool`

GetChangePasswordInResetMode returns the ChangePasswordInResetMode field if non-nil, zero value otherwise.

### GetChangePasswordInResetModeOk

`func (o *SecretUpdateConfiguration) GetChangePasswordInResetModeOk() (*bool, bool)`

GetChangePasswordInResetModeOk returns a tuple with the ChangePasswordInResetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordInResetMode

`func (o *SecretUpdateConfiguration) SetChangePasswordInResetMode(v bool)`

SetChangePasswordInResetMode sets ChangePasswordInResetMode field to given value.

### HasChangePasswordInResetMode

`func (o *SecretUpdateConfiguration) HasChangePasswordInResetMode() bool`

HasChangePasswordInResetMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


