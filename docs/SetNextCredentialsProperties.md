# SetNextCredentialsProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCredentials** | **string** | The new account credentials that will be allocated to the account in the Vault. | 
**ChangeImmediately** | Pointer to **bool** | Whether or not the password will be changed immediately in the Vault. | [optional] 

## Methods

### NewSetNextCredentialsProperties

`func NewSetNextCredentialsProperties(newCredentials string, ) *SetNextCredentialsProperties`

NewSetNextCredentialsProperties instantiates a new SetNextCredentialsProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNextCredentialsPropertiesWithDefaults

`func NewSetNextCredentialsPropertiesWithDefaults() *SetNextCredentialsProperties`

NewSetNextCredentialsPropertiesWithDefaults instantiates a new SetNextCredentialsProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewCredentials

`func (o *SetNextCredentialsProperties) GetNewCredentials() string`

GetNewCredentials returns the NewCredentials field if non-nil, zero value otherwise.

### GetNewCredentialsOk

`func (o *SetNextCredentialsProperties) GetNewCredentialsOk() (*string, bool)`

GetNewCredentialsOk returns a tuple with the NewCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCredentials

`func (o *SetNextCredentialsProperties) SetNewCredentials(v string)`

SetNewCredentials sets NewCredentials field to given value.


### GetChangeImmediately

`func (o *SetNextCredentialsProperties) GetChangeImmediately() bool`

GetChangeImmediately returns the ChangeImmediately field if non-nil, zero value otherwise.

### GetChangeImmediatelyOk

`func (o *SetNextCredentialsProperties) GetChangeImmediatelyOk() (*bool, bool)`

GetChangeImmediatelyOk returns a tuple with the ChangeImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeImmediately

`func (o *SetNextCredentialsProperties) SetChangeImmediately(v bool)`

SetChangeImmediately sets ChangeImmediately field to given value.

### HasChangeImmediately

`func (o *SetNextCredentialsProperties) HasChangeImmediately() bool`

HasChangeImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


