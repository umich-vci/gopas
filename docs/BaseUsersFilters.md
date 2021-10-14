# BaseUsersFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserType** | Pointer to **string** | The type of the user. | [optional] 
**ComponentUser** | Pointer to **bool** | Whether the user is a known component or not. | [optional] 

## Methods

### NewBaseUsersFilters

`func NewBaseUsersFilters() *BaseUsersFilters`

NewBaseUsersFilters instantiates a new BaseUsersFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseUsersFiltersWithDefaults

`func NewBaseUsersFiltersWithDefaults() *BaseUsersFilters`

NewBaseUsersFiltersWithDefaults instantiates a new BaseUsersFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserType

`func (o *BaseUsersFilters) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *BaseUsersFilters) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *BaseUsersFilters) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *BaseUsersFilters) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetComponentUser

`func (o *BaseUsersFilters) GetComponentUser() bool`

GetComponentUser returns the ComponentUser field if non-nil, zero value otherwise.

### GetComponentUserOk

`func (o *BaseUsersFilters) GetComponentUserOk() (*bool, bool)`

GetComponentUserOk returns a tuple with the ComponentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUser

`func (o *BaseUsersFilters) SetComponentUser(v bool)`

SetComponentUser sets ComponentUser field to given value.

### HasComponentUser

`func (o *BaseUsersFilters) HasComponentUser() bool`

HasComponentUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


