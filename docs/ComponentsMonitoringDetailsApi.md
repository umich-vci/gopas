# \ComponentsMonitoringDetailsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentsMonitoringDetailsGetComponentsMonitoringDetails**](ComponentsMonitoringDetailsApi.md#ComponentsMonitoringDetailsGetComponentsMonitoringDetails) | **Get** /api/ComponentsMonitoringDetails/{componentId} | 



## ComponentsMonitoringDetailsGetComponentsMonitoringDetails

> ComponentsMonitoringDetailsData ComponentsMonitoringDetailsGetComponentsMonitoringDetails(ctx, componentId).Execute()





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
    componentId := "componentId_example" // string | The type of component for which data will be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsMonitoringDetailsApi.ComponentsMonitoringDetailsGetComponentsMonitoringDetails(context.Background(), componentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsMonitoringDetailsApi.ComponentsMonitoringDetailsGetComponentsMonitoringDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentsMonitoringDetailsGetComponentsMonitoringDetails`: ComponentsMonitoringDetailsData
    fmt.Fprintf(os.Stdout, "Response from `ComponentsMonitoringDetailsApi.ComponentsMonitoringDetailsGetComponentsMonitoringDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | The type of component for which data will be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentsMonitoringDetailsData**](ComponentsMonitoringDetailsData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

