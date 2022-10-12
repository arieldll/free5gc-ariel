/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TargetUeIdentification Identifies the UE to which the request applies.
type TargetUeIdentification struct {
	Supis *[]Object `json:"supis,omitempty"`
	InterGroupIds *[]Object `json:"interGroupIds,omitempty"`
	AnyUeId *bool `json:"anyUeId,omitempty"`
}

// NewTargetUeIdentification instantiates a new TargetUeIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetUeIdentification() *TargetUeIdentification {
	this := TargetUeIdentification{}
	return &this
}

// NewTargetUeIdentificationWithDefaults instantiates a new TargetUeIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetUeIdentificationWithDefaults() *TargetUeIdentification {
	this := TargetUeIdentification{}
	return &this
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *TargetUeIdentification) GetSupis() []Object {
	if o == nil || o.Supis == nil {
		var ret []Object
		return ret
	}
	return *o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeIdentification) GetSupisOk() (*[]Object, bool) {
	if o == nil || o.Supis == nil {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *TargetUeIdentification) HasSupis() bool {
	if o != nil && o.Supis != nil {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []Object and assigns it to the Supis field.
func (o *TargetUeIdentification) SetSupis(v []Object) {
	o.Supis = &v
}

// GetInterGroupIds returns the InterGroupIds field value if set, zero value otherwise.
func (o *TargetUeIdentification) GetInterGroupIds() []Object {
	if o == nil || o.InterGroupIds == nil {
		var ret []Object
		return ret
	}
	return *o.InterGroupIds
}

// GetInterGroupIdsOk returns a tuple with the InterGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeIdentification) GetInterGroupIdsOk() (*[]Object, bool) {
	if o == nil || o.InterGroupIds == nil {
		return nil, false
	}
	return o.InterGroupIds, true
}

// HasInterGroupIds returns a boolean if a field has been set.
func (o *TargetUeIdentification) HasInterGroupIds() bool {
	if o != nil && o.InterGroupIds != nil {
		return true
	}

	return false
}

// SetInterGroupIds gets a reference to the given []Object and assigns it to the InterGroupIds field.
func (o *TargetUeIdentification) SetInterGroupIds(v []Object) {
	o.InterGroupIds = &v
}

// GetAnyUeId returns the AnyUeId field value if set, zero value otherwise.
func (o *TargetUeIdentification) GetAnyUeId() bool {
	if o == nil || o.AnyUeId == nil {
		var ret bool
		return ret
	}
	return *o.AnyUeId
}

// GetAnyUeIdOk returns a tuple with the AnyUeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetUeIdentification) GetAnyUeIdOk() (*bool, bool) {
	if o == nil || o.AnyUeId == nil {
		return nil, false
	}
	return o.AnyUeId, true
}

// HasAnyUeId returns a boolean if a field has been set.
func (o *TargetUeIdentification) HasAnyUeId() bool {
	if o != nil && o.AnyUeId != nil {
		return true
	}

	return false
}

// SetAnyUeId gets a reference to the given bool and assigns it to the AnyUeId field.
func (o *TargetUeIdentification) SetAnyUeId(v bool) {
	o.AnyUeId = &v
}

func (o TargetUeIdentification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Supis != nil {
		toSerialize["supis"] = o.Supis
	}
	if o.InterGroupIds != nil {
		toSerialize["interGroupIds"] = o.InterGroupIds
	}
	if o.AnyUeId != nil {
		toSerialize["anyUeId"] = o.AnyUeId
	}
	return json.Marshal(toSerialize)
}

type NullableTargetUeIdentification struct {
	value *TargetUeIdentification
	isSet bool
}

func (v NullableTargetUeIdentification) Get() *TargetUeIdentification {
	return v.value
}

func (v *NullableTargetUeIdentification) Set(val *TargetUeIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetUeIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetUeIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetUeIdentification(val *TargetUeIdentification) *NullableTargetUeIdentification {
	return &NullableTargetUeIdentification{value: val, isSet: true}
}

func (v NullableTargetUeIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetUeIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


