# AutomaticOnboardingRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **int64** | A numeric identifier for the rule, assigned by the system. | [optional] 
**RuleName** | Pointer to **string** | Name of the rule. This is either auto-generated or specified by the user when the rule is created. | [optional] 
**RuleDescription** | Pointer to **string** | A description of the rule. | [optional] 
**RulePrecedence** | Pointer to **int32** | The order in which the rules are run.  Rules are ordered based on creation time.The most recently created rule will have a precedence of 1, the next most recently created rule will have a precedence of 2, and so on.  During discovery, when a new account is discovered, it is first compared to the rule with precedence 1 to check if the account matches the rule&#39;s filters. If so, the account is onboarded according to the rule. If not, the account is compared to the next rule by precedence, and so on. | [optional] 
**TargetPlatformId** | Pointer to **string** | The ID of the platform that the onboarded account will be associated with. | [optional] 
**TargetDeviceType** | Pointer to **string** | Device type of the Target Platform | [optional] 
**TargetSafeName** | Pointer to **string** | The name of the Safe where the onboarded account will be stored.  Security requirement: If the user is not an owner of the Safe, a null string will be returned. | [optional] 
**IsAdminIDFilter** | Pointer to **bool** | Whether or not only accounts with the following admin ID will be onboarded automatically according to this rule.  Unix accounts whose UID is 0.  Windows accounts whose SID ends with 500.  If this value is set to false, the admin ID will not be considered and all accounts matching the rule will be onboarded. | [optional] 
**MachineTypeFilter** | Pointer to **string** | The Machine Type by which to filter. | [optional] 
**SystemTypeFilter** | Pointer to **string** | The System Type by which to filter. | [optional] 
**CreationTime** | Pointer to **int32** | The date and time when the rule was created. | [optional] 
**UserNameFilter** | Pointer to **string** | The name of the user by which to filter. | [optional] 
**UserNameMethod** | Pointer to **string** | The method to use when applying the username filter. | [optional] 
**AddressFilter** | Pointer to **string** | The IP address or DNS domain name of the machine by which to filter. | [optional] 
**AddressMethod** | Pointer to **string** | The method to use when applying the address filter. | [optional] 
**AccountTypeFilter** | Pointer to **string** |  | [optional] 
**AccountCategoryFilter** | Pointer to **string** | Filter for privileged or non-privileged accounts. | [optional] 
**ReconcileAccountId** | Pointer to **string** |  | [optional] 
**LastOnboardedTime** | Pointer to **int32** | The last time that an account was successfully onboarded using this rule. | [optional] 

## Methods

### NewAutomaticOnboardingRuleData

`func NewAutomaticOnboardingRuleData() *AutomaticOnboardingRuleData`

NewAutomaticOnboardingRuleData instantiates a new AutomaticOnboardingRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticOnboardingRuleDataWithDefaults

`func NewAutomaticOnboardingRuleDataWithDefaults() *AutomaticOnboardingRuleData`

NewAutomaticOnboardingRuleDataWithDefaults instantiates a new AutomaticOnboardingRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *AutomaticOnboardingRuleData) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AutomaticOnboardingRuleData) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AutomaticOnboardingRuleData) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AutomaticOnboardingRuleData) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *AutomaticOnboardingRuleData) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *AutomaticOnboardingRuleData) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *AutomaticOnboardingRuleData) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *AutomaticOnboardingRuleData) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetRuleDescription

`func (o *AutomaticOnboardingRuleData) GetRuleDescription() string`

GetRuleDescription returns the RuleDescription field if non-nil, zero value otherwise.

### GetRuleDescriptionOk

`func (o *AutomaticOnboardingRuleData) GetRuleDescriptionOk() (*string, bool)`

GetRuleDescriptionOk returns a tuple with the RuleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDescription

`func (o *AutomaticOnboardingRuleData) SetRuleDescription(v string)`

SetRuleDescription sets RuleDescription field to given value.

### HasRuleDescription

