# \MyRequestsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MyRequestsCreateRequest**](MyRequestsApi.md#MyRequestsCreateRequest) | **Post** /api/MyRequests | 
[**MyRequestsDeleteMyRequests**](MyRequestsApi.md#MyRequestsDeleteMyRequests) | **Delete** /api/MyRequests/{RequestID} | 
[**MyRequestsGetRequest**](MyRequestsApi.md#MyRequestsGetRequest) | **Get** /api/MyRequests/{RequestID} | 
[**MyRequestsGetRequests**](MyRequestsApi.md#MyRequestsGetRequests) | **Get** /api/MyRequests | 



## MyRequestsCreateRequest

> Request MyRequestsCreateRequest(ctx).Request(request).Execute()





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
    request := *openapiclient.NewNewRequestBody("AccountId_example") // NewRequestBody | The request to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MyRequestsApi.MyRequestsCreateRequest(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsApi.MyRequestsCreateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRequestsCreateRequest`: Request
    fmt.Fprintf(os.Stdout, "Response from `MyRequestsApi.MyRequestsCreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMyRequestsCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**NewRequestBody**](NewRequestBody.md) | The request to create. | 

### Return type

[**Request**](Request.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyRequestsDeleteMyRequests

> map[string]interface{} MyRequestsDeleteMyRequests(ctx, requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()





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
    requestID := "requestID_example" // string | The request's unique ID, composed of the SafeName and internal RequestID.
    requestID2 := "requestID_example" // string | The request's unique ID, composed of the SafeName and internal RequestID. (optional)
    safeName := "safeName_example" // string |  (optional)
    requestorUserName := "requestorUserName_example" // string |  (optional)
    requestorReason := "requestorReason_example" // string |  (optional)
    userReason := "userReason_example" // string |  (optional)
    creationDate := int64(789) // int64 |  (optional)
    operation := "operation_example" // string |  (optional)
    expirationDate := int64(789) // int64 |  (optional)
    operationType := int32(56) // int32 |  (optional)
    accessType := "accessType_example" // string |  (optional)
    confirmationsLeft := int64(789) // int64 |  (optional)
    accessFrom := int64(789) // int64 |  (optional)
    accessTo := int64(789) // int64 |  (optional)
    status := int32(56) // int32 |  (optional)
    statusTitle := "statusTitle_example" // string |  (optional)
    invalidRequestReason := int32(56) // int32 |  (optional)
    currentConfirmationLevel := int32(56) // int32 |  (optional)
    requiredConfirmersCountLevel2 := int32(56) // int32 |  (optional)
    ticketingSystemPropertiesName := "ticketingSystemPropertiesName_example" // string |  (optional)
    ticketingSystemPropertiesNumber := "ticketingSystemPropertiesNumber_example" // string |  (optional)
    ticketingSystemPropertiesStatus := "ticketingSystemPropertiesStatus_example" // string |  (optional)
    additionalInfo := TODO // map[string]interface{} |  (optional)
    accountDetailsAccountID := "accountDetailsAccountID_example" // string |  (optional)
    accountDetailsProperties := TODO // map[string]interface{} |  (optional)
    confirmers := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MyRequestsApi.MyRequestsDeleteMyRequests(context.Background(), requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsApi.MyRequestsDeleteMyRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRequestsDeleteMyRequests`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MyRequestsApi.MyRequestsDeleteMyRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | The request&#39;s unique ID, composed of the SafeName and internal RequestID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyRequestsDeleteMyRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestID2** | **string** | The request&#39;s unique ID, composed of the SafeName and internal RequestID. | 
 **safeName** | **string** |  | 
 **requestorUserName** | **string** |  | 
 **requestorReason** | **string** |  | 
 **userReason** | **string** |  | 
 **creationDate** | **int64** |  | 
 **operation** | **string** |  | 
 **expirationDate** | **int64** |  | 
 **operationType** | **int32** |  | 
 **accessType** | **string** |  | 
 **confirmationsLeft** | **int64** |  | 
 **accessFrom** | **int64** |  | 
 **accessTo** | **int64** |  | 
 **status** | **int32** |  | 
 **statusTitle** | **string** |  | 
 **invalidRequestReason** | **int32** |  | 
 **currentConfirmationLevel** | **int32** |  | 
 **requiredConfirmersCountLevel2** | **int32** |  | 
 **ticketingSystemPropertiesName** | **string** |  | 
 **ticketingSystemPropertiesNumber** | **string** |  | 
 **ticketingSystemPropertiesStatus** | **string** |  | 
 **additionalInfo** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **accountDetailsAccountID** | **string** |  | 
 **accountDetailsProperties** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **confirmers** | **[]map[string]interface{}** |  | 

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


## MyRequestsGetRequest

> MyRequest MyRequestsGetRequest(ctx, requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()





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
    requestID := "requestID_example" // string | 
    requestID2 := "requestID_example" // string |  (optional)
    safeName := "safeName_example" // string |  (optional)
    requestorUserName := "requestorUserName_example" // string |  (optional)
    requestorReason := "requestorReason_example" // string |  (optional)
    userReason := "userReason_example" // string |  (optional)
    creationDate := int64(789) // int64 |  (optional)
    operation := "operation_example" // string |  (optional)
    expirationDate := int64(789) // int64 |  (optional)
    operationType := int32(56) // int32 |  (optional)
    accessType := "accessType_example" // string |  (optional)
    confirmationsLeft := int64(789) // int64 |  (optional)
    accessFrom := int64(789) // int64 |  (optional)
    accessTo := int64(789) // int64 |  (optional)
    status := int32(56) // int32 |  (optional)
    statusTitle := "statusTitle_example" // string |  (optional)
    invalidRequestReason := int32(56) // int32 |  (optional)
    currentConfirmationLevel := int32(56) // int32 |  (optional)
    requiredConfirmersCountLevel2 := int32(56) // int32 |  (optional)
    ticketingSystemPropertiesName := "ticketingSystemPropertiesName_example" // string |  (optional)
    ticketingSystemPropertiesNumber := "ticketingSystemPropertiesNumber_example" // string |  (optional)
    ticketingSystemPropertiesStatus := "ticketingSystemPropertiesStatus_example" // string |  (optional)
    additionalInfo := TODO // map[string]interface{} |  (optional)
    accountDetailsAccountID := "accountDetailsAccountID_example" // string |  (optional)
    accountDetailsProperties := TODO // map[string]interface{} |  (optional)
    confirmers := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MyRequestsApi.MyRequestsGetRequest(context.Background(), requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsApi.MyRequestsGetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRequestsGetRequest`: MyRequest
    fmt.Fprintf(os.Stdout, "Response from `MyRequestsApi.MyRequestsGetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMyRequestsGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestID2** | **string** |  | 
 **safeName** | **string** |  | 
 **requestorUserName** | **string** |  | 
 **requestorReason** | **string** |  | 
 **userReason** | **string** |  | 
 **creationDate** | **int64** |  | 
 **operation** | **string** |  | 
 **expirationDate** | **int64** |  | 
 **operationType** | **int32** |  | 
 **accessType** | **string** |  | 
 **confirmationsLeft** | **int64** |  | 
 **accessFrom** | **int64** |  | 
 **accessTo** | **int64** |  | 
 **status** | **int32** |  | 
 **statusTitle** | **string** |  | 
 **invalidRequestReason** | **int32** |  | 
 **currentConfirmationLevel** | **int32** |  | 
 **requiredConfirmersCountLevel2** | **int32** |  | 
 **ticketingSystemPropertiesName** | **string** |  | 
 **ticketingSystemPropertiesNumber** | **string** |  | 
 **ticketingSystemPropertiesStatus** | **string** |  | 
 **additionalInfo** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **accountDetailsAccountID** | **string** |  | 
 **accountDetailsProperties** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **confirmers** | **[]map[string]interface{}** |  | 

### Return type

[**MyRequest**](MyRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MyRequestsGetRequests

> MyRequest MyRequestsGetRequests(ctx).OnlyWaiting(onlyWaiting).Expired(expired).Execute()





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
    onlyWaiting := true // bool | \"Only requests waiting for approval will be listed. (optional)
    expired := true // bool | Expired requests will be included in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MyRequestsApi.MyRequestsGetRequests(context.Background()).OnlyWaiting(onlyWaiting).Expired(expired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsApi.MyRequestsGetRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MyRequestsGetRequests`: MyRequest
    fmt.Fprintf(os.Stdout, "Response from `MyRequestsApi.MyRequestsGetRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMyRequestsGetRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyWaiting** | **bool** | \&quot;Only requests waiting for approval will be listed. | 
 **expired** | **bool** | Expired requests will be included in the list. | 

### Return type

[**MyRequest**](MyRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

