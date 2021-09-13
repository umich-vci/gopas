# \LDAPDirectoriesApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LDAPDirectoriesAddDirectory**](LDAPDirectoriesApi.md#LDAPDirectoriesAddDirectory) | **Post** /api/Configuration/LDAP/Directories | 
[**LDAPDirectoriesAddDirectoryMapping**](LDAPDirectoriesApi.md#LDAPDirectoriesAddDirectoryMapping) | **Post** /api/Configuration/LDAP/Directories/{directoryName}/Mappings | 
[**LDAPDirectoriesDeleteDirectory**](LDAPDirectoriesApi.md#LDAPDirectoriesDeleteDirectory) | **Delete** /api/Configuration/LDAP/Directories/{directoryName} | 
[**LDAPDirectoriesDeleteDirectoryMapping**](LDAPDirectoriesApi.md#LDAPDirectoriesDeleteDirectoryMapping) | **Delete** /api/Configuration/LDAP/Directories/{directoryName}/Mappings/{mappingID} | 
[**LDAPDirectoriesGetDirectories**](LDAPDirectoriesApi.md#LDAPDirectoriesGetDirectories) | **Get** /api/Configuration/LDAP/Directories | 
[**LDAPDirectoriesGetDirectory**](LDAPDirectoriesApi.md#LDAPDirectoriesGetDirectory) | **Get** /api/Configuration/LDAP/Directories/{directoryName} | 
[**LDAPDirectoriesGetDirectoryMapping**](LDAPDirectoriesApi.md#LDAPDirectoriesGetDirectoryMapping) | **Get** /api/Configuration/LDAP/Directories/{directoryName}/Mappings/{mappingID} | 
[**LDAPDirectoriesGetDirectoryMappings**](LDAPDirectoriesApi.md#LDAPDirectoriesGetDirectoryMappings) | **Get** /api/Configuration/LDAP/Directories/{directoryName}/Mappings | 
[**LDAPDirectoriesSetDirectoryMappingsOrder**](LDAPDirectoriesApi.md#LDAPDirectoriesSetDirectoryMappingsOrder) | **Post** /api/Configuration/LDAP/Directories/{directoryName}/Mappings/Reorder | 
[**LDAPDirectoriesUpdateDirectoryMapping**](LDAPDirectoriesApi.md#LDAPDirectoriesUpdateDirectoryMapping) | **Put** /api/Configuration/LDAP/Directories/{directoryName}/Mappings/{mappingID} | 



## LDAPDirectoriesAddDirectory

> LDAPDirectory LDAPDirectoriesAddDirectory(ctx).Directory(directory).Execute()





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
    directory := *openapiclient.NewLDAPDirectory("DirectoryType_example", []openapiclient.LDAPDomainController{*openapiclient.NewLDAPDomainController("Name_example")}, "DomainName_example", "DomainBaseContext_example") // LDAPDirectory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesAddDirectory(context.Background()).Directory(directory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesAddDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesAddDirectory`: LDAPDirectory
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesAddDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesAddDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directory** | [**LDAPDirectory**](LDAPDirectory.md) |  | 

### Return type

[**LDAPDirectory**](LDAPDirectory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesAddDirectoryMapping

> LDAPMappingData LDAPDirectoriesAddDirectoryMapping(ctx, directoryName).MappingData(mappingData).Execute()





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
    directoryName := "directoryName_example" // string | 
    mappingData := *openapiclient.NewLDAPMappingData("LDAPBranch_example", "MappingName_example") // LDAPMappingData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesAddDirectoryMapping(context.Background(), directoryName).MappingData(mappingData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesAddDirectoryMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesAddDirectoryMapping`: LDAPMappingData
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesAddDirectoryMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesAddDirectoryMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mappingData** | [**LDAPMappingData**](LDAPMappingData.md) |  | 

### Return type

[**LDAPMappingData**](LDAPMappingData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesDeleteDirectory

> LDAPDirectoriesDeleteDirectory(ctx, directoryName).Execute()





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
    directoryName := "directoryName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesDeleteDirectory(context.Background(), directoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesDeleteDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesDeleteDirectoryRequest struct via the builder pattern


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


## LDAPDirectoriesDeleteDirectoryMapping

> LDAPDirectoriesDeleteDirectoryMapping(ctx, directoryName, mappingID).Execute()





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
    directoryName := "directoryName_example" // string | The Directory name
    mappingID := int64(789) // int64 | The Mapping id to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesDeleteDirectoryMapping(context.Background(), directoryName, mappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesDeleteDirectoryMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | The Directory name | 
**mappingID** | **int64** | The Mapping id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesDeleteDirectoryMappingRequest struct via the builder pattern


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


## LDAPDirectoriesGetDirectories

> []LDAPDirectoryBase LDAPDirectoriesGetDirectories(ctx).Execute()





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
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesGetDirectories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesGetDirectories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesGetDirectories`: []LDAPDirectoryBase
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesGetDirectories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesGetDirectoriesRequest struct via the builder pattern


### Return type

[**[]LDAPDirectoryBase**](LDAPDirectoryBase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesGetDirectory

> LDAPDirectory LDAPDirectoriesGetDirectory(ctx, directoryName).Execute()





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
    directoryName := "directoryName_example" // string | The address/DNS name of the domain.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesGetDirectory(context.Background(), directoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesGetDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesGetDirectory`: LDAPDirectory
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesGetDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | The address/DNS name of the domain. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesGetDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LDAPDirectory**](LDAPDirectory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesGetDirectoryMapping

> LDAPMappingData LDAPDirectoriesGetDirectoryMapping(ctx, directoryName, mappingID).Execute()





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
    directoryName := "directoryName_example" // string | The address/DNS name of the domain.
    mappingID := int64(789) // int64 | Unique ID of the directory mapping.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMapping(context.Background(), directoryName, mappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesGetDirectoryMapping`: LDAPMappingData
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | The address/DNS name of the domain. | 
**mappingID** | **int64** | Unique ID of the directory mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesGetDirectoryMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LDAPMappingData**](LDAPMappingData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesGetDirectoryMappings

> []LDAPMappingData LDAPDirectoriesGetDirectoryMappings(ctx, directoryName).Execute()





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
    directoryName := "directoryName_example" // string | Directory name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMappings(context.Background(), directoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesGetDirectoryMappings`: []LDAPMappingData
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesGetDirectoryMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | Directory name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesGetDirectoryMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LDAPMappingData**](LDAPMappingData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LDAPDirectoriesSetDirectoryMappingsOrder

> LDAPDirectoriesSetDirectoryMappingsOrder(ctx, directoryName).MappingsOrder(mappingsOrder).Execute()





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
    directoryName := "directoryName_example" // string | Directory name
    mappingsOrder := *openapiclient.NewLDAPMappingsOrder([]int64{int64(123)}) // LDAPMappingsOrder | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesSetDirectoryMappingsOrder(context.Background(), directoryName).MappingsOrder(mappingsOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesSetDirectoryMappingsOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | Directory name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesSetDirectoryMappingsOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mappingsOrder** | [**LDAPMappingsOrder**](LDAPMappingsOrder.md) |  | 

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


## LDAPDirectoriesUpdateDirectoryMapping

> LDAPMappingData LDAPDirectoriesUpdateDirectoryMapping(ctx, directoryName, mappingID).MappingToEdit(mappingToEdit).Execute()





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
    directoryName := "directoryName_example" // string | Directory name
    mappingID := int64(789) // int64 | Mapping ID
    mappingToEdit := *openapiclient.NewLDAPMappingData("LDAPBranch_example", "MappingName_example") // LDAPMappingData | Updated Mapping data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPDirectoriesApi.LDAPDirectoriesUpdateDirectoryMapping(context.Background(), directoryName, mappingID).MappingToEdit(mappingToEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPDirectoriesApi.LDAPDirectoriesUpdateDirectoryMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LDAPDirectoriesUpdateDirectoryMapping`: LDAPMappingData
    fmt.Fprintf(os.Stdout, "Response from `LDAPDirectoriesApi.LDAPDirectoriesUpdateDirectoryMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryName** | **string** | Directory name | 
**mappingID** | **int64** | Mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLDAPDirectoriesUpdateDirectoryMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mappingToEdit** | [**LDAPMappingData**](LDAPMappingData.md) | Updated Mapping data | 

### Return type

[**LDAPMappingData**](LDAPMappingData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

