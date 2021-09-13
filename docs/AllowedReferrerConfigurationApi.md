# \AllowedReferrerConfigurationApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowedReferrerConfigurationAddAllowedReferrer**](AllowedReferrerConfigurationApi.md#AllowedReferrerConfigurationAddAllowedReferrer) | **Post** /api/Configuration/AccessRestriction/AllowedReferrers | 
[**AllowedReferrerConfigurationGetAllAllowedReferrers**](AllowedReferrerConfigurationApi.md#AllowedReferrerConfigurationGetAllAllowedReferrers) | **Get** /api/Configuration/AccessRestriction/AllowedReferrers | 



## AllowedReferrerConfigurationAddAllowedReferrer

> AllowedReferrerDto AllowedReferrerConfigurationAddAllowedReferrer(ctx).Item(item).Execute()





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
    item := *openapiclient.NewAllowedReferrerDto("ReferrerURL_example") // AllowedReferrerDto | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AllowedReferrerConfigurationApi.AllowedReferrerConfigurationAddAllowedReferrer(context.Background()).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedReferrerConfigurationApi.AllowedReferrerConfigurationAddAllowedReferrer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllowedReferrerConfigurationAddAllowedReferrer`: AllowedReferrerDto
    fmt.Fprintf(os.Stdout, "Response from `AllowedReferrerConfigurationApi.AllowedReferrerConfigurationAddAllowedReferrer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllowedReferrerConfigurationAddAllowedReferrerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**AllowedReferrerDto**](AllowedReferrerDto.md) |  | 

### Return type

[**AllowedReferrerDto**](AllowedReferrerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllowedReferrerConfigurationGetAllAllowedReferrers

> []AllowedReferrerDto AllowedReferrerConfigurationGetAllAllowedReferrers(ctx).Execute()





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
    resp, r, err := api_client.AllowedReferrerConfigurationApi.AllowedReferrerConfigurationGetAllAllowedReferrers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowedReferrerConfigurationApi.AllowedReferrerConfigurationGetAllAllowedReferrers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllowedReferrerConfigurationGetAllAllowedReferrers`: []AllowedReferrerDto
    fmt.Fprintf(os.Stdout, "Response from `AllowedReferrerConfigurationApi.AllowedReferrerConfigurationGetAllAllowedReferrers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAllowedReferrerConfigurationGetAllAllowedReferrersRequest struct via the builder pattern


### Return type

[**[]AllowedReferrerDto**](AllowedReferrerDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

