# \UserGroupsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGroupsAddMemberToGroup**](UserGroupsApi.md#UserGroupsAddMemberToGroup) | **Post** /api/UserGroups/{groupId}/Members | 
[**UserGroupsCreateUserGroup**](UserGroupsApi.md#UserGroupsCreateUserGroup) | **Post** /api/UserGroups | 
[**UserGroupsDeleteUserGroup**](UserGroupsApi.md#UserGroupsDeleteUserGroup) | **Delete** /api/UserGroups/{groupId} | 
[**UserGroupsEditUserGroup**](UserGroupsApi.md#UserGroupsEditUserGroup) | **Put** /api/UserGroups/{groupId} | 
[**UserGroupsGetUserGroups**](UserGroupsApi.md#UserGroupsGetUserGroups) | **Get** /api/UserGroups | 
[**UserGroupsRemoveUserFromGroup**](UserGroupsApi.md#UserGroupsRemoveUserFromGroup) | **Delete** /api/UserGroups/{groupId}/Members/{memberName} | 



## UserGroupsAddMemberToGroup

> GroupMember UserGroupsAddMemberToGroup(ctx, groupId).MemberToAdd(memberToAdd).Execute()





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
    groupId := int64(789) // int64 | The unique ID of the Vault group.
    memberToAdd := *openapiclient.NewGroupMember("MemberId_example", "DomainName_example") // GroupMember | The parameters for the member to be added

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsAddMemberToGroup(context.Background(), groupId).MemberToAdd(memberToAdd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsAddMemberToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsAddMemberToGroup`: GroupMember
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UserGroupsAddMemberToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The unique ID of the Vault group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsAddMemberToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberToAdd** | [**GroupMember**](GroupMember.md) | The parameters for the member to be added | 

### Return type

[**GroupMember**](GroupMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsCreateUserGroup

> map[string]interface{} UserGroupsCreateUserGroup(ctx).UserGroup(userGroup).Execute()





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
    userGroup := *openapiclient.NewBaseUserGroup("GroupName_example") // BaseUserGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsCreateUserGroup(context.Background()).UserGroup(userGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsCreateUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsCreateUserGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UserGroupsCreateUserGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsCreateUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userGroup** | [**BaseUserGroup**](BaseUserGroup.md) |  | 

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


## UserGroupsDeleteUserGroup

> map[string]interface{} UserGroupsDeleteUserGroup(ctx, groupId).Execute()





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
    groupId := int64(789) // int64 | The unique ID of the group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsDeleteUserGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsDeleteUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsDeleteUserGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UserGroupsDeleteUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The unique ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsDeleteUserGroupRequest struct via the builder pattern


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


## UserGroupsEditUserGroup

> map[string]interface{} UserGroupsEditUserGroup(ctx, groupId).Group(group).Execute()





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
    groupId := int64(789) // int64 | The unique ID of the group.
    group := *openapiclient.NewBaseUserGroup("GroupName_example") // BaseUserGroup | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsEditUserGroup(context.Background(), groupId).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsEditUserGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsEditUserGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UserGroupsEditUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The unique ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsEditUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**BaseUserGroup**](BaseUserGroup.md) |  | 

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


## UserGroupsGetUserGroups

> GetUserGroupsResponse UserGroupsGetUserGroups(ctx).Search(search).Sort(sort).Filter(filter).IncludeMembers(includeMembers).Execute()





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
    sort := []string{"Inner_example"} // []string |  (optional)
    filter := "filter_example" // string |  (optional)
    includeMembers := true // bool | Whether or not to return members for each User Group as part of the response. If not sent, the value will be False. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsGetUserGroups(context.Background()).Search(search).Sort(sort).Filter(filter).IncludeMembers(includeMembers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsGetUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGroupsGetUserGroups`: GetUserGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserGroupsApi.UserGroupsGetUserGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **sort** | **[]string** |  | 
 **filter** | **string** |  | 
 **includeMembers** | **bool** | Whether or not to return members for each User Group as part of the response. If not sent, the value will be False. | 

### Return type

[**GetUserGroupsResponse**](GetUserGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGroupsRemoveUserFromGroup

> UserGroupsRemoveUserFromGroup(ctx, groupId, memberName).Execute()





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
    groupId := int64(789) // int64 | The unique ID of the group.
    memberName := "memberName_example" // string | The name of the user/group to remove from the group members

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserGroupsApi.UserGroupsRemoveUserFromGroup(context.Background(), groupId, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserGroupsApi.UserGroupsRemoveUserFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The unique ID of the group. | 
**memberName** | **string** | The name of the user/group to remove from the group members | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGroupsRemoveUserFromGroupRequest struct via the builder pattern


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

