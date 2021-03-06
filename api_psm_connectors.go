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
)

// Linger please
var (
	_ _context.Context
)

// PSMConnectorsApiService PSMConnectorsApi service
type PSMConnectorsApiService service

type ApiPSMConnectorsGetPSMConnectorsControllerRequest struct {
	ctx        _context.Context
	ApiService *PSMConnectorsApiService
}

func (r ApiPSMConnectorsGetPSMConnectorsControllerRequest) Execute() ([]PSMConnector, *_nethttp.Response, error) {
	return r.ApiService.PSMConnectorsGetPSMConnectorsControllerExecute(r)
}

/*
PSMConnectorsGetPSMConnectorsController Method for PSMConnectorsGetPSMConnectorsController

This method allows vault admins to get a list of all connectors of an entire environment

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPSMConnectorsGetPSMConnectorsControllerRequest
*/
func (a *PSMConnectorsApiService) PSMConnectorsGetPSMConnectorsController(ctx _context.Context) ApiPSMConnectorsGetPSMConnectorsControllerRequest {
	return ApiPSMConnectorsGetPSMConnectorsControllerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []PSMConnector
func (a *PSMConnectorsApiService) PSMConnectorsGetPSMConnectorsControllerExecute(r ApiPSMConnectorsGetPSMConnectorsControllerRequest) ([]PSMConnector, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PSMConnector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PSMConnectorsApiService.PSMConnectorsGetPSMConnectorsController")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/PSM/Connectors"

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
