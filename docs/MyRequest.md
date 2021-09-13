# MyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestID** | Pointer to **string** |  | [optional] 
**SafeName** | Pointer to **string** |  | [optional] 
**RequestorUserName** | Pointer to **string** |  | [optional] 
**RequestorReason** | Pointer to **string** |  | [optional] 
**UserReason** | Pointer to **string** |  | [optional] 
**CreationDate** | Pointer to **int64** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **int64** |  | [optional] 
**OperationType** | Pointer to **int32** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**ConfirmationsLeft** | Pointer to **int64** |  | [optional] 
**AccessFrom** | Pointer to **int64** |  | [optional] 
**AccessTo** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StatusTitle** | Pointer to **string** |  | [optional] 
**InvalidRequestReason** | Pointer to **int32** |  | [optional] 
**CurrentConfirmationLevel** | Pointer to **int32** |  | [optional] 
**RequiredConfirmersCountLevel2** | Pointer to **int32** |  | [optional] 
**TicketingSystemProperties** | Pointer to [**TicketingSystem**](TicketingSystem.md) |  | [optional] 
**AdditionalInfo** | Pointer to **map[string]string** |  | [optional] 
**AccountDetails** | Pointer to [**Account**](Account.md) |  | [optional] 
**Confirmers** | Pointer to [**[]ConfirmerStatus**](ConfirmerStatus.md) |  | [optional] 

## Methods

### NewMyRequest

`func NewMyRequest() *MyRequest`

NewMyRequest instantiates a new MyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyRequestWithDefaults

`func NewMyRequestWithDefaults() *MyRequest`

NewMyRequestWithDefaults instantiates a new MyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestID

`func (o *MyRequest) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *MyRequest) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *MyRequest) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *MyRequest) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetSafeName

`func (o *MyRequest) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *MyRequest) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *MyRequest) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *MyRequest) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetRequestorUserName

`func (o *MyRequest) GetRequestorUserName() string`

GetRequestorUserName returns the RequestorUserName field if non-nil, zero value otherwise.

### GetRequestorUserNameOk

`func (o *MyRequest) GetRequestorUserNameOk() (*string, bool)`

GetRequestorUserNameOk returns a tuple with the RequestorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorUserName

`func (o *MyRequest) SetRequestorUserName(v string)`

SetRequestorUserName sets RequestorUserName field to given value.

### HasRequestorUserName

`func (o *MyRequest) HasRequestorUserName() bool`

HasRequestorUserName returns a boolean if a field has been set.

### GetRequestorReason

`func (o *MyRequest) GetRequestorReason() string`

GetRequestorReason returns the RequestorReason field if non-nil, zero value otherwise.

### GetRequestorReasonOk

`func (o *MyRequest) GetRequestorReasonOk() (*string, bool)`

GetRequestorReasonOk returns a tuple with the RequestorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorReason

`func (o *MyRequest) SetRequestorReason(v string)`

SetRequestorReason sets RequestorReason field to given value.

### HasRequestorReason

`func (o *MyRequest) HasRequestorReason() bool`

HasRequestorReason returns a boolean if a field has been set.

### GetUserReason

`func (o *MyRequest) GetUserReason() string`

GetUserReason returns the UserReason field if non-nil, zero value otherwise.

### GetUserReasonOk

`func (o *MyRequest) GetUserReasonOk() (*string, bool)`

GetUserReasonOk returns a tuple with the UserReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReason

`func (o *MyRequest) SetUserReason(v string)`

SetUserReason sets UserReason field to given value.

### HasUserReason

`func (o *MyRequest) HasUserReason() bool`

HasUserReason returns a boolean if a field has been set.

### GetCreationDate

`func (o *MyRequest) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *MyRequest) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *MyRequest) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *MyRequest) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetOperation

`func (o *MyRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *MyRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *MyRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *MyRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetExpirationDate

`func (o *MyRequest) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *MyRequest) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *MyRequest) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *MyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetOperationType

`func (o *MyRequest) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *MyRequest) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *MyRequest) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *MyRequest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetAccessType

