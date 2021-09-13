# \PlatformsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformsActivateGroupPlatform**](PlatformsApi.md#PlatformsActivateGroupPlatform) | **Post** /api/Platforms/Groups/{platformID}/activate | 
[**PlatformsActivateRotationalGroupPlatform**](PlatformsApi.md#PlatformsActivateRotationalGroupPlatform) | **Post** /api/Platforms/RotationalGroups/{platformID}/activate | 
[**PlatformsActivateTargetPlatform**](PlatformsApi.md#PlatformsActivateTargetPlatform) | **Post** /api/Platforms/Targets/{platformID}/activate | 
[**PlatformsDeactivateGroupPlatform**](PlatformsApi.md#PlatformsDeactivateGroupPlatform) | **Post** /api/Platforms/Groups/{platformID}/deactivate | 
[**PlatformsDeactivateRotationalGroupPlatform**](PlatformsApi.md#PlatformsDeactivateRotationalGroupPlatform) | **Post** /api/Platforms/RotationalGroups/{platformID}/deactivate | 
[**PlatformsDeactivateTargetPlatform**](PlatformsApi.md#PlatformsDeactivateTargetPlatform) | **Post** /api/Platforms/Targets/{platformID}/deactivate | 
[**PlatformsDeleteDependentPlatform**](PlatformsApi.md#PlatformsDeleteDependentPlatform) | **Delete** /api/Platforms/Dependents/{platformID} | 
[**PlatformsDeleteGroupPlatform**](PlatformsApi.md#PlatformsDeleteGroupPlatform) | **Delete** /api/Platforms/Groups/{platformID} | 
[**PlatformsDeleteRotationalGroupPlatform**](PlatformsApi.md#PlatformsDeleteRotationalGroupPlatform) | **Delete** /api/Platforms/RotationalGroups/{platformID} | 
[**PlatformsDeleteTargetPlatform**](PlatformsApi.md#PlatformsDeleteTargetPlatform) | **Delete** /api/Platforms/Targets/{platformID} | 
[**PlatformsDuplicateDependentPlatform**](PlatformsApi.md#PlatformsDuplicateDependentPlatform) | **Post** /api/Platforms/Dependents/{platformID}/Duplicate | 
[**PlatformsDuplicateGroupPlatform**](PlatformsApi.md#PlatformsDuplicateGroupPlatform) | **Post** /api/Platforms/Groups/{platformID}/Duplicate | 
[**PlatformsDuplicateRotationalGroupPlatform**](PlatformsApi.md#PlatformsDuplicateRotationalGroupPlatform) | **Post** /api/Platforms/RotationalGroups/{platformID}/Duplicate | 
[**PlatformsDuplicateTargetPlatform**](PlatformsApi.md#PlatformsDuplicateTargetPlatform) | **Post** /api/Platforms/Targets/{platformID}/Duplicate | 
[**PlatformsExport**](PlatformsApi.md#PlatformsExport) | **Post** /api/Platforms/{PlatformID}/Export | 
[**PlatformsExportDependentPlatform**](PlatformsApi.md#PlatformsExportDependentPlatform) | **Post** /api/Platforms/Dependents/{platformID}/Export | 
[**PlatformsExportGroupPlatform**](PlatformsApi.md#PlatformsExportGroupPlatform) | **Post** /api/Platforms/Groups/{platformID}/Export | 
[**PlatformsExportRotationalGroupPlatform**](PlatformsApi.md#PlatformsExportRotationalGroupPlatform) | **Post** /api/Platforms/RotationalGroups/{platformID}/Export | 
[**PlatformsExportTargetPlatform**](PlatformsApi.md#PlatformsExportTargetPlatform) | **Post** /api/Platforms/Targets/{platformID}/Export | 
[**PlatformsGetDependentPlaforms**](PlatformsApi.md#PlatformsGetDependentPlaforms) | **Get** /api/Platforms/Dependents | 
[**PlatformsGetGroupPlaforms**](PlatformsApi.md#PlatformsGetGroupPlaforms) | **Get** /api/Platforms/Groups | 
[**PlatformsGetPlaform**](PlatformsApi.md#PlatformsGetPlaform) | **Get** /api/Platforms/{id} | 
[**PlatformsGetPlatforms**](PlatformsApi.md#PlatformsGetPlatforms) | **Get** /api/Platforms | 
[**PlatformsGetPrivilegedSessionManagementDetails**](PlatformsApi.md#PlatformsGetPrivilegedSessionManagementDetails) | **Get** /api/Platforms/Targets/{platformID}/PrivilegedSessionManagement | 
[**PlatformsGetRotationalGroupPlaforms**](PlatformsApi.md#PlatformsGetRotationalGroupPlaforms) | **Get** /api/Platforms/RotationalGroups | 
[**PlatformsGetSafesByPlatform**](PlatformsApi.md#PlatformsGetSafesByPlatform) | **Get** /api/Platforms/{PlatformId}/Safes | 
[**PlatformsGetSystemTypes**](PlatformsApi.md#PlatformsGetSystemTypes) | **Get** /api/Platforms/Targets/SystemTypes | 
[**PlatformsGetTargetPlaforms**](PlatformsApi.md#PlatformsGetTargetPlaforms) | **Get** /api/Platforms/Targets | 
[**PlatformsImport**](PlatformsApi.md#PlatformsImport) | **Post** /api/Platforms/import | 
[**PlatformsSetPrivilegedSessionManagementDetails**](PlatformsApi.md#PlatformsSetPrivilegedSessionManagementDetails) | **Put** /api/Platforms/Targets/{platformID}/PrivilegedSessionManagement | 



## PlatformsActivateGroupPlatform

> map[string]interface{} PlatformsActivateGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsActivateGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsActivateGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsActivateGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsActivateGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsActivateGroupPlatformRequest struct via the builder pattern


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


## PlatformsActivateRotationalGroupPlatform

> map[string]interface{} PlatformsActivateRotationalGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the rotational group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsActivateRotationalGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsActivateRotationalGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsActivateRotationalGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsActivateRotationalGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the rotational group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsActivateRotationalGroupPlatformRequest struct via the builder pattern


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


## PlatformsActivateTargetPlatform

> map[string]interface{} PlatformsActivateTargetPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsActivateTargetPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsActivateTargetPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsActivateTargetPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsActivateTargetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsActivateTargetPlatformRequest struct via the builder pattern


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


## PlatformsDeactivateGroupPlatform

> map[string]interface{} PlatformsDeactivateGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeactivateGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeactivateGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeactivateGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeactivateGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeactivateGroupPlatformRequest struct via the builder pattern


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


## PlatformsDeactivateRotationalGroupPlatform

> map[string]interface{} PlatformsDeactivateRotationalGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the rotational group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeactivateRotationalGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeactivateRotationalGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeactivateRotationalGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeactivateRotationalGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the rotational group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeactivateRotationalGroupPlatformRequest struct via the builder pattern


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


## PlatformsDeactivateTargetPlatform

> map[string]interface{} PlatformsDeactivateTargetPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeactivateTargetPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeactivateTargetPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeactivateTargetPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeactivateTargetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeactivateTargetPlatformRequest struct via the builder pattern


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


## PlatformsDeleteDependentPlatform

> map[string]interface{} PlatformsDeleteDependentPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the dependent platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeleteDependentPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeleteDependentPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeleteDependentPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeleteDependentPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the dependent platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeleteDependentPlatformRequest struct via the builder pattern


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


## PlatformsDeleteGroupPlatform

> map[string]interface{} PlatformsDeleteGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeleteGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeleteGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeleteGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeleteGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeleteGroupPlatformRequest struct via the builder pattern


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


## PlatformsDeleteRotationalGroupPlatform

> map[string]interface{} PlatformsDeleteRotationalGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the rotational group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeleteRotationalGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeleteRotationalGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeleteRotationalGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeleteRotationalGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the rotational group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeleteRotationalGroupPlatformRequest struct via the builder pattern


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


## PlatformsDeleteTargetPlatform

> map[string]interface{} PlatformsDeleteTargetPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDeleteTargetPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDeleteTargetPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDeleteTargetPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDeleteTargetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDeleteTargetPlatformRequest struct via the builder pattern


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


## PlatformsDuplicateDependentPlatform

> map[string]interface{} PlatformsDuplicateDependentPlatform(ctx, platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the source dependent platform from which to duplicate
    duplicatePlatformDetails := *openapiclient.NewDuplicatePlatform("Name_example") // DuplicatePlatform | The details of the new platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDuplicateDependentPlatform(context.Background(), platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDuplicateDependentPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDuplicateDependentPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDuplicateDependentPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the source dependent platform from which to duplicate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDuplicateDependentPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlatformDetails** | [**DuplicatePlatform**](DuplicatePlatform.md) | The details of the new platform. | 

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


## PlatformsDuplicateGroupPlatform

> map[string]interface{} PlatformsDuplicateGroupPlatform(ctx, platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the source group platform from which to duplicate
    duplicatePlatformDetails := *openapiclient.NewDuplicatePlatform("Name_example") // DuplicatePlatform | The details of the new platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDuplicateGroupPlatform(context.Background(), platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDuplicateGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDuplicateGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDuplicateGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the source group platform from which to duplicate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDuplicateGroupPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlatformDetails** | [**DuplicatePlatform**](DuplicatePlatform.md) | The details of the new platform. | 

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


## PlatformsDuplicateRotationalGroupPlatform

> map[string]interface{} PlatformsDuplicateRotationalGroupPlatform(ctx, platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the source rotational group platform from which to duplicate
    duplicatePlatformDetails := *openapiclient.NewDuplicatePlatform("Name_example") // DuplicatePlatform | The details of the new platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDuplicateRotationalGroupPlatform(context.Background(), platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDuplicateRotationalGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDuplicateRotationalGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDuplicateRotationalGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the source rotational group platform from which to duplicate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDuplicateRotationalGroupPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlatformDetails** | [**DuplicatePlatform**](DuplicatePlatform.md) | The details of the new platform. | 

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


## PlatformsDuplicateTargetPlatform

> map[string]interface{} PlatformsDuplicateTargetPlatform(ctx, platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the source target platform from which to duplicate
    duplicatePlatformDetails := *openapiclient.NewDuplicatePlatform("Name_example") // DuplicatePlatform | The details of the new platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsDuplicateTargetPlatform(context.Background(), platformID).DuplicatePlatformDetails(duplicatePlatformDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsDuplicateTargetPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsDuplicateTargetPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsDuplicateTargetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the source target platform from which to duplicate | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsDuplicateTargetPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlatformDetails** | [**DuplicatePlatform**](DuplicatePlatform.md) | The details of the new platform. | 

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


## PlatformsExport

> PlatformsExport(ctx, platformID).Execute()





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
    platformID := "platformID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsExport(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsExportRequest struct via the builder pattern


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


## PlatformsExportDependentPlatform

> map[string]interface{} PlatformsExportDependentPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the dependent platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsExportDependentPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsExportDependentPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsExportDependentPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsExportDependentPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the dependent platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsExportDependentPlatformRequest struct via the builder pattern


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


## PlatformsExportGroupPlatform

> map[string]interface{} PlatformsExportGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsExportGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsExportGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsExportGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsExportGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsExportGroupPlatformRequest struct via the builder pattern


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


## PlatformsExportRotationalGroupPlatform

> map[string]interface{} PlatformsExportRotationalGroupPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the rotational group platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsExportRotationalGroupPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsExportRotationalGroupPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsExportRotationalGroupPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsExportRotationalGroupPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the rotational group platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsExportRotationalGroupPlatformRequest struct via the builder pattern


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


## PlatformsExportTargetPlatform

> map[string]interface{} PlatformsExportTargetPlatform(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsExportTargetPlatform(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsExportTargetPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsExportTargetPlatform`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsExportTargetPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsExportTargetPlatformRequest struct via the builder pattern


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


## PlatformsGetDependentPlaforms

> []DependentPlatform PlatformsGetDependentPlaforms(ctx).Search(search).Execute()





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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetDependentPlaforms(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetDependentPlaforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetDependentPlaforms`: []DependentPlatform
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetDependentPlaforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetDependentPlaformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]DependentPlatform**](DependentPlatform.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetGroupPlaforms

> []GroupPlatform PlatformsGetGroupPlaforms(ctx).Search(search).Execute()





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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetGroupPlaforms(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetGroupPlaforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetGroupPlaforms`: []GroupPlatform
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetGroupPlaforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetGroupPlaformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]GroupPlatform**](GroupPlatform.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetPlaform

> Platform PlatformsGetPlaform(ctx, id).Execute()





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
    id := "id_example" // string | The unique ID of the platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetPlaform(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetPlaform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetPlaform`: Platform
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetPlaform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetPlaformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Platform**](Platform.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetPlatforms

> []PlatformListModel PlatformsGetPlatforms(ctx).Active(active).PlatformType(platformType).Search(search).SystemType(systemType).Execute()





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
    active := true // bool |  (optional)
    platformType := int32(56) // int32 |  (optional)
    search := "search_example" // string | The search will be by Platform ID or Platform Name. (optional)
    systemType := "systemType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetPlatforms(context.Background()).Active(active).PlatformType(platformType).Search(search).SystemType(systemType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetPlatforms`: []PlatformListModel
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetPlatforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** |  | 
 **platformType** | **int32** |  | 
 **search** | **string** | The search will be by Platform ID or Platform Name. | 
 **systemType** | **string** |  | 

### Return type

[**[]PlatformListModel**](PlatformListModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetPrivilegedSessionManagementDetails

> PrivilegedSessionManagement PlatformsGetPrivilegedSessionManagementDetails(ctx, platformID).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetPrivilegedSessionManagementDetails(context.Background(), platformID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetPrivilegedSessionManagementDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetPrivilegedSessionManagementDetails`: PrivilegedSessionManagement
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetPrivilegedSessionManagementDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetPrivilegedSessionManagementDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivilegedSessionManagement**](PrivilegedSessionManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetRotationalGroupPlaforms

> []RotationalGroupPlatform PlatformsGetRotationalGroupPlaforms(ctx).Search(search).Execute()





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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetRotationalGroupPlaforms(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetRotationalGroupPlaforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetRotationalGroupPlaforms`: []RotationalGroupPlatform
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetRotationalGroupPlaforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetRotationalGroupPlaformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]RotationalGroupPlatform**](RotationalGroupPlatform.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetSafesByPlatform

> []string PlatformsGetSafesByPlatform(ctx, platformId).Safe(safe).Execute()





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
    platformId := "platformId_example" // string | 
    safe := "safe_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetSafesByPlatform(context.Background(), platformId).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetSafesByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetSafesByPlatform`: []string
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetSafesByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetSafesByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **safe** | **string** |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetSystemTypes

> []SystemType PlatformsGetSystemTypes(ctx).Execute()





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
    resp, r, err := api_client.PlatformsApi.PlatformsGetSystemTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetSystemTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetSystemTypes`: []SystemType
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetSystemTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetSystemTypesRequest struct via the builder pattern


### Return type

[**[]SystemType**](SystemType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsGetTargetPlaforms

> []TargetPlatform PlatformsGetTargetPlaforms(ctx).Filter(filter).Search(search).Execute()





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
    filter := "filter_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsGetTargetPlaforms(context.Background()).Filter(filter).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsGetTargetPlaforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsGetTargetPlaforms`: []TargetPlatform
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsGetTargetPlaforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsGetTargetPlaformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]TargetPlatform**](TargetPlatform.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsImport

> ImportPlatformResponse PlatformsImport(ctx).ImportPlatform(importPlatform).Execute()





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
    importPlatform := *openapiclient.NewImportPlatformData(string(123)) // ImportPlatformData | The ZIP file containing the platform and additional configuration parameters.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsImport(context.Background()).ImportPlatform(importPlatform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsImport`: ImportPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importPlatform** | [**ImportPlatformData**](ImportPlatformData.md) | The ZIP file containing the platform and additional configuration parameters. | 

### Return type

[**ImportPlatformResponse**](ImportPlatformResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformsSetPrivilegedSessionManagementDetails

> PrivilegedSessionManagement PlatformsSetPrivilegedSessionManagementDetails(ctx, platformID).PrivilegedSessionManagement(privilegedSessionManagement).Execute()





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
    platformID := int64(789) // int64 | Unique ID of the target platform.
    privilegedSessionManagement := *openapiclient.NewPrivilegedSessionManagement("PSMServerId_example") // PrivilegedSessionManagement | Unique ID of the psm server and list of connectors

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlatformsApi.PlatformsSetPrivilegedSessionManagementDetails(context.Background(), platformID).PrivilegedSessionManagement(privilegedSessionManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformsApi.PlatformsSetPrivilegedSessionManagementDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformsSetPrivilegedSessionManagementDetails`: PrivilegedSessionManagement
    fmt.Fprintf(os.Stdout, "Response from `PlatformsApi.PlatformsSetPrivilegedSessionManagementDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformID** | **int64** | Unique ID of the target platform. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformsSetPrivilegedSessionManagementDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privilegedSessionManagement** | [**PrivilegedSessionManagement**](PrivilegedSessionManagement.md) | Unique ID of the psm server and list of connectors | 

### Return type

[**PrivilegedSessionManagement**](PrivilegedSessionManagement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

