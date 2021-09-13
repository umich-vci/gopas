/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopas

import (
	"encoding/json"
)

// AccountModel struct for AccountModel
type AccountModel struct {
	CategoryModificationTime *int64  `json:"categoryModificationTime,omitempty"`
	Id                       *string `json:"id,omitempty"`
	// The name of the account.
	Name *string `json:"name,omitempty"`
	// The name or address of the machine where the account will be used.  Valid values: DNS/IP/URL where the account is managed
	Address *string `json:"address,omitempty"`
	// Account user's name.
	UserName *string `json:"userName,omitempty"`
	// The platform assigned to this account.  Valid values: Valid platform IDs, example: WinServerLocal
	PlatformId string `json:"platformId"`
	// The Safe where the account will be created.
	SafeName string `json:"safeName"`
	// Description The type of password.  Valid values password, key
	SecretType *string `json:"secretType,omitempty"`
	Secret     *string `json:"secret,omitempty"`
	// Object containing key-value pairs to associate with the account, as defined by the account platform.  These properties are validated against the mandatory and optional properties of the specified platform's definition.  Optional properties that do not exist on the account will not be returned here.  Internal properties are not returned.  Valid values example: {\"Location\": \"IT\", \"OwnerName\":\"MSSPAdmin\"}
	PlatformAccountProperties *map[string]string         `json:"platformAccountProperties,omitempty"`
	SecretManagement          *AutomaticSecretManagement `json:"secretManagement,omitempty"`
	RemoteMachinesAccess      *RemoteMachinesAccess      `json:"remoteMachinesAccess,omitempty"`
	// Date and time account was created
	CreatedTime *int64 `json:"createdTime,omitempty"`
}

// NewAccountModel instantiates a new AccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountModel(platformId string, safeName string) *AccountModel {
	this := AccountModel{}
	this.PlatformId = platformId
	this.SafeName = safeName
	return &this
}

// NewAccountModelWithDefaults instantiates a new AccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountModelWithDefaults() *AccountModel {
	this := AccountModel{}
	return &this
}

// GetCategoryModificationTime returns the CategoryModificationTime field value if set, zero value otherwise.
func (o *AccountModel) GetCategoryModificationTime() int64 {
	if o == nil || o.CategoryModificationTime == nil {
		var ret int64
		return ret
	}
	return *o.CategoryModificationTime
}

// GetCategoryModificationTimeOk returns a tuple with the CategoryModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetCategoryModificationTimeOk() (*int64, bool) {
	if o == nil || o.CategoryModificationTime == nil {
		return nil, false
	}
	return o.CategoryModificationTime, true
}

// HasCategoryModificationTime returns a boolean if a field has been set.
func (o *AccountModel) HasCategoryModificationTime() bool {
	if o != nil && o.CategoryModificationTime != nil {
		return true
	}

	return false
}

// SetCategoryModificationTime gets a reference to the given int64 and assigns it to the CategoryModificationTime field.
func (o *AccountModel) SetCategoryModificationTime(v int64) {
	o.CategoryModificationTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountModel) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AccountModel) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountModel) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AccountModel) SetAddress(v string) {
	o.Address = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AccountModel) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AccountModel) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AccountModel) SetUserName(v string) {
	o.UserName = &v
}

// GetPlatformId returns the PlatformId field value
func (o *AccountModel) GetPlatformId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value
// and a boolean to check if the value has been set.
func (o *AccountModel) GetPlatformIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlatformId, true
}

// SetPlatformId sets field value
func (o *AccountModel) SetPlatformId(v string) {
	o.PlatformId = v
}

// GetSafeName returns the SafeName field value
func (o *AccountModel) GetSafeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value
// and a boolean to check if the value has been set.
func (o *AccountModel) GetSafeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SafeName, true
}

// SetSafeName sets field value
func (o *AccountModel) SetSafeName(v string) {
	o.SafeName = v
}

// GetSecretType returns the SecretType field value if set, zero value otherwise.
func (o *AccountModel) GetSecretType() string {
	if o == nil || o.SecretType == nil {
		var ret string
		return ret
	}
	return *o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetSecretTypeOk() (*string, bool) {
	if o == nil || o.SecretType == nil {
		return nil, false
	}
	return o.SecretType, true
}

// HasSecretType returns a boolean if a field has been set.
func (o *AccountModel) HasSecretType() bool {
	if o != nil && o.SecretType != nil {
		return true
	}

	return false
}

// SetSecretType gets a reference to the given string and assigns it to the SecretType field.
func (o *AccountModel) SetSecretType(v string) {
	o.SecretType = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AccountModel) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AccountModel) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AccountModel) SetSecret(v string) {
	o.Secret = &v
}

