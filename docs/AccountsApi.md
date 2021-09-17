# \AccountsApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsAdHocConnect**](AccountsApi.md#AccountsAdHocConnect) | **Post** /api/Accounts/AdHocConnect | 
[**AccountsAddAccount**](AccountsApi.md#AccountsAddAccount) | **Post** /api/Accounts | 
[**AccountsCPMChangeNow**](AccountsApi.md#AccountsCPMChangeNow) | **Post** /api/Accounts/{accountId}/Change | 
[**AccountsCPMSetNext**](AccountsApi.md#AccountsCPMSetNext) | **Post** /api/Accounts/{accountId}/SetNextPassword | 
[**AccountsChangeCredentialsInTheVault**](AccountsApi.md#AccountsChangeCredentialsInTheVault) | **Post** /api/Accounts/{accountId}/Password/Update | 
[**AccountsCheckIn**](AccountsApi.md#AccountsCheckIn) | **Post** /api/Accounts/{accountId}/CheckIn | 
[**AccountsDeleteAccount**](AccountsApi.md#AccountsDeleteAccount) | **Delete** /api/Accounts/{id} | 
[**AccountsDownload**](AccountsApi.md#AccountsDownload) | **Post** /api/Accounts/{accountId}/Secret/Retrieve | 
[**AccountsGeneratePassword**](AccountsApi.md#AccountsGeneratePassword) | **Post** /api/Accounts/{accountId}/Secret/Generate | 
[**AccountsGetAccount**](AccountsApi.md#AccountsGetAccount) | **Get** /api/Accounts/{id} | 
[**AccountsGetAccountLegacy**](AccountsApi.md#AccountsGetAccountLegacy) | **Get** /WebServices/PIMServices.svc/Accounts | 
[**AccountsGetAccounts**](AccountsApi.md#AccountsGetAccounts) | **Get** /api/Accounts | 
[**AccountsGetVersions**](AccountsApi.md#AccountsGetVersions) | **Get** /api/Accounts/{accountId}/Secret/Versions | 
[**AccountsGrantAccess**](AccountsApi.md#AccountsGrantAccess) | **Post** /api/Accounts/{accountId}/GrantAdministrativeAccess | 
[**AccountsLinkAccount**](AccountsApi.md#AccountsLinkAccount) | **Post** /api/Accounts/{id}/LinkAccount | 
[**AccountsPSMConnect**](AccountsApi.md#AccountsPSMConnect) | **Post** /api/Accounts/{accountId}/PSMConnect | 
[**AccountsReconcile**](AccountsApi.md#AccountsReconcile) | **Post** /api/Accounts/{accountId}/Reconcile | 
[**AccountsRetrieve**](AccountsApi.md#AccountsRetrieve) | **Post** /api/Accounts/{accountId}/Password/Retrieve | 
[**AccountsRevokeAccess**](AccountsApi.md#AccountsRevokeAccess) | **Post** /api/Accounts/{accountId}/RevokeAdministrativeAccess | 
[**AccountsUpdateAccount**](AccountsApi.md#AccountsUpdateAccount) | **Patch** /api/Accounts/{accountId} | 
[**AccountsVerify**](AccountsApi.md#AccountsVerify) | **Post** /api/Accounts/{accountId}/Verify | 



## AccountsAdHocConnect

> AccountsAdHocConnect(ctx).ConnectionData(connectionData).Execute()





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
    connectionData := *openapiclient.NewAdHocConnectionData("Secret_example", "Address_example", "UserName_example", "PlatformId_example", *openapiclient.NewAccountPSMConnectPrerequisites("ConnectionComponent_example")) // AdHocConnectionData | The information for the connection request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsAdHocConnect(context.Background()).ConnectionData(connectionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsAdHocConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAdHocConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionData** | [**AdHocConnectionData**](AdHocConnectionData.md) | The information for the connection request. | 

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


## AccountsAddAccount

> map[string]interface{} AccountsAddAccount(ctx).Account(account).Execute()





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
    account := *openapiclient.NewAccountModel("PlatformId_example", "SafeName_example") // AccountModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsAddAccount(context.Background()).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsAddAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsAddAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsAddAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsAddAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**AccountModel**](AccountModel.md) |  | 

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


## AccountsCPMChangeNow

> map[string]interface{} AccountsCPMChangeNow(ctx, accountId).ChangeProperties(changeProperties).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    changeProperties := *openapiclient.NewChangeCredentialsNowProperties() // ChangeCredentialsNowProperties | The change credentials properties.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsCPMChangeNow(context.Background(), accountId).ChangeProperties(changeProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsCPMChangeNow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsCPMChangeNow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsCPMChangeNow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsCPMChangeNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeProperties** | [**ChangeCredentialsNowProperties**](ChangeCredentialsNowProperties.md) | The change credentials properties. | 

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


## AccountsCPMSetNext

> map[string]interface{} AccountsCPMSetNext(ctx, accountId).ChangeProperties(changeProperties).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    changeProperties := *openapiclient.NewSetNextCredentialsProperties("NewCredentials_example") // SetNextCredentialsProperties | change properties data.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsCPMSetNext(context.Background(), accountId).ChangeProperties(changeProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsCPMSetNext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsCPMSetNext`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsCPMSetNext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsCPMSetNextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeProperties** | [**SetNextCredentialsProperties**](SetNextCredentialsProperties.md) | change properties data. | 

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


## AccountsChangeCredentialsInTheVault

> map[string]interface{} AccountsChangeCredentialsInTheVault(ctx, accountId).ChangeProperties(changeProperties).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    changeProperties := *openapiclient.NewChangeInVaultProperties("NewCredentials_example") // ChangeInVaultProperties | The change properties data.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsChangeCredentialsInTheVault(context.Background(), accountId).ChangeProperties(changeProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsChangeCredentialsInTheVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsChangeCredentialsInTheVault`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsChangeCredentialsInTheVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsChangeCredentialsInTheVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeProperties** | [**ChangeInVaultProperties**](ChangeInVaultProperties.md) | The change properties data. | 

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


## AccountsCheckIn

> map[string]interface{} AccountsCheckIn(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsCheckIn(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsCheckIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsCheckIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsCheckIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsCheckInRequest struct via the builder pattern


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


## AccountsDeleteAccount

> map[string]interface{} AccountsDeleteAccount(ctx, id).Execute()





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
    id := "id_example" // string | The account's unique ID, composed of the SafeID and internal AccountID of the account to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsDeleteAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsDeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsDeleteAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsDeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account&#39;s unique ID, composed of the SafeID and internal AccountID of the account to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsDeleteAccountRequest struct via the builder pattern


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


## AccountsDownload

> map[string]interface{} AccountsDownload(ctx, accountId).AccountContentPrerequisites(accountContentPrerequisites).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    accountContentPrerequisites := *openapiclient.NewAccountContentPrerequsites() // AccountContentPrerequsites | The account content prerequisites.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsDownload(context.Background(), accountId).AccountContentPrerequisites(accountContentPrerequisites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsDownload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountContentPrerequisites** | [**AccountContentPrerequsites**](AccountContentPrerequsites.md) | The account content prerequisites. | 

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


## AccountsGeneratePassword

> []GeneratedPassword AccountsGeneratePassword(ctx, accountId).AccDetails(accDetails).Execute()





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
    accountId := "accountId_example" // string | Unique ID of the account.
    accDetails := *openapiclient.NewGeneratePasswordAccountDetailsJson("AccountId_example") // GeneratePasswordAccountDetailsJson | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGeneratePassword(context.Background(), accountId).AccDetails(accDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGeneratePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGeneratePassword`: []GeneratedPassword
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGeneratePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGeneratePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accDetails** | [**GeneratePasswordAccountDetailsJson**](GeneratePasswordAccountDetailsJson.md) |  | 

### Return type

[**[]GeneratedPassword**](GeneratedPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetAccount

> AccountModel AccountsGetAccount(ctx, id).Execute()





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
    id := "id_example" // string | The account's unique ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGetAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGetAccount`: AccountModel
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account&#39;s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountModel**](AccountModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetAccountLegacy

> AccountModelv1 AccountsGetAccountLegacy(ctx).Keywords(keywords).Safe(safe).Execute()





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
    keywords := "keywords_example" // string | Keywords to identify the account (optional)
    safe := "safe_example" // string | Safe to search in for the account (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGetAccountLegacy(context.Background()).Keywords(keywords).Safe(safe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGetAccountLegacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGetAccountLegacy`: AccountModelv1
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGetAccountLegacy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetAccountLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywords** | **string** | Keywords to identify the account | 
 **safe** | **string** | Safe to search in for the account | 

### Return type

[**AccountModelv1**](AccountModelv1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetAccounts

> GetAccountsResponse AccountsGetAccounts(ctx).Search(search).Sort(sort).Offset(offset).Limit(limit).Filter(filter).SearchType(searchType).Execute()





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
    search := "search_example" // string | List of keywords separated with space to search in accounts. (optional)
    sort := []string{"Inner_example"} // []string | Property or properties to sort returned accounts, followed by asc (default) or desc to control sort direction. Multiple sorts are comma-separated. Maximum number of properties is 3. (optional)
    offset := int32(56) // int32 | Offset of the first returned account into the collection of results. (optional)
    limit := int32(56) // int32 | Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000. (optional)
    filter := "filter_example" // string | Search for accounts filtered by specific Safe. Filter=safename eq [safe name] (optional)
    searchType := "searchType_example" // string | Type of search to perform - if keywords are contained in the relevant account properties values or in the start of the properties values (the later enhances performance) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGetAccounts(context.Background()).Search(search).Sort(sort).Offset(offset).Limit(limit).Filter(filter).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGetAccounts`: GetAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | List of keywords separated with space to search in accounts. | 
 **sort** | **[]string** | Property or properties to sort returned accounts, followed by asc (default) or desc to control sort direction. Multiple sorts are comma-separated. Maximum number of properties is 3. | 
 **offset** | **int32** | Offset of the first returned account into the collection of results. | 
 **limit** | **int32** | Maximum number of returned accounts. If not specified, the default value is 50. The maximum number that can be specified is 1000. | 
 **filter** | **string** | Search for accounts filtered by specific Safe. Filter&#x3D;safename eq [safe name] | 
 **searchType** | **string** | Type of search to perform - if keywords are contained in the relevant account properties values or in the start of the properties values (the later enhances performance) | 

### Return type

[**GetAccountsResponse**](GetAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGetVersions

> []PVFileVersion AccountsGetVersions(ctx, accountId).ShowTemporary(showTemporary).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    showTemporary := true // bool | Whether to return both real and temporary password version or only real. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGetVersions(context.Background(), accountId).ShowTemporary(showTemporary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGetVersions`: []PVFileVersion
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showTemporary** | **bool** | Whether to return both real and temporary password version or only real. | 

### Return type

[**[]PVFileVersion**](PVFileVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGrantAccess

> map[string]interface{} AccountsGrantAccess(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | The local account that will be used to add the logged on user to the Administrators group on the machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsGrantAccess(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsGrantAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGrantAccess`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsGrantAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The local account that will be used to add the logged on user to the Administrators group on the machine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGrantAccessRequest struct via the builder pattern


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


## AccountsLinkAccount

> AccountsLinkAccount(ctx, id).LinkAccount(linkAccount).Execute()





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
    id := "id_example" // string | The unique ID of the account.
    linkAccount := *openapiclient.NewLinkAccountData("Name_example", "Safe_example", "Folder_example", int32(123)) // LinkAccountData | The details of the linked account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsLinkAccount(context.Background(), id).LinkAccount(linkAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsLinkAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsLinkAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkAccount** | [**LinkAccountData**](LinkAccountData.md) | The details of the linked account. | 

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


## AccountsPSMConnect

> AccountsPSMConnect(ctx, accountId).Request(request).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account to retrieve and use to connect to the target system through PSM.
    request := *openapiclient.NewAccountPSMConnectPrerequisites("ConnectionComponent_example") // AccountPSMConnectPrerequisites | The information of the request to connect.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsPSMConnect(context.Background(), accountId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsPSMConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to retrieve and use to connect to the target system through PSM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsPSMConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**AccountPSMConnectPrerequisites**](AccountPSMConnectPrerequisites.md) | The information of the request to connect. | 

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


## AccountsReconcile

> map[string]interface{} AccountsReconcile(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsReconcile(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsReconcile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsReconcile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsReconcile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsReconcileRequest struct via the builder pattern


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


## AccountsRetrieve

> map[string]interface{} AccountsRetrieve(ctx, accountId).AccountContentPrerequisites(accountContentPrerequisites).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.
    accountContentPrerequisites := *openapiclient.NewAccountContentPrerequsites() // AccountContentPrerequsites | The account content prerequisites.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsRetrieve(context.Background(), accountId).AccountContentPrerequisites(accountContentPrerequisites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsRetrieve`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountContentPrerequisites** | [**AccountContentPrerequsites**](AccountContentPrerequsites.md) | The account content prerequisites. | 

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


## AccountsRevokeAccess

> AccountsRevokeAccess(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | The local account for the target machine on which the logged on user will be removed from the Administrators group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsRevokeAccess(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsRevokeAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The local account for the target machine on which the logged on user will be removed from the Administrators group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsRevokeAccessRequest struct via the builder pattern


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


## AccountsUpdateAccount

> AccountModel AccountsUpdateAccount(ctx, accountId).AccountPatch(accountPatch).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account to update. This is retrieved by the Get Account Service.
    accountPatch := []openapiclient.OperationAccountModel{*openapiclient.NewOperationAccountModel()} // []OperationAccountModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsUpdateAccount(context.Background(), accountId).AccountPatch(accountPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsUpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsUpdateAccount`: AccountModel
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsUpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to update. This is retrieved by the Get Account Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountPatch** | [**[]OperationAccountModel**](OperationAccountModel.md) |  | 

### Return type

[**AccountModel**](AccountModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsVerify

> map[string]interface{} AccountsVerify(ctx, accountId).Execute()





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
    accountId := "accountId_example" // string | The unique ID of the account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.AccountsVerify(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.AccountsVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsVerify`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.AccountsVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsVerifyRequest struct via the builder pattern


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

