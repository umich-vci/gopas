# \IncomingRequestsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IncomingRequestsConfirmRequests**](IncomingRequestsApi.md#IncomingRequestsConfirmRequests) | **Post** /api/IncomingRequests/{RequestID}/confirm | 
[**IncomingRequestsGetRequest**](IncomingRequestsApi.md#IncomingRequestsGetRequest) | **Get** /api/IncomingRequests/{RequestID} | 
[**IncomingRequestsGetRequests**](IncomingRequestsApi.md#IncomingRequestsGetRequests) | **Get** /api/IncomingRequests | 
[**IncomingRequestsRejectRequests**](IncomingRequestsApi.md#IncomingRequestsRejectRequests) | **Post** /api/IncomingRequests/{RequestID}/reject | 



## IncomingRequestsConfirmRequests

> IncomingRequestsConfirmRequests(ctx, requestID).ConfirmationInfo(confirmationInfo).RequestorFullName(requestorFullName).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()





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
    requestID := "requestID_example" // string | The unique ID of the request to confirm.
    confirmationInfo := *openapiclient.NewConfirmRejectData() // ConfirmRejectData | The confirmation info.
    requestorFullName := "requestorFullName_example" // string |  (optional)
    requestID2 := "requestID_example" // string | The unique ID of the request to confirm. (optional)
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
    resp, r, err := api_client.IncomingRequestsApi.IncomingRequestsConfirmRequests(context.Background(), requestID).ConfirmationInfo(confirmationInfo).RequestorFullName(requestorFullName).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingRequestsApi.IncomingRequestsConfirmRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | The unique ID of the request to confirm. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncomingRequestsConfirmRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmationInfo** | [**ConfirmRejectData**](ConfirmRejectData.md) | The confirmation info. | 
 **requestorFullName** | **string** |  | 
 **requestID2** | **string** | The unique ID of the request to confirm. | 
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomingRequestsGetRequest

> IncomingRequest IncomingRequestsGetRequest(ctx, requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()





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
    requestID := "requestID_example" // string | The unique ID of the request.
    requestID2 := "requestID_example" // string | The unique ID of the request. (optional)
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
    resp, r, err := api_client.IncomingRequestsApi.IncomingRequestsGetRequest(context.Background(), requestID).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingRequestsApi.IncomingRequestsGetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomingRequestsGetRequest`: IncomingRequest
    fmt.Fprintf(os.Stdout, "Response from `IncomingRequestsApi.IncomingRequestsGetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | The unique ID of the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncomingRequestsGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestID2** | **string** | The unique ID of the request. | 
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

[**IncomingRequest**](IncomingRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomingRequestsGetRequests

> IncomingRequest IncomingRequestsGetRequests(ctx).OnlyWaiting(onlyWaiting).Expired(expired).Execute()





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
    onlyWaiting := true // bool | Only requests waiting for approval will be listed. (optional)
    expired := true // bool | Expired requests will be included in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncomingRequestsApi.IncomingRequestsGetRequests(context.Background()).OnlyWaiting(onlyWaiting).Expired(expired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingRequestsApi.IncomingRequestsGetRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomingRequestsGetRequests`: IncomingRequest
    fmt.Fprintf(os.Stdout, "Response from `IncomingRequestsApi.IncomingRequestsGetRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomingRequestsGetRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyWaiting** | **bool** | Only requests waiting for approval will be listed. | 
 **expired** | **bool** | Expired requests will be included in the list. | 

### Return type

[**IncomingRequest**](IncomingRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomingRequestsRejectRequests

> IncomingRequestsRejectRequests(ctx, requestID).RejectionInfo(rejectionInfo).RequestorFullName(requestorFullName).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()





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
    requestID := "requestID_example" // string | The unique ID of the request to reject.
    rejectionInfo := *openapiclient.NewConfirmRejectData() // ConfirmRejectData | The confirmation info.
    requestorFullName := "requestorFullName_example" // string |  (optional)
    requestID2 := "requestID_example" // string | The unique ID of the request to reject. (optional)
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
    resp, r, err := api_client.IncomingRequestsApi.IncomingRequestsRejectRequests(context.Background(), requestID).RejectionInfo(rejectionInfo).RequestorFullName(requestorFullName).RequestID2(requestID2).SafeName(safeName).RequestorUserName(requestorUserName).RequestorReason(requestorReason).UserReason(userReason).CreationDate(creationDate).Operation(operation).ExpirationDate(expirationDate).OperationType(operationType).AccessType(accessType).ConfirmationsLeft(confirmationsLeft).AccessFrom(accessFrom).AccessTo(accessTo).Status(status).StatusTitle(statusTitle).InvalidRequestReason(invalidRequestReason).CurrentConfirmationLevel(currentConfirmationLevel).RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2).TicketingSystemPropertiesName(ticketingSystemPropertiesName).TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber).TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus).AdditionalInfo(additionalInfo).AccountDetailsAccountID(accountDetailsAccountID).AccountDetailsProperties(accountDetailsProperties).Confirmers(confirmers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncomingRequestsApi.IncomingRequestsRejectRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestID** | **string** | The unique ID of the request to reject. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncomingRequestsRejectRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectionInfo** | [**ConfirmRejectData**](ConfirmRejectData.md) | The confirmation info. | 
 **requestorFullName** | **string** |  | 
 **requestID2** | **string** | The unique ID of the request to reject. | 
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

