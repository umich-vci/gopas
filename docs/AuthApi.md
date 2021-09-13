# \AuthApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthLogoff**](AuthApi.md#AuthLogoff) | **Post** /api/Auth/Logoff | 
[**AuthLogon**](AuthApi.md#AuthLogon) | **Post** /api/Auth/{type}/Logon | 



## AuthLogoff

> LogoffData AuthLogoff(ctx).Execute()





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
    resp, r, err := api_client.AuthApi.AuthLogoff(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLogoff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthLogoff`: LogoffData
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthLogoff`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogoffRequest struct via the builder pattern


### Return type

[**LogoffData**](LogoffData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogon

> AuthLogon(ctx, type_).Data(data).Execute()





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
    type_ := "type__example" // string | The authentication type. (for example: cyberark)
    data := *openapiclient.NewLogonData() // LogonData | The login information.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthLogon(context.Background(), type_).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthLogon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The authentication type. (for example: cyberark) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**LogonData**](LogonData.md) | The login information. | 

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

