# AuthenticationMethodsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to [**[]AuthenticationMethod**](AuthenticationMethod.md) |  | [optional] 

## Methods

### NewAuthenticationMethodsResponse

`func NewAuthenticationMethodsResponse() *AuthenticationMethodsResponse`

NewAuthenticationMethodsResponse instantiates a new AuthenticationMethodsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodsResponseWithDefaults

`func NewAuthenticationMethodsResponseWithDefaults() *AuthenticationMethodsResponse`

NewAuthenticationMethodsResponseWithDefaults instantiates a new AuthenticationMethodsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *AuthenticationMethodsResponse) GetMethods() []AuthenticationMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *AuthenticationMethodsResponse) GetMethodsOk() (*[]AuthenticationMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *AuthenticationMethodsResponse) SetMethods(v []AuthenticationMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *AuthenticationMethodsResponse) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


