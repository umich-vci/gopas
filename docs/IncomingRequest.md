# IncomingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestorFullName** | Pointer to **string** |  | [optional] 
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

### NewIncomingRequest

`func NewIncomingRequest() *IncomingRequest`

NewIncomingRequest instantiates a new IncomingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingRequestWithDefaults

`func NewIncomingRequestWithDefaults() *IncomingRequest`

NewIncomingRequestWithDefaults instantiates a new IncomingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestorFullName

`func (o *IncomingRequest) GetRequestorFullName() string`

GetRequestorFullName returns the RequestorFullName field if non-nil, zero value otherwise.

### GetRequestorFullNameOk

`func (o *IncomingRequest) GetRequestorFullNameOk() (*string, bool)`

GetRequestorFullNameOk returns a tuple with the RequestorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorFullName

`func (o *IncomingRequest) SetRequestorFullName(v string)`

SetRequestorFullName sets RequestorFullName field to given value.

### HasRequestorFullName

`func (o *IncomingRequest) HasRequestorFullName() bool`

HasRequestorFullName returns a boolean if a field has been set.

### GetRequestID

`func (o *IncomingRequest) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *IncomingRequest) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *IncomingRequest) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *IncomingRequest) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetSafeName

`func (o *IncomingRequest) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *IncomingRequest) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *IncomingRequest) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.

### HasSafeName

`func (o *IncomingRequest) HasSafeName() bool`

HasSafeName returns a boolean if a field has been set.

### GetRequestorUserName

`func (o *IncomingRequest) GetRequestorUserName() string`

GetRequestorUserName returns the RequestorUserName field if non-nil, zero value otherwise.

### GetRequestorUserNameOk

`func (o *IncomingRequest) GetRequestorUserNameOk() (*string, bool)`

GetRequestorUserNameOk returns a tuple with the RequestorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorUserName

`func (o *IncomingRequest) SetRequestorUserName(v string)`

SetRequestorUserName sets RequestorUserName field to given value.

### HasRequestorUserName

`func (o *IncomingRequest) HasRequestorUserName() bool`

HasRequestorUserName returns a boolean if a field has been set.

### GetRequestorReason

`func (o *IncomingRequest) GetRequestorReason() string`

GetRequestorReason returns the RequestorReason field if non-nil, zero value otherwise.

### GetRequestorReasonOk

`func (o *IncomingRequest) GetRequestorReasonOk() (*string, bool)`

GetRequestorReasonOk returns a tuple with the RequestorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestorReason

`func (o *IncomingRequest) SetRequestorReason(v string)`

SetRequestorReason sets RequestorReason field to given value.

### HasRequestorReason

`func (o *IncomingRequest) HasRequestorReason() bool`

HasRequestorReason returns a boolean if a field has been set.

### GetUserReason

`func (o *IncomingRequest) GetUserReason() string`

GetUserReason returns the UserReason field if non-nil, zero value otherwise.

### GetUserReasonOk

`func (o *IncomingRequest) GetUserReasonOk() (*string, bool)`

GetUserReasonOk returns a tuple with the UserReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReason

`func (o *IncomingRequest) SetUserReason(v string)`

SetUserReason sets UserReason field to given value.

### HasUserReason

`func (o *IncomingRequest) HasUserReason() bool`

HasUserReason returns a boolean if a field has been set.

### GetCreationDate

`func (o *IncomingRequest) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *IncomingRequest) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *IncomingRequest) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *IncomingRequest) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetOperation

`func (o *IncomingRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IncomingRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IncomingRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *IncomingRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetExpirationDate

`func (o *IncomingRequest) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *IncomingRequest) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *IncomingRequest) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *IncomingRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetOperationType

`func (o *IncomingRequest) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *IncomingRequest) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *IncomingRequest) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *IncomingRequest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetAccessType

