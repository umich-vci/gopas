# ResetUserPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** | The user&#39;s new password.  The new password maximum length is up to 39 (including) characters and it must meet the password policy requirements. | 

## Methods

### NewResetUserPassword

`func NewResetUserPassword(newPassword string, ) *ResetUserPassword`

NewResetUserPassword instantiates a new ResetUserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetUserPasswordWithDefaults

`func NewResetUserPasswordWithDefaults() *ResetUserPassword`

NewResetUserPasswordWithDefaults instantiates a new ResetUserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *ResetUserPassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetUserPassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetUserPassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


