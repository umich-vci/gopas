# DiscoveredDependencyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the on-boarded or pending dependency. | [optional] 
**Status** | Pointer to **string** | The status of the dependency. The possible values are:  addedDependency - the dependency was on-boarded to the Vault.  addedAsPending - the dependency was added to the Pending Accounts list.  alreadyExists - the dependency is already exists or already pending.  failed - there was a failure during dependency creation.  skip - the dependency wasn&#39;t added since there was a failure in previous dependency creation. | [optional] 
**Name** | Pointer to **string** | The dependency name. | [optional] 
**Address** | Pointer to **string** | The dependency address. | [optional] 
**Type** | Pointer to **string** | The dependency type. | [optional] 
**FailureReason** | Pointer to **string** | The failure message, in case the dependency failed to create. | [optional] 

## Methods

### NewDiscoveredDependencyResponse

`func NewDiscoveredDependencyResponse() *DiscoveredDependencyResponse`

NewDiscoveredDependencyResponse instantiates a new DiscoveredDependencyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredDependencyResponseWithDefaults

`func NewDiscoveredDependencyResponseWithDefaults() *DiscoveredDependencyResponse`

NewDiscoveredDependencyResponseWithDefaults instantiates a new DiscoveredDependencyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveredDependencyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveredDependencyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveredDependencyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscoveredDependencyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DiscoveredDependencyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveredDependencyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveredDependencyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveredDependencyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *DiscoveredDependencyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveredDependencyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveredDependencyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveredDependencyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveredDependencyResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveredDependencyResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveredDependencyResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveredDependencyResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetType

`func (o *DiscoveredDependencyResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveredDependencyResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveredDependencyResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveredDependencyResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFailureReason

`func (o *DiscoveredDependencyResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DiscoveredDependencyResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DiscoveredDependencyResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DiscoveredDependencyResponse) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


