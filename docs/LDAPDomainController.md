# LDAPDomainController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the server machine where the external directory is installed. | 
**Port** | Pointer to **int32** | The port number through which the external directory can be accessed. | [optional] 
**SSLConnect** | Pointer to **bool** | Whether or not to connect to the external directory with SSL. | [optional] 

## Methods

### NewLDAPDomainController

`func NewLDAPDomainController(name string, ) *LDAPDomainController`

NewLDAPDomainController instantiates a new LDAPDomainController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPDomainControllerWithDefaults

`func NewLDAPDomainControllerWithDefaults() *LDAPDomainController`

NewLDAPDomainControllerWithDefaults instantiates a new LDAPDomainController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LDAPDomainController) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPDomainController) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPDomainController) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *LDAPDomainController) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPDomainController) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPDomainController) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPDomainController) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSSLConnect

`func (o *LDAPDomainController) GetSSLConnect() bool`

GetSSLConnect returns the SSLConnect field if non-nil, zero value otherwise.

### GetSSLConnectOk

`func (o *LDAPDomainController) GetSSLConnectOk() (*bool, bool)`

GetSSLConnectOk returns a tuple with the SSLConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLConnect

`func (o *LDAPDomainController) SetSSLConnect(v bool)`

SetSSLConnect sets SSLConnect field to given value.

### HasSSLConnect

`func (o *LDAPDomainController) HasSSLConnect() bool`

HasSSLConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


