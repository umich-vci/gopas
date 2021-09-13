# Request

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

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestID

`func (o *Request) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *Request) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *Request) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *Request) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetSafeName

`func (o *Request) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *Request) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *Request) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *Request) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetRequestorUserName

`func (o *Request) GetRequestorUserName() string`

GetRequestorUserName returns the RequestorUserName field if non-nil, zero value otherwise.

### GetRequestorUserNameOk

`func (o *Request) GetRequestorUserNameOk() (*string, bool)`

GetRequestorUserNameOk returns a tuple with the RequestorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorUserName

`func (o *Request) SetRequestorUserName(v string)`

SetRequestorUserName sets RequestorUserName field to given value.

### HasRequestorUserName

`func (o *Request) HasRequestorUserName() bool`

HasRequestorUserName returns a boolean if a field has been set.

### GetRequestorReason

`func (o *Request) GetRequestorReason() string`

GetRequestorReason returns the RequestorReason field if non-nil, zero value otherwise.

### GetRequestorReasonOk

`func (o *Request) GetRequestorReasonOk() (*string, bool)`

GetRequestorReasonOk returns a tuple with the RequestorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorReason

`func (o *Request) SetRequestorReason(v string)`

SetRequestorReason sets RequestorReason field to given value.

### HasRequestorReason

`func (o *Request) HasRequestorReason() bool`

HasRequestorReason returns a boolean if a field has been set.

### GetUserReason

`func (o *Request) GetUserReason() string`

GetUserReason returns the UserReason field if non-nil, zero value otherwise.

### GetUserReasonOk

`func (o *Request) GetUserReasonOk() (*string, bool)`

GetUserReasonOk returns a tuple with the UserReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReason

`func (o *Request) SetUserReason(v string)`

SetUserReason sets UserReason field to given value.

### HasUserReason

`func (o *Request) HasUserReason() bool`

HasUserReason returns a boolean if a field has been set.

### GetCreationDate

`func (o *Request) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Request) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Request) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Request) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetOperation

`func (o *Request) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Request) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Request) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Request) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Request) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Request) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Request) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Request) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetOperationType

`func (o *Request) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *Request) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *Request) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *Request) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetAccessType

`func (o *Request) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Request) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Request) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *Request) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetConfirmationsLeft

`func (o *Request) GetConfirmationsLeft() int64`

GetConfirmationsLeft returns the ConfirmationsLeft field if non-nil, zero value otherwise.

### GetConfirmationsLeftOk

`func (o *Request) GetConfirmationsLeftOk() (*int64, bool)`

GetConfirmationsLeftOk returns a tuple with the ConfirmationsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsLeft

`func (o *Request) SetConfirmationsLeft(v int64)`

SetConfirmationsLeft sets ConfirmationsLeft field to given value.

### HasConfirmationsLeft

`func (o *Request) HasConfirmationsLeft() bool`

HasConfirmationsLeft returns a boolean if a field has been set.

### GetAccessFrom

`func (o *Request) GetAccessFrom() int64`

GetAccessFrom returns the AccessFrom field if non-nil, zero value otherwise.

### GetAccessFromOk

`func (o *Request) GetAccessFromOk() (*int64, bool)`

GetAccessFromOk returns a tuple with the AccessFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFrom

`func (o *Request) SetAccessFrom(v int64)`

SetAccessFrom sets AccessFrom field to given value.

### HasAccessFrom

`func (o *Request) HasAccessFrom() bool`

HasAccessFrom returns a boolean if a field has been set.

### GetAccessTo

`func (o *Request) GetAccessTo() int64`

GetAccessTo returns the AccessTo field if non-nil, zero value otherwise.

### GetAccessToOk

`func (o *Request) GetAccessToOk() (*int64, bool)`

GetAccessToOk returns a tuple with the AccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTo

`func (o *Request) SetAccessTo(v int64)`

SetAccessTo sets AccessTo field to given value.

### HasAccessTo

`func (o *Request) HasAccessTo() bool`

HasAccessTo returns a boolean if a field has been set.

### GetStatus