`func (o *AutomaticOnboardingRuleData) HasRuleDescription() bool`

HasRuleDescription returns a boolean if a field has been set.

### GetRulePrecedence

`func (o *AutomaticOnboardingRuleData) GetRulePrecedence() int32`

GetRulePrecedence returns the RulePrecedence field if non-nil, zero value otherwise.

### GetRulePrecedenceOk

`func (o *AutomaticOnboardingRuleData) GetRulePrecedenceOk() (*int32, bool)`

GetRulePrecedenceOk returns a tuple with the RulePrecedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePrecedence

`func (o *AutomaticOnboardingRuleData) SetRulePrecedence(v int32)`

SetRulePrecedence sets RulePrecedence field to given value.

### HasRulePrecedence

`func (o *AutomaticOnboardingRuleData) HasRulePrecedence() bool`

HasRulePrecedence returns a boolean if a field has been set.

### GetTargetPlatformId

`func (o *AutomaticOnboardingRuleData) GetTargetPlatformId() string`

GetTargetPlatformId returns the TargetPlatformId field if non-nil, zero value otherwise.

### GetTargetPlatformIdOk

`func (o *AutomaticOnboardingRuleData) GetTargetPlatformIdOk() (*string, bool)`

GetTargetPlatformIdOk returns a tuple with the TargetPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatformId

`func (o *AutomaticOnboardingRuleData) SetTargetPlatformId(v string)`

SetTargetPlatformId sets TargetPlatformId field to given value.

### HasTargetPlatformId

`func (o *AutomaticOnboardingRuleData) HasTargetPlatformId() bool`

HasTargetPlatformId returns a boolean if a field has been set.

### GetTargetDeviceType

`func (o *AutomaticOnboardingRuleData) GetTargetDeviceType() string`

GetTargetDeviceType returns the TargetDeviceType field if non-nil, zero value otherwise.

### GetTargetDeviceTypeOk

`func (o *AutomaticOnboardingRuleData) GetTargetDeviceTypeOk() (*string, bool)`

GetTargetDeviceTypeOk returns a tuple with the TargetDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeviceType

`func (o *AutomaticOnboardingRuleData) SetTargetDeviceType(v string)`

SetTargetDeviceType sets TargetDeviceType field to given value.

### HasTargetDeviceType

`func (o *AutomaticOnboardingRuleData) HasTargetDeviceType() bool`

HasTargetDeviceType returns a boolean if a field has been set.

### GetTargetSafeName

`func (o *AutomaticOnboardingRuleData) GetTargetSafeName() string`

GetTargetSafeName returns the TargetSafeName field if non-nil, zero value otherwise.

### GetTargetSafeNameOk

`func (o *AutomaticOnboardingRuleData) GetTargetSafeNameOk() (*string, bool)`

GetTargetSafeNameOk returns a tuple with the TargetSafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSafeName

`func (o *AutomaticOnboardingRuleData) SetTargetSafeName(v string)`

SetTargetSafeName sets TargetSafeName field to given value.

### HasTargetSafeName

`func (o *AutomaticOnboardingRuleData) HasTargetSafeName() bool`

HasTargetSafeName returns a boolean if a field has been set.

### GetIsAdminIDFilter

`func (o *AutomaticOnboardingRuleData) GetIsAdminIDFilter() bool`

GetIsAdminIDFilter returns the IsAdminIDFilter field if non-nil, zero value otherwise.

### GetIsAdminIDFilterOk

`func (o *AutomaticOnboardingRuleData) GetIsAdminIDFilterOk() (*bool, bool)`

GetIsAdminIDFilterOk returns a tuple with the IsAdminIDFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminIDFilter

`func (o *AutomaticOnboardingRuleData) SetIsAdminIDFilter(v bool)`

SetIsAdminIDFilter sets IsAdminIDFilter field to given value.

### HasIsAdminIDFilter

`func (o *AutomaticOnboardingRuleData) HasIsAdminIDFilter() bool`

HasIsAdminIDFilter returns a boolean if a field has been set.

