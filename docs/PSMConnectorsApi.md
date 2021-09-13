# \PSMConnectorsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PSMConnectorsGetPSMConnectorsController**](PSMConnectorsApi.md#PSMConnectorsGetPSMConnectorsController) | **Get** /api/PSM/Connectors | 



## PSMConnectorsGetPSMConnectorsController

> []PSMConnector PSMConnectorsGetPSMConnectorsController(ctx).Execute()





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
    resp, r, err := api_client.PSMConnectorsApi.PSMConnectorsGetPSMConnectorsController(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PSMConnectorsApi.PSMConnectorsGetPSMConnectorsController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PSMConnectorsGetPSMConnectorsController`: []PSMConnector
    fmt.Fprintf(os.Stdout, "Response from `PSMConnectorsApi.PSMConnectorsGetPSMConnectorsController`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPSMConnectorsGetPSMConnectorsControllerRequest struct via the builder pattern


### Return type

[**[]PSMConnector**](PSMConnector.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

