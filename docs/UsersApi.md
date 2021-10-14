# \UsersApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActivateUser**](UsersApi.md#UsersActivateUser) | **Post** /api/Users/{userID}/Activate | 
[**UsersAddUserSSHKey**](UsersApi.md#UsersAddUserSSHKey) | **Post** /api/Users/{userID}/Secret/SSHKeys | 
[**UsersCreateUser**](UsersApi.md#UsersCreateUser) | **Post** /api/Users | 
[**UsersDeleteUser**](UsersApi.md#UsersDeleteUser) | **Delete** /api/Users/{userID} | 
[**UsersDeleteUserSSHKey**](UsersApi.md#UsersDeleteUserSSHKey) | **Delete** /api/Users/{userID}/Secret/SSHKeys | 
[**UsersDestroyAllUsersCachedSSHKey**](UsersApi.md#UsersDestroyAllUsersCachedSSHKey) | **Delete** /api/Users/Secret/SSHKeys/ClearCache | 
[**UsersDestroyCurrentUserCachedSSHKey**](UsersApi.md#UsersDestroyCurrentUserCachedSSHKey) | **Delete** /api/Users/Secret/SSHKeys/Cache | 
[**UsersDestroyUserCachedSSHKey**](UsersApi.md#UsersDestroyUserCachedSSHKey) | **Delete** /api/Users/{userID}/Secret/SSHKeys/Cache | 
[**UsersEditUser**](UsersApi.md#UsersEditUser) | **Put** /api/Users/{userID} | 
[**UsersGenerateCurrentUserCachedSSHKey**](UsersApi.md#UsersGenerateCurrentUserCachedSSHKey) | **Post** /api/Users/Secret/SSHKeys/Cache | 
[**UsersGenerateUserCachedSSHKey**](UsersApi.md#UsersGenerateUserCachedSSHKey) | **Post** /api/Users/{userID}/Secret/SSHKeys/Cache | 
[**UsersGetUserDetails**](UsersApi.md#UsersGetUserDetails) | **Get** /api/Users/{userID} | 
[**UsersGetUserSSHKeys**](UsersApi.md#UsersGetUserSSHKeys) | **Get** /api/Users/{userID}/Secret/SSHKeys | 
[**UsersGetUsers**](UsersApi.md#UsersGetUsers) | **Get** /api/Users | 
[**UsersResetUserPassword**](UsersApi.md#UsersResetUserPassword) | **Post** /api/Users/{userID}/ResetPassword | 



## UsersActivateUser

> UsersActivateUser(ctx, userID).Execute()





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
    userID := "userID_example" // string | The ID of the user to activate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersActivateUser(context.Background(), userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersActivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user to activate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivateUserRequest struct via the builder pattern


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


## UsersAddUserSSHKey

> PublicSSHKeyModel UsersAddUserSSHKey(ctx, userID).SshKeyToAdd(sshKeyToAdd).Execute()





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
    userID := "userID_example" // string | The ID of the user to add the ssh key to.
    sshKeyToAdd := *openapiclient.NewAddUserSSHKeyModel("PublicKey_example") // AddUserSSHKeyModel | The ssh key to add to the user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersAddUserSSHKey(context.Background(), userID).SshKeyToAdd(sshKeyToAdd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersAddUserSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAddUserSSHKey`: PublicSSHKeyModel
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersAddUserSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user to add the ssh key to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAddUserSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeyToAdd** | [**AddUserSSHKeyModel**](AddUserSSHKeyModel.md) | The ssh key to add to the user | 

### Return type

[**PublicSSHKeyModel**](PublicSSHKeyModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersCreateUser

> User UsersCreateUser(ctx).User(user).Execute()





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
    user := *openapiclient.NewUser("Username_example") // User | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersCreateUser(context.Background()).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersCreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersCreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteUser

> UsersDeleteUser(ctx, userID).Execute()





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
    userID := "userID_example" // string | The ID of the User to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersDeleteUser(context.Background(), userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the User to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteUserRequest struct via the builder pattern


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


## UsersDeleteUserSSHKey

> UsersDeleteUserSSHKey(ctx, userID).SshKeyToDel(sshKeyToDel).Execute()





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
    userID := "userID_example" // string | The ID of the user to delete the public ssh key from.
    sshKeyToDel := *openapiclient.NewDeleteUserSSHKeyModel("KeyFingerprint_example") // DeleteUserSSHKeyModel | The SSH Key fingerprint to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersDeleteUserSSHKey(context.Background(), userID).SshKeyToDel(sshKeyToDel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDeleteUserSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user to delete the public ssh key from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteUserSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sshKeyToDel** | [**DeleteUserSSHKeyModel**](DeleteUserSSHKeyModel.md) | The SSH Key fingerprint to delete | 

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


## UsersDestroyAllUsersCachedSSHKey

> UsersDestroyAllUsersCachedSSHKey(ctx).Search(search).UserType(userType).ComponentUser(componentUser).Execute()





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
    userType := "userType_example" // string | The type of the user. (optional)
    componentUser := true // bool | Whether the user is a known component or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersDestroyAllUsersCachedSSHKey(context.Background()).Search(search).UserType(userType).ComponentUser(componentUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDestroyAllUsersCachedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroyAllUsersCachedSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **userType** | **string** | The type of the user. | 
 **componentUser** | **bool** | Whether the user is a known component or not. | 

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


## UsersDestroyCurrentUserCachedSSHKey

> UsersDestroyCurrentUserCachedSSHKey(ctx).Execute()





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
    resp, r, err := api_client.UsersApi.UsersDestroyCurrentUserCachedSSHKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDestroyCurrentUserCachedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroyCurrentUserCachedSSHKeyRequest struct via the builder pattern


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


## UsersDestroyUserCachedSSHKey

> UsersDestroyUserCachedSSHKey(ctx, userID).Execute()





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
    userID := "userID_example" // string | The ID of the User for which the cached ssh key will be destroyed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersDestroyUserCachedSSHKey(context.Background(), userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDestroyUserCachedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the User for which the cached ssh key will be destroyed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroyUserCachedSSHKeyRequest struct via the builder pattern


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


## UsersEditUser

> User UsersEditUser(ctx, userID).User(user).Execute()





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
    userID := "userID_example" // string | The ID of the user to update.
    user := *openapiclient.NewUser("Username_example") // User | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersEditUser(context.Background(), userID).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersEditUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersEditUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersEditUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersEditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGenerateCurrentUserCachedSSHKey

> GenerateUserCachedSSHKeyResponse UsersGenerateCurrentUserCachedSSHKey(ctx).GenModel(genModel).Execute()





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
    genModel := *openapiclient.NewGenerateUserCachedSSHKeyModel() // GenerateUserCachedSSHKeyModel | The information as to how to format and output the generated ssh key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGenerateCurrentUserCachedSSHKey(context.Background()).GenModel(genModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGenerateCurrentUserCachedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGenerateCurrentUserCachedSSHKey`: GenerateUserCachedSSHKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGenerateCurrentUserCachedSSHKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGenerateCurrentUserCachedSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genModel** | [**GenerateUserCachedSSHKeyModel**](GenerateUserCachedSSHKeyModel.md) | The information as to how to format and output the generated ssh key | 

### Return type

[**GenerateUserCachedSSHKeyResponse**](GenerateUserCachedSSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGenerateUserCachedSSHKey

> GenerateUserCachedSSHKeyResponse UsersGenerateUserCachedSSHKey(ctx, userID).GenModel(genModel).Execute()





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
    userID := "userID_example" // string | The ID of the User for which the key will be generated
    genModel := *openapiclient.NewGenerateUserCachedSSHKeyModel() // GenerateUserCachedSSHKeyModel | The information as to how to format and output the generated ssh key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGenerateUserCachedSSHKey(context.Background(), userID).GenModel(genModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGenerateUserCachedSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGenerateUserCachedSSHKey`: GenerateUserCachedSSHKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGenerateUserCachedSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the User for which the key will be generated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGenerateUserCachedSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genModel** | [**GenerateUserCachedSSHKeyModel**](GenerateUserCachedSSHKeyModel.md) | The information as to how to format and output the generated ssh key | 

### Return type

[**GenerateUserCachedSSHKeyResponse**](GenerateUserCachedSSHKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUserDetails

> User UsersGetUserDetails(ctx, userID).Execute()





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
    userID := "userID_example" // string | The ID of the User for which information is returned.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGetUserDetails(context.Background(), userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetUserDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetUserDetails`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetUserDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the User for which information is returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUserSSHKeys

> GetUserSSHKeysResponse UsersGetUserSSHKeys(ctx, userID).KeyStoreTypeFilter(keyStoreTypeFilter).Execute()





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
    userID := "userID_example" // string | The ID of the user for which public ssh keys are returned.
    keyStoreTypeFilter := int32(56) // int32 | The key type filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGetUserSSHKeys(context.Background(), userID).KeyStoreTypeFilter(keyStoreTypeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetUserSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetUserSSHKeys`: GetUserSSHKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetUserSSHKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user for which public ssh keys are returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetUserSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyStoreTypeFilter** | **int32** | The key type filter | 

### Return type

[**GetUserSSHKeysResponse**](GetUserSSHKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetUsers

> []BaseUser UsersGetUsers(ctx).ExtendedDetails(extendedDetails).Search(search).Sort(sort).UserName(userName).UserType(userType).ComponentUser(componentUser).Execute()





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
    extendedDetails := true // bool | returns users group membership if true (optional)
    search := "search_example" // string |  (optional)
    sort := []string{"Inner_example"} // []string |  (optional)
    userName := "userName_example" // string | The name of the user. (optional)
    userType := "userType_example" // string | The type of the user. (optional)
    componentUser := true // bool | Whether the user is a known component or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersGetUsers(context.Background()).ExtendedDetails(extendedDetails).Search(search).Sort(sort).UserName(userName).UserType(userType).ComponentUser(componentUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetUsers`: []BaseUser
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedDetails** | **bool** | returns users group membership if true | 
 **search** | **string** |  | 
 **sort** | **[]string** |  | 
 **userName** | **string** | The name of the user. | 
 **userType** | **string** | The type of the user. | 
 **componentUser** | **bool** | Whether the user is a known component or not. | 

### Return type

[**[]BaseUser**](BaseUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersResetUserPassword

> UsersResetUserPassword(ctx, userID).ResetUserPassword(resetUserPassword).Execute()





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
    userID := "userID_example" // string | The ID of the user whose password will be reset.
    resetUserPassword := *openapiclient.NewResetUserPassword("NewPassword_example") // ResetUserPassword | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersResetUserPassword(context.Background(), userID).ResetUserPassword(resetUserPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersResetUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The ID of the user whose password will be reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersResetUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetUserPassword** | [**ResetUserPassword**](ResetUserPassword.md) |  | 

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

