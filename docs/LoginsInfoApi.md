# \LoginsInfoApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginsInfoGetLoginsInfo**](LoginsInfoApi.md#LoginsInfoGetLoginsInfo) | **Get** /api/LoginsInfo | 



## LoginsInfoGetLoginsInfo

> LoginsInfoData LoginsInfoGetLoginsInfo(ctx).Execute()



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
    resp, r, err := api_client.LoginsInfoApi.LoginsInfoGetLoginsInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginsInfoApi.LoginsInfoGetLoginsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginsInfoGetLoginsInfo`: LoginsInfoData
    fmt.Fprintf(os.Stdout, "Response from `LoginsInfoApi.LoginsInfoGetLoginsInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginsInfoGetLoginsInfoRequest struct via the builder pattern


### Return type

[**LoginsInfoData**](LoginsInfoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

