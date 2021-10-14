# \SafesApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SafesAddSafe**](SafesApi.md#SafesAddSafe) | **Post** /api/Safes | 
[**SafesAddSafeMember**](SafesApi.md#SafesAddSafeMember) | **Post** /api/Safes/{safeUrlId}/members | 
[**SafesDeleteSafe**](SafesApi.md#SafesDeleteSafe) | **Delete** /api/Safes/{safeUrlId} | 
[**SafesDeleteSafeMember**](SafesApi.md#SafesDeleteSafeMember) | **Delete** /api/Safes/{safeUrlId}/members/{memberName} | 
[**SafesGetGroups**](SafesApi.md#SafesGetGroups) | **Get** /api/Safes/{safeName}/accountgroups | 
[**SafesGetSafeDetails**](SafesApi.md#SafesGetSafeDetails) | **Get** /api/Safes/{safeUrlId} | 
[**SafesGetSafeMember**](SafesApi.md#SafesGetSafeMember) | **Get** /api/Safes/{safeUrlId}/members/{memberName} | 
[**SafesGetSafeMembers**](SafesApi.md#SafesGetSafeMembers) | **Get** /api/Safes/{safeUrlId}/members | 
[**SafesGetSafes**](SafesApi.md#SafesGetSafes) | **Get** /api/Safes | 
[**SafesUpdateSafe**](SafesApi.md#SafesUpdateSafe) | **Put** /api/Safes/{safeUrlId} | 
[**SafesUpdateSafeMember**](SafesApi.md#SafesUpdateSafeMember) | **Put** /api/Safes/{safeUrlId}/members/{memberName} | 



## SafesAddSafe

> AddSafeResponse SafesAddSafe(ctx).AddSafeRequest(addSafeRequest).Execute()





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
    addSafeRequest := *openapiclient.NewAddSafeRequest() // AddSafeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesAddSafe(context.Background()).AddSafeRequest(addSafeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesAddSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesAddSafe`: AddSafeResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesAddSafe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSafesAddSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSafeRequest** | [**AddSafeRequest**](AddSafeRequest.md) |  | 

### Return type

[**AddSafeResponse**](AddSafeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesAddSafeMember

> AddSafeMemberResponse SafesAddSafeMember(ctx, safeUrlId).AddSafeMemberRequestBody(addSafeMemberRequestBody).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe used when calling Safe APIs.
    addSafeMemberRequestBody := *openapiclient.NewAddSafeMemberRequestBody("MemberName_example", map[string]bool{"key": false}) // AddSafeMemberRequestBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesAddSafeMember(context.Background(), safeUrlId).AddSafeMemberRequestBody(addSafeMemberRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesAddSafeMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesAddSafeMember`: AddSafeMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesAddSafeMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe used when calling Safe APIs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesAddSafeMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSafeMemberRequestBody** | [**AddSafeMemberRequestBody**](AddSafeMemberRequestBody.md) |  | 

### Return type

[**AddSafeMemberResponse**](AddSafeMemberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesDeleteSafe

> map[string]interface{} SafesDeleteSafe(ctx, safeUrlId).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesDeleteSafe(context.Background(), safeUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesDeleteSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesDeleteSafe`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesDeleteSafe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesDeleteSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SafesDeleteSafeMember

> map[string]interface{} SafesDeleteSafeMember(ctx, safeUrlId, memberName).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe from which to delete the member.
    memberName := "memberName_example" // string | The name of the Safe member to delete from the list of Safe members.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesDeleteSafeMember(context.Background(), safeUrlId, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesDeleteSafeMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesDeleteSafeMember`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesDeleteSafeMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe from which to delete the member. | 
**memberName** | **string** | The name of the Safe member to delete from the list of Safe members. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesDeleteSafeMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SafesGetSafeDetails

> GetSafeDetailsResponse SafesGetSafeDetails(ctx, safeUrlId).IncludeAccounts(includeAccounts).UseCache(useCache).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe.
    includeAccounts := true // bool | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be False. (optional)
    useCache := true // bool | Whether to retrieve from session or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesGetSafeDetails(context.Background(), safeUrlId).IncludeAccounts(includeAccounts).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesGetSafeDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesGetSafeDetails`: GetSafeDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesGetSafeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesGetSafeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAccounts** | **bool** | Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be False. | 
 **useCache** | **bool** | Whether to retrieve from session or not. | 

### Return type

[**GetSafeDetailsResponse**](GetSafeDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesGetSafeMember

> GetSafeMemberResponse SafesGetSafeMember(ctx, safeUrlId, memberName).UseCache(useCache).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe used when calling Safe APIs.
    memberName := "memberName_example" // string | The Vault user name, Domain user name or group name of the Safe member.
    useCache := true // bool | Whether to retrieve from session or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesGetSafeMember(context.Background(), safeUrlId, memberName).UseCache(useCache).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesGetSafeMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesGetSafeMember`: GetSafeMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesGetSafeMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe used when calling Safe APIs. | 
**memberName** | **string** | The Vault user name, Domain user name or group name of the Safe member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesGetSafeMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useCache** | **bool** | Whether to retrieve from session or not. | 

### Return type

[**GetSafeMemberResponse**](GetSafeMemberResponse.md)

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
    safeUrlId := "safeUrlId_example" // string | The name of the safe to return all its members
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
**safeUrlId** | **string** | The name of the safe to return all its members | 

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


## SafesUpdateSafe

> UpdateSafeResponse SafesUpdateSafe(ctx, safeUrlId).UpdateSafeRequestBody(updateSafeRequestBody).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe used when calling Safe APIs.
    updateSafeRequestBody := *openapiclient.NewUpdateSafeRequestBody() // UpdateSafeRequestBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesUpdateSafe(context.Background(), safeUrlId).UpdateSafeRequestBody(updateSafeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesUpdateSafe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesUpdateSafe`: UpdateSafeResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesUpdateSafe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe used when calling Safe APIs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesUpdateSafeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSafeRequestBody** | [**UpdateSafeRequestBody**](UpdateSafeRequestBody.md) |  | 

### Return type

[**UpdateSafeResponse**](UpdateSafeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafesUpdateSafeMember

> UpdateSafeMemberResponse SafesUpdateSafeMember(ctx, safeUrlId, memberName).UpdateSafeMemberRequestBody(updateSafeMemberRequestBody).Execute()





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
    safeUrlId := "safeUrlId_example" // string | The name of the Safe used when calling Safe APIs.
    memberName := "memberName_example" // string | The Vault user name, Domain user name or group name of the Safe member.
    updateSafeMemberRequestBody := *openapiclient.NewUpdateSafeMemberRequestBody(map[string]bool{"key": false}) // UpdateSafeMemberRequestBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SafesApi.SafesUpdateSafeMember(context.Background(), safeUrlId, memberName).UpdateSafeMemberRequestBody(updateSafeMemberRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SafesApi.SafesUpdateSafeMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SafesUpdateSafeMember`: UpdateSafeMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `SafesApi.SafesUpdateSafeMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safeUrlId** | **string** | The name of the Safe used when calling Safe APIs. | 
**memberName** | **string** | The Vault user name, Domain user name or group name of the Safe member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSafesUpdateSafeMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSafeMemberRequestBody** | [**UpdateSafeMemberRequestBody**](UpdateSafeMemberRequestBody.md) |  | 

### Return type

[**UpdateSafeMemberResponse**](UpdateSafeMemberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

