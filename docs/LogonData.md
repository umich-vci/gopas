# LogonData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | The name of the user who will logon to the Vault | [optional] 
**Password** | Pointer to **string** | The password of the user | [optional] 
**NewPassword** | Pointer to **string** | The new password of the user. This parameter is optional, and enables you to change a password  This parameter can be used with the following authentication methods:  CyberArk,  LDAP | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AdditionalInfo** | Pointer to **string** |  | [optional] 
**SecureMode** | Pointer to **bool** |  | [optional] 
**ConcurrentSession** | Pointer to **bool** | Set this parameter to true In order to allow more than one connection for the same user simultaneously.    Up to 300 concurrent sessions. | [optional] 

## Methods

### NewLogonData

`func NewLogonData() *LogonData`

NewLogonData instantiates a new LogonData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogonDataWithDefaults

`func NewLogonDataWithDefaults() *LogonData`

NewLogonDataWithDefaults instantiates a new LogonData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *LogonData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *LogonData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *LogonData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *LogonData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *LogonData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LogonData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LogonData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LogonData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *LogonData) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *LogonData) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *LogonData) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *LogonData) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetType

`func (o *LogonData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogonData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogonData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogonData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *LogonData) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *LogonData) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *LogonData) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *LogonData) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetSecureMode

`func (o *LogonData) GetSecureMode() bool`

GetSecureMode returns the SecureMode field if non-nil, zero value otherwise.

### GetSecureModeOk

`func (o *LogonData) GetSecureModeOk() (*bool, bool)`

GetSecureModeOk returns a tuple with the SecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureMode

`func (o *LogonData) SetSecureMode(v bool)`

SetSecureMode sets SecureMode field to given value.

### HasSecureMode

`func (o *LogonData) HasSecureMode() bool`

HasSecureMode returns a boolean if a field has been set.

### GetConcurrentSession

`func (o *LogonData) GetConcurrentSession() bool`

GetConcurrentSession returns the ConcurrentSession field if non-nil, zero value otherwise.

### GetConcurrentSessionOk

`func (o *LogonData) GetConcurrentSessionOk() (*bool, bool)`

GetConcurrentSessionOk returns a tuple with the ConcurrentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSession

`func (o *LogonData) SetConcurrentSession(v bool)`

SetConcurrentSession sets ConcurrentSession field to given value.

### HasConcurrentSession

`func (o *LogonData) HasConcurrentSession() bool`

HasConcurrentSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


