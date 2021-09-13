# AllowedReferrerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferrerURL** | **string** | The URL from where users will be able to access PVWA. You can specify either of the following:  A portal URL that allows access from any page or sub-directory under the company’s portal URL.  For example, https://CompanyA/portal/.  The exact name of a URL that only allows access from a specific page.For example, https://CompanyB/management/dashboard. | 
**RegularExpression** | Pointer to **bool** | Defines whether or not the URL is a regular expression. | [optional] 

## Methods

### NewAllowedReferrerDto

`func NewAllowedReferrerDto(referrerURL string, ) *AllowedReferrerDto`

NewAllowedReferrerDto instantiates a new AllowedReferrerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedReferrerDtoWithDefaults

`func NewAllowedReferrerDtoWithDefaults() *AllowedReferrerDto`

NewAllowedReferrerDtoWithDefaults instantiates a new AllowedReferrerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferrerURL

`func (o *AllowedReferrerDto) GetReferrerURL() string`

GetReferrerURL returns the ReferrerURL field if non-nil, zero value otherwise.

### GetReferrerURLOk

`func (o *AllowedReferrerDto) GetReferrerURLOk() (*string, bool)`

GetReferrerURLOk returns a tuple with the ReferrerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerURL

`func (o *AllowedReferrerDto) SetReferrerURL(v string)`

SetReferrerURL sets ReferrerURL field to given value.


### GetRegularExpression

`func (o *AllowedReferrerDto) GetRegularExpression() bool`

GetRegularExpression returns the RegularExpression field if non-nil, zero value otherwise.

### GetRegularExpressionOk

`func (o *AllowedReferrerDto) GetRegularExpressionOk() (*bool, bool)`

GetRegularExpressionOk returns a tuple with the RegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularExpression

`func (o *AllowedReferrerDto) SetRegularExpression(v bool)`

SetRegularExpression sets RegularExpression field to given value.

### HasRegularExpression

`func (o *AllowedReferrerDto) HasRegularExpression() bool`

HasRegularExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