// GetPlatformAccountProperties returns the PlatformAccountProperties field value if set, zero value otherwise.
func (o *AccountModel) GetPlatformAccountProperties() map[string]string {
	if o == nil || o.PlatformAccountProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.PlatformAccountProperties
}

// GetPlatformAccountPropertiesOk returns a tuple with the PlatformAccountProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetPlatformAccountPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.PlatformAccountProperties == nil {
		return nil, false
	}
	return o.PlatformAccountProperties, true
}

// HasPlatformAccountProperties returns a boolean if a field has been set.
func (o *AccountModel) HasPlatformAccountProperties() bool {
	if o != nil && o.PlatformAccountProperties != nil {
		return true
	}

	return false
}

// SetPlatformAccountProperties gets a reference to the given map[string]string and assigns it to the PlatformAccountProperties field.
func (o *AccountModel) SetPlatformAccountProperties(v map[string]string) {
	o.PlatformAccountProperties = &v
}

// GetSecretManagement returns the SecretManagement field value if set, zero value otherwise.
func (o *AccountModel) GetSecretManagement() AutomaticSecretManagement {
	if o == nil || o.SecretManagement == nil {
		var ret AutomaticSecretManagement
		return ret
	}
	return *o.SecretManagement
}

// GetSecretManagementOk returns a tuple with the SecretManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetSecretManagementOk() (*AutomaticSecretManagement, bool) {
	if o == nil || o.SecretManagement == nil {
		return nil, false
	}
	return o.SecretManagement, true
}

// HasSecretManagement returns a boolean if a field has been set.
func (o *AccountModel) HasSecretManagement() bool {
	if o != nil && o.SecretManagement != nil {
		return true
	}

	return false
}

// SetSecretManagement gets a reference to the given AutomaticSecretManagement and assigns it to the SecretManagement field.
func (o *AccountModel) SetSecretManagement(v AutomaticSecretManagement) {
	o.SecretManagement = &v
}

// GetRemoteMachinesAccess returns the RemoteMachinesAccess field value if set, zero value otherwise.
func (o *AccountModel) GetRemoteMachinesAccess() RemoteMachinesAccess {
	if o == nil || o.RemoteMachinesAccess == nil {
		var ret RemoteMachinesAccess
		return ret
	}
	return *o.RemoteMachinesAccess
}

// GetRemoteMachinesAccessOk returns a tuple with the RemoteMachinesAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetRemoteMachinesAccessOk() (*RemoteMachinesAccess, bool) {
	if o == nil || o.RemoteMachinesAccess == nil {
		return nil, false
	}
	return o.RemoteMachinesAccess, true
}

// HasRemoteMachinesAccess returns a boolean if a field has been set.
func (o *AccountModel) HasRemoteMachinesAccess() bool {
	if o != nil && o.RemoteMachinesAccess != nil {
		return true
	}

	return false
}

// SetRemoteMachinesAccess gets a reference to the given RemoteMachinesAccess and assigns it to the RemoteMachinesAccess field.
func (o *AccountModel) SetRemoteMachinesAccess(v RemoteMachinesAccess) {
	o.RemoteMachinesAccess = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *AccountModel) GetCreatedTime() int64 {
	if o == nil || o.CreatedTime == nil {
		var ret int64
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountModel) GetCreatedTimeOk() (*int64, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *AccountModel) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int64 and assigns it to the CreatedTime field.
func (o *AccountModel) SetCreatedTime(v int64) {
	o.CreatedTime = &v
}

func (o AccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryModificationTime != nil {
		toSerialize["categoryModificationTime"] = o.CategoryModificationTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["platformId"] = o.PlatformId
	}
	if true {
		toSerialize["safeName"] = o.SafeName
	}
	if o.SecretType != nil {
		toSerialize["secretType"] = o.SecretType
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.PlatformAccountProperties != nil {
		toSerialize["platformAccountProperties"] = o.PlatformAccountProperties
	}
	if o.SecretManagement != nil {
		toSerialize["secretManagement"] = o.SecretManagement
	}
	if o.RemoteMachinesAccess != nil {
		toSerialize["remoteMachinesAccess"] = o.RemoteMachinesAccess
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	return json.Marshal(toSerialize)
}

type NullableAccountModel struct {
	value *AccountModel
	isSet bool
}

func (v NullableAccountModel) Get() *AccountModel {
	return v.value
}

func (v *NullableAccountModel) Set(val *AccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountModel(val *AccountModel) *NullableAccountModel {
	return &NullableAccountModel{value: val, isSet: true}
}

func (v NullableAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