### GetMachineTypeFilter

`func (o *AutomaticOnboardingRuleData) GetMachineTypeFilter() string`

GetMachineTypeFilter returns the MachineTypeFilter field if non-nil, zero value otherwise.

### GetMachineTypeFilterOk

`func (o *AutomaticOnboardingRuleData) GetMachineTypeFilterOk() (*string, bool)`

GetMachineTypeFilterOk returns a tuple with the MachineTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeFilter

`func (o *AutomaticOnboardingRuleData) SetMachineTypeFilter(v string)`

SetMachineTypeFilter sets MachineTypeFilter field to given value.

### HasMachineTypeFilter

`func (o *AutomaticOnboardingRuleData) HasMachineTypeFilter() bool`

HasMachineTypeFilter returns a boolean if a field has been set.

### GetSystemTypeFilter

`func (o *AutomaticOnboardingRuleData) GetSystemTypeFilter() string`

GetSystemTypeFilter returns the SystemTypeFilter field if non-nil, zero value otherwise.

### GetSystemTypeFilterOk

`func (o *AutomaticOnboardingRuleData) GetSystemTypeFilterOk() (*string, bool)`

GetSystemTypeFilterOk returns a tuple with the SystemTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTypeFilter

`func (o *AutomaticOnboardingRuleData) SetSystemTypeFilter(v string)`

SetSystemTypeFilter sets SystemTypeFilter field to given value.

### HasSystemTypeFilter

`func (o *AutomaticOnboardingRuleData) HasSystemTypeFilter() bool`

HasSystemTypeFilter returns a boolean if a field has been set.

### GetCreationTime

`func (o *AutomaticOnboardingRuleData) GetCreationTime() int32`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AutomaticOnboardingRuleData) GetCreationTimeOk() (*int32, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AutomaticOnboardingRuleData) SetCreationTime(v int32)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AutomaticOnboardingRuleData) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUserNameFilter

`func (o *AutomaticOnboardingRuleData) GetUserNameFilter() string`

GetUserNameFilter returns the UserNameFilter field if non-nil, zero value otherwise.

### GetUserNameFilterOk

`func (o *AutomaticOnboardingRuleData) GetUserNameFilterOk() (*string, bool)`

GetUserNameFilterOk returns a tuple with the UserNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameFilter

`func (o *AutomaticOnboardingRuleData) SetUserNameFilter(v string)`

SetUserNameFilter sets UserNameFilter field to given value.

### HasUserNameFilter

`func (o *AutomaticOnboardingRuleData) HasUserNameFilter() bool`

HasUserNameFilter returns a boolean if a field has been set.

### GetUserNameMethod

`func (o *AutomaticOnboardingRuleData) GetUserNameMethod() string`

GetUserNameMethod returns the UserNameMethod field if non-nil, zero value otherwise.

### GetUserNameMethodOk

`func (o *AutomaticOnboardingRuleData) GetUserNameMethodOk() (*string, bool)`

GetUserNameMethodOk returns a tuple with the UserNameMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameMethod

`func (o *AutomaticOnboardingRuleData) SetUserNameMethod(v string)`

SetUserNameMethod sets UserNameMethod field to given value.

### HasUserNameMethod

`func (o *AutomaticOnboardingRuleData) HasUserNameMethod() bool`

HasUserNameMethod returns a boolean if a field has been set.

### GetAddressFilter

`func (o *AutomaticOnboardingRuleData) GetAddressFilter() string`

GetAddressFilter returns the AddressFilter field if non-nil, zero value otherwise.

### GetAddressFilterOk

`func (o *AutomaticOnboardingRuleData) GetAddressFilterOk() (*string, bool)`

GetAddressFilterOk returns a tuple with the AddressFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFilter

`func (o *AutomaticOnboardingRuleData) SetAddressFilter(v string)`

SetAddressFilter sets AddressFilter field to given value.

### HasAddressFilter

`func (o *AutomaticOnboardingRuleData) HasAddressFilter() bool`

