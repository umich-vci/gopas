# \AutomaticOnboardingRulesApi

All URIs are relative to *https://pas.example.com/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomaticOnboardingRulesDeleteAutomaticOnboardingRule**](AutomaticOnboardingRulesApi.md#AutomaticOnboardingRulesDeleteAutomaticOnboardingRule) | **Delete** /api/AutomaticOnboardingRules/{id} | 
[**AutomaticOnboardingRulesGetAutomaticOnboardingRules**](AutomaticOnboardingRulesApi.md#AutomaticOnboardingRulesGetAutomaticOnboardingRules) | **Get** /api/AutomaticOnboardingRules | 
[**AutomaticOnboardingRulesPostAutomaticOnboardingRule**](AutomaticOnboardingRulesApi.md#AutomaticOnboardingRulesPostAutomaticOnboardingRule) | **Post** /api/AutomaticOnboardingRules | 
[**AutomaticOnboardingRulesUpdateAutomaticOnboardingRule**](AutomaticOnboardingRulesApi.md#AutomaticOnboardingRulesUpdateAutomaticOnboardingRule) | **Put** /api/AutomaticOnboardingRules/{id} | 



## AutomaticOnboardingRulesDeleteAutomaticOnboardingRule

> AutomaticOnboardingRulesDeleteAutomaticOnboardingRule(ctx, id).Execute()





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
    id := int64(789) // int64 | The unique ID of the rule to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticOnboardingRulesApi.AutomaticOnboardingRulesDeleteAutomaticOnboardingRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesDeleteAutomaticOnboardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The unique ID of the rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest struct via the builder pattern


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


## AutomaticOnboardingRulesGetAutomaticOnboardingRules

> AutomaticOnboardingRuleData AutomaticOnboardingRulesGetAutomaticOnboardingRules(ctx).Names(names).Execute()





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
    names := []string{"Inner_example"} // []string | List of names (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticOnboardingRulesApi.AutomaticOnboardingRulesGetAutomaticOnboardingRules(context.Background()).Names(names).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesGetAutomaticOnboardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutomaticOnboardingRulesGetAutomaticOnboardingRules`: AutomaticOnboardingRuleData
    fmt.Fprintf(os.Stdout, "Response from `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesGetAutomaticOnboardingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | List of names | 

### Return type

[**AutomaticOnboardingRuleData**](AutomaticOnboardingRuleData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomaticOnboardingRulesPostAutomaticOnboardingRule

> AutomaticOnboardingRuleData AutomaticOnboardingRulesPostAutomaticOnboardingRule(ctx).Data(data).Execute()





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
    data := *openapiclient.NewAutomaticOnboardingRuleDataIn("TargetPlatformId_example", "TargetSafeName_example", int32(123)) // AutomaticOnboardingRuleDataIn | The onboarding rule to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticOnboardingRulesApi.AutomaticOnboardingRulesPostAutomaticOnboardingRule(context.Background()).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesPostAutomaticOnboardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutomaticOnboardingRulesPostAutomaticOnboardingRule`: AutomaticOnboardingRuleData
    fmt.Fprintf(os.Stdout, "Response from `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesPostAutomaticOnboardingRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**AutomaticOnboardingRuleDataIn**](AutomaticOnboardingRuleDataIn.md) | The onboarding rule to create. | 

### Return type

[**AutomaticOnboardingRuleData**](AutomaticOnboardingRuleData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomaticOnboardingRulesUpdateAutomaticOnboardingRule

> AutomaticOnboardingRuleData AutomaticOnboardingRulesUpdateAutomaticOnboardingRule(ctx, id).Data(data).Execute()





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
    id := int64(789) // int64 | The unique ID of the rule to update
    data := *openapiclient.NewAutomaticOnboardingRuleDataIn("TargetPlatformId_example", "TargetSafeName_example", int32(123)) // AutomaticOnboardingRuleDataIn | The onboarding rule to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomaticOnboardingRulesApi.AutomaticOnboardingRulesUpdateAutomaticOnboardingRule(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesUpdateAutomaticOnboardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutomaticOnboardingRulesUpdateAutomaticOnboardingRule`: AutomaticOnboardingRuleData
    fmt.Fprintf(os.Stdout, "Response from `AutomaticOnboardingRulesApi.AutomaticOnboardingRulesUpdateAutomaticOnboardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The unique ID of the rule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**AutomaticOnboardingRuleDataIn**](AutomaticOnboardingRuleDataIn.md) | The onboarding rule to create | 

### Return type

[**AutomaticOnboardingRuleData**](AutomaticOnboardingRuleData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

