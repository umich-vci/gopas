# PrivilegedAccessWorkflows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireDualControlPasswordAccessApproval** | Pointer to **bool** | Indicates whether accounts associated with this platform require dual control. valid values: true\\false | [optional] 
**EnforceCheckinCheckoutExclusiveAccess** | Pointer to **bool** | Indicates whether the Enforce check-in/check-out exclusive access rule is active or inactive for this platform. valid values: true\\false | [optional] 
**EnforceOnetimePasswordAccess** | Pointer to **bool** | Indicates whether the Enforce one-time password access rule is active for this platform. valid values: true\\false | [optional] 

## Methods

### NewPrivilegedAccessWorkflows

`func NewPrivilegedAccessWorkflows() *PrivilegedAccessWorkflows`

NewPrivilegedAccessWorkflows instantiates a new PrivilegedAccessWorkflows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedAccessWorkflowsWithDefaults

`func NewPrivilegedAccessWorkflowsWithDefaults() *PrivilegedAccessWorkflows`

NewPrivilegedAccessWorkflowsWithDefaults instantiates a new PrivilegedAccessWorkflows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireDualControlPasswordAccessApproval

`func (o *PrivilegedAccessWorkflows) GetRequireDualControlPasswordAccessApproval() bool`

GetRequireDualControlPasswordAccessApproval returns the RequireDualControlPasswordAccessApproval field if non-nil, zero value otherwise.

### GetRequireDualControlPasswordAccessApprovalOk

`func (o *PrivilegedAccessWorkflows) GetRequireDualControlPasswordAccessApprovalOk() (*bool, bool)`

GetRequireDualControlPasswordAccessApprovalOk returns a tuple with the RequireDualControlPasswordAccessApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireDualControlPasswordAccessApproval

`func (o *PrivilegedAccessWorkflows) SetRequireDualControlPasswordAccessApproval(v bool)`

SetRequireDualControlPasswordAccessApproval sets RequireDualControlPasswordAccessApproval field to given value.

### HasRequireDualControlPasswordAccessApproval

`func (o *PrivilegedAccessWorkflows) HasRequireDualControlPasswordAccessApproval() bool`

HasRequireDualControlPasswordAccessApproval returns a boolean if a field has been set.

### GetEnforceCheckinCheckoutExclusiveAccess

`func (o *PrivilegedAccessWorkflows) GetEnforceCheckinCheckoutExclusiveAccess() bool`

GetEnforceCheckinCheckoutExclusiveAccess returns the EnforceCheckinCheckoutExclusiveAccess field if non-nil, zero value otherwise.

### GetEnforceCheckinCheckoutExclusiveAccessOk

`func (o *PrivilegedAccessWorkflows) GetEnforceCheckinCheckoutExclusiveAccessOk() (*bool, bool)`

GetEnforceCheckinCheckoutExclusiveAccessOk returns a tuple with the EnforceCheckinCheckoutExclusiveAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceCheckinCheckoutExclusiveAccess

`func (o *PrivilegedAccessWorkflows) SetEnforceCheckinCheckoutExclusiveAccess(v bool)`

SetEnforceCheckinCheckoutExclusiveAccess sets EnforceCheckinCheckoutExclusiveAccess field to given value.

### HasEnforceCheckinCheckoutExclusiveAccess

`func (o *PrivilegedAccessWorkflows) HasEnforceCheckinCheckoutExclusiveAccess() bool`

HasEnforceCheckinCheckoutExclusiveAccess returns a boolean if a field has been set.

### GetEnforceOnetimePasswordAccess

`func (o *PrivilegedAccessWorkflows) GetEnforceOnetimePasswordAccess() bool`

GetEnforceOnetimePasswordAccess returns the EnforceOnetimePasswordAccess field if non-nil, zero value otherwise.

### GetEnforceOnetimePasswordAccessOk

`func (o *PrivilegedAccessWorkflows) GetEnforceOnetimePasswordAccessOk() (*bool, bool)`

GetEnforceOnetimePasswordAccessOk returns a tuple with the EnforceOnetimePasswordAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceOnetimePasswordAccess

`func (o *PrivilegedAccessWorkflows) SetEnforceOnetimePasswordAccess(v bool)`

SetEnforceOnetimePasswordAccess sets EnforceOnetimePasswordAccess field to given value.

### HasEnforceOnetimePasswordAccess

`func (o *PrivilegedAccessWorkflows) HasEnforceOnetimePasswordAccess() bool`

HasEnforceOnetimePasswordAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


