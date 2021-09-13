# LDAPMappingsOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingsOrder** | **[]int64** | The new order of the mappings, sorted by importance.   Must include all mappings.  Each ID can only appear once. | 

## Methods

### NewLDAPMappingsOrder

`func NewLDAPMappingsOrder(mappingsOrder []int64, ) *LDAPMappingsOrder`

NewLDAPMappingsOrder instantiates a new LDAPMappingsOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPMappingsOrderWithDefaults

`func NewLDAPMappingsOrderWithDefaults() *LDAPMappingsOrder`

NewLDAPMappingsOrderWithDefaults instantiates a new LDAPMappingsOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappingsOrder

`func (o *LDAPMappingsOrder) GetMappingsOrder() []int64`

GetMappingsOrder returns the MappingsOrder field if non-nil, zero value otherwise.

### GetMappingsOrderOk

`func (o *LDAPMappingsOrder) GetMappingsOrderOk() (*[]int64, bool)`

GetMappingsOrderOk returns a tuple with the MappingsOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingsOrder

`func (o *LDAPMappingsOrder) SetMappingsOrder(v []int64)`

SetMappingsOrder sets MappingsOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


