# \RecordingsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordingsGetRecording**](RecordingsApi.md#RecordingsGetRecording) | **Get** /api/Recordings/{recordingId} | 
[**RecordingsGetRecordingActivities**](RecordingsApi.md#RecordingsGetRecordingActivities) | **Get** /api/Recordings/{recordingId}/activities | 
[**RecordingsGetRecordingFileValidityStatus**](RecordingsApi.md#RecordingsGetRecordingFileValidityStatus) | **Get** /api/Recordings/{recordingId}/valid | 
[**RecordingsGetRecordingProperties**](RecordingsApi.md#RecordingsGetRecordingProperties) | **Get** /api/Recordings/{recordingId}/properties | 
[**RecordingsGetRecordings**](RecordingsApi.md#RecordingsGetRecordings) | **Get** /api/Recordings | 
[**RecordingsPlayRecording**](RecordingsApi.md#RecordingsPlayRecording) | **Post** /api/Recordings/{recordingId}/Play | 



## RecordingsGetRecording

> map[string]interface{} RecordingsGetRecording(ctx, recordingId).ReturnURL(returnURL).Execute()





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
    recordingId := "recordingId_example" // string | The unique ID of the PSM Recorded Session.
    returnURL := "returnURL_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsGetRecording(context.Background(), recordingId).ReturnURL(returnURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsGetRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsGetRecording`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsGetRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | The unique ID of the PSM Recorded Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsGetRecordingRequest struct via the builder pattern


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


## RecordingsGetRecordingActivities

> map[string]interface{} RecordingsGetRecordingActivities(ctx, recordingId).Execute()





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
    recordingId := "recordingId_example" // string | The unique ID of the PSM Recorded Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsGetRecordingActivities(context.Background(), recordingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsGetRecordingActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsGetRecordingActivities`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsGetRecordingActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | The unique ID of the PSM Recorded Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsGetRecordingActivitiesRequest struct via the builder pattern


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


## RecordingsGetRecordingFileValidityStatus

> map[string]interface{} RecordingsGetRecordingFileValidityStatus(ctx, recordingId).Execute()



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
    recordingId := "recordingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsGetRecordingFileValidityStatus(context.Background(), recordingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsGetRecordingFileValidityStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsGetRecordingFileValidityStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsGetRecordingFileValidityStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsGetRecordingFileValidityStatusRequest struct via the builder pattern


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


## RecordingsGetRecordingProperties

> map[string]interface{} RecordingsGetRecordingProperties(ctx, recordingId).ReturnURL(returnURL).Execute()





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
    recordingId := "recordingId_example" // string | The unique ID of the PSM Recorded Session.
    returnURL := "returnURL_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsGetRecordingProperties(context.Background(), recordingId).ReturnURL(returnURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsGetRecordingProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsGetRecordingProperties`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsGetRecordingProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | The unique ID of the PSM Recorded Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsGetRecordingPropertiesRequest struct via the builder pattern


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


## RecordingsGetRecordings

> SessionData RecordingsGetRecordings(ctx).ReturnURL(returnURL).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Activities(activities).FromTime(fromTime).ToTime(toTime).Safe(safe).Execute()





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
    limit := int64(789) // int64 | Determines the number of recordings that are returned in the list. The maximum value is defined in the MaxRecords property in Options -> Privileged Session Management -> General Setting -> Search Properties. (optional)
    offset := int64(789) // int64 | Determines which recording results will be returned, according to a specific place in the returned list. This value defines the recording's place in the list and how many results will be skipped. (optional)
    useCache := true // bool |  (optional)
    sort := []string{"Inner_example"} // []string | The sort can be done by each property on the recording file: RiskScore, FileName, SafeName, FolderName, PSMVaultUserName, FromIP, RemoteMachine, Client, Protocol, AccountUserName, AccountAddress, AccountPlatformID, PSMStartTime, TicketID. The sort can be in ascending or descending order. To sort in descending order, specify ' - ' (dash) before the recording property by which to sort. (optional)
    search := "search_example" // string | Returns recordings that are filtered by properties that contain the specified search text. (optional)
    activities := "activities_example" // string | Returns recordings with specific activities. (optional)
    fromTime := int64(789) // int64 | Returns recordings from a specific date. (optional)
    toTime := int64(789) // int64 | Returns recordings before a specific date. (optional)
    safe := "safe_example" // string | Returns recordings from a specific Safe. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsGetRecordings(context.Background()).ReturnURL(returnURL).Limit(limit).Offset(offset).UseCache(useCache).Sort(sort).Search(search).Activities(activities).FromTime(fromTime).ToTime(toTime).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsGetRecordings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsGetRecordings`: SessionData
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsGetRecordings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsGetRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnURL** | **string** | The returned URL. | 
 **limit** | **int64** | Determines the number of recordings that are returned in the list. The maximum value is defined in the MaxRecords property in Options -&gt; Privileged Session Management -&gt; General Setting -&gt; Search Properties. | 
 **offset** | **int64** | Determines which recording results will be returned, according to a specific place in the returned list. This value defines the recording&#39;s place in the list and how many results will be skipped. | 
 **useCache** | **bool** |  | 
 **sort** | **[]string** | The sort can be done by each property on the recording file: RiskScore, FileName, SafeName, FolderName, PSMVaultUserName, FromIP, RemoteMachine, Client, Protocol, AccountUserName, AccountAddress, AccountPlatformID, PSMStartTime, TicketID. The sort can be in ascending or descending order. To sort in descending order, specify &#39; - &#39; (dash) before the recording property by which to sort. | 
 **search** | **string** | Returns recordings that are filtered by properties that contain the specified search text. | 
 **activities** | **string** | Returns recordings with specific activities. | 
 **fromTime** | **int64** | Returns recordings from a specific date. | 
 **toTime** | **int64** | Returns recordings before a specific date. | 
 **safe** | **string** | Returns recordings from a specific Safe. | 

### Return type

[**SessionData**](SessionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordingsPlayRecording

> map[string]interface{} RecordingsPlayRecording(ctx, recordingId).Execute()





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
    recordingId := "recordingId_example" // string | The unique ID of the PSM Recorded Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordingsApi.RecordingsPlayRecording(context.Background(), recordingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordingsApi.RecordingsPlayRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordingsPlayRecording`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordingsApi.RecordingsPlayRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string** | The unique ID of the PSM Recorded Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordingsPlayRecordingRequest struct via the builder pattern


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

