# LiveSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTerminate** | Pointer to **bool** |  | [optional] 
**CanMonitor** | Pointer to **bool** |  | [optional] 
**CanSuspend** | Pointer to **bool** |  | [optional] 
**SessionID** | Pointer to **string** |  | [optional] 
**SessionGuid** | Pointer to **string** |  | [optional] 
**SafeName** | Pointer to **string** |  | [optional] 
**FolderName** | Pointer to **string** |  | [optional] 
**IsLive** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**End** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**RemoteMachine** | Pointer to **string** |  | [optional] 
**ProtectionDate** | Pointer to **int64** |  | [optional] 
**ProtectedBy** | Pointer to **string** |  | [optional] 
**ProtectionEnabled** | Pointer to **bool** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**AccountPlatformID** | Pointer to **string** |  | [optional] 
**AccountAddress** | Pointer to **string** |  | [optional] 
**PIMSuCommand** | Pointer to **string** |  | [optional] 
**PIMSuCWD** | Pointer to **string** |  | [optional] 
**ConnectionComponentID** | Pointer to **string** |  | [optional] 
**PSMRecordingEntity** | Pointer to **string** |  | [optional] 
**TicketID** | Pointer to **string** |  | [optional] 
**FromIP** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**IncidentDetails** | Pointer to [**PTAData**](PTAData.md) |  | [optional] 
**RawProperties** | Pointer to **map[string]string** |  | [optional] 
**RecordingFiles** | Pointer to [**[]RecordingFile**](RecordingFile.md) |  | [optional] 
**RecordedActivities** | Pointer to [**[]RecordedCommand**](RecordedCommand.md) |  | [optional] 
**VideoSize** | Pointer to **int64** |  | [optional] 
**TextSize** | Pointer to **int64** |  | [optional] 
**DetailsUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewLiveSession

`func NewLiveSession() *LiveSession`

NewLiveSession instantiates a new LiveSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveSessionWithDefaults

`func NewLiveSessionWithDefaults() *LiveSession`

NewLiveSessionWithDefaults instantiates a new LiveSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTerminate

`func (o *LiveSession) GetCanTerminate() bool`

GetCanTerminate returns the CanTerminate field if non-nil, zero value otherwise.

### GetCanTerminateOk

`func (o *LiveSession) GetCanTerminateOk() (*bool, bool)`

GetCanTerminateOk returns a tuple with the CanTerminate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTerminate

`func (o *LiveSession) SetCanTerminate(v bool)`

SetCanTerminate sets CanTerminate field to given value.

### HasCanTerminate

`func (o *LiveSession) HasCanTerminate() bool`

HasCanTerminate returns a boolean if a field has been set.

### GetCanMonitor

`func (o *LiveSession) GetCanMonitor() bool`

GetCanMonitor returns the CanMonitor field if non-nil, zero value otherwise.

### GetCanMonitorOk

`func (o *LiveSession) GetCanMonitorOk() (*bool, bool)`

GetCanMonitorOk returns a tuple with the CanMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMonitor

`func (o *LiveSession) SetCanMonitor(v bool)`

SetCanMonitor sets CanMonitor field to given value.

### HasCanMonitor

`func (o *LiveSession) HasCanMonitor() bool`

HasCanMonitor returns a boolean if a field has been set.

### GetCanSuspend

`func (o *LiveSession) GetCanSuspend() bool`

GetCanSuspend returns the CanSuspend field if non-nil, zero value otherwise.

### GetCanSuspendOk

`func (o *LiveSession) GetCanSuspendOk() (*bool, bool)`

GetCanSuspendOk returns a tuple with the CanSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSuspend

`func (o *LiveSession) SetCanSuspend(v bool)`

SetCanSuspend sets CanSuspend field to given value.

### HasCanSuspend

`func (o *LiveSession) HasCanSuspend() bool`

HasCanSuspend returns a boolean if a field has been set.

### GetSessionID

`func (o *LiveSession) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *LiveSession) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *LiveSession) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.

### HasSessionID

`func (o *LiveSession) HasSessionID() bool`

HasSessionID returns a boolean if a field has been set.

### GetSessionGuid

`func (o *LiveSession) GetSessionGuid() string`

GetSessionGuid returns the SessionGuid field if non-nil, zero value otherwise.

### GetSessionGuidOk

`func (o *LiveSession) GetSessionGuidOk() (*string, bool)`

GetSessionGuidOk returns a tuple with the SessionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionGuid

`func (o *LiveSession) SetSessionGuid(v string)`

SetSessionGuid sets SessionGuid field to given value.

### HasSessionGuid

`func (o *LiveSession) HasSessionGuid() bool`

HasSessionGuid returns a boolean if a field has been set.

### GetSafeName

`func (o *LiveSession) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *LiveSession) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *LiveSession) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *LiveSession) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetFolderName

`func (o *LiveSession) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *LiveSession) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *LiveSession) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *LiveSession) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetIsLive

`func (o *LiveSession) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *LiveSession) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *LiveSession) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *LiveSession) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetFileName

`func (o *LiveSession) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LiveSession) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LiveSession) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LiveSession) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetStart

`func (o *LiveSession) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LiveSession) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LiveSession) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *LiveSession) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *LiveSession) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LiveSession) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LiveSession) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LiveSession) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDuration

`func (o *LiveSession) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LiveSession) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LiveSession) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LiveSession) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetUser

`func (o *LiveSession) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LiveSession) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LiveSession) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LiveSession) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRemoteMachine

`func (o *LiveSession) GetRemoteMachine() string`

GetRemoteMachine returns the RemoteMachine field if non-nil, zero value otherwise.

### GetRemoteMachineOk

`func (o *LiveSession) GetRemoteMachineOk() (*string, bool)`

GetRemoteMachineOk returns a tuple with the RemoteMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMachine

`func (o *LiveSession) SetRemoteMachine(v string)`

SetRemoteMachine sets RemoteMachine field to given value.

### HasRemoteMachine

`func (o *LiveSession) HasRemoteMachine() bool`

HasRemoteMachine returns a boolean if a field has been set.

### GetProtectionDate

`func (o *LiveSession) GetProtectionDate() int64`

GetProtectionDate returns the ProtectionDate field if non-nil, zero value otherwise.

### GetProtectionDateOk

`func (o *LiveSession) GetProtectionDateOk() (*int64, bool)`

GetProtectionDateOk returns a tuple with the ProtectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDate

`func (o *LiveSession) SetProtectionDate(v int64)`

SetProtectionDate sets ProtectionDate field to given value.

### HasProtectionDate

`func (o *LiveSession) HasProtectionDate() bool`

HasProtectionDate returns a boolean if a field has been set.

### GetProtectedBy

`func (o *LiveSession) GetProtectedBy() string`

GetProtectedBy returns the ProtectedBy field if non-nil, zero value otherwise.

### GetProtectedByOk

`func (o *LiveSession) GetProtectedByOk() (*string, bool)`

GetProtectedByOk returns a tuple with the ProtectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedBy

`func (o *LiveSession) SetProtectedBy(v string)`

SetProtectedBy sets ProtectedBy field to given value.

### HasProtectedBy

`func (o *LiveSession) HasProtectedBy() bool`

HasProtectedBy returns a boolean if a field has been set.

### GetProtectionEnabled

`func (o *LiveSession) GetProtectionEnabled() bool`

GetProtectionEnabled returns the ProtectionEnabled field if non-nil, zero value otherwise.

### GetProtectionEnabledOk

`func (o *LiveSession) GetProtectionEnabledOk() (*bool, bool)`

GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnabled

`func (o *LiveSession) SetProtectionEnabled(v bool)`

SetProtectionEnabled sets ProtectionEnabled field to given value.

### HasProtectionEnabled

`func (o *LiveSession) HasProtectionEnabled() bool`

HasProtectionEnabled returns a boolean if a field has been set.

### GetAccountUsername

