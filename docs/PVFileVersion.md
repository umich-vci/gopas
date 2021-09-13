# PVFileVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputStream** | Pointer to [**VaultFileStream**](VaultFileStream.md) |  | [optional] 
**OutputStream** | Pointer to [**VaultFileStream**](VaultFileStream.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Folder** | Pointer to **string** |  | [optional] [readonly] 
**Safe** | Pointer to **string** |  | [optional] [readonly] 
**InternalName** | Pointer to **string** |  | [optional] [readonly] 
**CreationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**Deletiondate** | Pointer to **time.Time** |  | [optional] [readonly] 
**DeletedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastUsedDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUsedBy** | Pointer to **string** |  | [optional] [readonly] 
**Size** | Pointer to **float64** |  | [optional] [readonly] 
**History** | Pointer to **int64** |  | [optional] [readonly] 
**Expdt** | Pointer to **time.Time** |  | [optional] [readonly] 
**FileAccessMarkFlags** | Pointer to **int32** |  | [optional] [readonly] 
**OwnerFlags** | Pointer to **int32** |  | [optional] [readonly] 
**LockDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**LockedBy** | Pointer to **string** |  | [optional] [readonly] 
**ObjectType** | Pointer to **int32** |  | [optional] [readonly] 
**FileNumber** | Pointer to **int32** |  | [optional] 
**IsDraft** | Pointer to **bool** |  | [optional] [readonly] 
**IsRetrieveLock** | Pointer to **bool** |  | [optional] [readonly] 
**FileCreationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**FileCreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**AgentLockedBy** | Pointer to **string** |  | [optional] [readonly] 
**Options** | Pointer to **int64** |  | [optional] [readonly] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**RealModificationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**RealModifiedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastSuccessfulBiosScan** | Pointer to **time.Time** |  | [optional] 
**LastFailureBiosScan** | Pointer to **time.Time** |  | [optional] 
**LastFailureBiosScanReason** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastHumanUsedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastHumanUsedDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastHumanRetrievedBy** | Pointer to **string** |  | [optional] [readonly] 
**LastHumanRetrievedDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**SafeID** | Pointer to **int64** |  | [optional] 
**AccessLevel** | Pointer to **int64** |  | [optional] [readonly] 
**MonitorAuth** | Pointer to **bool** |  | [optional] [readonly] 
**ManagerAuth** | Pointer to **bool** |  | [optional] [readonly] 
**GetAuth** | Pointer to **bool** |  | [optional] [readonly] 
**PutAuth** | Pointer to **bool** |  | [optional] [readonly] 
**DeleteAuth** | Pointer to **bool** |  | [optional] [readonly] 
**SupervisorAuth** | Pointer to **bool** |  | [optional] [readonly] 
**BackupAuth** | Pointer to **bool** |  | [optional] [readonly] 
**OwnersManagerAuth** | Pointer to **bool** |  | [optional] [readonly] 
**NoConfirmRequiredAuth** | Pointer to **bool** |  | [optional] [readonly] 
**ValidateSafeContentAuth** | Pointer to **bool** |  | [optional] [readonly] 
**ListAuth** | Pointer to **bool** |  | [optional] [readonly] 
**UsePasswordAuth** | Pointer to **bool** |  | [optional] [readonly] 
**UpdateObjectPropertiesAuth** | Pointer to **bool** |  | [optional] [readonly] 
**InitiateCPMChangeAuth** | Pointer to **bool** |  | [optional] [readonly] 
**InitiateCPMChangeWithManualPasswordAuth** | Pointer to **bool** |  | [optional] [readonly] 
**CreateFolderAuth** | Pointer to **bool** |  | [optional] [readonly] 
**DeleteFolderAuth** | Pointer to **bool** |  | [optional] [readonly] 
**MoveFromAuth** | Pointer to **bool** |  | [optional] [readonly] 
**MoveIntoAuth** | Pointer to **bool** |  | [optional] [readonly] 
**ViewAuditAuth** | Pointer to **bool** |  | [optional] [readonly] 
**ViewPermissionsAuth** | Pointer to **bool** |  | [optional] [readonly] 
**EventsListAuth** | Pointer to **bool** |  | [optional] [readonly] 
**EventsAddAuth** | Pointer to **bool** |  | [optional] [readonly] 
**CreateObjectAuth** | Pointer to **bool** |  | [optional] [readonly] 
**UnlockObjectAuth** | Pointer to **bool** |  | [optional] [readonly] 
**RenameObjectAuth** | Pointer to **bool** |  | [optional] [readonly] 
**AllAuth** | Pointer to **bool** |  | [optional] [readonly] 
**NeedConfirmStatus** | Pointer to **int64** |  | [optional] [readonly] 
**IsTempVersion** | Pointer to **bool** |  | [optional] [readonly] 
**CategoryModificationDate** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPVFileVersion

`func NewPVFileVersion() *PVFileVersion`

NewPVFileVersion instantiates a new PVFileVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPVFileVersionWithDefaults

`func NewPVFileVersionWithDefaults() *PVFileVersion`

NewPVFileVersionWithDefaults instantiates a new PVFileVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputStream

`func (o *PVFileVersion) GetInputStream() VaultFileStream`

GetInputStream returns the InputStream field if non-nil, zero value otherwise.

### GetInputStreamOk

`func (o *PVFileVersion) GetInputStreamOk() (*VaultFileStream, bool)`

GetInputStreamOk returns a tuple with the InputStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputStream

`func (o *PVFileVersion) SetInputStream(v VaultFileStream)`

SetInputStream sets InputStream field to given value.

### HasInputStream

`func (o *PVFileVersion) HasInputStream() bool`

HasInputStream returns a boolean if a field has been set.

### GetOutputStream

`func (o *PVFileVersion) GetOutputStream() VaultFileStream`

GetOutputStream returns the OutputStream field if non-nil, zero value otherwise.

### GetOutputStreamOk

`func (o *PVFileVersion) GetOutputStreamOk() (*VaultFileStream, bool)`

GetOutputStreamOk returns a tuple with the OutputStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStream

`func (o *PVFileVersion) SetOutputStream(v VaultFileStream)`

SetOutputStream sets OutputStream field to given value.

### HasOutputStream

`func (o *PVFileVersion) HasOutputStream() bool`

HasOutputStream returns a boolean if a field has been set.

### GetName

`func (o *PVFileVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PVFileVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PVFileVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PVFileVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFolder

`func (o *PVFileVersion) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *PVFileVersion) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *PVFileVersion) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *PVFileVersion) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetSafe

