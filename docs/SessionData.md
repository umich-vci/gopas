# SessionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSessionData

`func NewSessionData() *SessionData`

NewSessionData instantiates a new SessionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionDataWithDefaults

`func NewSessionDataWithDefaults() *SessionData`

NewSessionDataWithDefaults instantiates a new SessionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionID

`func (o *SessionData) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *SessionData) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *SessionData) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.

### HasSessionID

`func (o *SessionData) HasSessionID() bool`

HasSessionID returns a boolean if a field has been set.

### GetSessionGuid

`func (o *SessionData) GetSessionGuid() string`

GetSessionGuid returns the SessionGuid field if non-nil, zero value otherwise.

### GetSessionGuidOk

`func (o *SessionData) GetSessionGuidOk() (*string, bool)`

GetSessionGuidOk returns a tuple with the SessionGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionGuid

`func (o *SessionData) SetSessionGuid(v string)`

SetSessionGuid sets SessionGuid field to given value.

### HasSessionGuid

`func (o *SessionData) HasSessionGuid() bool`

HasSessionGuid returns a boolean if a field has been set.

### GetSafeName

`func (o *SessionData) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *SessionData) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *SessionData) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *SessionData) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetFolderName

`func (o *SessionData) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *SessionData) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *SessionData) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *SessionData) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetIsLive

`func (o *SessionData) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *SessionData) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *SessionData) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *SessionData) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetFileName

`func (o *SessionData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SessionData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SessionData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SessionData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetStart

`func (o *SessionData) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SessionData) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SessionData) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *SessionData) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SessionData) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SessionData) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SessionData) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SessionData) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDuration

`func (o *SessionData) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SessionData) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SessionData) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SessionData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetUser

`func (o *SessionData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionData) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionData) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRemoteMachine

`func (o *SessionData) GetRemoteMachine() string`

GetRemoteMachine returns the RemoteMachine field if non-nil, zero value otherwise.

### GetRemoteMachineOk

`func (o *SessionData) GetRemoteMachineOk() (*string, bool)`

GetRemoteMachineOk returns a tuple with the RemoteMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMachine

`func (o *SessionData) SetRemoteMachine(v string)`

SetRemoteMachine sets RemoteMachine field to given value.

### HasRemoteMachine

`func (o *SessionData) HasRemoteMachine() bool`

HasRemoteMachine returns a boolean if a field has been set.

### GetProtectionDate

`func (o *SessionData) GetProtectionDate() int64`

GetProtectionDate returns the ProtectionDate field if non-nil, zero value otherwise.

### GetProtectionDateOk

`func (o *SessionData) GetProtectionDateOk() (*int64, bool)`

GetProtectionDateOk returns a tuple with the ProtectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionDate

`func (o *SessionData) SetProtectionDate(v int64)`

SetProtectionDate sets ProtectionDate field to given value.

### HasProtectionDate

`func (o *SessionData) HasProtectionDate() bool`

HasProtectionDate returns a boolean if a field has been set.

### GetProtectedBy

`func (o *SessionData) GetProtectedBy() string`

GetProtectedBy returns the ProtectedBy field if non-nil, zero value otherwise.

### GetProtectedByOk

`func (o *SessionData) GetProtectedByOk() (*string, bool)`

GetProtectedByOk returns a tuple with the ProtectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedBy

`func (o *SessionData) SetProtectedBy(v string)`

SetProtectedBy sets ProtectedBy field to given value.

### HasProtectedBy

`func (o *SessionData) HasProtectedBy() bool`

HasProtectedBy returns a boolean if a field has been set.

### GetProtectionEnabled

`func (o *SessionData) GetProtectionEnabled() bool`

GetProtectionEnabled returns the ProtectionEnabled field if non-nil, zero value otherwise.

### GetProtectionEnabledOk

`func (o *SessionData) GetProtectionEnabledOk() (*bool, bool)`

GetProtectionEnabledOk returns a tuple with the ProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionEnabled

`func (o *SessionData) SetProtectionEnabled(v bool)`

SetProtectionEnabled sets ProtectionEnabled field to given value.

### HasProtectionEnabled

`func (o *SessionData) HasProtectionEnabled() bool`

HasProtectionEnabled returns a boolean if a field has been set.

### GetAccountUsername

