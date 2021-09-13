# AutomaticOnboardingRuleDataIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetPlatformId** | **string** | The ID of the platform that will be associated to the onboarded account. | 
**TargetSafeName** | **string** | The name of the Safe where the onboarded account will be stored. | 
**IsAdminIDFilter** | Pointer to **bool** |  | [optional] 
**MachineTypeFilter** | Pointer to **int32** |  | [optional] 
**SystemTypeFilter** | **int32** | The System Type by which to filter. | 
**UserNameFilter** | Pointer to **string** | The name of the user by which to filter. | [optional] 
**RuleName** | Pointer to **string** |  | [optional] 
**RuleDescription** | Pointer to **string** |  | [optional] 
**UserNameMethod** | Pointer to **int32** |  | [optional] 
**AddressFilter** | Pointer to **string** | The Machine Type by which to filter. | [optional] 
**AddressMethod** | Pointer to **int32** |  | [optional] 
**AccountCategoryFilter** | Pointer to **int32** |  | [optional] 

## Methods

### NewAutomaticOnboardingRuleDataIn

`func NewAutomaticOnboardingRuleDataIn(targetPlatformId string, targetSafeName string, systemTypeFilter int32, ) *AutomaticOnboardingRuleDataIn`

NewAutomaticOnboardingRuleDataIn instantiates a new AutomaticOnboardingRuleDataIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticOnboardingRuleDataInWithDefaults

`func NewAutomaticOnboardingRuleDataInWithDefaults() *AutomaticOnboardingRuleDataIn`

NewAutomaticOnboardingRuleDataInWithDefaults instantiates a new AutomaticOnboardingRuleDataIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetPlatformId

`func (o *AutomaticOnboardingRuleDataIn) GetTargetPlatformId() string`

GetTargetPlatformId returns the TargetPlatformId field if non-nil, zero value otherwise.

### GetTargetPlatformIdOk

`func (o *AutomaticOnboardingRuleDataIn) GetTargetPlatformIdOk() (*string, bool)`

GetTargetPlatformIdOk returns a tuple with the TargetPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatformId

`func (o *AutomaticOnboardingRuleDataIn) SetTargetPlatformId(v string)`

SetTargetPlatformId sets TargetPlatformId field to given value.


### GetTargetSafeName

`func (o *AutomaticOnboardingRuleDataIn) GetTargetSafeName() string`

GetTargetSafeName returns the TargetSafeName field if non-nil, zero value otherwise.

### GetTargetSafeNameOk

`func (o *AutomaticOnboardingRuleDataIn) GetTargetSafeNameOk() (*string, bool)`

GetTargetSafeNameOk returns a tuple with the TargetSafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSafeName

`func (o *AutomaticOnboardingRuleDataIn) SetTargetSafeName(v string)`

SetTargetSafeName sets TargetSafeName field to given value.


### GetIsAdminIDFilter

`func (o *AutomaticOnboardingRuleDataIn) GetIsAdminIDFilter() bool`

GetIsAdminIDFilter returns the IsAdminIDFilter field if non-nil, zero value otherwise.

### GetIsAdminIDFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetIsAdminIDFilterOk() (*bool, bool)`

GetIsAdminIDFilterOk returns a tuple with the IsAdminIDFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminIDFilter

`func (o *AutomaticOnboardingRuleDataIn) SetIsAdminIDFilter(v bool)`

SetIsAdminIDFilter sets IsAdminIDFilter field to given value.

### HasIsAdminIDFilter

`func (o *AutomaticOnboardingRuleDataIn) HasIsAdminIDFilter() bool`

HasIsAdminIDFilter returns a boolean if a field has been set.

### GetMachineTypeFilter

`func (o *AutomaticOnboardingRuleDataIn) GetMachineTypeFilter() int32`

GetMachineTypeFilter returns the MachineTypeFilter field if non-nil, zero value otherwise.

### GetMachineTypeFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetMachineTypeFilterOk() (*int32, bool)`

GetMachineTypeFilterOk returns a tuple with the MachineTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeFilter

`func (o *AutomaticOnboardingRuleDataIn) SetMachineTypeFilter(v int32)`

SetMachineTypeFilter sets MachineTypeFilter field to given value.

### HasMachineTypeFilter

`func (o *AutomaticOnboardingRuleDataIn) HasMachineTypeFilter() bool`

HasMachineTypeFilter returns a boolean if a field has been set.

### GetSystemTypeFilter

`func (o *AutomaticOnboardingRuleDataIn) GetSystemTypeFilter() int32`

