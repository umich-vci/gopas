# \ConnectionComponentsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionComponentsImport**](ConnectionComponentsApi.md#ConnectionComponentsImport) | **Post** /api/ConnectionComponents/import | 



## ConnectionComponentsImport

> ImportConnectionComponentResponse ConnectionComponentsImport(ctx).ImportConnectionComponent(importConnectionComponent).Execute()





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
    importConnectionComponent := *openapiclient.NewImportConnectionComponentData(string(123)) // ImportConnectionComponentData | The ZIP file containing the connection component and additional configuration parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionComponentsApi.ConnectionComponentsImport(context.Background()).ImportConnectionComponent(importConnectionComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionComponentsApi.ConnectionComponentsImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionComponentsImport`: ImportConnectionComponentResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionComponentsApi.ConnectionComponentsImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionComponentsImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importConnectionComponent** | [**ImportConnectionComponentData**](ImportConnectionComponentData.md) | The ZIP file containing the connection component and additional configuration parameters. | 

### Return type

[**ImportConnectionComponentResponse**](ImportConnectionComponentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