HasAddressFilter returns a boolean if a field has been set.

### GetAddressMethod

`func (o *AutomaticOnboardingRuleData) GetAddressMethod() string`

GetAddressMethod returns the AddressMethod field if non-nil, zero value otherwise.

### GetAddressMethodOk

`func (o *AutomaticOnboardingRuleData) GetAddressMethodOk() (*string, bool)`

GetAddressMethodOk returns a tuple with the AddressMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMethod

`func (o *AutomaticOnboardingRuleData) SetAddressMethod(v string)`

SetAddressMethod sets AddressMethod field to given value.

### HasAddressMethod

`func (o *AutomaticOnboardingRuleData) HasAddressMethod() bool`

HasAddressMethod returns a boolean if a field has been set.

### GetAccountTypeFilter

`func (o *AutomaticOnboardingRuleData) GetAccountTypeFilter() string`

GetAccountTypeFilter returns the AccountTypeFilter field if non-nil, zero value otherwise.

### GetAccountTypeFilterOk

`func (o *AutomaticOnboardingRuleData) GetAccountTypeFilterOk() (*string, bool)`

GetAccountTypeFilterOk returns a tuple with the AccountTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeFilter

`func (o *AutomaticOnboardingRuleData) SetAccountTypeFilter(v string)`

SetAccountTypeFilter sets AccountTypeFilter field to given value.

### HasAccountTypeFilter

`func (o *AutomaticOnboardingRuleData) HasAccountTypeFilter() bool`

HasAccountTypeFilter returns a boolean if a field has been set.

### GetAccountCategoryFilter

`func (o *AutomaticOnboardingRuleData) GetAccountCategoryFilter() string`

GetAccountCategoryFilter returns the AccountCategoryFilter field if non-nil, zero value otherwise.

### GetAccountCategoryFilterOk

`func (o *AutomaticOnboardingRuleData) GetAccountCategoryFilterOk() (*string, bool)`

GetAccountCategoryFilterOk returns a tuple with the AccountCategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryFilter

`func (o *AutomaticOnboardingRuleData) SetAccountCategoryFilter(v string)`

SetAccountCategoryFilter sets AccountCategoryFilter field to given value.

### HasAccountCategoryFilter

`func (o *AutomaticOnboardingRuleData) HasAccountCategoryFilter() bool`

HasAccountCategoryFilter returns a boolean if a field has been set.

### GetReconcileAccountId

`func (o *AutomaticOnboardingRuleData) GetReconcileAccountId() string`

GetReconcileAccountId returns the ReconcileAccountId field if non-nil, zero value otherwise.

### GetReconcileAccountIdOk

`func (o *AutomaticOnboardingRuleData) GetReconcileAccountIdOk() (*string, bool)`

GetReconcileAccountIdOk returns a tuple with the ReconcileAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconcileAccountId

`func (o *AutomaticOnboardingRuleData) SetReconcileAccountId(v string)`

SetReconcileAccountId sets ReconcileAccountId field to given value.

### HasReconcileAccountId

`func (o *AutomaticOnboardingRuleData) HasReconcileAccountId() bool`

HasReconcileAccountId returns a boolean if a field has been set.

### GetLastOnboardedTime

`func (o *AutomaticOnboardingRuleData) GetLastOnboardedTime() int32`

GetLastOnboardedTime returns the LastOnboardedTime field if non-nil, zero value otherwise.

### GetLastOnboardedTimeOk

`func (o *AutomaticOnboardingRuleData) GetLastOnboardedTimeOk() (*int32, bool)`

GetLastOnboardedTimeOk returns a tuple with the LastOnboardedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnboardedTime

`func (o *AutomaticOnboardingRuleData) SetLastOnboardedTime(v int32)`

SetLastOnboardedTime sets LastOnboardedTime field to given value.

### HasLastOnboardedTime

`func (o *AutomaticOnboardingRuleData) HasLastOnboardedTime() bool`

HasLastOnboardedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


