# ConfirmerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **int32** |  | [optional] 
**ID** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ActionDate** | Pointer to **int64** |  | [optional] 
**AdditionalDetails** | Pointer to **map[string]string** |  | [optional] 
**Members** | Pointer to [**[]ConfirmerMember**](ConfirmerMember.md) |  | [optional] 

## Methods

### NewConfirmerStatus

`func NewConfirmerStatus() *ConfirmerStatus`

NewConfirmerStatus instantiates a new ConfirmerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmerStatusWithDefaults

`func NewConfirmerStatusWithDefaults() *ConfirmerStatus`

NewConfirmerStatusWithDefaults instantiates a new ConfirmerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConfirmerStatus) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfirmerStatus) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfirmerStatus) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ConfirmerStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetID

`func (o *ConfirmerStatus) GetID() int64`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ConfirmerStatus) GetIDOk() (*int64, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ConfirmerStatus) SetID(v int64)`

SetID sets ID field to given value.

### HasID

`func (o *ConfirmerStatus) HasID() bool`

HasID returns a boolean if a field has been set.

### GetName

`func (o *ConfirmerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfirmerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfirmerStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfirmerStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAction

`func (o *ConfirmerStatus) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ConfirmerStatus) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ConfirmerStatus) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *ConfirmerStatus) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetReason

`func (o *ConfirmerStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConfirmerStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConfirmerStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConfirmerStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetActionDate

`func (o *ConfirmerStatus) GetActionDate() int64`

GetActionDate returns the ActionDate field if non-nil, zero value otherwise.

### GetActionDateOk

`func (o *ConfirmerStatus) GetActionDateOk() (*int64, bool)`

GetActionDateOk returns a tuple with the ActionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDate

`func (o *ConfirmerStatus) SetActionDate(v int64)`

SetActionDate sets ActionDate field to given value.

### HasActionDate

`func (o *ConfirmerStatus) HasActionDate() bool`

HasActionDate returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *ConfirmerStatus) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ConfirmerStatus) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ConfirmerStatus) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ConfirmerStatus) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetMembers

`func (o *ConfirmerStatus) GetMembers() []ConfirmerMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ConfirmerStatus) GetMembersOk() (*[]ConfirmerMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ConfirmerStatus) SetMembers(v []ConfirmerMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ConfirmerStatus) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


