# \SafesApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SafesAddSafe**](SafesApi.md#SafesAddSafe) | **Post** /api/Safes | 
[**SafesAddSafeOwner**](SafesApi.md#SafesAddSafeOwner) | **Post** /api/Safes/{safeUrlId}/members | 
[**SafesDeleteSafeDetails**](SafesApi.md#SafesDeleteSafeDetails) | **Delete** /api/Safes/{safeUrlId} | 
[**SafesGetGroups**](SafesApi.md#SafesGetGroups) | **Get** /api/Safes/{safeName}/accountgroups | 
[**SafesGetSafeMembers**](SafesApi.md#SafesGetSafeMembers) | **Get** /api/Safes/{safeUrlId}/members | 
[**SafesGetSafes**](SafesApi.md#SafesGetSafes) | **Get** /api/Safes | 



## SafesAddSafe

> map[string]interface{} SafesAddSafe(ctx).Safe(safe).Execute()





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
    safe := *openapiclient.NewAddSafeData() // AddSafeData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesAddSafe(context.Background()).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesAddSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesAddSafe`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesAddSafe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSafesAddSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safe** | [**AddSafeData**](AddSafeData.md) |  | 

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


## SafesAddSafeOwner

> map[string]interface{} SafesAddSafeOwner(ctx, safeUrlId).Member(member).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The unique ID of the Safe.
    member := *openapiclient.NewSafeMemberItem("MemberName_example", map[string]bool{"key": false}) // SafeMemberItem | An existing user to add as a Safe member.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesAddSafeOwner(context.Background(), safeUrlId).Member(member).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesAddSafeOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesAddSafeOwner`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesAddSafeOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The unique ID of the Safe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesAddSafeOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **member** | [**SafeMemberItem**](SafeMemberItem.md) | An existing user to add as a Safe member. | 

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


## SafesDeleteSafeDetails

> SafesDeleteSafeDetails(ctx, safeUrlId).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The unique ID of the Safe.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesDeleteSafeDetails(context.Background(), safeUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesDeleteSafeDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The unique ID of the Safe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesDeleteSafeDetailsRequest struct via the builder pattern


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


## SafesGetGroups

> AccountGroup SafesGetGroups(ctx, safeName).Execute()





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
    safeName := "safeName_example" // string | The name of the Safe where the account groups are.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesGetGroups(context.Background(), safeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesGetGroups`: AccountGroup
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesGetGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeName** | **string** | The name of the Safe where the account groups are. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountGroup**](AccountGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesGetSafeMembers

> SafeMemberResponse SafesGetSafeMembers(ctx, safeUrlId).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Filter(filter).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The unique ID of the safe to return all its members
    limit := int64(789) // int64 |  (optional)
    offset := int64(789) // int64 |  (optional)
    useCache := true // bool |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    search := "search_example" // string |  (optional)
    filter := "filter_example" // string | <para>Filtering according to REST standard. </para>  <para>memberType - Return only members of single type (user or group)</para>  <para>membershipExpired - Return only members that have or don't have an expired membership</para>  <para>includePredefinedUsers - Include predefined users in the returned list.</para> (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesGetSafeMembers(context.Background(), safeUrlId).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesGetSafeMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesGetSafeMembers`: SafeMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesGetSafeMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The unique ID of the safe to return all its members | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesGetSafeMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** |  | 
 **offset** | **int64** |  | 
 **useCache** | **bool** |  | 
 **sort** | **[]string** |  | 
 **search** | **string** |  | 
 **filter** | **string** | &lt;para&gt;Filtering according to REST standard. &lt;/para&gt;  &lt;para&gt;memberType - Return only members of single type (user or group)&lt;/para&gt;  &lt;para&gt;membershipExpired - Return only members that have or don&#39;t have an expired membership&lt;/para&gt;  &lt;para&gt;includePredefinedUsers - Include predefined users in the returned list.&lt;/para&gt; | 

### Return type

[**SafeMemberResponse**](SafeMemberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesGetSafes

> []SafeListItem SafesGetSafes(ctx).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).IncludeAccounts(includeAccounts).ExtendedDetails(extendedDetails).Execute()





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
    limit := int64(789) // int64 |  (optional)
    offset := int64(789) // int64 |  (optional)
    useCache := true // bool |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    search := "search_example" // string |  (optional)
    includeAccounts := true // bool | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be false. (optional)
    extendedDetails := true // bool | Whether or not to return all Safe data or only SafeName as part of the response. If not sent, the value will be true. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesGetSafes(context.Background()).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).IncludeAccounts(includeAccounts).ExtendedDetails(extendedDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesGetSafes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesGetSafes`: []SafeListItem
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesGetSafes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSafesGetSafesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** |  | 
 **offset** | **int64** |  | 
 **useCache** | **bool** |  | 
 **sort** | **[]string** |  | 
 **search** | **string** |  | 
 **includeAccounts** | **bool** | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be false. | 
 **extendedDetails** | **bool** | Whether or not to return all Safe data or only SafeName as part of the response. If not sent, the value will be true. | 

### Return type

[**[]SafeListItem**](SafeListItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

