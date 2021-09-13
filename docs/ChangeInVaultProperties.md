# ChangeInVaultProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeEntireGroup** | Pointer to **bool** | Whether or not to change the password in all accounts that belong to the same group.  This parameter is only relevant for accounts that belong to an account group.If this parameter does not belong to a group, it will be ignored. | [optional] 
**NewCredentials** | **string** | Whether or not the password will be generated according to the password policy rules. If the CPM is not configured to enforce a password policy rule, this parameter is irrelevant.  If the NewCredentails parameter contains a value, this parameter will be ignored. | 

## Methods

### NewChangeInVaultProperties

`func NewChangeInVaultProperties(newCredentials string, ) *ChangeInVaultProperties`

NewChangeInVaultProperties instantiates a new ChangeInVaultProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeInVaultPropertiesWithDefaults

`func NewChangeInVaultPropertiesWithDefaults() *ChangeInVaultProperties`

NewChangeInVaultPropertiesWithDefaults instantiates a new ChangeInVaultProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeEntireGroup

`func (o *ChangeInVaultProperties) GetChangeEntireGroup() bool`

GetChangeEntireGroup returns the ChangeEntireGroup field if non-nil, zero value otherwise.

### GetChangeEntireGroupOk

`func (o *ChangeInVaultProperties) GetChangeEntireGroupOk() (*bool, bool)`

GetChangeEntireGroupOk returns a tuple with the ChangeEntireGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeEntireGroup

`func (o *ChangeInVaultProperties) SetChangeEntireGroup(v bool)`

SetChangeEntireGroup sets ChangeEntireGroup field to given value.

### HasChangeEntireGroup

`func (o *ChangeInVaultProperties) HasChangeEntireGroup() bool`

HasChangeEntireGroup returns a boolean if a field has been set.

### GetNewCredentials

`func (o *ChangeInVaultProperties) GetNewCredentials() string`

GetNewCredentials returns the NewCredentials field if non-nil, zero value otherwise.

### GetNewCredentialsOk

`func (o *ChangeInVaultProperties) GetNewCredentialsOk() (*string, bool)`

GetNewCredentialsOk returns a tuple with the NewCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCredentials

`func (o *ChangeInVaultProperties) SetNewCredentials(v string)`

SetNewCredentials sets NewCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


