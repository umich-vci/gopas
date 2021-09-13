# \AuthenticationMethodConfigurationsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationMethodConfigurationsAddConfigurationValue**](AuthenticationMethodConfigurationsApi.md#AuthenticationMethodConfigurationsAddConfigurationValue) | **Post** /api/Configuration/AuthenticationMethods | 
[**AuthenticationMethodConfigurationsDeleteAuthenticationMethod**](AuthenticationMethodConfigurationsApi.md#AuthenticationMethodConfigurationsDeleteAuthenticationMethod) | **Delete** /api/Configuration/AuthenticationMethods/{id} | 
[**AuthenticationMethodConfigurationsGetAuthenticationMethod**](AuthenticationMethodConfigurationsApi.md#AuthenticationMethodConfigurationsGetAuthenticationMethod) | **Get** /api/Configuration/AuthenticationMethods/{id} | 
[**AuthenticationMethodConfigurationsGetAuthenticationMethods**](AuthenticationMethodConfigurationsApi.md#AuthenticationMethodConfigurationsGetAuthenticationMethods) | **Get** /api/Configuration/AuthenticationMethods | 
[**AuthenticationMethodConfigurationsUpdateAuthenticationMethod**](AuthenticationMethodConfigurationsApi.md#AuthenticationMethodConfigurationsUpdateAuthenticationMethod) | **Put** /api/Configuration/AuthenticationMethods/{id} | 



## AuthenticationMethodConfigurationsAddConfigurationValue

> AuthenticationMethod AuthenticationMethodConfigurationsAddConfigurationValue(ctx).AuthenticationMethod(authenticationMethod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authenticationMethod := *openapiclient.NewAuthenticationMethod("Id_example") // AuthenticationMethod | The authentication method information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsAddConfigurationValue(context.Background()).AuthenticationMethod(authenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsAddConfigurationValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsAddConfigurationValue`: AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsAddConfigurationValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAddConfigurationValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationMethod** | [**AuthenticationMethod**](AuthenticationMethod.md) | The authentication method information | 

### Return type

[**AuthenticationMethod**](AuthenticationMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodConfigurationsDeleteAuthenticationMethod

> AuthenticationMethodConfigurationsDeleteAuthenticationMethod(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The authentication method unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsDeleteAuthenticationMethod(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsDeleteAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The authentication method unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodConfigurationsGetAuthenticationMethod

> AuthenticationMethod AuthenticationMethodConfigurationsGetAuthenticationMethod(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The authentication method unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethod(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsGetAuthenticationMethod`: AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The authentication method unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthenticationMethod**](AuthenticationMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodConfigurationsGetAuthenticationMethods

> AuthenticationMethodsResponse AuthenticationMethodConfigurationsGetAuthenticationMethods(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsGetAuthenticationMethods`: AuthenticationMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsGetAuthenticationMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest struct via the builder pattern


### Return type

[**AuthenticationMethodsResponse**](AuthenticationMethodsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodConfigurationsUpdateAuthenticationMethod

> AuthenticationMethod AuthenticationMethodConfigurationsUpdateAuthenticationMethod(ctx, id).Value(value).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The authentication method unique identifier.
    value := *openapiclient.NewAuthenticationMethodBase() // AuthenticationMethodBase | The authentication method information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsUpdateAuthenticationMethod(context.Background(), id).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsUpdateAuthenticationMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsUpdateAuthenticationMethod`: AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsApi.AuthenticationMethodConfigurationsUpdateAuthenticationMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The authentication method unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | [**AuthenticationMethodBase**](AuthenticationMethodBase.md) | The authentication method information | 

### Return type

[**AuthenticationMethod**](AuthenticationMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

