# \AccountGroupsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountGroupsAddAccountGroup**](AccountGroupsApi.md#AccountGroupsAddAccountGroup) | **Post** /api/AccountGroups | 
[**AccountGroupsAddAccountToGroup**](AccountGroupsApi.md#AccountGroupsAddAccountToGroup) | **Post** /api/AccountGroups/{groupId}/members | 
[**AccountGroupsDeleteAccountFromGroup**](AccountGroupsApi.md#AccountGroupsDeleteAccountFromGroup) | **Delete** /api/AccountGroups/{groupId}/members/{accountId} | 
[**AccountGroupsGetGroupAccounts**](AccountGroupsApi.md#AccountGroupsGetGroupAccounts) | **Get** /api/AccountGroups/{groupId}/members | 
[**AccountGroupsGetGroups**](AccountGroupsApi.md#AccountGroupsGetGroups) | **Get** /api/AccountGroups | 



## AccountGroupsAddAccountGroup

> AccountGroup AccountGroupsAddAccountGroup(ctx).AddAccountGroupData(addAccountGroupData).Execute()





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
    addAccountGroupData := *openapiclient.NewAddAccountGroupData("GroupName_example", "GroupPlatformId_example", "Safe_example") // AddAccountGroupData | The account group data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AccountGroupsAddAccountGroup(context.Background()).AddAccountGroupData(addAccountGroupData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AccountGroupsAddAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGroupsAddAccountGroup`: AccountGroup
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsApi.AccountGroupsAddAccountGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountGroupsAddAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAccountGroupData** | [**AddAccountGroupData**](AddAccountGroupData.md) | The account group data | 

### Return type

[**AccountGroup**](AccountGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGroupsAddAccountToGroup

> AddMemberToAccountGroup AccountGroupsAddAccountToGroup(ctx, groupId).AddMemberData(addMemberData).Execute()





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
    groupId := "groupId_example" // string | The unique ID of account group.
    addMemberData := *openapiclient.NewAddMemberToAccountGroup("AccountId_example") // AddMemberToAccountGroup | The member to add

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AccountGroupsAddAccountToGroup(context.Background(), groupId).AddMemberData(addMemberData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AccountGroupsAddAccountToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGroupsAddAccountToGroup`: AddMemberToAccountGroup
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsApi.AccountGroupsAddAccountToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique ID of account group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGroupsAddAccountToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addMemberData** | [**AddMemberToAccountGroup**](AddMemberToAccountGroup.md) | The member to add | 

### Return type

[**AddMemberToAccountGroup**](AddMemberToAccountGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGroupsDeleteAccountFromGroup

> map[string]interface{} AccountGroupsDeleteAccountFromGroup(ctx, groupId, accountId).Execute()





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
    groupId := "groupId_example" // string | The unique ID of the group.
    accountId := "accountId_example" // string | The unique ID of the account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AccountGroupsDeleteAccountFromGroup(context.Background(), groupId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AccountGroupsDeleteAccountFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGroupsDeleteAccountFromGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsApi.AccountGroupsDeleteAccountFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique ID of the group. | 
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGroupsDeleteAccountFromGroupRequest struct via the builder pattern


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


## AccountGroupsGetGroupAccounts

> AccountGroupMember AccountGroupsGetGroupAccounts(ctx, groupId).Execute()





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
    groupId := "groupId_example" // string | The unique ID of the group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AccountGroupsGetGroupAccounts(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AccountGroupsGetGroupAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGroupsGetGroupAccounts`: AccountGroupMember
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsApi.AccountGroupsGetGroupAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountGroupsGetGroupAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountGroupMember**](AccountGroupMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGroupsGetGroups

> AccountGroup AccountGroupsGetGroups(ctx).Safe(safe).Execute()





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
    safe := "safe_example" // string | The name of the Safe where the account groups are. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountGroupsApi.AccountGroupsGetGroups(context.Background()).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsApi.AccountGroupsGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGroupsGetGroups`: AccountGroup
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsApi.AccountGroupsGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountGroupsGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safe** | **string** | The name of the Safe where the account groups are. | 

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