`func (o *LiveSession) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *LiveSession) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *LiveSession) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *LiveSession) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetAccountPlatformID

`func (o *LiveSession) GetAccountPlatformID() string`

GetAccountPlatformID returns the AccountPlatformID field if non-nil, zero value otherwise.

### GetAccountPlatformIDOk

`func (o *LiveSession) GetAccountPlatformIDOk() (*string, bool)`

GetAccountPlatformIDOk returns a tuple with the AccountPlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPlatformID

`func (o *LiveSession) SetAccountPlatformID(v string)`

SetAccountPlatformID sets AccountPlatformID field to given value.

### HasAccountPlatformID

`func (o *LiveSession) HasAccountPlatformID() bool`

HasAccountPlatformID returns a boolean if a field has been set.

### GetAccountAddress

`func (o *LiveSession) GetAccountAddress() string`

GetAccountAddress returns the AccountAddress field if non-nil, zero value otherwise.

### GetAccountAddressOk

`func (o *LiveSession) GetAccountAddressOk() (*string, bool)`

GetAccountAddressOk returns a tuple with the AccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAddress

`func (o *LiveSession) SetAccountAddress(v string)`

SetAccountAddress sets AccountAddress field to given value.

### HasAccountAddress

`func (o *LiveSession) HasAccountAddress() bool`

HasAccountAddress returns a boolean if a field has been set.

### GetPIMSuCommand

`func (o *LiveSession) GetPIMSuCommand() string`

GetPIMSuCommand returns the PIMSuCommand field if non-nil, zero value otherwise.

### GetPIMSuCommandOk

`func (o *LiveSession) GetPIMSuCommandOk() (*string, bool)`

GetPIMSuCommandOk returns a tuple with the PIMSuCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPIMSuCommand

`func (o *LiveSession) SetPIMSuCommand(v string)`

SetPIMSuCommand sets PIMSuCommand field to given value.

### HasPIMSuCommand

`func (o *LiveSession) HasPIMSuCommand() bool`

HasPIMSuCommand returns a boolean if a field has been set.

### GetPIMSuCWD

`func (o *LiveSession) GetPIMSuCWD() string`

GetPIMSuCWD returns the PIMSuCWD field if non-nil, zero value otherwise.

### GetPIMSuCWDOk

`func (o *LiveSession) GetPIMSuCWDOk() (*string, bool)`

GetPIMSuCWDOk returns a tuple with the PIMSuCWD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPIMSuCWD

`func (o *LiveSession) SetPIMSuCWD(v string)`

SetPIMSuCWD sets PIMSuCWD field to given value.

### HasPIMSuCWD

`func (o *LiveSession) HasPIMSuCWD() bool`

HasPIMSuCWD returns a boolean if a field has been set.

### GetConnectionComponentID

`func (o *LiveSession) GetConnectionComponentID() string`

GetConnectionComponentID returns the ConnectionComponentID field if non-nil, zero value otherwise.

### GetConnectionComponentIDOk

`func (o *LiveSession) GetConnectionComponentIDOk() (*string, bool)`

GetConnectionComponentIDOk returns a tuple with the ConnectionComponentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionComponentID

`func (o *LiveSession) SetConnectionComponentID(v string)`

SetConnectionComponentID sets ConnectionComponentID field to given value.

### HasConnectionComponentID

`func (o *LiveSession) HasConnectionComponentID() bool`

HasConnectionComponentID returns a boolean if a field has been set.

### GetPSMRecordingEntity

`func (o *LiveSession) GetPSMRecordingEntity() string`

GetPSMRecordingEntity returns the PSMRecordingEntity field if non-nil, zero value otherwise.

### GetPSMRecordingEntityOk

`func (o *LiveSession) GetPSMRecordingEntityOk() (*string, bool)`

GetPSMRecordingEntityOk returns a tuple with the PSMRecordingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMRecordingEntity

`func (o *LiveSession) SetPSMRecordingEntity(v string)`

SetPSMRecordingEntity sets PSMRecordingEntity field to given value.

### HasPSMRecordingEntity

`func (o *LiveSession) HasPSMRecordingEntity() bool`

HasPSMRecordingEntity returns a boolean if a field has been set.

### GetTicketID

`func (o *LiveSession) GetTicketID() string`

GetTicketID returns the TicketID field if non-nil, zero value otherwise.

### GetTicketIDOk

`func (o *LiveSession) GetTicketIDOk() (*string, bool)`

GetTicketIDOk returns a tuple with the TicketID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketID

`func (o *LiveSession) SetTicketID(v string)`

SetTicketID sets TicketID field to given value.

### HasTicketID

`func (o *LiveSession) HasTicketID() bool`

HasTicketID returns a boolean if a field has been set.

### GetFromIP

`func (o *LiveSession) GetFromIP() string`

GetFromIP returns the FromIP field if non-nil, zero value otherwise.

### GetFromIPOk

`func (o *LiveSession) GetFromIPOk() (*string, bool)`

GetFromIPOk returns a tuple with the FromIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIP

`func (o *LiveSession) SetFromIP(v string)`

SetFromIP sets FromIP field to given value.

### HasFromIP

`func (o *LiveSession) HasFromIP() bool`

HasFromIP returns a boolean if a field has been set.

### GetProtocol

`func (o *LiveSession) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LiveSession) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LiveSession) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LiveSession) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetClient

`func (o *LiveSession) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *LiveSession) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *LiveSession) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *LiveSession) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetRiskScore

`func (o *LiveSession) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *LiveSession) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *LiveSession) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *LiveSession) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetSeverity

`func (o *LiveSession) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *LiveSession) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *LiveSession) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *LiveSession) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetIncidentDetails

`func (o *LiveSession) GetIncidentDetails() PTAData`

GetIncidentDetails returns the IncidentDetails field if non-nil, zero value otherwise.

### GetIncidentDetailsOk

`func (o *LiveSession) GetIncidentDetailsOk() (*PTAData, bool)`

GetIncidentDetailsOk returns a tuple with the IncidentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentDetails

`func (o *LiveSession) SetIncidentDetails(v PTAData)`

SetIncidentDetails sets IncidentDetails field to given value.

### HasIncidentDetails

`func (o *LiveSession) HasIncidentDetails() bool`

HasIncidentDetails returns a boolean if a field has been set.

### GetRawProperties

`func (o *LiveSession) GetRawProperties() map[string]string`

GetRawProperties returns the RawProperties field if non-nil, zero value otherwise.

### GetRawPropertiesOk

`func (o *LiveSession) GetRawPropertiesOk() (*map[string]string, bool)`

GetRawPropertiesOk returns a tuple with the RawProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawProperties

`func (o *LiveSession) SetRawProperties(v map[string]string)`

SetRawProperties sets RawProperties field to given value.

### HasRawProperties

`func (o *LiveSession) HasRawProperties() bool`

HasRawProperties returns a boolean if a field has been set.

### GetRecordingFiles

`func (o *LiveSession) GetRecordingFiles() []RecordingFile`

GetRecordingFiles returns the RecordingFiles field if non-nil, zero value otherwise.

### GetRecordingFilesOk

`func (o *LiveSession) GetRecordingFilesOk() (*[]RecordingFile, bool)`

GetRecordingFilesOk returns a tuple with the RecordingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingFiles

`func (o *LiveSession) SetRecordingFiles(v []RecordingFile)`

SetRecordingFiles sets RecordingFiles field to given value.

### HasRecordingFiles

`func (o *LiveSession) HasRecordingFiles() bool`

HasRecordingFiles returns a boolean if a field has been set.

### GetRecordedActivities

`func (o *LiveSession) GetRecordedActivities() []RecordedCommand`

GetRecordedActivities returns the RecordedActivities field if non-nil, zero value otherwise.

### GetRecordedActivitiesOk

`func (o *LiveSession) GetRecordedActivitiesOk() (*[]RecordedCommand, bool)`

GetRecordedActivitiesOk returns a tuple with the RecordedActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedActivities

`func (o *LiveSession) SetRecordedActivities(v []RecordedCommand)`

SetRecordedActivities sets RecordedActivities field to given value.

### HasRecordedActivities

`func (o *LiveSession) HasRecordedActivities() bool`

HasRecordedActivities returns a boolean if a field has been set.

### GetVideoSize

`func (o *LiveSession) GetVideoSize() int64`

GetVideoSize returns the VideoSize field if non-nil, zero value otherwise.

### GetVideoSizeOk

`func (o *LiveSession) GetVideoSizeOk() (*int64, bool)`

GetVideoSizeOk returns a tuple with the VideoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSize

`func (o *LiveSession) SetVideoSize(v int64)`

SetVideoSize sets VideoSize field to given value.

### HasVideoSize

`func (o *LiveSession) HasVideoSize() bool`

HasVideoSize returns a boolean if a field has been set.

### GetTextSize

`func (o *LiveSession) GetTextSize() int64`

GetTextSize returns the TextSize field if non-nil, zero value otherwise.

### GetTextSizeOk

`func (o *LiveSession) GetTextSizeOk() (*int64, bool)`

GetTextSizeOk returns a tuple with the TextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSize

`func (o *LiveSession) SetTextSize(v int64)`

SetTextSize sets TextSize field to given value.

### HasTextSize

`func (o *LiveSession) HasTextSize() bool`

HasTextSize returns a boolean if a field has been set.

### GetDetailsUrl

`func (o *LiveSession) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *LiveSession) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *LiveSession) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *LiveSession) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


