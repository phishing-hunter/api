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

// checks if the PostApikeyCreate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostApikeyCreate200Response{}

// PostApikeyCreate200Response struct for PostApikeyCreate200Response
type PostApikeyCreate200Response struct {
	ApiKey *string `json:"api_key,omitempty"`
}

// NewPostApikeyCreate200Response instantiates a new PostApikeyCreate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostApikeyCreate200Response() *PostApikeyCreate200Response {
	this := PostApikeyCreate200Response{}
	return &this
}

// NewPostApikeyCreate200ResponseWithDefaults instantiates a new PostApikeyCreate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostApikeyCreate200ResponseWithDefaults() *PostApikeyCreate200Response {
	this := PostApikeyCreate200Response{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *PostApikeyCreate200Response) GetApiKey() string {
	if o == nil || isNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostApikeyCreate200Response) GetApiKeyOk() (*string, bool) {
	if o == nil || isNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PostApikeyCreate200Response) HasApiKey() bool {
	if o != nil && !isNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *PostApikeyCreate200Response) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o PostApikeyCreate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostApikeyCreate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	return toSerialize, nil
}

type NullablePostApikeyCreate200Response struct {
	value *PostApikeyCreate200Response
	isSet bool
}

func (v NullablePostApikeyCreate200Response) Get() *PostApikeyCreate200Response {
	return v.value
}

func (v *NullablePostApikeyCreate200Response) Set(val *PostApikeyCreate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostApikeyCreate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostApikeyCreate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostApikeyCreate200Response(val *PostApikeyCreate200Response) *NullablePostApikeyCreate200Response {
	return &NullablePostApikeyCreate200Response{value: val, isSet: true}
}

func (v NullablePostApikeyCreate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostApikeyCreate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


