# SafeMembersFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | &lt;para&gt;Filtering according to REST standard. &lt;/para&gt;  &lt;para&gt;memberType - Return only members of single type (user or group)&lt;/para&gt;  &lt;para&gt;membershipExpired - Return only members that have or don&#39;t have an expired membership&lt;/para&gt;  &lt;para&gt;includePredefinedUsers - Include predefined users in the returned list.&lt;/para&gt; | [optional] 

## Methods

### NewSafeMembersFilter

`func NewSafeMembersFilter() *SafeMembersFilter`

NewSafeMembersFilter instantiates a new SafeMembersFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeMembersFilterWithDefaults

`func NewSafeMembersFilterWithDefaults() *SafeMembersFilter`

NewSafeMembersFilterWithDefaults instantiates a new SafeMembersFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SafeMembersFilter) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SafeMembersFilter) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SafeMembersFilter) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SafeMembersFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