GetSystemTypeFilter returns the SystemTypeFilter field if non-nil, zero value otherwise.

### GetSystemTypeFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetSystemTypeFilterOk() (*int32, bool)`

GetSystemTypeFilterOk returns a tuple with the SystemTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTypeFilter

`func (o *AutomaticOnboardingRuleDataIn) SetSystemTypeFilter(v int32)`

SetSystemTypeFilter sets SystemTypeFilter field to given value.


### GetUserNameFilter

`func (o *AutomaticOnboardingRuleDataIn) GetUserNameFilter() string`

GetUserNameFilter returns the UserNameFilter field if non-nil, zero value otherwise.

### GetUserNameFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetUserNameFilterOk() (*string, bool)`

GetUserNameFilterOk returns a tuple with the UserNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameFilter

`func (o *AutomaticOnboardingRuleDataIn) SetUserNameFilter(v string)`

SetUserNameFilter sets UserNameFilter field to given value.

### HasUserNameFilter

`func (o *AutomaticOnboardingRuleDataIn) HasUserNameFilter() bool`

HasUserNameFilter returns a boolean if a field has been set.

### GetRuleName

`func (o *AutomaticOnboardingRuleDataIn) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *AutomaticOnboardingRuleDataIn) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *AutomaticOnboardingRuleDataIn) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *AutomaticOnboardingRuleDataIn) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *AutomaticOnboardingRuleDataIn) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *AutomaticOnboardingRuleDataIn) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *AutomaticOnboardingRuleDataIn) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *AutomaticOnboardingRuleDataIn) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetUserNameMethod

`func (o *AutomaticOnboardingRuleDataIn) GetUserNameMethod() int32`

GetUserNameMethod returns the UserNameMethod field if non-nil, zero value otherwise.

### GetUserNameMethodOk

`func (o *AutomaticOnboardingRuleDataIn) GetUserNameMethodOk() (*int32, bool)`

GetUserNameMethodOk returns a tuple with the UserNameMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameMethod

`func (o *AutomaticOnboardingRuleDataIn) SetUserNameMethod(v int32)`

SetUserNameMethod sets UserNameMethod field to given value.

### HasUserNameMethod

`func (o *AutomaticOnboardingRuleDataIn) HasUserNameMethod() bool`

HasUserNameMethod returns a boolean if a field has been set.

### GetAddressFilter

`func (o *AutomaticOnboardingRuleDataIn) GetAddressFilter() string`

GetAddressFilter returns the AddressFilter field if non-nil, zero value otherwise.

### GetAddressFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetAddressFilterOk() (*string, bool)`

GetAddressFilterOk returns a tuple with the AddressFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFilter

`func (o *AutomaticOnboardingRuleDataIn) SetAddressFilter(v string)`

SetAddressFilter sets AddressFilter field to given value.

### HasAddressFilter

`func (o *AutomaticOnboardingRuleDataIn) HasAddressFilter() bool`

HasAddressFilter returns a boolean if a field has been set.

### GetAddressMethod

`func (o *AutomaticOnboardingRuleDataIn) GetAddressMethod() int32`

GetAddressMethod returns the AddressMethod field if non-nil, zero value otherwise.

### GetAddressMethodOk

`func (o *AutomaticOnboardingRuleDataIn) GetAddressMethodOk() (*int32, bool)`

GetAddressMethodOk returns a tuple with the AddressMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMethod

`func (o *AutomaticOnboardingRuleDataIn) SetAddressMethod(v int32)`

SetAddressMethod sets AddressMethod field to given value.

### HasAddressMethod

`func (o *AutomaticOnboardingRuleDataIn) HasAddressMethod() bool`

HasAddressMethod returns a boolean if a field has been set.

### GetAccountCategoryFilter

`func (o *AutomaticOnboardingRuleDataIn) GetAccountCategoryFilter() int32`

GetAccountCategoryFilter returns the AccountCategoryFilter field if non-nil, zero value otherwise.

### GetAccountCategoryFilterOk

`func (o *AutomaticOnboardingRuleDataIn) GetAccountCategoryFilterOk() (*int32, bool)`

GetAccountCategoryFilterOk returns a tuple with the AccountCategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryFilter

`func (o *AutomaticOnboardingRuleDataIn) SetAccountCategoryFilter(v int32)`

SetAccountCategoryFilter sets AccountCategoryFilter field to given value.

### HasAccountCategoryFilter

`func (o *AutomaticOnboardingRuleDataIn) HasAccountCategoryFilter() bool`

HasAccountCategoryFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


