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

// ComponentsMonitoringDetailsApiService ComponentsMonitoringDetailsApi service
type ComponentsMonitoringDetailsApiService service

type ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest struct {
	ctx         _context.Context
	ApiService  *ComponentsMonitoringDetailsApiService
	componentId string
}

func (r ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest) Execute() (ComponentsMonitoringDetailsData, *_nethttp.Response, error) {
	return r.ApiService.ComponentsMonitoringDetailsGetComponentsMonitoringDetailsExecute(r)
}

/*
ComponentsMonitoringDetailsGetComponentsMonitoringDetails Method for ComponentsMonitoringDetailsGetComponentsMonitoringDetails

This method returns details about specific components and all their installed instances, and system health information for each one.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param componentId The type of component for which data will be retrieved.
 @return ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest
*/
func (a *ComponentsMonitoringDetailsApiService) ComponentsMonitoringDetailsGetComponentsMonitoringDetails(ctx _context.Context, componentId string) ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest {
	return ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest{
		ApiService:  a,
		ctx:         ctx,
		componentId: componentId,
	}
}

// Execute executes the request
//  @return ComponentsMonitoringDetailsData
func (a *ComponentsMonitoringDetailsApiService) ComponentsMonitoringDetailsGetComponentsMonitoringDetailsExecute(r ApiComponentsMonitoringDetailsGetComponentsMonitoringDetailsRequest) (ComponentsMonitoringDetailsData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ComponentsMonitoringDetailsData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsMonitoringDetailsApiService.ComponentsMonitoringDetailsGetComponentsMonitoringDetails")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/ComponentsMonitoringDetails/{componentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"componentId"+"}", _neturl.PathEscape(parameterToString(r.componentId, "")), -1)

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