# \DiscoveredAccountsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveredAccountsAddDiscoveredAccount**](DiscoveredAccountsApi.md#DiscoveredAccountsAddDiscoveredAccount) | **Post** /api/DiscoveredAccounts | 
[**DiscoveredAccountsDeleteDiscoveredAccounts**](DiscoveredAccountsApi.md#DiscoveredAccountsDeleteDiscoveredAccounts) | **Delete** /api/DiscoveredAccounts | 
[**DiscoveredAccountsGetDiscoveredAccount**](DiscoveredAccountsApi.md#DiscoveredAccountsGetDiscoveredAccount) | **Get** /api/DiscoveredAccounts/{id} | 
[**DiscoveredAccountsGetDiscoveredAccounts**](DiscoveredAccountsApi.md#DiscoveredAccountsGetDiscoveredAccounts) | **Get** /api/DiscoveredAccounts | 



## DiscoveredAccountsAddDiscoveredAccount

> DiscoveredAccountResponse DiscoveredAccountsAddDiscoveredAccount(ctx).DiscoveredAccount(discoveredAccount).Execute()





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
    discoveredAccount := *openapiclient.NewDiscoveredAccount("UserName_example", "Address_example", "PlatformType_example") // DiscoveredAccount | The discovered account to add

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DiscoveredAccountsApi.DiscoveredAccountsAddDiscoveredAccount(context.Background()).DiscoveredAccount(discoveredAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscoveredAccountsApi.DiscoveredAccountsAddDiscoveredAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoveredAccountsAddDiscoveredAccount`: DiscoveredAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscoveredAccountsApi.DiscoveredAccountsAddDiscoveredAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveredAccountsAddDiscoveredAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoveredAccount** | [**DiscoveredAccount**](DiscoveredAccount.md) | The discovered account to add | 

### Return type

[**DiscoveredAccountResponse**](DiscoveredAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveredAccountsDeleteDiscoveredAccounts

> map[string]interface{} DiscoveredAccountsDeleteDiscoveredAccounts(ctx).Execute()





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
    resp, r, err := api_client.DiscoveredAccountsApi.DiscoveredAccountsDeleteDiscoveredAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscoveredAccountsApi.DiscoveredAccountsDeleteDiscoveredAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoveredAccountsDeleteDiscoveredAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DiscoveredAccountsApi.DiscoveredAccountsDeleteDiscoveredAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveredAccountsDeleteDiscoveredAccountsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveredAccountsGetDiscoveredAccount

> GetDiscoveredAccount DiscoveredAccountsGetDiscoveredAccount(ctx, id).Execute()





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
    id := "id_example" // string | The account's unique ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoveredAccountsGetDiscoveredAccount`: GetDiscoveredAccount
    fmt.Fprintf(os.Stdout, "Response from `DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account&#39;s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveredAccountsGetDiscoveredAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDiscoveredAccount**](GetDiscoveredAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoveredAccountsGetDiscoveredAccounts

> GetDiscoveredAccountsResponse DiscoveredAccountsGetDiscoveredAccounts(ctx).Search(search).SearchType(searchType).Offset(offset).Limit(limit).Filter(filter).Execute()





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
    search := "search_example" // string | Searches according to REST standard (search=\"search word\").    Search is supported for userName and address. (optional)
    searchType := "searchType_example" // string | The type of search to perform.     The keyword can either be contained within the account property values, or at the beginning of the value  specified in the search parameter.    When using a keyword at the beginning of a value, performance is enhanced.    Valid values: contains (default) or startswith (optional)
    offset := int32(56) // int32 | The offset of the first returned accounts into the list of results. (optional)
    limit := int32(56) // int32 | The maximum number of returned accounts.     If not specified, the server limits the results to 100.    The maximum number that can be specified is 1000. (optional)
    filter := "filter_example" // string | Search for accounts filtered by:    • platformType - Get discovered accounts of a specific platform, using the platform type name.        Usage: filter=platformType eq Windows Server Local            Type: string            Valid values:                o Windows Server Local                    o Windows Desktop Local                    o Windows Domain                    o Unix                    o Unix SSH Key                    o AWS                    o AWS Access Keys                    o Azure Password Management        • privileged - Get only privileged or non-privileged discovered accounts.        Usage: filter=privileged eq true            Type: boolean            Valid values: true/false        • accountEnabled - Get only enabled or disabled discovered accounts.           Usage: filter=accountEnabled eq true            Type: boolean            Valid values: true/false                  Note: To use more than one filter, you can use the AND operator.         filter=platformType eq Windows Server Local AND privileged eq true (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccounts(context.Background()).Search(search).SearchType(searchType).Offset(offset).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoveredAccountsGetDiscoveredAccounts`: GetDiscoveredAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscoveredAccountsApi.DiscoveredAccountsGetDiscoveredAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoveredAccountsGetDiscoveredAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Searches according to REST standard (search&#x3D;\&quot;search word\&quot;).    Search is supported for userName and address. | 
 **searchType** | **string** | The type of search to perform.     The keyword can either be contained within the account property values, or at the beginning of the value  specified in the search parameter.    When using a keyword at the beginning of a value, performance is enhanced.    Valid values: contains (default) or startswith | 
 **offset** | **int32** | The offset of the first returned accounts into the list of results. | 
 **limit** | **int32** | The maximum number of returned accounts.     If not specified, the server limits the results to 100.    The maximum number that can be specified is 1000. | 
 **filter** | **string** | Search for accounts filtered by:    • platformType - Get discovered accounts of a specific platform, using the platform type name.        Usage: filter&#x3D;platformType eq Windows Server Local            Type: string            Valid values:                o Windows Server Local                    o Windows Desktop Local                    o Windows Domain                    o Unix                    o Unix SSH Key                    o AWS                    o AWS Access Keys                    o Azure Password Management        • privileged - Get only privileged or non-privileged discovered accounts.        Usage: filter&#x3D;privileged eq true            Type: boolean            Valid values: true/false        • accountEnabled - Get only enabled or disabled discovered accounts.           Usage: filter&#x3D;accountEnabled eq true            Type: boolean            Valid values: true/false                  Note: To use more than one filter, you can use the AND operator.         filter&#x3D;platformType eq Windows Server Local AND privileged eq true | 

### Return type

[**GetDiscoveredAccountsResponse**](GetDiscoveredAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