`func (o *IncomingRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *IncomingRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *IncomingRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *IncomingRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetConfirmationsLeft

`func (o *IncomingRequest) GetConfirmationsLeft() int64`

GetConfirmationsLeft returns the ConfirmationsLeft field if non-nil, zero value otherwise.

### GetConfirmationsLeftOk

`func (o *IncomingRequest) GetConfirmationsLeftOk() (*int64, bool)`

GetConfirmationsLeftOk returns a tuple with the ConfirmationsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsLeft

`func (o *IncomingRequest) SetConfirmationsLeft(v int64)`

SetConfirmationsLeft sets ConfirmationsLeft field to given value.

### HasConfirmationsLeft

`func (o *IncomingRequest) HasConfirmationsLeft() bool`

HasConfirmationsLeft returns a boolean if a field has been set.

### GetAccessFrom

`func (o *IncomingRequest) GetAccessFrom() int64`

GetAccessFrom returns the AccessFrom field if non-nil, zero value otherwise.

### GetAccessFromOk

`func (o *IncomingRequest) GetAccessFromOk() (*int64, bool)`

GetAccessFromOk returns a tuple with the AccessFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessFrom

`func (o *IncomingRequest) SetAccessFrom(v int64)`

SetAccessFrom sets AccessFrom field to given value.

### HasAccessFrom

`func (o *IncomingRequest) HasAccessFrom() bool`

HasAccessFrom returns a boolean if a field has been set.

### GetAccessTo

`func (o *IncomingRequest) GetAccessTo() int64`

GetAccessTo returns the AccessTo field if non-nil, zero value otherwise.

### GetAccessToOk

`func (o *IncomingRequest) GetAccessToOk() (*int64, bool)`

GetAccessToOk returns a tuple with the AccessTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTo

`func (o *IncomingRequest) SetAccessTo(v int64)`

SetAccessTo sets AccessTo field to given value.

### HasAccessTo

`func (o *IncomingRequest) HasAccessTo() bool`

HasAccessTo returns a boolean if a field has been set.

### GetStatus

`func (o *IncomingRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncomingRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncomingRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncomingRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTitle

`func (o *IncomingRequest) GetStatusTitle() string`

GetStatusTitle returns the StatusTitle field if non-nil, zero value otherwise.

### GetStatusTitleOk

`func (o *IncomingRequest) GetStatusTitleOk() (*string, bool)`

GetStatusTitleOk returns a tuple with the StatusTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTitle

`func (o *IncomingRequest) SetStatusTitle(v string)`

SetStatusTitle sets StatusTitle field to given value.

### HasStatusTitle

`func (o *IncomingRequest) HasStatusTitle() bool`

HasStatusTitle returns a boolean if a field has been set.

### GetInvalidRequestReason

`func (o *IncomingRequest) GetInvalidRequestReason() int32`

GetInvalidRequestReason returns the InvalidRequestReason field if non-nil, zero value otherwise.

### GetInvalidRequestReasonOk

`func (o *IncomingRequest) GetInvalidRequestReasonOk() (*int32, bool)`

GetInvalidRequestReasonOk returns a tuple with the InvalidRequestReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRequestReason

`func (o *IncomingRequest) SetInvalidRequestReason(v int32)`

SetInvalidRequestReason sets InvalidRequestReason field to given value.

### HasInvalidRequestReason

`func (o *IncomingRequest) HasInvalidRequestReason() bool`

HasInvalidRequestReason returns a boolean if a field has been set.

### GetCurrentConfirmationLevel

`func (o *IncomingRequest) GetCurrentConfirmationLevel() int32`

GetCurrentConfirmationLevel returns the CurrentConfirmationLevel field if non-nil, zero value otherwise.

### GetCurrentConfirmationLevelOk

`func (o *IncomingRequest) GetCurrentConfirmationLevelOk() (*int32, bool)`

GetCurrentConfirmationLevelOk returns a tuple with the CurrentConfirmationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfirmationLevel

`func (o *IncomingRequest) SetCurrentConfirmationLevel(v int32)`

SetCurrentConfirmationLevel sets CurrentConfirmationLevel field to given value.

### HasCurrentConfirmationLevel

`func (o *IncomingRequest) HasCurrentConfirmationLevel() bool`

HasCurrentConfirmationLevel returns a boolean if a field has been set.

### GetRequiredConfirmersCountLevel2

`func (o *IncomingRequest) GetRequiredConfirmersCountLevel2() int32`

GetRequiredConfirmersCountLevel2 returns the RequiredConfirmersCountLevel2 field if non-nil, zero value otherwise.

### GetRequiredConfirmersCountLevel2Ok

`func (o *IncomingRequest) GetRequiredConfirmersCountLevel2Ok() (*int32, bool)`

GetRequiredConfirmersCountLevel2Ok returns a tuple with the RequiredConfirmersCountLevel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfirmersCountLevel2

`func (o *IncomingRequest) SetRequiredConfirmersCountLevel2(v int32)`

SetRequiredConfirmersCountLevel2 sets RequiredConfirmersCountLevel2 field to given value.

### HasRequiredConfirmersCountLevel2

`func (o *IncomingRequest) HasRequiredConfirmersCountLevel2() bool`

HasRequiredConfirmersCountLevel2 returns a boolean if a field has been set.

### GetTicketingSystemProperties

`func (o *IncomingRequest) GetTicketingSystemProperties() TicketingSystem`

GetTicketingSystemProperties returns the TicketingSystemProperties field if non-nil, zero value otherwise.

### GetTicketingSystemPropertiesOk

`func (o *IncomingRequest) GetTicketingSystemPropertiesOk() (*TicketingSystem, bool)`

GetTicketingSystemPropertiesOk returns a tuple with the TicketingSystemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketingSystemProperties

`func (o *IncomingRequest) SetTicketingSystemProperties(v TicketingSystem)`

SetTicketingSystemProperties sets TicketingSystemProperties field to given value.

### HasTicketingSystemProperties

`func (o *IncomingRequest) HasTicketingSystemProperties() bool`

HasTicketingSystemProperties returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *IncomingRequest) GetAdditionalInfo() map[string]string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *IncomingRequest) GetAdditionalInfoOk() (*map[string]string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *IncomingRequest) SetAdditionalInfo(v map[string]string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *IncomingRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAccountDetails

`func (o *IncomingRequest) GetAccountDetails() Account`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *IncomingRequest) GetAccountDetailsOk() (*Account, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *IncomingRequest) SetAccountDetails(v Account)`

SetAccountDetails sets AccountDetails field to given value.

### HasAccountDetails

`func (o *IncomingRequest) HasAccountDetails() bool`

HasAccountDetails returns a boolean if a field has been set.

### GetConfirmers

`func (o *IncomingRequest) GetConfirmers() []ConfirmerStatus`

GetConfirmers returns the Confirmers field if non-nil, zero value otherwise.

### GetConfirmersOk

`func (o *IncomingRequest) GetConfirmersOk() (*[]ConfirmerStatus, bool)`

GetConfirmersOk returns a tuple with the Confirmers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmers

`func (o *IncomingRequest) SetConfirmers(v []ConfirmerStatus)`

SetConfirmers sets Confirmers field to given value.

### HasConfirmers

`func (o *IncomingRequest) HasConfirmers() bool`

HasConfirmers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