`func (o *PVFileVersion) GetSafe() string`

GetSafe returns the Safe field if non-nil, zero value otherwise.

### GetSafeOk

`func (o *PVFileVersion) GetSafeOk() (*string, bool)`

GetSafeOk returns a tuple with the Safe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafe

`func (o *PVFileVersion) SetSafe(v string)`

SetSafe sets Safe field to given value.

### HasSafe

`func (o *PVFileVersion) HasSafe() bool`

HasSafe returns a boolean if a field has been set.

### GetInternalName

`func (o *PVFileVersion) GetInternalName() string`

GetInternalName returns the InternalName field if non-nil, zero value otherwise.

### GetInternalNameOk

`func (o *PVFileVersion) GetInternalNameOk() (*string, bool)`

GetInternalNameOk returns a tuple with the InternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalName

`func (o *PVFileVersion) SetInternalName(v string)`

SetInternalName sets InternalName field to given value.

### HasInternalName

`func (o *PVFileVersion) HasInternalName() bool`

HasInternalName returns a boolean if a field has been set.

### GetCreationDate

`func (o *PVFileVersion) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PVFileVersion) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PVFileVersion) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PVFileVersion) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PVFileVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PVFileVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PVFileVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PVFileVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletiondate

`func (o *PVFileVersion) GetDeletiondate() time.Time`

GetDeletiondate returns the Deletiondate field if non-nil, zero value otherwise.

### GetDeletiondateOk

`func (o *PVFileVersion) GetDeletiondateOk() (*time.Time, bool)`

GetDeletiondateOk returns a tuple with the Deletiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletiondate

`func (o *PVFileVersion) SetDeletiondate(v time.Time)`

SetDeletiondate sets Deletiondate field to given value.

### HasDeletiondate

`func (o *PVFileVersion) HasDeletiondate() bool`

HasDeletiondate returns a boolean if a field has been set.

### GetDeletedBy

`func (o *PVFileVersion) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *PVFileVersion) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *PVFileVersion) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *PVFileVersion) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetLastUsedDate

`func (o *PVFileVersion) GetLastUsedDate() time.Time`

GetLastUsedDate returns the LastUsedDate field if non-nil, zero value otherwise.

### GetLastUsedDateOk

`func (o *PVFileVersion) GetLastUsedDateOk() (*time.Time, bool)`

GetLastUsedDateOk returns a tuple with the LastUsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDate

`func (o *PVFileVersion) SetLastUsedDate(v time.Time)`

SetLastUsedDate sets LastUsedDate field to given value.

### HasLastUsedDate

`func (o *PVFileVersion) HasLastUsedDate() bool`

HasLastUsedDate returns a boolean if a field has been set.

### GetLastUsedBy

`func (o *PVFileVersion) GetLastUsedBy() string`

GetLastUsedBy returns the LastUsedBy field if non-nil, zero value otherwise.

### GetLastUsedByOk

`func (o *PVFileVersion) GetLastUsedByOk() (*string, bool)`

GetLastUsedByOk returns a tuple with the LastUsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedBy

`func (o *PVFileVersion) SetLastUsedBy(v string)`

SetLastUsedBy sets LastUsedBy field to given value.

### HasLastUsedBy

`func (o *PVFileVersion) HasLastUsedBy() bool`

HasLastUsedBy returns a boolean if a field has been set.

### GetSize

`func (o *PVFileVersion) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PVFileVersion) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PVFileVersion) SetSize(v float64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PVFileVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetHistory

`func (o *PVFileVersion) GetHistory() int64`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PVFileVersion) GetHistoryOk() (*int64, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PVFileVersion) SetHistory(v int64)`

SetHistory sets History field to given value.

### HasHistory

`func (o *PVFileVersion) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetExpdt

`func (o *PVFileVersion) GetExpdt() time.Time`

GetExpdt returns the Expdt field if non-nil, zero value otherwise.

### GetExpdtOk

`func (o *PVFileVersion) GetExpdtOk() (*time.Time, bool)`

GetExpdtOk returns a tuple with the Expdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpdt

`func (o *PVFileVersion) SetExpdt(v time.Time)`

SetExpdt sets Expdt field to given value.

### HasExpdt

`func (o *PVFileVersion) HasExpdt() bool`

HasExpdt returns a boolean if a field has been set.

### GetFileAccessMarkFlags

`func (o *PVFileVersion) GetFileAccessMarkFlags() int32`

GetFileAccessMarkFlags returns the FileAccessMarkFlags field if non-nil, zero value otherwise.

### GetFileAccessMarkFlagsOk

`func (o *PVFileVersion) GetFileAccessMarkFlagsOk() (*int32, bool)`

GetFileAccessMarkFlagsOk returns a tuple with the FileAccessMarkFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAccessMarkFlags

`func (o *PVFileVersion) SetFileAccessMarkFlags(v int32)`

SetFileAccessMarkFlags sets FileAccessMarkFlags field to given value.

### HasFileAccessMarkFlags

`func (o *PVFileVersion) HasFileAccessMarkFlags() bool`

HasFileAccessMarkFlags returns a boolean if a field has been set.

### GetOwnerFlags

`func (o *PVFileVersion) GetOwnerFlags() int32`

GetOwnerFlags returns the OwnerFlags field if non-nil, zero value otherwise.

### GetOwnerFlagsOk

`func (o *PVFileVersion) GetOwnerFlagsOk() (*int32, bool)`

GetOwnerFlagsOk returns a tuple with the OwnerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFlags

`func (o *PVFileVersion) SetOwnerFlags(v int32)`

SetOwnerFlags sets OwnerFlags field to given value.

### HasOwnerFlags

`func (o *PVFileVersion) HasOwnerFlags() bool`

HasOwnerFlags returns a boolean if a field has been set.

### GetLockDate

`func (o *PVFileVersion) GetLockDate() time.Time`

GetLockDate returns the LockDate field if non-nil, zero value otherwise.

### GetLockDateOk

`func (o *PVFileVersion) GetLockDateOk() (*time.Time, bool)`

GetLockDateOk returns a tuple with the LockDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDate

`func (o *PVFileVersion) SetLockDate(v time.Time)`

SetLockDate sets LockDate field to given value.

### HasLockDate

`func (o *PVFileVersion) HasLockDate() bool`

HasLockDate returns a boolean if a field has been set.

### GetLockedBy

`func (o *PVFileVersion) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *PVFileVersion) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *PVFileVersion) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *PVFileVersion) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetObjectType

