# LDAPDirectoryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | **string** | The address of the domain. | 
**DomainBaseContext** | **string** | The base context of the External Directory. | 

## Methods

### NewLDAPDirectoryBase

`func NewLDAPDirectoryBase(domainName string, domainBaseContext string, ) *LDAPDirectoryBase`

NewLDAPDirectoryBase instantiates a new LDAPDirectoryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDirectoryBaseWithDefaults

`func NewLDAPDirectoryBaseWithDefaults() *LDAPDirectoryBase`

NewLDAPDirectoryBaseWithDefaults instantiates a new LDAPDirectoryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *LDAPDirectoryBase) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *LDAPDirectoryBase) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *LDAPDirectoryBase) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetDomainBaseContext

`func (o *LDAPDirectoryBase) GetDomainBaseContext() string`

GetDomainBaseContext returns the DomainBaseContext field if non-nil, zero value otherwise.

### GetDomainBaseContextOk

`func (o *LDAPDirectoryBase) GetDomainBaseContextOk() (*string, bool)`

GetDomainBaseContextOk returns a tuple with the DomainBaseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainBaseContext

`func (o *LDAPDirectoryBase) SetDomainBaseContext(v string)`

SetDomainBaseContext sets DomainBaseContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


