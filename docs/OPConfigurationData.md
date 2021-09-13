# OPConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the provider.  This ID is used to identify the OIDC Identity Provider in PVWA | 
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

### NewOPConfigurationData

`func NewOPConfigurationData(id string, clientId string, clientSecretMethod string, discoveryEndpointUrl string, ) *OPConfigurationData`

NewOPConfigurationData instantiates a new OPConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOPConfigurationDataWithDefaults

`func NewOPConfigurationDataWithDefaults() *OPConfigurationData`

NewOPConfigurationDataWithDefaults instantiates a new OPConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OPConfigurationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OPConfigurationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OPConfigurationData) SetId(v string)`

SetId sets Id field to given value.


### GetAuthenticationFlow

`func (o *OPConfigurationData) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OPConfigurationData) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OPConfigurationData) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OPConfigurationData) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### GetAuthenticationEndpointUrl

`func (o *OPConfigurationData) GetAuthenticationEndpointUrl() string`

GetAuthenticationEndpointUrl returns the AuthenticationEndpointUrl field if non-nil, zero value otherwise.

### GetAuthenticationEndpointUrlOk

`func (o *OPConfigurationData) GetAuthenticationEndpointUrlOk() (*string, bool)`

GetAuthenticationEndpointUrlOk returns a tuple with the AuthenticationEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEndpointUrl

`func (o *OPConfigurationData) SetAuthenticationEndpointUrl(v string)`

SetAuthenticationEndpointUrl sets AuthenticationEndpointUrl field to given value.

### HasAuthenticationEndpointUrl

`func (o *OPConfigurationData) HasAuthenticationEndpointUrl() bool`

HasAuthenticationEndpointUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *OPConfigurationData) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OPConfigurationData) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OPConfigurationData) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OPConfigurationData) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetDescription

`func (o *OPConfigurationData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OPConfigurationData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OPConfigurationData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OPConfigurationData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJwkSet

`func (o *OPConfigurationData) GetJwkSet() string`

GetJwkSet returns the JwkSet field if non-nil, zero value otherwise.

### GetJwkSetOk

`func (o *OPConfigurationData) GetJwkSetOk() (*string, bool)`

GetJwkSetOk returns a tuple with the JwkSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwkSet

`func (o *OPConfigurationData) SetJwkSet(v string)`

SetJwkSet sets JwkSet field to given value.

### HasJwkSet

`func (o *OPConfigurationData) HasJwkSet() bool`

HasJwkSet returns a boolean if a field has been set.

### GetClientId

`func (o *OPConfigurationData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OPConfigurationData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OPConfigurationData) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OPConfigurationData) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OPConfigurationData) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OPConfigurationData) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OPConfigurationData) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretMethod

`func (o *OPConfigurationData) GetClientSecretMethod() string`

GetClientSecretMethod returns the ClientSecretMethod field if non-nil, zero value otherwise.

### GetClientSecretMethodOk

`func (o *OPConfigurationData) GetClientSecretMethodOk() (*string, bool)`

GetClientSecretMethodOk returns a tuple with the ClientSecretMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretMethod

`func (o *OPConfigurationData) SetClientSecretMethod(v string)`

SetClientSecretMethod sets ClientSecretMethod field to given value.


### GetDiscoveryEndpointUrl

`func (o *OPConfigurationData) GetDiscoveryEndpointUrl() string`

GetDiscoveryEndpointUrl returns the DiscoveryEndpointUrl field if non-nil, zero value otherwise.

### GetDiscoveryEndpointUrlOk

`func (o *OPConfigurationData) GetDiscoveryEndpointUrlOk() (*string, bool)`

GetDiscoveryEndpointUrlOk returns a tuple with the DiscoveryEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpointUrl

`func (o *OPConfigurationData) SetDiscoveryEndpointUrl(v string)`

SetDiscoveryEndpointUrl sets DiscoveryEndpointUrl field to given value.


### GetUserNameClaim

`func (o *OPConfigurationData) GetUserNameClaim() string`

GetUserNameClaim returns the UserNameClaim field if non-nil, zero value otherwise.

### GetUserNameClaimOk

`func (o *OPConfigurationData) GetUserNameClaimOk() (*string, bool)`

GetUserNameClaimOk returns a tuple with the UserNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameClaim

`func (o *OPConfigurationData) SetUserNameClaim(v string)`

SetUserNameClaim sets UserNameClaim field to given value.

### HasUserNameClaim

`func (o *OPConfigurationData) HasUserNameClaim() bool`

HasUserNameClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


