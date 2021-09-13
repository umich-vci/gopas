/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopas

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OIDCConfigurationApiService OIDCConfigurationApi service
type OIDCConfigurationApiService service

type ApiOIDCConfigurationAddProviderRequest struct {
	ctx               _context.Context
	ApiService        *OIDCConfigurationApiService
	configurationData *OPConfigurationData
}

func (r ApiOIDCConfigurationAddProviderRequest) ConfigurationData(configurationData OPConfigurationData) ApiOIDCConfigurationAddProviderRequest {
	r.configurationData = &configurationData
	return r
}

func (r ApiOIDCConfigurationAddProviderRequest) Execute() (OPConfigurationData, *_nethttp.Response, error) {
	return r.ApiService.OIDCConfigurationAddProviderExecute(r)
}

/*
OIDCConfigurationAddProvider Method for OIDCConfigurationAddProvider

This method creates an OIDC Identity Provider in the Vault.

Any user who is a member of the Vault Admins group can run this web service.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOIDCConfigurationAddProviderRequest
*/
func (a *OIDCConfigurationApiService) OIDCConfigurationAddProvider(ctx _context.Context) ApiOIDCConfigurationAddProviderRequest {
	return ApiOIDCConfigurationAddProviderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return OPConfigurationData
func (a *OIDCConfigurationApiService) OIDCConfigurationAddProviderExecute(r ApiOIDCConfigurationAddProviderRequest) (OPConfigurationData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OPConfigurationData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OIDCConfigurationApiService.OIDCConfigurationAddProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/OIDC/Providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.configurationData == nil {
		return localVarReturnValue, nil, reportError("configurationData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.configurationData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOIDCConfigurationDeleteProviderRequest struct {
	ctx        _context.Context
	ApiService *OIDCConfigurationApiService
	id         string
}

func (r ApiOIDCConfigurationDeleteProviderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OIDCConfigurationDeleteProviderExecute(r)
}

/*
OIDCConfigurationDeleteProvider Method for OIDCConfigurationDeleteProvider

This method deletes a specific OIDC Identity Provider.

Any user who is a member of the Vault Admins group can run this web service.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.
 @return ApiOIDCConfigurationDeleteProviderRequest
*/
func (a *OIDCConfigurationApiService) OIDCConfigurationDeleteProvider(ctx _context.Context, id string) ApiOIDCConfigurationDeleteProviderRequest {
	return ApiOIDCConfigurationDeleteProviderRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OIDCConfigurationApiService) OIDCConfigurationDeleteProviderExecute(r ApiOIDCConfigurationDeleteProviderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OIDCConfigurationApiService.OIDCConfigurationDeleteProvider")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/OIDC/Providers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiOIDCConfigurationGetProviderRequest struct {
	ctx        _context.Context
	ApiService *OIDCConfigurationApiService
	id         string
}

func (r ApiOIDCConfigurationGetProviderRequest) Execute() (OPConfigurationData, *_nethttp.Response, error) {
	return r.ApiService.OIDCConfigurationGetProviderExecute(r)
}

/*
OIDCConfigurationGetProvider Method for OIDCConfigurationGetProvider

This method returns a specific OIDC Identity Provider.

Any user who is a member of the Vault Admins group can run this web service.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.
 @return ApiOIDCConfigurationGetProviderRequest
*/
func (a *OIDCConfigurationApiService) OIDCConfigurationGetProvider(ctx _context.Context, id string) ApiOIDCConfigurationGetProviderRequest {
	return ApiOIDCConfigurationGetProviderRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return OPConfigurationData
func (a *OIDCConfigurationApiService) OIDCConfigurationGetProviderExecute(r ApiOIDCConfigurationGetProviderRequest) (OPConfigurationData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OPConfigurationData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OIDCConfigurationApiService.OIDCConfigurationGetProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/OIDC/Providers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOIDCConfigurationGetProvidersRequest struct {
	ctx        _context.Context
	ApiService *OIDCConfigurationApiService
}

func (r ApiOIDCConfigurationGetProvidersRequest) Execute() (OPConfigurations, *_nethttp.Response, error) {
	return r.ApiService.OIDCConfigurationGetProvidersExecute(r)
}

/*
OIDCConfigurationGetProviders Method for OIDCConfigurationGetProviders

This method return a list of all OIDC Identity Providers.

Any user who is a member of the Vault Admins group can run this web service.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOIDCConfigurationGetProvidersRequest
*/
func (a *OIDCConfigurationApiService) OIDCConfigurationGetProviders(ctx _context.Context) ApiOIDCConfigurationGetProvidersRequest {
	return ApiOIDCConfigurationGetProvidersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return OPConfigurations
func (a *OIDCConfigurationApiService) OIDCConfigurationGetProvidersExecute(r ApiOIDCConfigurationGetProvidersRequest) (OPConfigurations, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OPConfigurations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OIDCConfigurationApiService.OIDCConfigurationGetProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/OIDC/Providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOIDCConfigurationUpdateProviderRequest struct {
	ctx               _context.Context
	ApiService        *OIDCConfigurationApiService
	id                string
	configurationData *OPConfigurationBase
}

func (r ApiOIDCConfigurationUpdateProviderRequest) ConfigurationData(configurationData OPConfigurationBase) ApiOIDCConfigurationUpdateProviderRequest {
	r.configurationData = &configurationData
	return r
}

func (r ApiOIDCConfigurationUpdateProviderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OIDCConfigurationUpdateProviderExecute(r)
}

/*
OIDCConfigurationUpdateProvider Method for OIDCConfigurationUpdateProvider

This method updates an existing OIDC Identity Provider.

Any user who is a member of the Vault Admins group can run this web service.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier of the provider.              This ID is used to identify the OIDC Identity Provider in PVWA.
 @return ApiOIDCConfigurationUpdateProviderRequest
*/
func (a *OIDCConfigurationApiService) OIDCConfigurationUpdateProvider(ctx _context.Context, id string) ApiOIDCConfigurationUpdateProviderRequest {
	return ApiOIDCConfigurationUpdateProviderRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OIDCConfigurationApiService) OIDCConfigurationUpdateProviderExecute(r ApiOIDCConfigurationUpdateProviderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OIDCConfigurationApiService.OIDCConfigurationUpdateProvider")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/OIDC/Providers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.configurationData == nil {
		return nil, reportError("configurationData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.configurationData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
