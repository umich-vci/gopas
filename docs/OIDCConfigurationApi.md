# \OIDCConfigurationApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OIDCConfigurationAddProvider**](OIDCConfigurationApi.md#OIDCConfigurationAddProvider) | **Post** /api/Configuration/OIDC/Providers | 
[**OIDCConfigurationDeleteProvider**](OIDCConfigurationApi.md#OIDCConfigurationDeleteProvider) | **Delete** /api/Configuration/OIDC/Providers/{id} | 
[**OIDCConfigurationGetProvider**](OIDCConfigurationApi.md#OIDCConfigurationGetProvider) | **Get** /api/Configuration/OIDC/Providers/{id} | 
[**OIDCConfigurationGetProviders**](OIDCConfigurationApi.md#OIDCConfigurationGetProviders) | **Get** /api/Configuration/OIDC/Providers | 
[**OIDCConfigurationUpdateProvider**](OIDCConfigurationApi.md#OIDCConfigurationUpdateProvider) | **Put** /api/Configuration/OIDC/Providers/{id} | 



## OIDCConfigurationAddProvider

> OPConfigurationData OIDCConfigurationAddProvider(ctx).ConfigurationData(configurationData).Execute()





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
    configurationData := *openapiclient.NewOPConfigurationData("Id_example", "ClientId_example", "ClientSecretMethod_example", "DiscoveryEndpointUrl_example") // OPConfigurationData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OIDCConfigurationApi.OIDCConfigurationAddProvider(context.Background()).ConfigurationData(configurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OIDCConfigurationApi.OIDCConfigurationAddProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OIDCConfigurationAddProvider`: OPConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `OIDCConfigurationApi.OIDCConfigurationAddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOIDCConfigurationAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationData** | [**OPConfigurationData**](OPConfigurationData.md) |  | 

### Return type

[**OPConfigurationData**](OPConfigurationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OIDCConfigurationDeleteProvider

> OIDCConfigurationDeleteProvider(ctx, id).Execute()





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
    id := "id_example" // string | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OIDCConfigurationApi.OIDCConfigurationDeleteProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OIDCConfigurationApi.OIDCConfigurationDeleteProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOIDCConfigurationDeleteProviderRequest struct via the builder pattern


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


## OIDCConfigurationGetProvider

> OPConfigurationData OIDCConfigurationGetProvider(ctx, id).Execute()





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
    id := "id_example" // string | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OIDCConfigurationApi.OIDCConfigurationGetProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OIDCConfigurationApi.OIDCConfigurationGetProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OIDCConfigurationGetProvider`: OPConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `OIDCConfigurationApi.OIDCConfigurationGetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOIDCConfigurationGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OPConfigurationData**](OPConfigurationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OIDCConfigurationGetProviders

> OPConfigurations OIDCConfigurationGetProviders(ctx).Execute()





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
    resp, r, err := api_client.OIDCConfigurationApi.OIDCConfigurationGetProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OIDCConfigurationApi.OIDCConfigurationGetProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OIDCConfigurationGetProviders`: OPConfigurations
    fmt.Fprintf(os.Stdout, "Response from `OIDCConfigurationApi.OIDCConfigurationGetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOIDCConfigurationGetProvidersRequest struct via the builder pattern


### Return type

[**OPConfigurations**](OPConfigurations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OIDCConfigurationUpdateProvider

> OIDCConfigurationUpdateProvider(ctx, id).ConfigurationData(configurationData).Execute()





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
    id := "id_example" // string | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.
    configurationData := *openapiclient.NewOPConfigurationBase("ClientId_example", "ClientSecretMethod_example", "DiscoveryEndpointUrl_example") // OPConfigurationBase | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OIDCConfigurationApi.OIDCConfigurationUpdateProvider(context.Background(), id).ConfigurationData(configurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OIDCConfigurationApi.OIDCConfigurationUpdateProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOIDCConfigurationUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationData** | [**OPConfigurationBase**](OPConfigurationBase.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

