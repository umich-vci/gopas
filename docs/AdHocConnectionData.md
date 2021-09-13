# AdHocConnectionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | **string** |  | 
**Address** | **string** |  | 
**UserName** | **string** |  | 
**PlatformId** | **string** |  | 
**PSMConnectPrerequisites** | [**AccountPSMConnectPrerequisites**](AccountPSMConnectPrerequisites.md) |  | 
**ExtraFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAdHocConnectionData

`func NewAdHocConnectionData(secret string, address string, userName string, platformId string, pSMConnectPrerequisites AccountPSMConnectPrerequisites, ) *AdHocConnectionData`

NewAdHocConnectionData instantiates a new AdHocConnectionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdHocConnectionDataWithDefaults

`func NewAdHocConnectionDataWithDefaults() *AdHocConnectionData`

NewAdHocConnectionDataWithDefaults instantiates a new AdHocConnectionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *AdHocConnectionData) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AdHocConnectionData) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AdHocConnectionData) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetAddress

`func (o *AdHocConnectionData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AdHocConnectionData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AdHocConnectionData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetUserName

`func (o *AdHocConnectionData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AdHocConnectionData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AdHocConnectionData) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPlatformId

`func (o *AdHocConnectionData) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *AdHocConnectionData) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *AdHocConnectionData) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetPSMConnectPrerequisites

`func (o *AdHocConnectionData) GetPSMConnectPrerequisites() AccountPSMConnectPrerequisites`

GetPSMConnectPrerequisites returns the PSMConnectPrerequisites field if non-nil, zero value otherwise.

### GetPSMConnectPrerequisitesOk

`func (o *AdHocConnectionData) GetPSMConnectPrerequisitesOk() (*AccountPSMConnectPrerequisites, bool)`

GetPSMConnectPrerequisitesOk returns a tuple with the PSMConnectPrerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMConnectPrerequisites

`func (o *AdHocConnectionData) SetPSMConnectPrerequisites(v AccountPSMConnectPrerequisites)`

SetPSMConnectPrerequisites sets PSMConnectPrerequisites field to given value.


### GetExtraFields

`func (o *AdHocConnectionData) GetExtraFields() map[string]string`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *AdHocConnectionData) GetExtraFieldsOk() (*map[string]string, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *AdHocConnectionData) SetExtraFields(v map[string]string)`

SetExtraFields sets ExtraFields field to given value.

### HasExtraFields

`func (o *AdHocConnectionData) HasExtraFields() bool`

HasExtraFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


