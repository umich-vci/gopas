# \BulkAccountsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAccountsAddBulkAccounts**](BulkAccountsApi.md#BulkAccountsAddBulkAccounts) | **Post** /api/BulkActions/Accounts | 
[**BulkAccountsGetBulkAction**](BulkAccountsApi.md#BulkAccountsGetBulkAction) | **Get** /api/BulkActions/Accounts/{bulkId} | 
[**BulkAccountsGetUserBulkAccountsActions**](BulkAccountsApi.md#BulkAccountsGetUserBulkAccountsActions) | **Get** /api/BulkActions/Accounts | 



## BulkAccountsAddBulkAccounts

> map[string]interface{} BulkAccountsAddBulkAccounts(ctx).BulkAccountsModel(bulkAccountsModel).Execute()





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
    bulkAccountsModel := *openapiclient.NewBulkAccountsModel() // BulkAccountsModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkAccountsApi.BulkAccountsAddBulkAccounts(context.Background()).BulkAccountsModel(bulkAccountsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkAccountsApi.BulkAccountsAddBulkAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkAccountsAddBulkAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BulkAccountsApi.BulkAccountsAddBulkAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAccountsAddBulkAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAccountsModel** | [**BulkAccountsModel**](BulkAccountsModel.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAccountsGetBulkAction

> BulkAccountsOperationResultExtended BulkAccountsGetBulkAction(ctx, bulkId).Execute()





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
    bulkId := int64(789) // int64 | The identifier of the bulk account upload.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkAccountsApi.BulkAccountsGetBulkAction(context.Background(), bulkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkAccountsApi.BulkAccountsGetBulkAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkAccountsGetBulkAction`: BulkAccountsOperationResultExtended
    fmt.Fprintf(os.Stdout, "Response from `BulkAccountsApi.BulkAccountsGetBulkAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bulkId** | **int64** | The identifier of the bulk account upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkAccountsGetBulkActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkAccountsOperationResultExtended**](BulkAccountsOperationResultExtended.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAccountsGetUserBulkAccountsActions

> GetBulkAccountsActionsResult BulkAccountsGetUserBulkAccountsActions(ctx).Limit(limit).Filter(filter).Execute()





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
    limit := int32(56) // int32 | Maximum number of returned bulk account actions. If not specified, the default value is 10. The maximum number that can be specified is 1000. (optional)
    filter := "filter_example" // string | Search for bulk account actions filtered by spesific status. e.g. Filter=status eq InProgress (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BulkAccountsApi.BulkAccountsGetUserBulkAccountsActions(context.Background()).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkAccountsApi.BulkAccountsGetUserBulkAccountsActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkAccountsGetUserBulkAccountsActions`: GetBulkAccountsActionsResult
    fmt.Fprintf(os.Stdout, "Response from `BulkAccountsApi.BulkAccountsGetUserBulkAccountsActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAccountsGetUserBulkAccountsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned bulk account actions. If not specified, the default value is 10. The maximum number that can be specified is 1000. | 
 **filter** | **string** | Search for bulk account actions filtered by spesific status. e.g. Filter&#x3D;status eq InProgress | 

### Return type

[**GetBulkAccountsActionsResult**](GetBulkAccountsActionsResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