`func (o *PVFileVersion) GetObjectType() int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PVFileVersion) GetObjectTypeOk() (*int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PVFileVersion) SetObjectType(v int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PVFileVersion) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetFileNumber

`func (o *PVFileVersion) GetFileNumber() int32`

GetFileNumber returns the FileNumber field if non-nil, zero value otherwise.

### GetFileNumberOk

`func (o *PVFileVersion) GetFileNumberOk() (*int32, bool)`

GetFileNumberOk returns a tuple with the FileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNumber

`func (o *PVFileVersion) SetFileNumber(v int32)`

SetFileNumber sets FileNumber field to given value.

### HasFileNumber

`func (o *PVFileVersion) HasFileNumber() bool`

HasFileNumber returns a boolean if a field has been set.

### GetIsDraft

`func (o *PVFileVersion) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *PVFileVersion) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *PVFileVersion) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *PVFileVersion) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetIsRetrieveLock

`func (o *PVFileVersion) GetIsRetrieveLock() bool`

GetIsRetrieveLock returns the IsRetrieveLock field if non-nil, zero value otherwise.

### GetIsRetrieveLockOk

`func (o *PVFileVersion) GetIsRetrieveLockOk() (*bool, bool)`

GetIsRetrieveLockOk returns a tuple with the IsRetrieveLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRetrieveLock

`func (o *PVFileVersion) SetIsRetrieveLock(v bool)`

SetIsRetrieveLock sets IsRetrieveLock field to given value.

### HasIsRetrieveLock

`func (o *PVFileVersion) HasIsRetrieveLock() bool`

HasIsRetrieveLock returns a boolean if a field has been set.

### GetFileCreationDate

`func (o *PVFileVersion) GetFileCreationDate() time.Time`

GetFileCreationDate returns the FileCreationDate field if non-nil, zero value otherwise.

### GetFileCreationDateOk

`func (o *PVFileVersion) GetFileCreationDateOk() (*time.Time, bool)`

GetFileCreationDateOk returns a tuple with the FileCreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreationDate

`func (o *PVFileVersion) SetFileCreationDate(v time.Time)`

SetFileCreationDate sets FileCreationDate field to given value.

### HasFileCreationDate

`func (o *PVFileVersion) HasFileCreationDate() bool`

HasFileCreationDate returns a boolean if a field has been set.

### GetFileCreatedBy

`func (o *PVFileVersion) GetFileCreatedBy() string`

GetFileCreatedBy returns the FileCreatedBy field if non-nil, zero value otherwise.

### GetFileCreatedByOk

`func (o *PVFileVersion) GetFileCreatedByOk() (*string, bool)`

GetFileCreatedByOk returns a tuple with the FileCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCreatedBy

`func (o *PVFileVersion) SetFileCreatedBy(v string)`

SetFileCreatedBy sets FileCreatedBy field to given value.

### HasFileCreatedBy

`func (o *PVFileVersion) HasFileCreatedBy() bool`

HasFileCreatedBy returns a boolean if a field has been set.

### GetAgentLockedBy

`func (o *PVFileVersion) GetAgentLockedBy() string`

GetAgentLockedBy returns the AgentLockedBy field if non-nil, zero value otherwise.

### GetAgentLockedByOk

`func (o *PVFileVersion) GetAgentLockedByOk() (*string, bool)`

GetAgentLockedByOk returns a tuple with the AgentLockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLockedBy

`func (o *PVFileVersion) SetAgentLockedBy(v string)`

SetAgentLockedBy sets AgentLockedBy field to given value.

### HasAgentLockedBy

`func (o *PVFileVersion) HasAgentLockedBy() bool`

HasAgentLockedBy returns a boolean if a field has been set.

### GetOptions

`func (o *PVFileVersion) GetOptions() int64`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PVFileVersion) GetOptionsOk() (*int64, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PVFileVersion) SetOptions(v int64)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PVFileVersion) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetModificationDate

`func (o *PVFileVersion) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *PVFileVersion) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *PVFileVersion) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *PVFileVersion) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetRealModificationDate

`func (o *PVFileVersion) GetRealModificationDate() time.Time`

GetRealModificationDate returns the RealModificationDate field if non-nil, zero value otherwise.

### GetRealModificationDateOk

`func (o *PVFileVersion) GetRealModificationDateOk() (*time.Time, bool)`

GetRealModificationDateOk returns a tuple with the RealModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealModificationDate

`func (o *PVFileVersion) SetRealModificationDate(v time.Time)`

SetRealModificationDate sets RealModificationDate field to given value.

### HasRealModificationDate

`func (o *PVFileVersion) HasRealModificationDate() bool`

HasRealModificationDate returns a boolean if a field has been set.

### GetRealModifiedBy

`func (o *PVFileVersion) GetRealModifiedBy() string`

GetRealModifiedBy returns the RealModifiedBy field if non-nil, zero value otherwise.

### GetRealModifiedByOk

`func (o *PVFileVersion) GetRealModifiedByOk() (*string, bool)`

GetRealModifiedByOk returns a tuple with the RealModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealModifiedBy

`func (o *PVFileVersion) SetRealModifiedBy(v string)`

SetRealModifiedBy sets RealModifiedBy field to given value.

### HasRealModifiedBy

`func (o *PVFileVersion) HasRealModifiedBy() bool`

HasRealModifiedBy returns a boolean if a field has been set.

### GetLastSuccessfulBiosScan

`func (o *PVFileVersion) GetLastSuccessfulBiosScan() time.Time`

GetLastSuccessfulBiosScan returns the LastSuccessfulBiosScan field if non-nil, zero value otherwise.

### GetLastSuccessfulBiosScanOk

`func (o *PVFileVersion) GetLastSuccessfulBiosScanOk() (*time.Time, bool)`

GetLastSuccessfulBiosScanOk returns a tuple with the LastSuccessfulBiosScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulBiosScan

`func (o *PVFileVersion) SetLastSuccessfulBiosScan(v time.Time)`

SetLastSuccessfulBiosScan sets LastSuccessfulBiosScan field to given value.

### HasLastSuccessfulBiosScan

`func (o *PVFileVersion) HasLastSuccessfulBiosScan() bool`

HasLastSuccessfulBiosScan returns a boolean if a field has been set.

### GetLastFailureBiosScan

`func (o *PVFileVersion) GetLastFailureBiosScan() time.Time`

GetLastFailureBiosScan returns the LastFailureBiosScan field if non-nil, zero value otherwise.

### GetLastFailureBiosScanOk

`func (o *PVFileVersion) GetLastFailureBiosScanOk() (*time.Time, bool)`

GetLastFailureBiosScanOk returns a tuple with the LastFailureBiosScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureBiosScan

`func (o *PVFileVersion) SetLastFailureBiosScan(v time.Time)`

SetLastFailureBiosScan sets LastFailureBiosScan field to given value.

### HasLastFailureBiosScan

`func (o *PVFileVersion) HasLastFailureBiosScan() bool`

HasLastFailureBiosScan returns a boolean if a field has been set.

### GetLastFailureBiosScanReason

`func (o *PVFileVersion) GetLastFailureBiosScanReason() string`

GetLastFailureBiosScanReason returns the LastFailureBiosScanReason field if non-nil, zero value otherwise.

### GetLastFailureBiosScanReasonOk

`func (o *PVFileVersion) GetLastFailureBiosScanReasonOk() (*string, bool)`

GetLastFailureBiosScanReasonOk returns a tuple with the LastFailureBiosScanReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureBiosScanReason

`func (o *PVFileVersion) SetLastFailureBiosScanReason(v string)`

SetLastFailureBiosScanReason sets LastFailureBiosScanReason field to given value.

### HasLastFailureBiosScanReason

`func (o *PVFileVersion) HasLastFailureBiosScanReason() bool`

HasLastFailureBiosScanReason returns a boolean if a field has been set.

### GetModifiedBy

`func (o *PVFileVersion) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PVFileVersion) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PVFileVersion) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *PVFileVersion) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetLastHumanUsedBy

`func (o *PVFileVersion) GetLastHumanUsedBy() string`

GetLastHumanUsedBy returns the LastHumanUsedBy field if non-nil, zero value otherwise.

### GetLastHumanUsedByOk

`func (o *PVFileVersion) GetLastHumanUsedByOk() (*string, bool)`

GetLastHumanUsedByOk returns a tuple with the LastHumanUsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHumanUsedBy

`func (o *PVFileVersion) SetLastHumanUsedBy(v string)`

SetLastHumanUsedBy sets LastHumanUsedBy field to given value.

### HasLastHumanUsedBy

`func (o *PVFileVersion) HasLastHumanUsedBy() bool`

HasLastHumanUsedBy returns a boolean if a field has been set.

### GetLastHumanUsedDate

`func (o *PVFileVersion) GetLastHumanUsedDate() time.Time`

GetLastHumanUsedDate returns the LastHumanUsedDate field if non-nil, zero value otherwise.

### GetLastHumanUsedDateOk

`func (o *PVFileVersion) GetLastHumanUsedDateOk() (*time.Time, bool)`

GetLastHumanUsedDateOk returns a tuple with the LastHumanUsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHumanUsedDate

`func (o *PVFileVersion) SetLastHumanUsedDate(v time.Time)`

SetLastHumanUsedDate sets LastHumanUsedDate field to given value.

### HasLastHumanUsedDate

`func (o *PVFileVersion) HasLastHumanUsedDate() bool`

HasLastHumanUsedDate returns a boolean if a field has been set.

### GetLastHumanRetrievedBy

`func (o *PVFileVersion) GetLastHumanRetrievedBy() string`

GetLastHumanRetrievedBy returns the LastHumanRetrievedBy field if non-nil, zero value otherwise.

### GetLastHumanRetrievedByOk

`func (o *PVFileVersion) GetLastHumanRetrievedByOk() (*string, bool)`

GetLastHumanRetrievedByOk returns a tuple with the LastHumanRetrievedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHumanRetrievedBy

`func (o *PVFileVersion) SetLastHumanRetrievedBy(v string)`

SetLastHumanRetrievedBy sets LastHumanRetrievedBy field to given value.

### HasLastHumanRetrievedBy

`func (o *PVFileVersion) HasLastHumanRetrievedBy() bool`

HasLastHumanRetrievedBy returns a boolean if a field has been set.

### GetLastHumanRetrievedDate

`func (o *PVFileVersion) GetLastHumanRetrievedDate() time.Time`

