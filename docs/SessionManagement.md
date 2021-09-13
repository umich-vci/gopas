# SessionManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirePrivilegedSessionMonitoringAndIsolation** | Pointer to **bool** | Indicates whether the Require privileged session monitoring and isolation rule is active for this platform. valid values: true\\false | [optional] 
**RecordAndSaveSessionActivity** | Pointer to **bool** | Indicates whether the Record and save session activity rule is active for this platform. valid values: true\\false | [optional] 
**PSMServerID** | Pointer to **string** | The unique ID of a PSM Server | [optional] 

## Methods

### NewSessionManagement

`func NewSessionManagement() *SessionManagement`

NewSessionManagement instantiates a new SessionManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionManagementWithDefaults

`func NewSessionManagementWithDefaults() *SessionManagement`

NewSessionManagementWithDefaults instantiates a new SessionManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirePrivilegedSessionMonitoringAndIsolation

`func (o *SessionManagement) GetRequirePrivilegedSessionMonitoringAndIsolation() bool`

GetRequirePrivilegedSessionMonitoringAndIsolation returns the RequirePrivilegedSessionMonitoringAndIsolation field if non-nil, zero value otherwise.

### GetRequirePrivilegedSessionMonitoringAndIsolationOk

`func (o *SessionManagement) GetRequirePrivilegedSessionMonitoringAndIsolationOk() (*bool, bool)`

GetRequirePrivilegedSessionMonitoringAndIsolationOk returns a tuple with the RequirePrivilegedSessionMonitoringAndIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePrivilegedSessionMonitoringAndIsolation

`func (o *SessionManagement) SetRequirePrivilegedSessionMonitoringAndIsolation(v bool)`

SetRequirePrivilegedSessionMonitoringAndIsolation sets RequirePrivilegedSessionMonitoringAndIsolation field to given value.

### HasRequirePrivilegedSessionMonitoringAndIsolation

`func (o *SessionManagement) HasRequirePrivilegedSessionMonitoringAndIsolation() bool`

HasRequirePrivilegedSessionMonitoringAndIsolation returns a boolean if a field has been set.

### GetRecordAndSaveSessionActivity

`func (o *SessionManagement) GetRecordAndSaveSessionActivity() bool`

GetRecordAndSaveSessionActivity returns the RecordAndSaveSessionActivity field if non-nil, zero value otherwise.

### GetRecordAndSaveSessionActivityOk

`func (o *SessionManagement) GetRecordAndSaveSessionActivityOk() (*bool, bool)`

GetRecordAndSaveSessionActivityOk returns a tuple with the RecordAndSaveSessionActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAndSaveSessionActivity

`func (o *SessionManagement) SetRecordAndSaveSessionActivity(v bool)`

SetRecordAndSaveSessionActivity sets RecordAndSaveSessionActivity field to given value.

### HasRecordAndSaveSessionActivity

`func (o *SessionManagement) HasRecordAndSaveSessionActivity() bool`

HasRecordAndSaveSessionActivity returns a boolean if a field has been set.

### GetPSMServerID

`func (o *SessionManagement) GetPSMServerID() string`

GetPSMServerID returns the PSMServerID field if non-nil, zero value otherwise.

### GetPSMServerIDOk

`func (o *SessionManagement) GetPSMServerIDOk() (*string, bool)`

GetPSMServerIDOk returns a tuple with the PSMServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPSMServerID

`func (o *SessionManagement) SetPSMServerID(v string)`

SetPSMServerID sets PSMServerID field to given value.

### HasPSMServerID

`func (o *SessionManagement) HasPSMServerID() bool`

HasPSMServerID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


