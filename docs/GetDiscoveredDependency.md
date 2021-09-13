# GetDiscoveredDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The dependency name | [optional] 
**Address** | Pointer to **string** | The dependency address | [optional] 
**Type** | Pointer to **string** | The dependency type.  One of the following: COM+ Application, IIS Anonymous Authentication, IIS Application Pool, Windows Scheduled Task, Windows Service | [optional] 
**TaskFolder** | Pointer to **string** | The dependency task folder - For Windows Scheduled Task | [optional] 

## Methods

### NewGetDiscoveredDependency

`func NewGetDiscoveredDependency() *GetDiscoveredDependency`

NewGetDiscoveredDependency instantiates a new GetDiscoveredDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveredDependencyWithDefaults

`func NewGetDiscoveredDependencyWithDefaults() *GetDiscoveredDependency`

NewGetDiscoveredDependencyWithDefaults instantiates a new GetDiscoveredDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDiscoveredDependency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveredDependency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveredDependency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveredDependency) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *GetDiscoveredDependency) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDiscoveredDependency) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDiscoveredDependency) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDiscoveredDependency) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetType

`func (o *GetDiscoveredDependency) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDiscoveredDependency) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDiscoveredDependency) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDiscoveredDependency) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTaskFolder

`func (o *GetDiscoveredDependency) GetTaskFolder() string`

GetTaskFolder returns the TaskFolder field if non-nil, zero value otherwise.

### GetTaskFolderOk

`func (o *GetDiscoveredDependency) GetTaskFolderOk() (*string, bool)`

GetTaskFolderOk returns a tuple with the TaskFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFolder

`func (o *GetDiscoveredDependency) SetTaskFolder(v string)`

SetTaskFolder sets TaskFolder field to given value.

### HasTaskFolder

`func (o *GetDiscoveredDependency) HasTaskFolder() bool`

HasTaskFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


