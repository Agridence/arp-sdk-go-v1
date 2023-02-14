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

// checks if the OfferPanelFilterPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferPanelFilterPort{}

// OfferPanelFilterPort struct for OfferPanelFilterPort
type OfferPanelFilterPort struct {
	Id *int32 `json:"id,omitempty"`
	PortCode *string `json:"port_code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewOfferPanelFilterPort instantiates a new OfferPanelFilterPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferPanelFilterPort() *OfferPanelFilterPort {
	this := OfferPanelFilterPort{}
	return &this
}

// NewOfferPanelFilterPortWithDefaults instantiates a new OfferPanelFilterPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferPanelFilterPortWithDefaults() *OfferPanelFilterPort {
	this := OfferPanelFilterPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfferPanelFilterPort) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPort) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfferPanelFilterPort) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OfferPanelFilterPort) SetId(v int32) {
	o.Id = &v
}

// GetPortCode returns the PortCode field value if set, zero value otherwise.
func (o *OfferPanelFilterPort) GetPortCode() string {
	if o == nil || isNil(o.PortCode) {
		var ret string
		return ret
	}
	return *o.PortCode
}

// GetPortCodeOk returns a tuple with the PortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPort) GetPortCodeOk() (*string, bool) {
	if o == nil || isNil(o.PortCode) {
		return nil, false
	}
	return o.PortCode, true
}

// HasPortCode returns a boolean if a field has been set.
func (o *OfferPanelFilterPort) HasPortCode() bool {
	if o != nil && !isNil(o.PortCode) {
		return true
	}

	return false
}

// SetPortCode gets a reference to the given string and assigns it to the PortCode field.
func (o *OfferPanelFilterPort) SetPortCode(v string) {
	o.PortCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OfferPanelFilterPort) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPanelFilterPort) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OfferPanelFilterPort) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OfferPanelFilterPort) SetName(v string) {
	o.Name = &v
}

func (o OfferPanelFilterPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferPanelFilterPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: port_code is readOnly
	// skip: name is readOnly
	return toSerialize, nil
}

type NullableOfferPanelFilterPort struct {
	value *OfferPanelFilterPort
	isSet bool
}

func (v NullableOfferPanelFilterPort) Get() *OfferPanelFilterPort {
	return v.value
}

func (v *NullableOfferPanelFilterPort) Set(val *OfferPanelFilterPort) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferPanelFilterPort) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferPanelFilterPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferPanelFilterPort(val *OfferPanelFilterPort) *NullableOfferPanelFilterPort {
	return &NullableOfferPanelFilterPort{value: val, isSet: true}
}

func (v NullableOfferPanelFilterPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferPanelFilterPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

