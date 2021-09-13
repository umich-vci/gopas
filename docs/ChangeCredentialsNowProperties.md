# ChangeCredentialsNowProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeEntireGroup** | Pointer to **bool** | Whether or not the CPM will change the credentials in all the accounts that belong to the same group.   This parameter is only relevant for accounts that belong to an account group, and if this parameter does not belong to a group,   it will be ignored.If this account is part of an account group and this value is not specified, the default value will be applied. | [optional] 

## Methods

### NewChangeCredentialsNowProperties

`func NewChangeCredentialsNowProperties() *ChangeCredentialsNowProperties`

NewChangeCredentialsNowProperties instantiates a new ChangeCredentialsNowProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeCredentialsNowPropertiesWithDefaults

`func NewChangeCredentialsNowPropertiesWithDefaults() *ChangeCredentialsNowProperties`

NewChangeCredentialsNowPropertiesWithDefaults instantiates a new ChangeCredentialsNowProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeEntireGroup

`func (o *ChangeCredentialsNowProperties) GetChangeEntireGroup() bool`

GetChangeEntireGroup returns the ChangeEntireGroup field if non-nil, zero value otherwise.

### GetChangeEntireGroupOk

`func (o *ChangeCredentialsNowProperties) GetChangeEntireGroupOk() (*bool, bool)`

GetChangeEntireGroupOk returns a tuple with the ChangeEntireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeEntireGroup

`func (o *ChangeCredentialsNowProperties) SetChangeEntireGroup(v bool)`

SetChangeEntireGroup sets ChangeEntireGroup field to given value.

### HasChangeEntireGroup

`func (o *ChangeCredentialsNowProperties) HasChangeEntireGroup() bool`

HasChangeEntireGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