`func (o *SessionData) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *SessionData) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *SessionData) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *SessionData) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetAccountPlatformID

`func (o *SessionData) GetAccountPlatformID() string`

GetAccountPlatformID returns the AccountPlatformID field if non-nil, zero value otherwise.

### GetAccountPlatformIDOk

`func (o *SessionData) GetAccountPlatformIDOk() (*string, bool)`

GetAccountPlatformIDOk returns a tuple with the AccountPlatformID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPlatformID

`func (o *SessionData) SetAccountPlatformID(v string)`

SetAccountPlatformID sets AccountPlatformID field to given value.

### HasAccountPlatformID

`func (o *SessionData) HasAccountPlatformID() bool`

HasAccountPlatformID returns a boolean if a field has been set.

### GetAccountAddress

`func (o *SessionData) GetAccountAddress() string`

GetAccountAddress returns the AccountAddress field if non-nil, zero value otherwise.

### GetAccountAddressOk

`func (o *SessionData) GetAccountAddressOk() (*string, bool)`

GetAccountAddressOk returns a tuple with the AccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAddress

`func (o *SessionData) SetAccountAddress(v string)`

SetAccountAddress sets AccountAddress field to given value.

### HasAccountAddress

`func (o *SessionData) HasAccountAddress() bool`

HasAccountAddress returns a boolean if a field has been set.

### GetPIMSuCommand

`func (o *SessionData) GetPIMSuCommand() string`

GetPIMSuCommand returns the PIMSuCommand field if non-nil, zero value otherwise.

### GetPIMSuCommandOk

`func (o *SessionData) GetPIMSuCommandOk() (*string, bool)`

GetPIMSuCommandOk returns a tuple with the PIMSuCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPIMSuCommand

`func (o *SessionData) SetPIMSuCommand(v string)`

SetPIMSuCommand sets PIMSuCommand field to given value.

### HasPIMSuCommand

`func (o *SessionData) HasPIMSuCommand() bool`

HasPIMSuCommand returns a boolean if a field has been set.

### GetPIMSuCWD

`func (o *SessionData) GetPIMSuCWD() string`

GetPIMSuCWD returns the PIMSuCWD field if non-nil, zero value otherwise.

### GetPIMSuCWDOk

`func (o *SessionData) GetPIMSuCWDOk() (*string, bool)`

GetPIMSuCWDOk returns a tuple with the PIMSuCWD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPIMSuCWD

`func (o *SessionData) SetPIMSuCWD(v string)`

SetPIMSuCWD sets PIMSuCWD field to given value.

### HasPIMSuCWD

`func (o *SessionData) HasPIMSuCWD() bool`

HasPIMSuCWD returns a boolean if a field has been set.

### GetConnectionComponentID

`func (o *SessionData) GetConnectionComponentID() string`

GetConnectionComponentID returns the ConnectionComponentID field if non-nil, zero value otherwise.

### GetConnectionComponentIDOk

`func (o *SessionData) GetConnectionComponentIDOk() (*string, bool)`

GetConnectionComponentIDOk returns a tuple with the ConnectionComponentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionComponentID

`func (o *SessionData) SetConnectionComponentID(v string)`

SetConnectionComponentID sets ConnectionComponentID field to given value.

### HasConnectionComponentID

`func (o *SessionData) HasConnectionComponentID() bool`

HasConnectionComponentID returns a boolean if a field has been set.

### GetPSMRecordingEntity

`func (o *SessionData) GetPSMRecordingEntity() string`

GetPSMRecordingEntity returns the PSMRecordingEntity field if non-nil, zero value otherwise.

### GetPSMRecordingEntityOk

`func (o *SessionData) GetPSMRecordingEntityOk() (*string, bool)`

GetPSMRecordingEntityOk returns a tuple with the PSMRecordingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMRecordingEntity

`func (o *SessionData) SetPSMRecordingEntity(v string)`

SetPSMRecordingEntity sets PSMRecordingEntity field to given value.

### HasPSMRecordingEntity

`func (o *SessionData) HasPSMRecordingEntity() bool`

HasPSMRecordingEntity returns a boolean if a field has been set.

### GetTicketID

`func (o *SessionData) GetTicketID() string`

GetTicketID returns the TicketID field if non-nil, zero value otherwise.

### GetTicketIDOk

`func (o *SessionData) GetTicketIDOk() (*string, bool)`

GetTicketIDOk returns a tuple with the TicketID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketID

`func (o *SessionData) SetTicketID(v string)`

SetTicketID sets TicketID field to given value.

### HasTicketID

`func (o *SessionData) HasTicketID() bool`

HasTicketID returns a boolean if a field has been set.

### GetFromIP

`func (o *SessionData) GetFromIP() string`

GetFromIP returns the FromIP field if non-nil, zero value otherwise.

### GetFromIPOk

`func (o *SessionData) GetFromIPOk() (*string, bool)`

GetFromIPOk returns a tuple with the FromIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIP

`func (o *SessionData) SetFromIP(v string)`

SetFromIP sets FromIP field to given value.

### HasFromIP

`func (o *SessionData) HasFromIP() bool`

HasFromIP returns a boolean if a field has been set.

### GetProtocol

`func (o *SessionData) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SessionData) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SessionData) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SessionData) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetClient

`func (o *SessionData) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionData) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionData) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionData) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetRiskScore

`func (o *SessionData) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *SessionData) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *SessionData) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *SessionData) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetSeverity

`func (o *SessionData) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SessionData) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SessionData) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SessionData) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetIncidentDetails

`func (o *SessionData) GetIncidentDetails() PTAData`

GetIncidentDetails returns the IncidentDetails field if non-nil, zero value otherwise.

### GetIncidentDetailsOk

`func (o *SessionData) GetIncidentDetailsOk() (*PTAData, bool)`

GetIncidentDetailsOk returns a tuple with the IncidentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentDetails

`func (o *SessionData) SetIncidentDetails(v PTAData)`

SetIncidentDetails sets IncidentDetails field to given value.

### HasIncidentDetails

`func (o *SessionData) HasIncidentDetails() bool`

HasIncidentDetails returns a boolean if a field has been set.

### GetRawProperties

`func (o *SessionData) GetRawProperties() map[string]string`

GetRawProperties returns the RawProperties field if non-nil, zero value otherwise.

### GetRawPropertiesOk

`func (o *SessionData) GetRawPropertiesOk() (*map[string]string, bool)`

GetRawPropertiesOk returns a tuple with the RawProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawProperties

`func (o *SessionData) SetRawProperties(v map[string]string)`

SetRawProperties sets RawProperties field to given value.

### HasRawProperties

`func (o *SessionData) HasRawProperties() bool`

HasRawProperties returns a boolean if a field has been set.

### GetRecordingFiles

`func (o *SessionData) GetRecordingFiles() []RecordingFile`

GetRecordingFiles returns the RecordingFiles field if non-nil, zero value otherwise.

### GetRecordingFilesOk

`func (o *SessionData) GetRecordingFilesOk() (*[]RecordingFile, bool)`

GetRecordingFilesOk returns a tuple with the RecordingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingFiles

`func (o *SessionData) SetRecordingFiles(v []RecordingFile)`

SetRecordingFiles sets RecordingFiles field to given value.

### HasRecordingFiles

`func (o *SessionData) HasRecordingFiles() bool`

HasRecordingFiles returns a boolean if a field has been set.

### GetRecordedActivities

`func (o *SessionData) GetRecordedActivities() []RecordedCommand`

GetRecordedActivities returns the RecordedActivities field if non-nil, zero value otherwise.

### GetRecordedActivitiesOk

`func (o *SessionData) GetRecordedActivitiesOk() (*[]RecordedCommand, bool)`

GetRecordedActivitiesOk returns a tuple with the RecordedActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedActivities

`func (o *SessionData) SetRecordedActivities(v []RecordedCommand)`

SetRecordedActivities sets RecordedActivities field to given value.

### HasRecordedActivities

`func (o *SessionData) HasRecordedActivities() bool`

HasRecordedActivities returns a boolean if a field has been set.

### GetVideoSize

`func (o *SessionData) GetVideoSize() int64`

GetVideoSize returns the VideoSize field if non-nil, zero value otherwise.

### GetVideoSizeOk

`func (o *SessionData) GetVideoSizeOk() (*int64, bool)`

GetVideoSizeOk returns a tuple with the VideoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSize

`func (o *SessionData) SetVideoSize(v int64)`

SetVideoSize sets VideoSize field to given value.

### HasVideoSize

`func (o *SessionData) HasVideoSize() bool`

HasVideoSize returns a boolean if a field has been set.

### GetTextSize

`func (o *SessionData) GetTextSize() int64`

GetTextSize returns the TextSize field if non-nil, zero value otherwise.

### GetTextSizeOk

`func (o *SessionData) GetTextSizeOk() (*int64, bool)`

GetTextSizeOk returns a tuple with the TextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSize

`func (o *SessionData) SetTextSize(v int64)`

SetTextSize sets TextSize field to given value.

### HasTextSize

`func (o *SessionData) HasTextSize() bool`

HasTextSize returns a boolean if a field has been set.

### GetDetailsUrl

`func (o *SessionData) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *SessionData) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *SessionData) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.

### HasDetailsUrl

`func (o *SessionData) HasDetailsUrl() bool`

HasDetailsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


