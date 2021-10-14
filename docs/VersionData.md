# VersionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionID** | Pointer to **int64** | Secret&#39;s version unique ID | [optional] 
**ModifiedBy** | Pointer to **string** | The user that modified the secret | [optional] 
**ModificationDate** | Pointer to **int64** | Secret&#39;s modification date | [optional] 
**IsTemporary** | Pointer to **bool** | Indication whether it&#39;s temporary or real password | [optional] 

## Methods

### NewVersionData

`func NewVersionData() *VersionData`

NewVersionData instantiates a new VersionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionDataWithDefaults

`func NewVersionDataWithDefaults() *VersionData`

NewVersionDataWithDefaults instantiates a new VersionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionID

`func (o *VersionData) GetVersionID() int64`

GetVersionID returns the VersionID field if non-nil, zero value otherwise.

### GetVersionIDOk

`func (o *VersionData) GetVersionIDOk() (*int64, bool)`

GetVersionIDOk returns a tuple with the VersionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionID

`func (o *VersionData) SetVersionID(v int64)`

SetVersionID sets VersionID field to given value.

### HasVersionID

`func (o *VersionData) HasVersionID() bool`

HasVersionID returns a boolean if a field has been set.

### GetModifiedBy

`func (o *VersionData) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VersionData) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VersionData) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *VersionData) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModificationDate

`func (o *VersionData) GetModificationDate() int64`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *VersionData) GetModificationDateOk() (*int64, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *VersionData) SetModificationDate(v int64)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *VersionData) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetIsTemporary

`func (o *VersionData) GetIsTemporary() bool`

GetIsTemporary returns the IsTemporary field if non-nil, zero value otherwise.

### GetIsTemporaryOk

`func (o *VersionData) GetIsTemporaryOk() (*bool, bool)`

GetIsTemporaryOk returns a tuple with the IsTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemporary

`func (o *VersionData) SetIsTemporary(v bool)`

SetIsTemporary sets IsTemporary field to given value.

### HasIsTemporary

`func (o *VersionData) HasIsTemporary() bool`

HasIsTemporary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


