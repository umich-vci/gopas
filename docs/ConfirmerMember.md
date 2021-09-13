# ConfirmerMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserID** | Pointer to **int64** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**AdditionalDetails** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfirmerMember

`func NewConfirmerMember() *ConfirmerMember`

NewConfirmerMember instantiates a new ConfirmerMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmerMemberWithDefaults

`func NewConfirmerMemberWithDefaults() *ConfirmerMember`

NewConfirmerMemberWithDefaults instantiates a new ConfirmerMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserID

`func (o *ConfirmerMember) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ConfirmerMember) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ConfirmerMember) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *ConfirmerMember) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *ConfirmerMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ConfirmerMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ConfirmerMember) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ConfirmerMember) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *ConfirmerMember) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ConfirmerMember) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ConfirmerMember) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ConfirmerMember) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


