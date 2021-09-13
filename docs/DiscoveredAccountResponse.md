# DiscoveredAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the on-boarded or pending account. | [optional] 
**Status** | Pointer to **string** | The status of the account. The possible values are:  addedAccount - the account was on-boarded to the Vault.  addedAsPending - the account was added to the Pending Accounts list.  updatedAccount - the account is already exists and new dependencies were added to it.  updatedPending - the pending account already exists and it was updated if needed.  alreadyExists - the account is already exists and no new dependencies were added. | [optional] 
**UserName** | Pointer to **string** | The name of the account user. | [optional] 
**Address** | Pointer to **string** | The name or address of the machine where the account is located. | [optional] 
**Safe** | Pointer to **string** | The name of the safe of the account. | [optional] 
**Dependencies** | Pointer to [**[]DiscoveredDependencyResponse**](DiscoveredDependencyResponse.md) | The dependencies of the account. | [optional] 

## Methods

### NewDiscoveredAccountResponse

`func NewDiscoveredAccountResponse() *DiscoveredAccountResponse`

NewDiscoveredAccountResponse instantiates a new DiscoveredAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredAccountResponseWithDefaults

`func NewDiscoveredAccountResponseWithDefaults() *DiscoveredAccountResponse`

NewDiscoveredAccountResponseWithDefaults instantiates a new DiscoveredAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveredAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveredAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DiscoveredAccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveredAccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveredAccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveredAccountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserName

`func (o *DiscoveredAccountResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DiscoveredAccountResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DiscoveredAccountResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DiscoveredAccountResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveredAccountResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveredAccountResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveredAccountResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveredAccountResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSafe

`func (o *DiscoveredAccountResponse) GetSafe() string`

GetSafe returns the Safe field if non-nil, zero value otherwise.

### GetSafeOk

`func (o *DiscoveredAccountResponse) GetSafeOk() (*string, bool)`

GetSafeOk returns a tuple with the Safe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafe

`func (o *DiscoveredAccountResponse) SetSafe(v string)`

SetSafe sets Safe field to given value.

### HasSafe

`func (o *DiscoveredAccountResponse) HasSafe() bool`

HasSafe returns a boolean if a field has been set.

### GetDependencies

`func (o *DiscoveredAccountResponse) GetDependencies() []DiscoveredDependencyResponse`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *DiscoveredAccountResponse) GetDependenciesOk() (*[]DiscoveredDependencyResponse, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *DiscoveredAccountResponse) SetDependencies(v []DiscoveredDependencyResponse)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *DiscoveredAccountResponse) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


