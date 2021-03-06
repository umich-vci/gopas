/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopas

import (
	"encoding/json"
)

// AuthenticationMethod struct for AuthenticationMethod
type AuthenticationMethod struct {
	// The authentication module identifier.
	Id string `json:"id"`
	// The display name of the authentication method.
	DisplayName *string `json:"displayName,omitempty"`
	// Whether or not the authentication method is enabled for use.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether or not the authentication method is available from the mobile application.
	MobileEnabled *bool `json:"mobileEnabled,omitempty"`
	// The logoff page URL of the 3rd party server.  The user is redirected to this page in order to complete the logoff.
	LogoffUrl *string `json:"logoffUrl,omitempty"`
	// Defines which second factor authentication to use when connecting to the Vault.  Empty value will disable the second factor authentication.  Valid values: cyberark, radius, ldap.
	SecondFactorAuth *string `json:"secondFactorAuth,omitempty"`
	// Defines the sign-in text for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods.
	SignInLabel *string `json:"signInLabel,omitempty"`
	// Defines the label of the username field for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods.
	UsernameFieldLabel *string `json:"usernameFieldLabel,omitempty"`
	// Defines the label of the password field for this authentication method.  Relevant only for CyberArk, RADIUS and LDAP authentication methods.
	PasswordFieldLabel *string `json:"passwordFieldLabel,omitempty"`
}

// NewAuthenticationMethod instantiates a new AuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethod(id string) *AuthenticationMethod {
	this := AuthenticationMethod{}
	this.Id = id
	return &this
}

// NewAuthenticationMethodWithDefaults instantiates a new AuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodWithDefaults() *AuthenticationMethod {
	this := AuthenticationMethod{}
	return &this
}

// GetId returns the Id field value
func (o *AuthenticationMethod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthenticationMethod) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthenticationMethod) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthenticationMethod) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMobileEnabled returns the MobileEnabled field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetMobileEnabled() bool {
	if o == nil || o.MobileEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MobileEnabled
}

// GetMobileEnabledOk returns a tuple with the MobileEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetMobileEnabledOk() (*bool, bool) {
	if o == nil || o.MobileEnabled == nil {
		return nil, false
	}
	return o.MobileEnabled, true
}

// HasMobileEnabled returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasMobileEnabled() bool {
	if o != nil && o.MobileEnabled != nil {
		return true
	}

	return false
}

// SetMobileEnabled gets a reference to the given bool and assigns it to the MobileEnabled field.
func (o *AuthenticationMethod) SetMobileEnabled(v bool) {
	o.MobileEnabled = &v
}

// GetLogoffUrl returns the LogoffUrl field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetLogoffUrl() string {
	if o == nil || o.LogoffUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoffUrl
}

// GetLogoffUrlOk returns a tuple with the LogoffUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetLogoffUrlOk() (*string, bool) {
	if o == nil || o.LogoffUrl == nil {
		return nil, false
	}
	return o.LogoffUrl, true
}

// HasLogoffUrl returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasLogoffUrl() bool {
	if o != nil && o.LogoffUrl != nil {
		return true
	}

	return false
}

// SetLogoffUrl gets a reference to the given string and assigns it to the LogoffUrl field.
func (o *AuthenticationMethod) SetLogoffUrl(v string) {
	o.LogoffUrl = &v
}

// GetSecondFactorAuth returns the SecondFactorAuth field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetSecondFactorAuth() string {
	if o == nil || o.SecondFactorAuth == nil {
		var ret string
		return ret
	}
	return *o.SecondFactorAuth
}

// GetSecondFactorAuthOk returns a tuple with the SecondFactorAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetSecondFactorAuthOk() (*string, bool) {
	if o == nil || o.SecondFactorAuth == nil {
		return nil, false
	}
	return o.SecondFactorAuth, true
}

// HasSecondFactorAuth returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasSecondFactorAuth() bool {
	if o != nil && o.SecondFactorAuth != nil {
		return true
	}

	return false
}

