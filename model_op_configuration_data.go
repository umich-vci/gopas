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

// OPConfigurationData struct for OPConfigurationData
type OPConfigurationData struct {
	// The unique identifier of the provider.  This ID is used to identify the OIDC Identity Provider in PVWA
	Id string `json:"id"`
	// The OIDC connection flow. Valid values
	AuthenticationFlow *string `json:"authenticationFlow,omitempty"`
	// The URL of the provider's authorization endpoint. Authentication requests will be sent to this URL.  Note: This is not relevant if the Discovery URL is provided
	AuthenticationEndpointUrl *string `json:"authenticationEndpointUrl,omitempty"`
	// The Issuer Identifier for the OpenID Provider.  This is used by the application to verify that the response was issued from a specific provider.  Note: This is not relevant if the Discovery URL is provided
	Issuer *string `json:"issuer,omitempty"`
	// A description of the provider
	Description *string `json:"description,omitempty"`
	// (JSON web key set) The set of keys provided by the OIDC Identity provider for validating JWT (JSON web tokens) during the authentication flow.  The JSON must include a \"keys\" parameter, which is an array of JWKs(JWT signing keys). Note: This is not relevant if the Discovery URL is provided
	JwkSet *string `json:"jwkSet,omitempty"`
	// The unique identifier for the client application.  This ID is created by the provider, and assigned to each client application upon registration
	ClientId string `json:"clientId"`
	// The client secret is only known to the application and the provider for secure communication during the authentication flow.                This secret is created by the provider, and assigned to each client application upon registration
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client authentication method for the client secret. Valid values
	ClientSecretMethod string `json:"clientSecretMethod"`
	// OIDC defines a discovery mechanism, called OpenID Connect Discovery, where an OIDC Identity provider publishes its metadata at a well-known URL.  This URL is metadata that describes the provider's configuration
	DiscoveryEndpointUrl string `json:"discoveryEndpointUrl"`
	// The property in the ID token provided by the OIDC Identity Provider that contains the user name.                Note: By default, the system will use the preferred_username claim in the ID token
	UserNameClaim *string `json:"userNameClaim,omitempty"`
}

// NewOPConfigurationData instantiates a new OPConfigurationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOPConfigurationData(id string, clientId string, clientSecretMethod string, discoveryEndpointUrl string) *OPConfigurationData {
	this := OPConfigurationData{}
	this.Id = id
	this.ClientId = clientId
	this.ClientSecretMethod = clientSecretMethod
	this.DiscoveryEndpointUrl = discoveryEndpointUrl
	return &this
}

// NewOPConfigurationDataWithDefaults instantiates a new OPConfigurationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOPConfigurationDataWithDefaults() *OPConfigurationData {
	this := OPConfigurationData{}
	return &this
}

// GetId returns the Id field value
func (o *OPConfigurationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OPConfigurationData) SetId(v string) {
	o.Id = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise.
func (o *OPConfigurationData) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil || o.AuthenticationFlow == nil {
		return nil, false
	}
	return o.AuthenticationFlow, true
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *OPConfigurationData) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow != nil {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given string and assigns it to the AuthenticationFlow field.
func (o *OPConfigurationData) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow = &v
}

// GetAuthenticationEndpointUrl returns the AuthenticationEndpointUrl field value if set, zero value otherwise.
func (o *OPConfigurationData) GetAuthenticationEndpointUrl() string {
	if o == nil || o.AuthenticationEndpointUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationEndpointUrl
}

// GetAuthenticationEndpointUrlOk returns a tuple with the AuthenticationEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetAuthenticationEndpointUrlOk() (*string, bool) {
	if o == nil || o.AuthenticationEndpointUrl == nil {
		return nil, false
	}
	return o.AuthenticationEndpointUrl, true
}

// HasAuthenticationEndpointUrl returns a boolean if a field has been set.
func (o *OPConfigurationData) HasAuthenticationEndpointUrl() bool {
	if o != nil && o.AuthenticationEndpointUrl != nil {
		return true
	}

	return false
}

// SetAuthenticationEndpointUrl gets a reference to the given string and assigns it to the AuthenticationEndpointUrl field.
func (o *OPConfigurationData) SetAuthenticationEndpointUrl(v string) {
	o.AuthenticationEndpointUrl = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OPConfigurationData) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OPConfigurationData) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OPConfigurationData) SetIssuer(v string) {
	o.Issuer = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OPConfigurationData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OPConfigurationData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OPConfigurationData) SetDescription(v string) {
	o.Description = &v
}