GetLastHumanRetrievedDate returns the LastHumanRetrievedDate field if non-nil, zero value otherwise.

### GetLastHumanRetrievedDateOk

`func (o *PVFileVersion) GetLastHumanRetrievedDateOk() (*time.Time, bool)`

GetLastHumanRetrievedDateOk returns a tuple with the LastHumanRetrievedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHumanRetrievedDate

`func (o *PVFileVersion) SetLastHumanRetrievedDate(v time.Time)`

SetLastHumanRetrievedDate sets LastHumanRetrievedDate field to given value.

### HasLastHumanRetrievedDate

`func (o *PVFileVersion) HasLastHumanRetrievedDate() bool`

HasLastHumanRetrievedDate returns a boolean if a field has been set.

### GetSafeID

`func (o *PVFileVersion) GetSafeID() int64`

GetSafeID returns the SafeID field if non-nil, zero value otherwise.

### GetSafeIDOk

`func (o *PVFileVersion) GetSafeIDOk() (*int64, bool)`

GetSafeIDOk returns a tuple with the SafeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeID

`func (o *PVFileVersion) SetSafeID(v int64)`

SetSafeID sets SafeID field to given value.

### HasSafeID

`func (o *PVFileVersion) HasSafeID() bool`

HasSafeID returns a boolean if a field has been set.

### GetAccessLevel

`func (o *PVFileVersion) GetAccessLevel() int64`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *PVFileVersion) GetAccessLevelOk() (*int64, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *PVFileVersion) SetAccessLevel(v int64)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *PVFileVersion) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetMonitorAuth

`func (o *PVFileVersion) GetMonitorAuth() bool`

GetMonitorAuth returns the MonitorAuth field if non-nil, zero value otherwise.

### GetMonitorAuthOk

`func (o *PVFileVersion) GetMonitorAuthOk() (*bool, bool)`

GetMonitorAuthOk returns a tuple with the MonitorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAuth

`func (o *PVFileVersion) SetMonitorAuth(v bool)`

SetMonitorAuth sets MonitorAuth field to given value.

### HasMonitorAuth

`func (o *PVFileVersion) HasMonitorAuth() bool`

HasMonitorAuth returns a boolean if a field has been set.

### GetManagerAuth

`func (o *PVFileVersion) GetManagerAuth() bool`

GetManagerAuth returns the ManagerAuth field if non-nil, zero value otherwise.

### GetManagerAuthOk

`func (o *PVFileVersion) GetManagerAuthOk() (*bool, bool)`

GetManagerAuthOk returns a tuple with the ManagerAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAuth

`func (o *PVFileVersion) SetManagerAuth(v bool)`

SetManagerAuth sets ManagerAuth field to given value.

### HasManagerAuth

`func (o *PVFileVersion) HasManagerAuth() bool`

HasManagerAuth returns a boolean if a field has been set.

### GetGetAuth

`func (o *PVFileVersion) GetGetAuth() bool`

GetGetAuth returns the GetAuth field if non-nil, zero value otherwise.

### GetGetAuthOk

`func (o *PVFileVersion) GetGetAuthOk() (*bool, bool)`

GetGetAuthOk returns a tuple with the GetAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetAuth

`func (o *PVFileVersion) SetGetAuth(v bool)`

SetGetAuth sets GetAuth field to given value.

### HasGetAuth

`func (o *PVFileVersion) HasGetAuth() bool`

HasGetAuth returns a boolean if a field has been set.

### GetPutAuth

`func (o *PVFileVersion) GetPutAuth() bool`

GetPutAuth returns the PutAuth field if non-nil, zero value otherwise.

### GetPutAuthOk

`func (o *PVFileVersion) GetPutAuthOk() (*bool, bool)`

GetPutAuthOk returns a tuple with the PutAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutAuth

`func (o *PVFileVersion) SetPutAuth(v bool)`

SetPutAuth sets PutAuth field to given value.

### HasPutAuth

`func (o *PVFileVersion) HasPutAuth() bool`

HasPutAuth returns a boolean if a field has been set.

### GetDeleteAuth

`func (o *PVFileVersion) GetDeleteAuth() bool`

GetDeleteAuth returns the DeleteAuth field if non-nil, zero value otherwise.

### GetDeleteAuthOk

`func (o *PVFileVersion) GetDeleteAuthOk() (*bool, bool)`

GetDeleteAuthOk returns a tuple with the DeleteAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAuth

`func (o *PVFileVersion) SetDeleteAuth(v bool)`

SetDeleteAuth sets DeleteAuth field to given value.

### HasDeleteAuth

`func (o *PVFileVersion) HasDeleteAuth() bool`

HasDeleteAuth returns a boolean if a field has been set.

### GetSupervisorAuth

`func (o *PVFileVersion) GetSupervisorAuth() bool`

GetSupervisorAuth returns the SupervisorAuth field if non-nil, zero value otherwise.

### GetSupervisorAuthOk

`func (o *PVFileVersion) GetSupervisorAuthOk() (*bool, bool)`

GetSupervisorAuthOk returns a tuple with the SupervisorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorAuth

`func (o *PVFileVersion) SetSupervisorAuth(v bool)`

SetSupervisorAuth sets SupervisorAuth field to given value.

### HasSupervisorAuth

`func (o *PVFileVersion) HasSupervisorAuth() bool`

HasSupervisorAuth returns a boolean if a field has been set.

### GetBackupAuth

`func (o *PVFileVersion) GetBackupAuth() bool`

GetBackupAuth returns the BackupAuth field if non-nil, zero value otherwise.

### GetBackupAuthOk

`func (o *PVFileVersion) GetBackupAuthOk() (*bool, bool)`

GetBackupAuthOk returns a tuple with the BackupAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAuth

`func (o *PVFileVersion) SetBackupAuth(v bool)`

SetBackupAuth sets BackupAuth field to given value.

### HasBackupAuth

`func (o *PVFileVersion) HasBackupAuth() bool`

