# LinkAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The linked account name. | 
**Safe** | **string** | The safe in which the linked account is stored. | 
**Folder** | **string** | The folder in which the linked account is stored. | 
**ExtraPasswordIndex** | **int32** | The linked account extra pass index.  The index can be of Reconcile Account, Logon Account or other linked account which is defined in the platform configuration. | 

## Methods

### NewLinkAccountData

`func NewLinkAccountData(name string, safe string, folder string, extraPasswordIndex int32, ) *LinkAccountData`

NewLinkAccountData instantiates a new LinkAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAccountDataWithDefaults

`func NewLinkAccountDataWithDefaults() *LinkAccountData`

NewLinkAccountDataWithDefaults instantiates a new LinkAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinkAccountData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkAccountData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkAccountData) SetName(v string)`

SetName sets Name field to given value.


### GetSafe

`func (o *LinkAccountData) GetSafe() string`

GetSafe returns the Safe field if non-nil, zero value otherwise.

### GetSafeOk

`func (o *LinkAccountData) GetSafeOk() (*string, bool)`

GetSafeOk returns a tuple with the Safe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafe

`func (o *LinkAccountData) SetSafe(v string)`

SetSafe sets Safe field to given value.


### GetFolder

`func (o *LinkAccountData) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *LinkAccountData) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *LinkAccountData) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetExtraPasswordIndex

`func (o *LinkAccountData) GetExtraPasswordIndex() int32`

GetExtraPasswordIndex returns the ExtraPasswordIndex field if non-nil, zero value otherwise.

### GetExtraPasswordIndexOk

`func (o *LinkAccountData) GetExtraPasswordIndexOk() (*int32, bool)`

GetExtraPasswordIndexOk returns a tuple with the ExtraPasswordIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPasswordIndex

`func (o *LinkAccountData) SetExtraPasswordIndex(v int32)`

SetExtraPasswordIndex sets ExtraPasswordIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


