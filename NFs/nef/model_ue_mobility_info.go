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

// UeMobilityInfo Contains UE mobility information associated with an application.
type UeMobilityInfo struct {
	Supi Object `json:"supi"`
	AppId *Object `json:"appId,omitempty"`
	UeTrajs []UeTrajectoryInfo `json:"ueTrajs"`
}

// NewUeMobilityInfo instantiates a new UeMobilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeMobilityInfo(supi Object, ueTrajs []UeTrajectoryInfo) *UeMobilityInfo {
	this := UeMobilityInfo{}
	this.Supi = supi
	this.UeTrajs = ueTrajs
	return &this
}

// NewUeMobilityInfoWithDefaults instantiates a new UeMobilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeMobilityInfoWithDefaults() *UeMobilityInfo {
	this := UeMobilityInfo{}
	return &this
}

// GetSupi returns the Supi field value
func (o *UeMobilityInfo) GetSupi() Object {
	if o == nil {
		var ret Object
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetSupiOk() (*Object, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *UeMobilityInfo) SetSupi(v Object) {
	o.Supi = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UeMobilityInfo) GetAppId() Object {
	if o == nil || o.AppId == nil {
		var ret Object
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetAppIdOk() (*Object, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UeMobilityInfo) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given Object and assigns it to the AppId field.
func (o *UeMobilityInfo) SetAppId(v Object) {
	o.AppId = &v
}

// GetUeTrajs returns the UeTrajs field value
func (o *UeMobilityInfo) GetUeTrajs() []UeTrajectoryInfo {
	if o == nil {
		var ret []UeTrajectoryInfo
		return ret
	}

	return o.UeTrajs
}

// GetUeTrajsOk returns a tuple with the UeTrajs field value
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetUeTrajsOk() (*[]UeTrajectoryInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UeTrajs, true
}

// SetUeTrajs sets field value
func (o *UeMobilityInfo) SetUeTrajs(v []UeTrajectoryInfo) {
	o.UeTrajs = v
}

func (o UeMobilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["ueTrajs"] = o.UeTrajs
	}
	return json.Marshal(toSerialize)
}

type NullableUeMobilityInfo struct {
	value *UeMobilityInfo
	isSet bool
}

func (v NullableUeMobilityInfo) Get() *UeMobilityInfo {
	return v.value
}

func (v *NullableUeMobilityInfo) Set(val *UeMobilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeMobilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeMobilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeMobilityInfo(val *UeMobilityInfo) *NullableUeMobilityInfo {
	return &NullableUeMobilityInfo{value: val, isSet: true}
}

func (v NullableUeMobilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeMobilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