HasBackupAuth returns a boolean if a field has been set.

### GetOwnersManagerAuth

`func (o *PVFileVersion) GetOwnersManagerAuth() bool`

GetOwnersManagerAuth returns the OwnersManagerAuth field if non-nil, zero value otherwise.

### GetOwnersManagerAuthOk

`func (o *PVFileVersion) GetOwnersManagerAuthOk() (*bool, bool)`

GetOwnersManagerAuthOk returns a tuple with the OwnersManagerAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnersManagerAuth

`func (o *PVFileVersion) SetOwnersManagerAuth(v bool)`

SetOwnersManagerAuth sets OwnersManagerAuth field to given value.

### HasOwnersManagerAuth

`func (o *PVFileVersion) HasOwnersManagerAuth() bool`

HasOwnersManagerAuth returns a boolean if a field has been set.

### GetNoConfirmRequiredAuth

`func (o *PVFileVersion) GetNoConfirmRequiredAuth() bool`

GetNoConfirmRequiredAuth returns the NoConfirmRequiredAuth field if non-nil, zero value otherwise.

### GetNoConfirmRequiredAuthOk

`func (o *PVFileVersion) GetNoConfirmRequiredAuthOk() (*bool, bool)`

GetNoConfirmRequiredAuthOk returns a tuple with the NoConfirmRequiredAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoConfirmRequiredAuth

`func (o *PVFileVersion) SetNoConfirmRequiredAuth(v bool)`

SetNoConfirmRequiredAuth sets NoConfirmRequiredAuth field to given value.

### HasNoConfirmRequiredAuth

`func (o *PVFileVersion) HasNoConfirmRequiredAuth() bool`

HasNoConfirmRequiredAuth returns a boolean if a field has been set.

### GetValidateSafeContentAuth

`func (o *PVFileVersion) GetValidateSafeContentAuth() bool`

GetValidateSafeContentAuth returns the ValidateSafeContentAuth field if non-nil, zero value otherwise.

### GetValidateSafeContentAuthOk

`func (o *PVFileVersion) GetValidateSafeContentAuthOk() (*bool, bool)`

GetValidateSafeContentAuthOk returns a tuple with the ValidateSafeContentAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSafeContentAuth

`func (o *PVFileVersion) SetValidateSafeContentAuth(v bool)`

SetValidateSafeContentAuth sets ValidateSafeContentAuth field to given value.

### HasValidateSafeContentAuth

`func (o *PVFileVersion) HasValidateSafeContentAuth() bool`

HasValidateSafeContentAuth returns a boolean if a field has been set.

### GetListAuth

`func (o *PVFileVersion) GetListAuth() bool`

GetListAuth returns the ListAuth field if non-nil, zero value otherwise.

### GetListAuthOk

`func (o *PVFileVersion) GetListAuthOk() (*bool, bool)`

GetListAuthOk returns a tuple with the ListAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAuth

`func (o *PVFileVersion) SetListAuth(v bool)`

SetListAuth sets ListAuth field to given value.

### HasListAuth

`func (o *PVFileVersion) HasListAuth() bool`

HasListAuth returns a boolean if a field has been set.

### GetUsePasswordAuth

`func (o *PVFileVersion) GetUsePasswordAuth() bool`

GetUsePasswordAuth returns the UsePasswordAuth field if non-nil, zero value otherwise.

### GetUsePasswordAuthOk

`func (o *PVFileVersion) GetUsePasswordAuthOk() (*bool, bool)`

GetUsePasswordAuthOk returns a tuple with the UsePasswordAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordAuth

`func (o *PVFileVersion) SetUsePasswordAuth(v bool)`

SetUsePasswordAuth sets UsePasswordAuth field to given value.

### HasUsePasswordAuth

`func (o *PVFileVersion) HasUsePasswordAuth() bool`

HasUsePasswordAuth returns a boolean if a field has been set.

### GetUpdateObjectPropertiesAuth

`func (o *PVFileVersion) GetUpdateObjectPropertiesAuth() bool`

GetUpdateObjectPropertiesAuth returns the UpdateObjectPropertiesAuth field if non-nil, zero value otherwise.

### GetUpdateObjectPropertiesAuthOk

`func (o *PVFileVersion) GetUpdateObjectPropertiesAuthOk() (*bool, bool)`

GetUpdateObjectPropertiesAuthOk returns a tuple with the UpdateObjectPropertiesAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateObjectPropertiesAuth

`func (o *PVFileVersion) SetUpdateObjectPropertiesAuth(v bool)`

SetUpdateObjectPropertiesAuth sets UpdateObjectPropertiesAuth field to given value.

### HasUpdateObjectPropertiesAuth

`func (o *PVFileVersion) HasUpdateObjectPropertiesAuth() bool`

HasUpdateObjectPropertiesAuth returns a boolean if a field has been set.

### GetInitiateCPMChangeAuth

`func (o *PVFileVersion) GetInitiateCPMChangeAuth() bool`

GetInitiateCPMChangeAuth returns the InitiateCPMChangeAuth field if non-nil, zero value otherwise.

### GetInitiateCPMChangeAuthOk

`func (o *PVFileVersion) GetInitiateCPMChangeAuthOk() (*bool, bool)`

GetInitiateCPMChangeAuthOk returns a tuple with the InitiateCPMChangeAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateCPMChangeAuth

`func (o *PVFileVersion) SetInitiateCPMChangeAuth(v bool)`

SetInitiateCPMChangeAuth sets InitiateCPMChangeAuth field to given value.

### HasInitiateCPMChangeAuth

`func (o *PVFileVersion) HasInitiateCPMChangeAuth() bool`

HasInitiateCPMChangeAuth returns a boolean if a field has been set.

### GetInitiateCPMChangeWithManualPasswordAuth

`func (o *PVFileVersion) GetInitiateCPMChangeWithManualPasswordAuth() bool`