`func (o *MyRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MyRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MyRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MyRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetConfirmationsLeft

`func (o *MyRequest) GetConfirmationsLeft() int64`

GetConfirmationsLeft returns the ConfirmationsLeft field if non-nil, zero value otherwise.

### GetConfirmationsLeftOk

`func (o *MyRequest) GetConfirmationsLeftOk() (*int64, bool)`

GetConfirmationsLeftOk returns a tuple with the ConfirmationsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsLeft

`func (o *MyRequest) SetConfirmationsLeft(v int64)`

SetConfirmationsLeft sets ConfirmationsLeft field to given value.

### HasConfirmationsLeft

`func (o *MyRequest) HasConfirmationsLeft() bool`

HasConfirmationsLeft returns a boolean if a field has been set.

### GetAccessFrom

`func (o *MyRequest) GetAccessFrom() int64`

GetAccessFrom returns the AccessFrom field if non-nil, zero value otherwise.

### GetAccessFromOk

`func (o *MyRequest) GetAccessFromOk() (*int64, bool)`

GetAccessFromOk returns a tuple with the AccessFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFrom

`func (o *MyRequest) SetAccessFrom(v int64)`

SetAccessFrom sets AccessFrom field to given value.

### HasAccessFrom

`func (o *MyRequest) HasAccessFrom() bool`

HasAccessFrom returns a boolean if a field has been set.

### GetAccessTo

`func (o *MyRequest) GetAccessTo() int64`

GetAccessTo returns the AccessTo field if non-nil, zero value otherwise.

### GetAccessToOk

`func (o *MyRequest) GetAccessToOk() (*int64, bool)`

GetAccessToOk returns a tuple with the AccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTo

`func (o *MyRequest) SetAccessTo(v int64)`

SetAccessTo sets AccessTo field to given value.

### HasAccessTo

`func (o *MyRequest) HasAccessTo() bool`

HasAccessTo returns a boolean if a field has been set.

### GetStatus

`func (o *MyRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTitle

`func (o *MyRequest) GetStatusTitle() string`

GetStatusTitle returns the StatusTitle field if non-nil, zero value otherwise.

### GetStatusTitleOk

`func (o *MyRequest) GetStatusTitleOk() (*string, bool)`

GetStatusTitleOk returns a tuple with the StatusTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTitle

`func (o *MyRequest) SetStatusTitle(v string)`

SetStatusTitle sets StatusTitle field to given value.

### HasStatusTitle

`func (o *MyRequest) HasStatusTitle() bool`

HasStatusTitle returns a boolean if a field has been set.

### GetInvalidRequestReason

`func (o *MyRequest) GetInvalidRequestReason() int32`

GetInvalidRequestReason returns the InvalidRequestReason field if non-nil, zero value otherwise.

### GetInvalidRequestReasonOk

`func (o *MyRequest) GetInvalidRequestReasonOk() (*int32, bool)`

GetInvalidRequestReasonOk returns a tuple with the InvalidRequestReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRequestReason

`func (o *MyRequest) SetInvalidRequestReason(v int32)`

SetInvalidRequestReason sets InvalidRequestReason field to given value.

### HasInvalidRequestReason

`func (o *MyRequest) HasInvalidRequestReason() bool`

HasInvalidRequestReason returns a boolean if a field has been set.

### GetCurrentConfirmationLevel

`func (o *MyRequest) GetCurrentConfirmationLevel() int32`

GetCurrentConfirmationLevel returns the CurrentConfirmationLevel field if non-nil, zero value otherwise.

### GetCurrentConfirmationLevelOk

`func (o *MyRequest) GetCurrentConfirmationLevelOk() (*int32, bool)`

GetCurrentConfirmationLevelOk returns a tuple with the CurrentConfirmationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmationLevel

`func (o *MyRequest) SetCurrentConfirmationLevel(v int32)`

SetCurrentConfirmationLevel sets CurrentConfirmationLevel field to given value.

### HasCurrentConfirmationLevel

`func (o *MyRequest) HasCurrentConfirmationLevel() bool`

HasCurrentConfirmationLevel returns a boolean if a field has been set.

### GetRequiredConfirmersCountLevel2

`func (o *MyRequest) GetRequiredConfirmersCountLevel2() int32`

GetRequiredConfirmersCountLevel2 returns the RequiredConfirmersCountLevel2 field if non-nil, zero value otherwise.

### GetRequiredConfirmersCountLevel2Ok

`func (o *MyRequest) GetRequiredConfirmersCountLevel2Ok() (*int32, bool)`

GetRequiredConfirmersCountLevel2Ok returns a tuple with the RequiredConfirmersCountLevel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfirmersCountLevel2

`func (o *MyRequest) SetRequiredConfirmersCountLevel2(v int32)`

SetRequiredConfirmersCountLevel2 sets RequiredConfirmersCountLevel2 field to given value.

### HasRequiredConfirmersCountLevel2

`func (o *MyRequest) HasRequiredConfirmersCountLevel2() bool`

HasRequiredConfirmersCountLevel2 returns a boolean if a field has been set.

### GetTicketingSystemProperties

`func (o *MyRequest) GetTicketingSystemProperties() TicketingSystem`

GetTicketingSystemProperties returns the TicketingSystemProperties field if non-nil, zero value otherwise.

### GetTicketingSystemPropertiesOk

`func (o *MyRequest) GetTicketingSystemPropertiesOk() (*TicketingSystem, bool)`

GetTicketingSystemPropertiesOk returns a tuple with the TicketingSystemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystemProperties

`func (o *MyRequest) SetTicketingSystemProperties(v TicketingSystem)`

SetTicketingSystemProperties sets TicketingSystemProperties field to given value.

### HasTicketingSystemProperties

`func (o *MyRequest) HasTicketingSystemProperties() bool`

HasTicketingSystemProperties returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *MyRequest) GetAdditionalInfo() map[string]string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *MyRequest) GetAdditionalInfoOk() (*map[string]string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *MyRequest) SetAdditionalInfo(v map[string]string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *MyRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAccountDetails

`func (o *MyRequest) GetAccountDetails() Account`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *MyRequest) GetAccountDetailsOk() (*Account, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *MyRequest) SetAccountDetails(v Account)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *MyRequest) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### GetConfirmers

`func (o *MyRequest) GetConfirmers() []ConfirmerStatus`

GetConfirmers returns the Confirmers field if non-nil, zero value otherwise.

### GetConfirmersOk

`func (o *MyRequest) GetConfirmersOk() (*[]ConfirmerStatus, bool)`

GetConfirmersOk returns a tuple with the Confirmers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmers

`func (o *MyRequest) SetConfirmers(v []ConfirmerStatus)`

SetConfirmers sets Confirmers field to given value.

### HasConfirmers

`func (o *MyRequest) HasConfirmers() bool`

HasConfirmers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


