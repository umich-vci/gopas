# \LiveSessionsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LiveSessionsGetLiveSession**](LiveSessionsApi.md#LiveSessionsGetLiveSession) | **Get** /api/LiveSessions/{liveSessionId} | 
[**LiveSessionsGetLiveSessionActivities**](LiveSessionsApi.md#LiveSessionsGetLiveSessionActivities) | **Get** /api/LiveSessions/{liveSessionId}/activities | 
[**LiveSessionsGetLiveSessionProperties**](LiveSessionsApi.md#LiveSessionsGetLiveSessionProperties) | **Get** /api/LiveSessions/{liveSessionId}/properties | 
[**LiveSessionsGetLiveSessions**](LiveSessionsApi.md#LiveSessionsGetLiveSessions) | **Get** /api/LiveSessions | 
[**LiveSessionsPSMAutoTermination**](LiveSessionsApi.md#LiveSessionsPSMAutoTermination) | **Post** /api/LiveSessions/{liveSessionId}/terminate | 
[**LiveSessionsPSMMonitor**](LiveSessionsApi.md#LiveSessionsPSMMonitor) | **Get** /api/LiveSessions/{liveSessionId}/monitor | 
[**LiveSessionsPSMSessionResume**](LiveSessionsApi.md#LiveSessionsPSMSessionResume) | **Post** /api/LiveSessions/{liveSessionId}/resume | 
[**LiveSessionsPSMSuspend**](LiveSessionsApi.md#LiveSessionsPSMSuspend) | **Post** /api/LiveSessions/{liveSessionId}/suspend | 



## LiveSessionsGetLiveSession

> map[string]interface{} LiveSessionsGetLiveSession(ctx, liveSessionId).ReturnURL(returnURL).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.
    returnURL := "returnURL_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsGetLiveSession(context.Background(), liveSessionId).ReturnURL(returnURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsGetLiveSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiveSessionsGetLiveSession`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LiveSessionsApi.LiveSessionsGetLiveSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsGetLiveSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnURL** | **string** |  | 

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


## LiveSessionsGetLiveSessionActivities

> map[string]interface{} LiveSessionsGetLiveSessionActivities(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsGetLiveSessionActivities(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsGetLiveSessionActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiveSessionsGetLiveSessionActivities`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LiveSessionsApi.LiveSessionsGetLiveSessionActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsGetLiveSessionActivitiesRequest struct via the builder pattern


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


## LiveSessionsGetLiveSessionProperties

> map[string]interface{} LiveSessionsGetLiveSessionProperties(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsGetLiveSessionProperties(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsGetLiveSessionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiveSessionsGetLiveSessionProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LiveSessionsApi.LiveSessionsGetLiveSessionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsGetLiveSessionPropertiesRequest struct via the builder pattern


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


## LiveSessionsGetLiveSessions

> LiveSession LiveSessionsGetLiveSessions(ctx).ReturnURL(returnURL).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Activities(activities).FromTime(fromTime).ToTime(toTime).Safe(safe).Execute()





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
    returnURL := "returnURL_example" // string | The returned URL. (optional)
    limit := int64(789) // int64 | Determines the number of lives sessions that are returned in the list. The maximum value is defined in the MaxRecords property in Options -> Privileged Session Management -> General Setting -> Search Properties. (optional)
    offset := int64(789) // int64 | Determines which recording results will be returned, according to a specific place in the returned list. This value defines the recording's place in the list and how many results will be skipped. (optional)
    useCache := true // bool |  (optional)
    sort := []string{"Inner_example"} // []string | The sort can be done by each property on the recording file: RiskScore, FileName, SafeName, FolderName, PSMVaultUserName, FromIP, RemoteMachine, Client, Protocol, AccountUserName, AccountAddress, AccountPlatformID, PSMStartTime, TicketID. The sort can be in ascending or descending order. To sort in descending order, specify ' - ' (dash) before the recording property by which to sort. (optional)
    search := "search_example" // string | Returns lives sessions that are filtered by properties that contain the specified search text. (optional)
    activities := "activities_example" // string | Returns lives sessions with specific activities. (optional)
    fromTime := int64(789) // int64 | Returns lives sessions from a specific date. (optional)
    toTime := int64(789) // int64 | Returns lives sessions before a specific date. (optional)
    safe := "safe_example" // string | Returns lives sessions that use accounts from a specific Safe. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsGetLiveSessions(context.Background()).ReturnURL(returnURL).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Activities(activities).FromTime(fromTime).ToTime(toTime).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsGetLiveSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiveSessionsGetLiveSessions`: LiveSession
    fmt.Fprintf(os.Stdout, "Response from `LiveSessionsApi.LiveSessionsGetLiveSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsGetLiveSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnURL** | **string** | The returned URL. | 
 **limit** | **int64** | Determines the number of lives sessions that are returned in the list. The maximum value is defined in the MaxRecords property in Options -&gt; Privileged Session Management -&gt; General Setting -&gt; Search Properties. | 
 **offset** | **int64** | Determines which recording results will be returned, according to a specific place in the returned list. This value defines the recording&#39;s place in the list and how many results will be skipped. | 
 **useCache** | **bool** |  | 
 **sort** | **[]string** | The sort can be done by each property on the recording file: RiskScore, FileName, SafeName, FolderName, PSMVaultUserName, FromIP, RemoteMachine, Client, Protocol, AccountUserName, AccountAddress, AccountPlatformID, PSMStartTime, TicketID. The sort can be in ascending or descending order. To sort in descending order, specify &#39; - &#39; (dash) before the recording property by which to sort. | 
 **search** | **string** | Returns lives sessions that are filtered by properties that contain the specified search text. | 
 **activities** | **string** | Returns lives sessions with specific activities. | 
 **fromTime** | **int64** | Returns lives sessions from a specific date. | 
 **toTime** | **int64** | Returns lives sessions before a specific date. | 
 **safe** | **string** | Returns lives sessions that use accounts from a specific Safe. | 

### Return type

[**LiveSession**](LiveSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveSessionsPSMAutoTermination

> LiveSessionsPSMAutoTermination(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsPSMAutoTermination(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsPSMAutoTermination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsPSMAutoTerminationRequest struct via the builder pattern


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


## LiveSessionsPSMMonitor

> LiveSessionsPSMMonitor(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsPSMMonitor(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsPSMMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsPSMMonitorRequest struct via the builder pattern


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


## LiveSessionsPSMSessionResume

> LiveSessionsPSMSessionResume(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsPSMSessionResume(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsPSMSessionResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsPSMSessionResumeRequest struct via the builder pattern


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


## LiveSessionsPSMSuspend

> LiveSessionsPSMSuspend(ctx, liveSessionId).Execute()





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
    liveSessionId := "liveSessionId_example" // string | The unique ID of the PSM Active Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LiveSessionsApi.LiveSessionsPSMSuspend(context.Background(), liveSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveSessionsApi.LiveSessionsPSMSuspend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveSessionId** | **string** | The unique ID of the PSM Active Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveSessionsPSMSuspendRequest struct via the builder pattern


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