GetInitiateCPMChangeWithManualPasswordAuth returns the InitiateCPMChangeWithManualPasswordAuth field if non-nil, zero value otherwise.

### GetInitiateCPMChangeWithManualPasswordAuthOk

`func (o *PVFileVersion) GetInitiateCPMChangeWithManualPasswordAuthOk() (*bool, bool)`

GetInitiateCPMChangeWithManualPasswordAuthOk returns a tuple with the InitiateCPMChangeWithManualPasswordAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateCPMChangeWithManualPasswordAuth

`func (o *PVFileVersion) SetInitiateCPMChangeWithManualPasswordAuth(v bool)`

SetInitiateCPMChangeWithManualPasswordAuth sets InitiateCPMChangeWithManualPasswordAuth field to given value.

### HasInitiateCPMChangeWithManualPasswordAuth

`func (o *PVFileVersion) HasInitiateCPMChangeWithManualPasswordAuth() bool`

HasInitiateCPMChangeWithManualPasswordAuth returns a boolean if a field has been set.

### GetCreateFolderAuth

`func (o *PVFileVersion) GetCreateFolderAuth() bool`

GetCreateFolderAuth returns the CreateFolderAuth field if non-nil, zero value otherwise.

### GetCreateFolderAuthOk

`func (o *PVFileVersion) GetCreateFolderAuthOk() (*bool, bool)`

GetCreateFolderAuthOk returns a tuple with the CreateFolderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFolderAuth

`func (o *PVFileVersion) SetCreateFolderAuth(v bool)`

SetCreateFolderAuth sets CreateFolderAuth field to given value.

### HasCreateFolderAuth

`func (o *PVFileVersion) HasCreateFolderAuth() bool`

HasCreateFolderAuth returns a boolean if a field has been set.

### GetDeleteFolderAuth

`func (o *PVFileVersion) GetDeleteFolderAuth() bool`

GetDeleteFolderAuth returns the DeleteFolderAuth field if non-nil, zero value otherwise.

### GetDeleteFolderAuthOk

`func (o *PVFileVersion) GetDeleteFolderAuthOk() (*bool, bool)`

GetDeleteFolderAuthOk returns a tuple with the DeleteFolderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFolderAuth

`func (o *PVFileVersion) SetDeleteFolderAuth(v bool)`

SetDeleteFolderAuth sets DeleteFolderAuth field to given value.

### HasDeleteFolderAuth

`func (o *PVFileVersion) HasDeleteFolderAuth() bool`

HasDeleteFolderAuth returns a boolean if a field has been set.

### GetMoveFromAuth

`func (o *PVFileVersion) GetMoveFromAuth() bool`

GetMoveFromAuth returns the MoveFromAuth field if non-nil, zero value otherwise.

### GetMoveFromAuthOk

`func (o *PVFileVersion) GetMoveFromAuthOk() (*bool, bool)`

GetMoveFromAuthOk returns a tuple with the MoveFromAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFromAuth

`func (o *PVFileVersion) SetMoveFromAuth(v bool)`

SetMoveFromAuth sets MoveFromAuth field to given value.

### HasMoveFromAuth

`func (o *PVFileVersion) HasMoveFromAuth() bool`

HasMoveFromAuth returns a boolean if a field has been set.

### GetMoveIntoAuth

`func (o *PVFileVersion) GetMoveIntoAuth() bool`

GetMoveIntoAuth returns the MoveIntoAuth field if non-nil, zero value otherwise.

### GetMoveIntoAuthOk

`func (o *PVFileVersion) GetMoveIntoAuthOk() (*bool, bool)`

GetMoveIntoAuthOk returns a tuple with the MoveIntoAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveIntoAuth

`func (o *PVFileVersion) SetMoveIntoAuth(v bool)`

SetMoveIntoAuth sets MoveIntoAuth field to given value.

### HasMoveIntoAuth

`func (o *PVFileVersion) HasMoveIntoAuth() bool`

HasMoveIntoAuth returns a boolean if a field has been set.

### GetViewAuditAuth

`func (o *PVFileVersion) GetViewAuditAuth() bool`

GetViewAuditAuth returns the ViewAuditAuth field if non-nil, zero value otherwise.

### GetViewAuditAuthOk

`func (o *PVFileVersion) GetViewAuditAuthOk() (*bool, bool)`

GetViewAuditAuthOk returns a tuple with the ViewAuditAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAuditAuth

`func (o *PVFileVersion) SetViewAuditAuth(v bool)`

SetViewAuditAuth sets ViewAuditAuth field to given value.

### HasViewAuditAuth

`func (o *PVFileVersion) HasViewAuditAuth() bool`

HasViewAuditAuth returns a boolean if a field has been set.

### GetViewPermissionsAuth

`func (o *PVFileVersion) GetViewPermissionsAuth() bool`

GetViewPermissionsAuth returns the ViewPermissionsAuth field if non-nil, zero value otherwise.

### GetViewPermissionsAuthOk

`func (o *PVFileVersion) GetViewPermissionsAuthOk() (*bool, bool)`

GetViewPermissionsAuthOk returns a tuple with the ViewPermissionsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPermissionsAuth

`func (o *PVFileVersion) SetViewPermissionsAuth(v bool)`

SetViewPermissionsAuth sets ViewPermissionsAuth field to given value.

### HasViewPermissionsAuth

`func (o *PVFileVersion) HasViewPermissionsAuth() bool`

HasViewPermissionsAuth returns a boolean if a field has been set.

### GetEventsListAuth

`func (o *PVFileVersion) GetEventsListAuth() bool`

GetEventsListAuth returns the EventsListAuth field if non-nil, zero value otherwise.

### GetEventsListAuthOk

`func (o *PVFileVersion) GetEventsListAuthOk() (*bool, bool)`

GetEventsListAuthOk returns a tuple with the EventsListAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsListAuth

`func (o *PVFileVersion) SetEventsListAuth(v bool)`

