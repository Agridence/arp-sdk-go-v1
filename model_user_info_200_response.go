/*
Agridence API

Agridence

API version: v1
Contact: it@agridence.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arp_sdk_go_v1

import (
	"encoding/json"
)

// checks if the UserInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfo200Response{}

// UserInfo200Response struct for UserInfo200Response
type UserInfo200Response struct {
	Company map[string]interface{} `json:"company"`
	Name string `json:"name"`
	SgxtsrSubscription NullableString `json:"sgxtsr_subscription"`
}

// NewUserInfo200Response instantiates a new UserInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo200Response(company map[string]interface{}, name string, sgxtsrSubscription NullableString) *UserInfo200Response {
	this := UserInfo200Response{}
	this.Company = company
	this.Name = name
	this.SgxtsrSubscription = sgxtsrSubscription
	return &this
}

// NewUserInfo200ResponseWithDefaults instantiates a new UserInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfo200ResponseWithDefaults() *UserInfo200Response {
	this := UserInfo200Response{}
	return &this
}

// GetCompany returns the Company field value
func (o *UserInfo200Response) GetCompany() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *UserInfo200Response) GetCompanyOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Company, true
}

// SetCompany sets field value
func (o *UserInfo200Response) SetCompany(v map[string]interface{}) {
	o.Company = v
}

// GetName returns the Name field value
func (o *UserInfo200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserInfo200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserInfo200Response) SetName(v string) {
	o.Name = v
}

// GetSgxtsrSubscription returns the SgxtsrSubscription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserInfo200Response) GetSgxtsrSubscription() string {
	if o == nil || o.SgxtsrSubscription.Get() == nil {
		var ret string
		return ret
	}

	return *o.SgxtsrSubscription.Get()
}

// GetSgxtsrSubscriptionOk returns a tuple with the SgxtsrSubscription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo200Response) GetSgxtsrSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SgxtsrSubscription.Get(), o.SgxtsrSubscription.IsSet()
}

// SetSgxtsrSubscription sets field value
func (o *UserInfo200Response) SetSgxtsrSubscription(v string) {
	o.SgxtsrSubscription.Set(&v)
}

func (o UserInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	toSerialize["name"] = o.Name
	toSerialize["sgxtsr_subscription"] = o.SgxtsrSubscription.Get()
	return toSerialize, nil
}

type NullableUserInfo200Response struct {
	value *UserInfo200Response
	isSet bool
}

func (v NullableUserInfo200Response) Get() *UserInfo200Response {
	return v.value
}

func (v *NullableUserInfo200Response) Set(val *UserInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo200Response(val *UserInfo200Response) *NullableUserInfo200Response {
	return &NullableUserInfo200Response{value: val, isSet: true}
}

func (v NullableUserInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


