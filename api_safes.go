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
	"reflect"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// SafesApiService SafesApi service
type SafesApiService service

type ApiSafesAddSafeRequest struct {
	ctx        _context.Context
	ApiService *SafesApiService
	safe       *AddSafeData
}

func (r ApiSafesAddSafeRequest) Safe(safe AddSafeData) ApiSafesAddSafeRequest {
	r.safe = &safe
	return r
}

func (r ApiSafesAddSafeRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.SafesAddSafeExecute(r)
}

/*
SafesAddSafe Method for SafesAddSafe

This method adds a new Safe to the Vault.
The user who runs this web service requires Add Safes permission in the Vault.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSafesAddSafeRequest
*/
func (a *SafesApiService) SafesAddSafe(ctx _context.Context) ApiSafesAddSafeRequest {
	return ApiSafesAddSafeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SafesApiService) SafesAddSafeExecute(r ApiSafesAddSafeRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesAddSafe")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.safe == nil {
		return localVarReturnValue, nil, reportError("safe is required and must be specified")
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
	localVarPostBody = r.safe
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

type ApiSafesAddSafeOwnerRequest struct {
	ctx        _context.Context
	ApiService *SafesApiService
	safeUrlId  string
	member     *SafeMemberItem
}

// An existing user to add as a Safe member.
func (r ApiSafesAddSafeOwnerRequest) Member(member SafeMemberItem) ApiSafesAddSafeOwnerRequest {
	r.member = &member
	return r
}

func (r ApiSafesAddSafeOwnerRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.SafesAddSafeOwnerExecute(r)
}

/*
SafesAddSafeOwner Method for SafesAddSafeOwner

This method adds an existing user as a Safe member.
The user who run this web service requires Manage and View Members permissions in the Safe.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param safeUrlId The unique ID of the Safe.
 @return ApiSafesAddSafeOwnerRequest
*/
func (a *SafesApiService) SafesAddSafeOwner(ctx _context.Context, safeUrlId string) ApiSafesAddSafeOwnerRequest {
	return ApiSafesAddSafeOwnerRequest{
		ApiService: a,
		ctx:        ctx,
		safeUrlId:  safeUrlId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *SafesApiService) SafesAddSafeOwnerExecute(r ApiSafesAddSafeOwnerRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesAddSafeOwner")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes/{safeUrlId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"safeUrlId"+"}", _neturl.PathEscape(parameterToString(r.safeUrlId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.member == nil {
		return localVarReturnValue, nil, reportError("member is required and must be specified")
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
	localVarPostBody = r.member
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

type ApiSafesDeleteSafeDetailsRequest struct {
	ctx        _context.Context
	ApiService *SafesApiService
	safeUrlId  string
}

func (r ApiSafesDeleteSafeDetailsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SafesDeleteSafeDetailsExecute(r)
}

/*
SafesDeleteSafeDetails Method for SafesDeleteSafeDetails

This method deletes a safe from the Vault.
The user who runs this web service requires Manage Safe permission on the required Safe.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param safeUrlId The unique ID of the Safe.
 @return ApiSafesDeleteSafeDetailsRequest
*/
func (a *SafesApiService) SafesDeleteSafeDetails(ctx _context.Context, safeUrlId string) ApiSafesDeleteSafeDetailsRequest {
	return ApiSafesDeleteSafeDetailsRequest{
		ApiService: a,
		ctx:        ctx,
		safeUrlId:  safeUrlId,
	}
}

// Execute executes the request
func (a *SafesApiService) SafesDeleteSafeDetailsExecute(r ApiSafesDeleteSafeDetailsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesDeleteSafeDetails")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes/{safeUrlId}"
	localVarPath = strings.Replace(localVarPath, "{"+"safeUrlId"+"}", _neturl.PathEscape(parameterToString(r.safeUrlId, "")), -1)

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

type ApiSafesGetGroupsRequest struct {
	ctx        _context.Context
	ApiService *SafesApiService
	safeName   string
}

func (r ApiSafesGetGroupsRequest) Execute() (AccountGroup, *_nethttp.Response, error) {
	return r.ApiService.SafesGetGroupsExecute(r)
}

/*
SafesGetGroups Method for SafesGetGroups

This method returns all the existing account groups in a specific Safe. The user performing this task must have the following permissions in the Safe:
List Accounts.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param safeName The name of the Safe where the account groups are.
 @return ApiSafesGetGroupsRequest
*/
func (a *SafesApiService) SafesGetGroups(ctx _context.Context, safeName string) ApiSafesGetGroupsRequest {
	return ApiSafesGetGroupsRequest{
		ApiService: a,
		ctx:        ctx,
		safeName:   safeName,
	}
}

// Execute executes the request
//  @return AccountGroup
func (a *SafesApiService) SafesGetGroupsExecute(r ApiSafesGetGroupsRequest) (AccountGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesGetGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes/{safeName}/accountgroups"
	localVarPath = strings.Replace(localVarPath, "{"+"safeName"+"}", _neturl.PathEscape(parameterToString(r.safeName, "")), -1)

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

type ApiSafesGetSafeMembersRequest struct {
	ctx        _context.Context
	ApiService *SafesApiService
	safeUrlId  string
	limit      *int64
	offset     *int64
	useCache   *bool
	sort       *[]string
	search     *string
	filter     *string
}

func (r ApiSafesGetSafeMembersRequest) Limit(limit int64) ApiSafesGetSafeMembersRequest {
	r.limit = &limit
	return r
}
func (r ApiSafesGetSafeMembersRequest) Offset(offset int64) ApiSafesGetSafeMembersRequest {
	r.offset = &offset
	return r
}
func (r ApiSafesGetSafeMembersRequest) UseCache(useCache bool) ApiSafesGetSafeMembersRequest {
	r.useCache = &useCache
	return r
}
func (r ApiSafesGetSafeMembersRequest) Sort(sort []string) ApiSafesGetSafeMembersRequest {
	r.sort = &sort
	return r
}
func (r ApiSafesGetSafeMembersRequest) Search(search string) ApiSafesGetSafeMembersRequest {
	r.search = &search
	return r
}

// &lt;para&gt;Filtering according to REST standard. &lt;/para&gt;  &lt;para&gt;memberType - Return only members of single type (user or group)&lt;/para&gt;  &lt;para&gt;membershipExpired - Return only members that have or don&#39;t have an expired membership&lt;/para&gt;  &lt;para&gt;includePredefinedUsers - Include predefined users in the returned list.&lt;/para&gt;
func (r ApiSafesGetSafeMembersRequest) Filter(filter string) ApiSafesGetSafeMembersRequest {
	r.filter = &filter
	return r
}

func (r ApiSafesGetSafeMembersRequest) Execute() (SafeMemberResponse, *_nethttp.Response, error) {
	return r.ApiService.SafesGetSafeMembersExecute(r)
}

/*
SafesGetSafeMembers Method for SafesGetSafeMembers

This method returns the list of members of a Safe.
The user who run this web service requires View Safe Members permission on the Safe.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param safeUrlId The unique ID of the safe to return all its members
 @return ApiSafesGetSafeMembersRequest
*/
func (a *SafesApiService) SafesGetSafeMembers(ctx _context.Context, safeUrlId string) ApiSafesGetSafeMembersRequest {
	return ApiSafesGetSafeMembersRequest{
		ApiService: a,
		ctx:        ctx,
		safeUrlId:  safeUrlId,
	}
}

// Execute executes the request
//  @return SafeMemberResponse
func (a *SafesApiService) SafesGetSafeMembersExecute(r ApiSafesGetSafeMembersRequest) (SafeMemberResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SafeMemberResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesGetSafeMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes/{safeUrlId}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"safeUrlId"+"}", _neturl.PathEscape(parameterToString(r.safeUrlId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.useCache != nil {
		localVarQueryParams.Add("useCache", parameterToString(*r.useCache, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
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

type ApiSafesGetSafesRequest struct {
	ctx             _context.Context
	ApiService      *SafesApiService
	limit           *int64
	offset          *int64
	useCache        *bool
	sort            *[]string
	search          *string
	includeAccounts *bool
	extendedDetails *bool
}

func (r ApiSafesGetSafesRequest) Limit(limit int64) ApiSafesGetSafesRequest {
	r.limit = &limit
	return r
}
func (r ApiSafesGetSafesRequest) Offset(offset int64) ApiSafesGetSafesRequest {
	r.offset = &offset
	return r
}
func (r ApiSafesGetSafesRequest) UseCache(useCache bool) ApiSafesGetSafesRequest {
	r.useCache = &useCache
	return r
}
func (r ApiSafesGetSafesRequest) Sort(sort []string) ApiSafesGetSafesRequest {
	r.sort = &sort
	return r
}
func (r ApiSafesGetSafesRequest) Search(search string) ApiSafesGetSafesRequest {
	r.search = &search
	return r
}

// Whether or not to return accounts for each Safe as part of the response. If not sent, the value will be false.
func (r ApiSafesGetSafesRequest) IncludeAccounts(includeAccounts bool) ApiSafesGetSafesRequest {
	r.includeAccounts = &includeAccounts
	return r
}

// Whether or not to return all Safe data or only SafeName as part of the response. If not sent, the value will be true.
func (r ApiSafesGetSafesRequest) ExtendedDetails(extendedDetails bool) ApiSafesGetSafesRequest {
	r.extendedDetails = &extendedDetails
	return r
}

func (r ApiSafesGetSafesRequest) Execute() ([]SafeListItem, *_nethttp.Response, error) {
	return r.ApiService.SafesGetSafesExecute(r)
}

/*
SafesGetSafes Method for SafesGetSafes

This method returns a list of all Safes in the Vault that the user has permissions for.
The user who runs this web service must be a member of the Safes in the Vault that are returned in the list.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSafesGetSafesRequest
*/
func (a *SafesApiService) SafesGetSafes(ctx _context.Context) ApiSafesGetSafesRequest {
	return ApiSafesGetSafesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []SafeListItem
func (a *SafesApiService) SafesGetSafesExecute(r ApiSafesGetSafesRequest) ([]SafeListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SafeListItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SafesApiService.SafesGetSafes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Safes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.useCache != nil {
		localVarQueryParams.Add("useCache", parameterToString(*r.useCache, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.includeAccounts != nil {
		localVarQueryParams.Add("includeAccounts", parameterToString(*r.includeAccounts, ""))
	}
	if r.extendedDetails != nil {
		localVarQueryParams.Add("extendedDetails", parameterToString(*r.extendedDetails, ""))
	}
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
