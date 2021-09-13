# ConfirmRejectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The confirmer&#39;s reason for confirming\\rejecting this request. | [optional] 

## Methods

### NewConfirmRejectData

`func NewConfirmRejectData() *ConfirmRejectData`

NewConfirmRejectData instantiates a new ConfirmRejectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmRejectDataWithDefaults

`func NewConfirmRejectDataWithDefaults() *ConfirmRejectData`

NewConfirmRejectDataWithDefaults instantiates a new ConfirmRejectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ConfirmRejectData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConfirmRejectData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConfirmRejectData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConfirmRejectData) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