SetEventsListAuth sets EventsListAuth field to given value.

### HasEventsListAuth

`func (o *PVFileVersion) HasEventsListAuth() bool`

HasEventsListAuth returns a boolean if a field has been set.

### GetEventsAddAuth

`func (o *PVFileVersion) GetEventsAddAuth() bool`

GetEventsAddAuth returns the EventsAddAuth field if non-nil, zero value otherwise.

### GetEventsAddAuthOk

`func (o *PVFileVersion) GetEventsAddAuthOk() (*bool, bool)`

GetEventsAddAuthOk returns a tuple with the EventsAddAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsAddAuth

`func (o *PVFileVersion) SetEventsAddAuth(v bool)`

SetEventsAddAuth sets EventsAddAuth field to given value.

### HasEventsAddAuth

`func (o *PVFileVersion) HasEventsAddAuth() bool`

HasEventsAddAuth returns a boolean if a field has been set.

### GetCreateObjectAuth

`func (o *PVFileVersion) GetCreateObjectAuth() bool`

GetCreateObjectAuth returns the CreateObjectAuth field if non-nil, zero value otherwise.

### GetCreateObjectAuthOk

`func (o *PVFileVersion) GetCreateObjectAuthOk() (*bool, bool)`

GetCreateObjectAuthOk returns a tuple with the CreateObjectAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateObjectAuth

`func (o *PVFileVersion) SetCreateObjectAuth(v bool)`

SetCreateObjectAuth sets CreateObjectAuth field to given value.

### HasCreateObjectAuth

`func (o *PVFileVersion) HasCreateObjectAuth() bool`

HasCreateObjectAuth returns a boolean if a field has been set.

### GetUnlockObjectAuth

`func (o *PVFileVersion) GetUnlockObjectAuth() bool`

GetUnlockObjectAuth returns the UnlockObjectAuth field if non-nil, zero value otherwise.

### GetUnlockObjectAuthOk

`func (o *PVFileVersion) GetUnlockObjectAuthOk() (*bool, bool)`

GetUnlockObjectAuthOk returns a tuple with the UnlockObjectAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockObjectAuth

`func (o *PVFileVersion) SetUnlockObjectAuth(v bool)`

SetUnlockObjectAuth sets UnlockObjectAuth field to given value.

### HasUnlockObjectAuth

`func (o *PVFileVersion) HasUnlockObjectAuth() bool`

HasUnlockObjectAuth returns a boolean if a field has been set.

### GetRenameObjectAuth

`func (o *PVFileVersion) GetRenameObjectAuth() bool`

GetRenameObjectAuth returns the RenameObjectAuth field if non-nil, zero value otherwise.

### GetRenameObjectAuthOk

`func (o *PVFileVersion) GetRenameObjectAuthOk() (*bool, bool)`

GetRenameObjectAuthOk returns a tuple with the RenameObjectAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameObjectAuth

`func (o *PVFileVersion) SetRenameObjectAuth(v bool)`

SetRenameObjectAuth sets RenameObjectAuth field to given value.

### HasRenameObjectAuth

`func (o *PVFileVersion) HasRenameObjectAuth() bool`

HasRenameObjectAuth returns a boolean if a field has been set.

### GetAllAuth

`func (o *PVFileVersion) GetAllAuth() bool`

GetAllAuth returns the AllAuth field if non-nil, zero value otherwise.

### GetAllAuthOk

`func (o *PVFileVersion) GetAllAuthOk() (*bool, bool)`

GetAllAuthOk returns a tuple with the AllAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAuth

`func (o *PVFileVersion) SetAllAuth(v bool)`

SetAllAuth sets AllAuth field to given value.

### HasAllAuth

`func (o *PVFileVersion) HasAllAuth() bool`

HasAllAuth returns a boolean if a field has been set.

### GetNeedConfirmStatus

`func (o *PVFileVersion) GetNeedConfirmStatus() int64`

GetNeedConfirmStatus returns the NeedConfirmStatus field if non-nil, zero value otherwise.

### GetNeedConfirmStatusOk

`func (o *PVFileVersion) GetNeedConfirmStatusOk() (*int64, bool)`

GetNeedConfirmStatusOk returns a tuple with the NeedConfirmStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedConfirmStatus

`func (o *PVFileVersion) SetNeedConfirmStatus(v int64)`

SetNeedConfirmStatus sets NeedConfirmStatus field to given value.

### HasNeedConfirmStatus

`func (o *PVFileVersion) HasNeedConfirmStatus() bool`

HasNeedConfirmStatus returns a boolean if a field has been set.

### GetIsTempVersion

`func (o *PVFileVersion) GetIsTempVersion() bool`

GetIsTempVersion returns the IsTempVersion field if non-nil, zero value otherwise.

### GetIsTempVersionOk

`func (o *PVFileVersion) GetIsTempVersionOk() (*bool, bool)`

GetIsTempVersionOk returns a tuple with the IsTempVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTempVersion

`func (o *PVFileVersion) SetIsTempVersion(v bool)`

SetIsTempVersion sets IsTempVersion field to given value.

### HasIsTempVersion

`func (o *PVFileVersion) HasIsTempVersion() bool`

HasIsTempVersion returns a boolean if a field has been set.

### GetCategoryModificationDate

`func (o *PVFileVersion) GetCategoryModificationDate() time.Time`

GetCategoryModificationDate returns the CategoryModificationDate field if non-nil, zero value otherwise.

### GetCategoryModificationDateOk

`func (o *PVFileVersion) GetCategoryModificationDateOk() (*time.Time, bool)`

GetCategoryModificationDateOk returns a tuple with the CategoryModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryModificationDate

`func (o *PVFileVersion) SetCategoryModificationDate(v time.Time)`

SetCategoryModificationDate sets CategoryModificationDate field to given value.

### HasCategoryModificationDate

`func (o *PVFileVersion) HasCategoryModificationDate() bool`

HasCategoryModificationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


