# PTAData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incident** | Pointer to [**Incident**](Incident.md) |  | [optional] 
**Sessions** | Pointer to [**[]Session**](Session.md) |  | [optional] 

## Methods

### NewPTAData

`func NewPTAData() *PTAData`

NewPTAData instantiates a new PTAData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTADataWithDefaults

`func NewPTADataWithDefaults() *PTAData`

NewPTADataWithDefaults instantiates a new PTAData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncident

`func (o *PTAData) GetIncident() Incident`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PTAData) GetIncidentOk() (*Incident, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PTAData) SetIncident(v Incident)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PTAData) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetSessions

`func (o *PTAData) GetSessions() []Session`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *PTAData) GetSessionsOk() (*[]Session, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *PTAData) SetSessions(v []Session)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *PTAData) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