`func (o *Request) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Request) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Request) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Request) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTitle

`func (o *Request) GetStatusTitle() string`

GetStatusTitle returns the StatusTitle field if non-nil, zero value otherwise.

### GetStatusTitleOk

`func (o *Request) GetStatusTitleOk() (*string, bool)`

GetStatusTitleOk returns a tuple with the StatusTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTitle

`func (o *Request) SetStatusTitle(v string)`

SetStatusTitle sets StatusTitle field to given value.

### HasStatusTitle

`func (o *Request) HasStatusTitle() bool`

HasStatusTitle returns a boolean if a field has been set.

### GetInvalidRequestReason

`func (o *Request) GetInvalidRequestReason() int32`

GetInvalidRequestReason returns the InvalidRequestReason field if non-nil, zero value otherwise.

### GetInvalidRequestReasonOk

`func (o *Request) GetInvalidRequestReasonOk() (*int32, bool)`

GetInvalidRequestReasonOk returns a tuple with the InvalidRequestReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRequestReason

`func (o *Request) SetInvalidRequestReason(v int32)`

SetInvalidRequestReason sets InvalidRequestReason field to given value.

### HasInvalidRequestReason

`func (o *Request) HasInvalidRequestReason() bool`

HasInvalidRequestReason returns a boolean if a field has been set.

### GetCurrentConfirmationLevel

`func (o *Request) GetCurrentConfirmationLevel() int32`

GetCurrentConfirmationLevel returns the CurrentConfirmationLevel field if non-nil, zero value otherwise.

### GetCurrentConfirmationLevelOk

`func (o *Request) GetCurrentConfirmationLevelOk() (*int32, bool)`

GetCurrentConfirmationLevelOk returns a tuple with the CurrentConfirmationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmationLevel

`func (o *Request) SetCurrentConfirmationLevel(v int32)`

SetCurrentConfirmationLevel sets CurrentConfirmationLevel field to given value.

### HasCurrentConfirmationLevel

`func (o *Request) HasCurrentConfirmationLevel() bool`

HasCurrentConfirmationLevel returns a boolean if a field has been set.

### GetRequiredConfirmersCountLevel2

`func (o *Request) GetRequiredConfirmersCountLevel2() int32`

GetRequiredConfirmersCountLevel2 returns the RequiredConfirmersCountLevel2 field if non-nil, zero value otherwise.

### GetRequiredConfirmersCountLevel2Ok

`func (o *Request) GetRequiredConfirmersCountLevel2Ok() (*int32, bool)`

GetRequiredConfirmersCountLevel2Ok returns a tuple with the RequiredConfirmersCountLevel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfirmersCountLevel2

`func (o *Request) SetRequiredConfirmersCountLevel2(v int32)`

SetRequiredConfirmersCountLevel2 sets RequiredConfirmersCountLevel2 field to given value.

### HasRequiredConfirmersCountLevel2

`func (o *Request) HasRequiredConfirmersCountLevel2() bool`

HasRequiredConfirmersCountLevel2 returns a boolean if a field has been set.

### GetTicketingSystemProperties

`func (o *Request) GetTicketingSystemProperties() TicketingSystem`

GetTicketingSystemProperties returns the TicketingSystemProperties field if non-nil, zero value otherwise.

### GetTicketingSystemPropertiesOk

`func (o *Request) GetTicketingSystemPropertiesOk() (*TicketingSystem, bool)`

GetTicketingSystemPropertiesOk returns a tuple with the TicketingSystemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystemProperties

`func (o *Request) SetTicketingSystemProperties(v TicketingSystem)`

SetTicketingSystemProperties sets TicketingSystemProperties field to given value.

### HasTicketingSystemProperties

`func (o *Request) HasTicketingSystemProperties() bool`

HasTicketingSystemProperties returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Request) GetAdditionalInfo() map[string]string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Request) GetAdditionalInfoOk() (*map[string]string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Request) SetAdditionalInfo(v map[string]string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Request) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAccountDetails

`func (o *Request) GetAccountDetails() Account`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *Request) GetAccountDetailsOk() (*Account, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *Request) SetAccountDetails(v Account)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *Request) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### GetConfirmers

`func (o *Request) GetConfirmers() []ConfirmerStatus`

GetConfirmers returns the Confirmers field if non-nil, zero value otherwise.

### GetConfirmersOk

`func (o *Request) GetConfirmersOk() (*[]ConfirmerStatus, bool)`

GetConfirmersOk returns a tuple with the Confirmers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmers

`func (o *Request) SetConfirmers(v []ConfirmerStatus)`

SetConfirmers sets Confirmers field to given value.

### HasConfirmers

`func (o *Request) HasConfirmers() bool`

HasConfirmers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


