# AuthenticationMethodBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the authentication method. | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the authentication method is enabled for use. | [optional] 
**MobileEnabled** | Pointer to **bool** | Whether or not the authentication method is available from the mobile application. | [optional] 
**LogoffUrl** | Pointer to **string** | The logoff page URL of the 3rd party server.  The user is redirected to this page in order to complete the logoff. | [optional] 
**SecondFactorAuth** | Pointer to **string** | Defines which second factor authentication to use when connecting to the Vault.  Empty value will disable the second factor authentication.  Valid values: cyberark, radius, ldap. | [optional] 
**SignInLabel** | Pointer to **string** | Defines the sign-in text for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods. | [optional] 
**UsernameFieldLabel** | Pointer to **string** | Defines the label of the username field for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods. | [optional] 
**PasswordFieldLabel** | Pointer to **string** | Defines the label of the password field for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods. | [optional] 

## Methods

### NewAuthenticationMethodBase

`func NewAuthenticationMethodBase() *AuthenticationMethodBase`

NewAuthenticationMethodBase instantiates a new AuthenticationMethodBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodBaseWithDefaults

`func NewAuthenticationMethodBaseWithDefaults() *AuthenticationMethodBase`

NewAuthenticationMethodBaseWithDefaults instantiates a new AuthenticationMethodBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AuthenticationMethodBase) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthenticationMethodBase) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthenticationMethodBase) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthenticationMethodBase) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthenticationMethodBase) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthenticationMethodBase) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthenticationMethodBase) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthenticationMethodBase) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMobileEnabled

`func (o *AuthenticationMethodBase) GetMobileEnabled() bool`

GetMobileEnabled returns the MobileEnabled field if non-nil, zero value otherwise.

### GetMobileEnabledOk

`func (o *AuthenticationMethodBase) GetMobileEnabledOk() (*bool, bool)`

GetMobileEnabledOk returns a tuple with the MobileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileEnabled

`func (o *AuthenticationMethodBase) SetMobileEnabled(v bool)`

SetMobileEnabled sets MobileEnabled field to given value.

### HasMobileEnabled

`func (o *AuthenticationMethodBase) HasMobileEnabled() bool`

HasMobileEnabled returns a boolean if a field has been set.

### GetLogoffUrl

`func (o *AuthenticationMethodBase) GetLogoffUrl() string`

GetLogoffUrl returns the LogoffUrl field if non-nil, zero value otherwise.

### GetLogoffUrlOk

`func (o *AuthenticationMethodBase) GetLogoffUrlOk() (*string, bool)`

GetLogoffUrlOk returns a tuple with the LogoffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffUrl

`func (o *AuthenticationMethodBase) SetLogoffUrl(v string)`

SetLogoffUrl sets LogoffUrl field to given value.

### HasLogoffUrl

`func (o *AuthenticationMethodBase) HasLogoffUrl() bool`

HasLogoffUrl returns a boolean if a field has been set.

### GetSecondFactorAuth

`func (o *AuthenticationMethodBase) GetSecondFactorAuth() string`

GetSecondFactorAuth returns the SecondFactorAuth field if non-nil, zero value otherwise.

### GetSecondFactorAuthOk

`func (o *AuthenticationMethodBase) GetSecondFactorAuthOk() (*string, bool)`

GetSecondFactorAuthOk returns a tuple with the SecondFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondFactorAuth

`func (o *AuthenticationMethodBase) SetSecondFactorAuth(v string)`

SetSecondFactorAuth sets SecondFactorAuth field to given value.

### HasSecondFactorAuth

`func (o *AuthenticationMethodBase) HasSecondFactorAuth() bool`

HasSecondFactorAuth returns a boolean if a field has been set.

### GetSignInLabel

`func (o *AuthenticationMethodBase) GetSignInLabel() string`

GetSignInLabel returns the SignInLabel field if non-nil, zero value otherwise.

### GetSignInLabelOk

`func (o *AuthenticationMethodBase) GetSignInLabelOk() (*string, bool)`

GetSignInLabelOk returns a tuple with the SignInLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInLabel

`func (o *AuthenticationMethodBase) SetSignInLabel(v string)`

SetSignInLabel sets SignInLabel field to given value.

### HasSignInLabel

`func (o *AuthenticationMethodBase) HasSignInLabel() bool`

HasSignInLabel returns a boolean if a field has been set.

### GetUsernameFieldLabel

`func (o *AuthenticationMethodBase) GetUsernameFieldLabel() string`

GetUsernameFieldLabel returns the UsernameFieldLabel field if non-nil, zero value otherwise.

### GetUsernameFieldLabelOk

`func (o *AuthenticationMethodBase) GetUsernameFieldLabelOk() (*string, bool)`

GetUsernameFieldLabelOk returns a tuple with the UsernameFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFieldLabel

`func (o *AuthenticationMethodBase) SetUsernameFieldLabel(v string)`

SetUsernameFieldLabel sets UsernameFieldLabel field to given value.

### HasUsernameFieldLabel

`func (o *AuthenticationMethodBase) HasUsernameFieldLabel() bool`

HasUsernameFieldLabel returns a boolean if a field has been set.

### GetPasswordFieldLabel

`func (o *AuthenticationMethodBase) GetPasswordFieldLabel() string`

GetPasswordFieldLabel returns the PasswordFieldLabel field if non-nil, zero value otherwise.

### GetPasswordFieldLabelOk

`func (o *AuthenticationMethodBase) GetPasswordFieldLabelOk() (*string, bool)`

GetPasswordFieldLabelOk returns a tuple with the PasswordFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordFieldLabel

`func (o *AuthenticationMethodBase) SetPasswordFieldLabel(v string)`

SetPasswordFieldLabel sets PasswordFieldLabel field to given value.

### HasPasswordFieldLabel

`func (o *AuthenticationMethodBase) HasPasswordFieldLabel() bool`

HasPasswordFieldLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


