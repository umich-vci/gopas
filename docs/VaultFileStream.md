# VaultFileStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StayLockAfterUpdate** | Pointer to **bool** |  | [optional] 
**SeekHandle** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**Exist** | Pointer to **bool** |  | [optional] [readonly] 
**CanRead** | Pointer to **bool** |  | [optional] [readonly] 
**CanSeek** | Pointer to **bool** |  | [optional] [readonly] 
**CanWrite** | Pointer to **bool** |  | [optional] [readonly] 
**Position** | Pointer to **int64** |  | [optional] 
**Length** | Pointer to **int64** |  | [optional] [readonly] 
**PerformedWrite** | Pointer to **bool** |  | [optional] 
**CommitChangesOnClose** | Pointer to **bool** |  | [optional] 
**PerformedSeekResume** | Pointer to **bool** |  | [optional] 
**CanTimeout** | Pointer to **bool** |  | [optional] [readonly] 
**ReadTimeout** | Pointer to **int32** |  | [optional] 
**WriteTimeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewVaultFileStream

`func NewVaultFileStream() *VaultFileStream`

NewVaultFileStream instantiates a new VaultFileStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultFileStreamWithDefaults

`func NewVaultFileStreamWithDefaults() *VaultFileStream`

NewVaultFileStreamWithDefaults instantiates a new VaultFileStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStayLockAfterUpdate

`func (o *VaultFileStream) GetStayLockAfterUpdate() bool`

GetStayLockAfterUpdate returns the StayLockAfterUpdate field if non-nil, zero value otherwise.

### GetStayLockAfterUpdateOk

`func (o *VaultFileStream) GetStayLockAfterUpdateOk() (*bool, bool)`

GetStayLockAfterUpdateOk returns a tuple with the StayLockAfterUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStayLockAfterUpdate

`func (o *VaultFileStream) SetStayLockAfterUpdate(v bool)`

SetStayLockAfterUpdate sets StayLockAfterUpdate field to given value.

### HasStayLockAfterUpdate

`func (o *VaultFileStream) HasStayLockAfterUpdate() bool`

HasStayLockAfterUpdate returns a boolean if a field has been set.

### GetSeekHandle

`func (o *VaultFileStream) GetSeekHandle() string`

GetSeekHandle returns the SeekHandle field if non-nil, zero value otherwise.

### GetSeekHandleOk

`func (o *VaultFileStream) GetSeekHandleOk() (*string, bool)`

GetSeekHandleOk returns a tuple with the SeekHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeekHandle

`func (o *VaultFileStream) SetSeekHandle(v string)`

SetSeekHandle sets SeekHandle field to given value.

### HasSeekHandle

`func (o *VaultFileStream) HasSeekHandle() bool`

HasSeekHandle returns a boolean if a field has been set.

### GetCreationDate

`func (o *VaultFileStream) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VaultFileStream) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VaultFileStream) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VaultFileStream) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetModificationDate

`func (o *VaultFileStream) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *VaultFileStream) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *VaultFileStream) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *VaultFileStream) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetExist

`func (o *VaultFileStream) GetExist() bool`

GetExist returns the Exist field if non-nil, zero value otherwise.

### GetExistOk

`func (o *VaultFileStream) GetExistOk() (*bool, bool)`

GetExistOk returns a tuple with the Exist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExist

`func (o *VaultFileStream) SetExist(v bool)`

SetExist sets Exist field to given value.

### HasExist

`func (o *VaultFileStream) HasExist() bool`

HasExist returns a boolean if a field has been set.

### GetCanRead

`func (o *VaultFileStream) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *VaultFileStream) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *VaultFileStream) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *VaultFileStream) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### GetCanSeek

`func (o *VaultFileStream) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *VaultFileStream) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *VaultFileStream) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *VaultFileStream) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetCanWrite

`func (o *VaultFileStream) GetCanWrite() bool`

GetCanWrite returns the CanWrite field if non-nil, zero value otherwise.

### GetCanWriteOk

`func (o *VaultFileStream) GetCanWriteOk() (*bool, bool)`

GetCanWriteOk returns a tuple with the CanWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWrite

`func (o *VaultFileStream) SetCanWrite(v bool)`

SetCanWrite sets CanWrite field to given value.

### HasCanWrite

`func (o *VaultFileStream) HasCanWrite() bool`

HasCanWrite returns a boolean if a field has been set.

### GetPosition

`func (o *VaultFileStream) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VaultFileStream) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VaultFileStream) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VaultFileStream) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLength

`func (o *VaultFileStream) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VaultFileStream) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VaultFileStream) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *VaultFileStream) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPerformedWrite

`func (o *VaultFileStream) GetPerformedWrite() bool`

GetPerformedWrite returns the PerformedWrite field if non-nil, zero value otherwise.

### GetPerformedWriteOk

`func (o *VaultFileStream) GetPerformedWriteOk() (*bool, bool)`

GetPerformedWriteOk returns a tuple with the PerformedWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedWrite

`func (o *VaultFileStream) SetPerformedWrite(v bool)`

SetPerformedWrite sets PerformedWrite field to given value.

### HasPerformedWrite

`func (o *VaultFileStream) HasPerformedWrite() bool`

HasPerformedWrite returns a boolean if a field has been set.

### GetCommitChangesOnClose

`func (o *VaultFileStream) GetCommitChangesOnClose() bool`

GetCommitChangesOnClose returns the CommitChangesOnClose field if non-nil, zero value otherwise.

### GetCommitChangesOnCloseOk

`func (o *VaultFileStream) GetCommitChangesOnCloseOk() (*bool, bool)`

GetCommitChangesOnCloseOk returns a tuple with the CommitChangesOnClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitChangesOnClose

`func (o *VaultFileStream) SetCommitChangesOnClose(v bool)`

SetCommitChangesOnClose sets CommitChangesOnClose field to given value.

### HasCommitChangesOnClose

`func (o *VaultFileStream) HasCommitChangesOnClose() bool`

HasCommitChangesOnClose returns a boolean if a field has been set.

### GetPerformedSeekResume

`func (o *VaultFileStream) GetPerformedSeekResume() bool`

GetPerformedSeekResume returns the PerformedSeekResume field if non-nil, zero value otherwise.

### GetPerformedSeekResumeOk

`func (o *VaultFileStream) GetPerformedSeekResumeOk() (*bool, bool)`

GetPerformedSeekResumeOk returns a tuple with the PerformedSeekResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedSeekResume

`func (o *VaultFileStream) SetPerformedSeekResume(v bool)`

SetPerformedSeekResume sets PerformedSeekResume field to given value.

### HasPerformedSeekResume

`func (o *VaultFileStream) HasPerformedSeekResume() bool`

HasPerformedSeekResume returns a boolean if a field has been set.

### GetCanTimeout

`func (o *VaultFileStream) GetCanTimeout() bool`

GetCanTimeout returns the CanTimeout field if non-nil, zero value otherwise.

### GetCanTimeoutOk

`func (o *VaultFileStream) GetCanTimeoutOk() (*bool, bool)`

GetCanTimeoutOk returns a tuple with the CanTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTimeout

`func (o *VaultFileStream) SetCanTimeout(v bool)`

SetCanTimeout sets CanTimeout field to given value.

### HasCanTimeout

`func (o *VaultFileStream) HasCanTimeout() bool`

HasCanTimeout returns a boolean if a field has been set.

### GetReadTimeout

`func (o *VaultFileStream) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *VaultFileStream) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *VaultFileStream) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *VaultFileStream) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetWriteTimeout

`func (o *VaultFileStream) GetWriteTimeout() int32`

GetWriteTimeout returns the WriteTimeout field if non-nil, zero value otherwise.

### GetWriteTimeoutOk

`func (o *VaultFileStream) GetWriteTimeoutOk() (*int32, bool)`

GetWriteTimeoutOk returns a tuple with the WriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteTimeout

`func (o *VaultFileStream) SetWriteTimeout(v int32)`

SetWriteTimeout sets WriteTimeout field to given value.

### HasWriteTimeout

`func (o *VaultFileStream) HasWriteTimeout() bool`

HasWriteTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