// GetJwkSet returns the JwkSet field value if set, zero value otherwise.
func (o *OPConfigurationData) GetJwkSet() string {
	if o == nil || o.JwkSet == nil {
		var ret string
		return ret
	}
	return *o.JwkSet
}

// GetJwkSetOk returns a tuple with the JwkSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetJwkSetOk() (*string, bool) {
	if o == nil || o.JwkSet == nil {
		return nil, false
	}
	return o.JwkSet, true
}

// HasJwkSet returns a boolean if a field has been set.
func (o *OPConfigurationData) HasJwkSet() bool {
	if o != nil && o.JwkSet != nil {
		return true
	}

	return false
}

// SetJwkSet gets a reference to the given string and assigns it to the JwkSet field.
func (o *OPConfigurationData) SetJwkSet(v string) {
	o.JwkSet = &v
}

// GetClientId returns the ClientId field value
func (o *OPConfigurationData) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OPConfigurationData) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *OPConfigurationData) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OPConfigurationData) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OPConfigurationData) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientSecretMethod returns the ClientSecretMethod field value
func (o *OPConfigurationData) GetClientSecretMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecretMethod
}

// GetClientSecretMethodOk returns a tuple with the ClientSecretMethod field value
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetClientSecretMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecretMethod, true
}

// SetClientSecretMethod sets field value
func (o *OPConfigurationData) SetClientSecretMethod(v string) {
	o.ClientSecretMethod = v
}

// GetDiscoveryEndpointUrl returns the DiscoveryEndpointUrl field value
func (o *OPConfigurationData) GetDiscoveryEndpointUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscoveryEndpointUrl
}

// GetDiscoveryEndpointUrlOk returns a tuple with the DiscoveryEndpointUrl field value
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetDiscoveryEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryEndpointUrl, true
}

// SetDiscoveryEndpointUrl sets field value
func (o *OPConfigurationData) SetDiscoveryEndpointUrl(v string) {
	o.DiscoveryEndpointUrl = v
}

// GetUserNameClaim returns the UserNameClaim field value if set, zero value otherwise.
func (o *OPConfigurationData) GetUserNameClaim() string {
	if o == nil || o.UserNameClaim == nil {
		var ret string
		return ret
	}
	return *o.UserNameClaim
}

// GetUserNameClaimOk returns a tuple with the UserNameClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OPConfigurationData) GetUserNameClaimOk() (*string, bool) {
	if o == nil || o.UserNameClaim == nil {
		return nil, false
	}
	return o.UserNameClaim, true
}

// HasUserNameClaim returns a boolean if a field has been set.
func (o *OPConfigurationData) HasUserNameClaim() bool {
	if o != nil && o.UserNameClaim != nil {
		return true
	}

	return false
}

// SetUserNameClaim gets a reference to the given string and assigns it to the UserNameClaim field.
func (o *OPConfigurationData) SetUserNameClaim(v string) {
	o.UserNameClaim = &v
}

func (o OPConfigurationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.AuthenticationFlow != nil {
		toSerialize["authenticationFlow"] = o.AuthenticationFlow
	}
	if o.AuthenticationEndpointUrl != nil {
		toSerialize["authenticationEndpointUrl"] = o.AuthenticationEndpointUrl
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.JwkSet != nil {
		toSerialize["jwkSet"] = o.JwkSet
	}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if true {
		toSerialize["clientSecretMethod"] = o.ClientSecretMethod
	}
	if true {
		toSerialize["discoveryEndpointUrl"] = o.DiscoveryEndpointUrl
	}
	if o.UserNameClaim != nil {
		toSerialize["userNameClaim"] = o.UserNameClaim
	}
	return json.Marshal(toSerialize)
}

type NullableOPConfigurationData struct {
	value *OPConfigurationData
	isSet bool
}

func (v NullableOPConfigurationData) Get() *OPConfigurationData {
	return v.value
}

func (v *NullableOPConfigurationData) Set(val *OPConfigurationData) {
	v.value = val
	v.isSet = true
}

func (v NullableOPConfigurationData) IsSet() bool {
	return v.isSet
}

func (v *NullableOPConfigurationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOPConfigurationData(val *OPConfigurationData) *NullableOPConfigurationData {
	return &NullableOPConfigurationData{value: val, isSet: true}
}

func (v NullableOPConfigurationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOPConfigurationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
