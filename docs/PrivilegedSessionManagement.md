# PrivilegedSessionManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PSMConnectors** | Pointer to [**[]A4b**](A4b.md) | List of PSM Connectors linked to the platform | [optional] 
**PSMServerId** | **string** | PSM server ID linked to the platform | 
**PSMServerName** | Pointer to **string** | Name of the PSM server linked to the platform | [optional] 

## Methods

### NewPrivilegedSessionManagement

`func NewPrivilegedSessionManagement(pSMServerId string, ) *PrivilegedSessionManagement`

NewPrivilegedSessionManagement instantiates a new PrivilegedSessionManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedSessionManagementWithDefaults

`func NewPrivilegedSessionManagementWithDefaults() *PrivilegedSessionManagement`

NewPrivilegedSessionManagementWithDefaults instantiates a new PrivilegedSessionManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPSMConnectors

`func (o *PrivilegedSessionManagement) GetPSMConnectors() []A4b`

GetPSMConnectors returns the PSMConnectors field if non-nil, zero value otherwise.

### GetPSMConnectorsOk

`func (o *PrivilegedSessionManagement) GetPSMConnectorsOk() (*[]A4b, bool)`

GetPSMConnectorsOk returns a tuple with the PSMConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMConnectors

`func (o *PrivilegedSessionManagement) SetPSMConnectors(v []A4b)`

SetPSMConnectors sets PSMConnectors field to given value.

### HasPSMConnectors

`func (o *PrivilegedSessionManagement) HasPSMConnectors() bool`

HasPSMConnectors returns a boolean if a field has been set.

### GetPSMServerId

`func (o *PrivilegedSessionManagement) GetPSMServerId() string`

GetPSMServerId returns the PSMServerId field if non-nil, zero value otherwise.

### GetPSMServerIdOk

`func (o *PrivilegedSessionManagement) GetPSMServerIdOk() (*string, bool)`

GetPSMServerIdOk returns a tuple with the PSMServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMServerId

`func (o *PrivilegedSessionManagement) SetPSMServerId(v string)`

SetPSMServerId sets PSMServerId field to given value.


### GetPSMServerName

`func (o *PrivilegedSessionManagement) GetPSMServerName() string`

GetPSMServerName returns the PSMServerName field if non-nil, zero value otherwise.

### GetPSMServerNameOk

`func (o *PrivilegedSessionManagement) GetPSMServerNameOk() (*string, bool)`

GetPSMServerNameOk returns a tuple with the PSMServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMServerName

`func (o *PrivilegedSessionManagement) SetPSMServerName(v string)`

SetPSMServerName sets PSMServerName field to given value.

### HasPSMServerName

`func (o *PrivilegedSessionManagement) HasPSMServerName() bool`

HasPSMServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


