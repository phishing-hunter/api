/*
Phishing Hunter API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package phishing_hunter

import (
	"encoding/json"
)

// checks if the GetHunting200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetHunting200Response{}

// GetHunting200Response struct for GetHunting200Response
type GetHunting200Response struct {
	Keywords *GetHunting200ResponseKeywords `json:"keywords,omitempty"`
	TargetScore *int32 `json:"target_score,omitempty"`
}

// NewGetHunting200Response instantiates a new GetHunting200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHunting200Response() *GetHunting200Response {
	this := GetHunting200Response{}
	return &this
}

// NewGetHunting200ResponseWithDefaults instantiates a new GetHunting200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHunting200ResponseWithDefaults() *GetHunting200Response {
	this := GetHunting200Response{}
	return &this
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *GetHunting200Response) GetKeywords() GetHunting200ResponseKeywords {
	if o == nil || isNil(o.Keywords) {
		var ret GetHunting200ResponseKeywords
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHunting200Response) GetKeywordsOk() (*GetHunting200ResponseKeywords, bool) {
	if o == nil || isNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *GetHunting200Response) HasKeywords() bool {
	if o != nil && !isNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given GetHunting200ResponseKeywords and assigns it to the Keywords field.
func (o *GetHunting200Response) SetKeywords(v GetHunting200ResponseKeywords) {
	o.Keywords = &v
}

// GetTargetScore returns the TargetScore field value if set, zero value otherwise.
func (o *GetHunting200Response) GetTargetScore() int32 {
	if o == nil || isNil(o.TargetScore) {
		var ret int32
		return ret
	}
	return *o.TargetScore
}

// GetTargetScoreOk returns a tuple with the TargetScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHunting200Response) GetTargetScoreOk() (*int32, bool) {
	if o == nil || isNil(o.TargetScore) {
		return nil, false
	}
	return o.TargetScore, true
}

// HasTargetScore returns a boolean if a field has been set.
func (o *GetHunting200Response) HasTargetScore() bool {
	if o != nil && !isNil(o.TargetScore) {
		return true
	}

	return false
}

// SetTargetScore gets a reference to the given int32 and assigns it to the TargetScore field.
func (o *GetHunting200Response) SetTargetScore(v int32) {
	o.TargetScore = &v
}

func (o GetHunting200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetHunting200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !isNil(o.TargetScore) {
		toSerialize["target_score"] = o.TargetScore
	}
	return toSerialize, nil
}

type NullableGetHunting200Response struct {
	value *GetHunting200Response
	isSet bool
}

func (v NullableGetHunting200Response) Get() *GetHunting200Response {
	return v.value
}

func (v *NullableGetHunting200Response) Set(val *GetHunting200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHunting200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHunting200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHunting200Response(val *GetHunting200Response) *NullableGetHunting200Response {
	return &NullableGetHunting200Response{value: val, isSet: true}
}

func (v NullableGetHunting200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHunting200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