// SetSecondFactorAuth gets a reference to the given string and assigns it to the SecondFactorAuth field.
func (o *AuthenticationMethod) SetSecondFactorAuth(v string) {
	o.SecondFactorAuth = &v
}

// GetSignInLabel returns the SignInLabel field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetSignInLabel() string {
	if o == nil || o.SignInLabel == nil {
		var ret string
		return ret
	}
	return *o.SignInLabel
}

// GetSignInLabelOk returns a tuple with the SignInLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetSignInLabelOk() (*string, bool) {
	if o == nil || o.SignInLabel == nil {
		return nil, false
	}
	return o.SignInLabel, true
}

// HasSignInLabel returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasSignInLabel() bool {
	if o != nil && o.SignInLabel != nil {
		return true
	}

	return false
}

// SetSignInLabel gets a reference to the given string and assigns it to the SignInLabel field.
func (o *AuthenticationMethod) SetSignInLabel(v string) {
	o.SignInLabel = &v
}

// GetUsernameFieldLabel returns the UsernameFieldLabel field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetUsernameFieldLabel() string {
	if o == nil || o.UsernameFieldLabel == nil {
		var ret string
		return ret
	}
	return *o.UsernameFieldLabel
}

// GetUsernameFieldLabelOk returns a tuple with the UsernameFieldLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetUsernameFieldLabelOk() (*string, bool) {
	if o == nil || o.UsernameFieldLabel == nil {
		return nil, false
	}
	return o.UsernameFieldLabel, true
}

// HasUsernameFieldLabel returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasUsernameFieldLabel() bool {
	if o != nil && o.UsernameFieldLabel != nil {
		return true
	}

	return false
}

// SetUsernameFieldLabel gets a reference to the given string and assigns it to the UsernameFieldLabel field.
func (o *AuthenticationMethod) SetUsernameFieldLabel(v string) {
	o.UsernameFieldLabel = &v
}

// GetPasswordFieldLabel returns the PasswordFieldLabel field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetPasswordFieldLabel() string {
	if o == nil || o.PasswordFieldLabel == nil {
		var ret string
		return ret
	}
	return *o.PasswordFieldLabel
}

// GetPasswordFieldLabelOk returns a tuple with the PasswordFieldLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetPasswordFieldLabelOk() (*string, bool) {
	if o == nil || o.PasswordFieldLabel == nil {
		return nil, false
	}
	return o.PasswordFieldLabel, true
}

// HasPasswordFieldLabel returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasPasswordFieldLabel() bool {
	if o != nil && o.PasswordFieldLabel != nil {
		return true
	}

	return false
}

// SetPasswordFieldLabel gets a reference to the given string and assigns it to the PasswordFieldLabel field.
func (o *AuthenticationMethod) SetPasswordFieldLabel(v string) {
	o.PasswordFieldLabel = &v
}

func (o AuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MobileEnabled != nil {
		toSerialize["mobileEnabled"] = o.MobileEnabled
	}
	if o.LogoffUrl != nil {
		toSerialize["logoffUrl"] = o.LogoffUrl
	}
	if o.SecondFactorAuth != nil {
		toSerialize["secondFactorAuth"] = o.SecondFactorAuth
	}
	if o.SignInLabel != nil {
		toSerialize["signInLabel"] = o.SignInLabel
	}
	if o.UsernameFieldLabel != nil {
		toSerialize["usernameFieldLabel"] = o.UsernameFieldLabel
	}
	if o.PasswordFieldLabel != nil {
		toSerialize["passwordFieldLabel"] = o.PasswordFieldLabel
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticationMethod struct {
	value *AuthenticationMethod
	isSet bool
}

func (v NullableAuthenticationMethod) Get() *AuthenticationMethod {
	return v.value
}

func (v *NullableAuthenticationMethod) Set(val *AuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethod(val *AuthenticationMethod) *NullableAuthenticationMethod {
	return &NullableAuthenticationMethod{value: val, isSet: true}
}

func (v NullableAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
