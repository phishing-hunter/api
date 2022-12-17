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

// checks if the GetObservations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetObservations200Response{}

// GetObservations200Response struct for GetObservations200Response
type GetObservations200Response struct {
	ObservationUrls map[string]interface{} `json:"observation_urls,omitempty"`
}

// NewGetObservations200Response instantiates a new GetObservations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetObservations200Response() *GetObservations200Response {
	this := GetObservations200Response{}
	return &this
}

// NewGetObservations200ResponseWithDefaults instantiates a new GetObservations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetObservations200ResponseWithDefaults() *GetObservations200Response {
	this := GetObservations200Response{}
	return &this
}

// GetObservationUrls returns the ObservationUrls field value if set, zero value otherwise.
func (o *GetObservations200Response) GetObservationUrls() map[string]interface{} {
	if o == nil || isNil(o.ObservationUrls) {
		var ret map[string]interface{}
		return ret
	}
	return o.ObservationUrls
}

// GetObservationUrlsOk returns a tuple with the ObservationUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetObservations200Response) GetObservationUrlsOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.ObservationUrls) {
		return map[string]interface{}{}, false
	}
	return o.ObservationUrls, true
}

// HasObservationUrls returns a boolean if a field has been set.
func (o *GetObservations200Response) HasObservationUrls() bool {
	if o != nil && !isNil(o.ObservationUrls) {
		return true
	}

	return false
}

// SetObservationUrls gets a reference to the given map[string]interface{} and assigns it to the ObservationUrls field.
func (o *GetObservations200Response) SetObservationUrls(v map[string]interface{}) {
	o.ObservationUrls = v
}

func (o GetObservations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetObservations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObservationUrls) {
		toSerialize["observation_urls"] = o.ObservationUrls
	}
	return toSerialize, nil
}

type NullableGetObservations200Response struct {
	value *GetObservations200Response
	isSet bool
}

func (v NullableGetObservations200Response) Get() *GetObservations200Response {
	return v.value
}

func (v *NullableGetObservations200Response) Set(val *GetObservations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetObservations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetObservations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetObservations200Response(val *GetObservations200Response) *NullableGetObservations200Response {
	return &NullableGetObservations200Response{value: val, isSet: true}
}

func (v NullableGetObservations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetObservations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

