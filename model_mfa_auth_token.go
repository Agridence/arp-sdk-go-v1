/*
Agridence API

Agridence

API version: v1.1
Contact: it@agridence.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arp_sdk_go_v1

import (
	"encoding/json"
)

// checks if the MFAAuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAAuthToken{}

// MFAAuthToken struct for MFAAuthToken
type MFAAuthToken struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token *string `json:"token,omitempty"`
}

// NewMFAAuthToken instantiates a new MFAAuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAAuthToken(email string, password string) *MFAAuthToken {
	this := MFAAuthToken{}
	this.Email = email
	this.Password = password
	return &this
}

// NewMFAAuthTokenWithDefaults instantiates a new MFAAuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAAuthTokenWithDefaults() *MFAAuthToken {
	this := MFAAuthToken{}
	return &this
}

// GetEmail returns the Email field value
func (o *MFAAuthToken) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *MFAAuthToken) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *MFAAuthToken) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *MFAAuthToken) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *MFAAuthToken) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *MFAAuthToken) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MFAAuthToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAAuthToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MFAAuthToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MFAAuthToken) SetToken(v string) {
	o.Token = &v
}

func (o MFAAuthToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAAuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	// skip: token is readOnly
	return toSerialize, nil
}

type NullableMFAAuthToken struct {
	value *MFAAuthToken
	isSet bool
}

func (v NullableMFAAuthToken) Get() *MFAAuthToken {
	return v.value
}

func (v *NullableMFAAuthToken) Set(val *MFAAuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAAuthToken(val *MFAAuthToken) *NullableMFAAuthToken {
	return &NullableMFAAuthToken{value: val, isSet: true}
}

func (v NullableMFAAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


