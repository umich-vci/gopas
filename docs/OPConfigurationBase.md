# OPConfigurationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationFlow** | Pointer to **string** | The OIDC connection flow. Valid values | [optional] 
**AuthenticationEndpointUrl** | Pointer to **string** | The URL of the provider&#39;s authorization endpoint. Authentication requests will be sent to this URL.  Note: This is not relevant if the Discovery URL is provided | [optional] 
**Issuer** | Pointer to **string** | The Issuer Identifier for the OpenID Provider.  This is used by the application to verify that the response was issued from a specific provider.  Note: This is not relevant if the Discovery URL is provided | [optional] 
**Description** | Pointer to **string** | A description of the provider | [optional] 
**JwkSet** | Pointer to **string** | (JSON web key set) The set of keys provided by the OIDC Identity provider for validating JWT (JSON web tokens) during the authentication flow.  The JSON must include a \&quot;keys\&quot; parameter, which is an array of JWKs(JWT signing keys). Note: This is not relevant if the Discovery URL is provided | [optional] 
**ClientId** | **string** | The unique identifier for the client application.  This ID is created by the provider, and assigned to each client application upon registration | 
**ClientSecret** | Pointer to **string** | The client secret is only known to the application and the provider for secure communication during the authentication flow.                This secret is created by the provider, and assigned to each client application upon registration | [optional] 
**ClientSecretMethod** | **string** | The client authentication method for the client secret. Valid values | 
**DiscoveryEndpointUrl** | **string** | OIDC defines a discovery mechanism, called OpenID Connect Discovery, where an OIDC Identity provider publishes its metadata at a well-known URL.  This URL is metadata that describes the provider&#39;s configuration | 
**UserNameClaim** | Pointer to **string** | The property in the ID token provided by the OIDC Identity Provider that contains the user name.                Note: By default, the system will use the preferred_username claim in the ID token | [optional] 

## Methods

### NewOPConfigurationBase

`func NewOPConfigurationBase(clientId string, clientSecretMethod string, discoveryEndpointUrl string, ) *OPConfigurationBase`

NewOPConfigurationBase instantiates a new OPConfigurationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPConfigurationBaseWithDefaults

`func NewOPConfigurationBaseWithDefaults() *OPConfigurationBase`

NewOPConfigurationBaseWithDefaults instantiates a new OPConfigurationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationFlow

`func (o *OPConfigurationBase) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OPConfigurationBase) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OPConfigurationBase) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OPConfigurationBase) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### GetAuthenticationEndpointUrl

`func (o *OPConfigurationBase) GetAuthenticationEndpointUrl() string`

GetAuthenticationEndpointUrl returns the AuthenticationEndpointUrl field if non-nil, zero value otherwise.

### GetAuthenticationEndpointUrlOk

`func (o *OPConfigurationBase) GetAuthenticationEndpointUrlOk() (*string, bool)`

GetAuthenticationEndpointUrlOk returns a tuple with the AuthenticationEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEndpointUrl

`func (o *OPConfigurationBase) SetAuthenticationEndpointUrl(v string)`

SetAuthenticationEndpointUrl sets AuthenticationEndpointUrl field to given value.

### HasAuthenticationEndpointUrl

`func (o *OPConfigurationBase) HasAuthenticationEndpointUrl() bool`

HasAuthenticationEndpointUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *OPConfigurationBase) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OPConfigurationBase) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OPConfigurationBase) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OPConfigurationBase) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetDescription

`func (o *OPConfigurationBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OPConfigurationBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OPConfigurationBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OPConfigurationBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJwkSet

`func (o *OPConfigurationBase) GetJwkSet() string`

GetJwkSet returns the JwkSet field if non-nil, zero value otherwise.

### GetJwkSetOk

`func (o *OPConfigurationBase) GetJwkSetOk() (*string, bool)`

GetJwkSetOk returns a tuple with the JwkSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwkSet

`func (o *OPConfigurationBase) SetJwkSet(v string)`

SetJwkSet sets JwkSet field to given value.

### HasJwkSet

`func (o *OPConfigurationBase) HasJwkSet() bool`

HasJwkSet returns a boolean if a field has been set.

### GetClientId

`func (o *OPConfigurationBase) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OPConfigurationBase) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OPConfigurationBase) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OPConfigurationBase) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OPConfigurationBase) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OPConfigurationBase) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OPConfigurationBase) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretMethod

`func (o *OPConfigurationBase) GetClientSecretMethod() string`

GetClientSecretMethod returns the ClientSecretMethod field if non-nil, zero value otherwise.

### GetClientSecretMethodOk

`func (o *OPConfigurationBase) GetClientSecretMethodOk() (*string, bool)`

GetClientSecretMethodOk returns a tuple with the ClientSecretMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretMethod

`func (o *OPConfigurationBase) SetClientSecretMethod(v string)`

SetClientSecretMethod sets ClientSecretMethod field to given value.


### GetDiscoveryEndpointUrl

`func (o *OPConfigurationBase) GetDiscoveryEndpointUrl() string`

GetDiscoveryEndpointUrl returns the DiscoveryEndpointUrl field if non-nil, zero value otherwise.

### GetDiscoveryEndpointUrlOk

`func (o *OPConfigurationBase) GetDiscoveryEndpointUrlOk() (*string, bool)`

GetDiscoveryEndpointUrlOk returns a tuple with the DiscoveryEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpointUrl

`func (o *OPConfigurationBase) SetDiscoveryEndpointUrl(v string)`

SetDiscoveryEndpointUrl sets DiscoveryEndpointUrl field to given value.


### GetUserNameClaim

`func (o *OPConfigurationBase) GetUserNameClaim() string`

GetUserNameClaim returns the UserNameClaim field if non-nil, zero value otherwise.

### GetUserNameClaimOk

`func (o *OPConfigurationBase) GetUserNameClaimOk() (*string, bool)`

GetUserNameClaimOk returns a tuple with the UserNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameClaim

`func (o *OPConfigurationBase) SetUserNameClaim(v string)`

SetUserNameClaim sets UserNameClaim field to given value.

### HasUserNameClaim

`func (o *OPConfigurationBase) HasUserNameClaim() bool`

HasUserNameClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


